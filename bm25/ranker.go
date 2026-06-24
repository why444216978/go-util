package bm25

import (
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	crawlabbm25 "github.com/crawlab-team/bm25"
	"github.com/go-ego/gse"
)

// dictRelPath gse 字典在项目内的相对路径（相对于当前工作目录）
const dictRelPath = "data/dict/zh"

// resolveDictFiles 基于当前工作目录定位字典文件，返回 gse 接受的逗号分隔格式。
func resolveDictFiles() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("bm25: get cwd failed: %w", err)
	}
	dir := filepath.Join(cwd, dictRelPath)
	s := filepath.Join(dir, "s_1.txt")
	t := filepath.Join(dir, "t_1.txt")
	if _, err := os.Stat(s); err != nil {
		return "", fmt.Errorf("bm25: dict file not found: %s: %w", s, err)
	}
	if _, err := os.Stat(t); err != nil {
		return "", fmt.Errorf("bm25: dict file not found: %s: %w", t, err)
	}
	return s + ", " + t, nil
}

// IndexHandle BM25 索引句柄（不透明标识）
type IndexHandle string

// Document 待排序的文档
type Document struct {
	ID      string // 文档唯一标识
	Content string // 文档文本内容
}

// RankResult BM25 排序结果
type RankResult struct {
	ID    string  // 文档唯一标识
	Score float64 // BM25 分数
}

// Ranker BM25 排序器接口
type Ranker interface {
	// BuildIndex 基于 documents 构建 BM25 索引并缓存，返回句柄
	BuildIndex(documents []Document) (IndexHandle, error)

	// RankByIndex 基于已缓存的索引进行 BM25 排序，返回 TopN
	RankByIndex(handle IndexHandle, query string, topN int) ([]RankResult, error)

	// RemoveIndex 手动移除缓存的索引
	RemoveIndex(handle IndexHandle)
}

var (
	ErrIndexNotFound  = errors.New("bm25: index not found or expired")
	ErrEmptyDocuments = errors.New("bm25: documents cannot be empty")
)

// cacheEntry 缓存条目
type cacheEntry struct {
	mu        sync.Mutex
	bm25      *crawlabbm25.BM25Okapi
	documents []Document
	createdAt time.Time
}

// rankerImpl Ranker 实现
type rankerImpl struct {
	seg   gse.Segmenter
	cache sync.Map // map[IndexHandle]*cacheEntry
	ttl   time.Duration
	k1    float64
	b     float64
}

// tokenize 合并精确模式与搜索模式分词结果，保留完整词的同时产生子词
func (r *rankerImpl) tokenize(s string) []string {
	c := strings.ReplaceAll(s, " ", "")
	precise := r.seg.Cut(c, true)
	search := r.seg.CutSearch(c)
	seen := make(map[string]struct{}, len(precise)+len(search))
	result := make([]string, 0, len(precise)+len(search))
	for _, t := range precise {
		if _, ok := seen[t]; !ok {
			seen[t] = struct{}{}
			result = append(result, t)
		}
	}
	for _, t := range search {
		if _, ok := seen[t]; !ok {
			seen[t] = struct{}{}
			result = append(result, t)
		}
	}
	return result
}

// NewRankerWithParams
func NewRanker(dictFiles []string, ttl time.Duration, k1, b float64) (Ranker, error) {
	seg, err := gse.New(dictFiles...)
	if err != nil {
		return nil, fmt.Errorf("bm25: gse.New failed with dict %q: %w", dictFiles, err)
	}
	r := &rankerImpl{seg: seg, ttl: ttl, k1: k1, b: b}
	go r.cleanupLoop()
	return r, nil
}

// computeHandle 计算 documents 的 IndexHandle
// 按 ID 升序排序后，拼接 ID+"\x00"+Content 计算 MD5
func computeHandle(documents []Document) IndexHandle {
	sorted := make([]Document, len(documents))
	copy(sorted, documents)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].ID < sorted[j].ID
	})

	h := md5.New()
	for _, doc := range sorted {
		h.Write([]byte(doc.ID))
		h.Write([]byte{0})
		h.Write([]byte(doc.Content))
		h.Write([]byte{0})
	}
	return IndexHandle(fmt.Sprintf("%x", h.Sum(nil)))
}

// BuildIndex 基于 documents 构建 BM25 索引并缓存
func (r *rankerImpl) BuildIndex(documents []Document) (IndexHandle, error) {
	if len(documents) == 0 {
		return "", ErrEmptyDocuments
	}

	handle := computeHandle(documents)

	// 幂等：已存在则直接返回
	if _, ok := r.cache.Load(handle); ok {
		return handle, nil
	}

	// 构建 corpus
	corpus := make([]string, len(documents))
	for i, doc := range documents {
		corpus[i] = doc.Content
	}

	// 创建 BM25 Okapi 实例。
	bm25Inst, err := crawlabbm25.NewBM25Okapi(corpus, r.tokenize, r.k1, r.b, nil)
	if err != nil {
		return "", err
	}

	// 保存 documents 副本用于结果映射（保持传入顺序，与 corpus 一致）
	docsCopy := make([]Document, len(documents))
	copy(docsCopy, documents)

	entry := &cacheEntry{
		bm25:      bm25Inst,
		documents: docsCopy,
		createdAt: time.Now(),
	}

	// LoadOrStore 保证并发 BuildIndex 只有一个写入成功
	r.cache.LoadOrStore(handle, entry)
	return handle, nil
}

// RankByIndex 基于已缓存的索引进行 BM25 排序
func (r *rankerImpl) RankByIndex(handle IndexHandle, query string, topN int) ([]RankResult, error) {
	if query == "" || topN <= 0 {
		return nil, nil
	}

	val, ok := r.cache.Load(handle)
	if !ok {
		return nil, ErrIndexNotFound
	}
	entry, ok := val.(*cacheEntry)
	if !ok || entry == nil {
		return nil, ErrIndexNotFound
	}

	// 检查是否已过期
	if time.Since(entry.createdAt) > r.ttl {
		r.cache.Delete(handle)
		return nil, ErrIndexNotFound
	}

	// 对 query 分词
	queryTokens := r.tokenize(query)
	if len(queryTokens) == 0 {
		return nil, nil
	}

	// 加锁保护 GetScores（idfCache 非并发安全）
	entry.mu.Lock()
	if entry.bm25 == nil {
		entry.mu.Unlock()
		return nil, ErrIndexNotFound
	}
	scores, err := entry.bm25.GetScores(queryTokens)
	entry.mu.Unlock()
	if err != nil {
		return nil, err
	}

	// 构建结果并按分数降序排序
	type indexedScore struct {
		index int
		score float64
	}
	scored := make([]indexedScore, 0, len(scores))
	for i, s := range scores {
		scored = append(scored, indexedScore{index: i, score: s})
	}

	sort.Slice(scored, func(i, j int) bool {
		return scored[i].score > scored[j].score
	})

	if topN > len(scored) {
		topN = len(scored)
	}

	results := make([]RankResult, topN)
	for i := 0; i < topN; i++ {
		results[i] = RankResult{
			ID:    entry.documents[scored[i].index].ID,
			Score: scored[i].score,
		}
	}

	return results, nil
}

// RemoveIndex 手动移除缓存的索引
func (r *rankerImpl) RemoveIndex(handle IndexHandle) {
	r.cache.Delete(handle)
}

// cleanupLoop 后台定期清理过期缓存
func (r *rankerImpl) cleanupLoop() {
	ticker := time.NewTicker(r.ttl / 2)
	defer ticker.Stop()
	for range ticker.C {
		r.cache.Range(func(key, value any) bool {
			entry, ok := value.(*cacheEntry)
			if !ok || entry == nil {
				r.cache.Delete(key)
				return true
			}
			if time.Since(entry.createdAt) > r.ttl {
				r.cache.Delete(key)
			}
			return true
		})
	}
}
