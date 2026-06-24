package bm25

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestRanker(t *testing.T, ttl ...time.Duration) Ranker {
	t.Helper()

	var r Ranker
	var err error
	if len(ttl) > 0 {
		r, err = NewRanker([]string{"./dict.txt"}, ttl[0], 1.5, 0.5)
	} else {
		r, err = NewRanker([]string{"./dict.txt"}, 5*time.Minute, 1.5, 0.5)
	}

	require.NoError(t, err)
	return r
}

// ===== BuildIndex Tests (Cases 1-5) =====

func TestBuildIndex_Basic(t *testing.T) {
	r := newTestRanker(t)
	docs := []Document{
		{ID: "1", Content: "你好世界"},
		{ID: "2", Content: "测试文档"},
	}
	handle, err := r.BuildIndex(docs)
	assert.NoError(t, err)
	assert.NotEmpty(t, string(handle))
}

func TestBuildIndex_Idempotent(t *testing.T) {
	r := newTestRanker(t)
	docs := []Document{
		{ID: "1", Content: "你好世界"},
		{ID: "2", Content: "测试文档"},
	}

	h1, err := r.BuildIndex(docs)
	assert.NoError(t, err)
	h2, err := r.BuildIndex(docs)
	assert.NoError(t, err)
	assert.Equal(t, h1, h2)
}

func TestBuildIndex_DifferentDocs(t *testing.T) {
	r := newTestRanker(t)
	docs1 := []Document{{ID: "1", Content: "你好世界"}}
	docs2 := []Document{{ID: "2", Content: "不同内容"}}

	h1, err := r.BuildIndex(docs1)
	assert.NoError(t, err)
	h2, err := r.BuildIndex(docs2)
	assert.NoError(t, err)
	assert.NotEqual(t, h1, h2)
}

func TestBuildIndex_OrderIndependent(t *testing.T) {
	r := newTestRanker(t)
	docs1 := []Document{
		{ID: "1", Content: "你好世界"},
		{ID: "2", Content: "测试文档"},
	}
	docs2 := []Document{
		{ID: "2", Content: "测试文档"},
		{ID: "1", Content: "你好世界"},
	}

	h1, err := r.BuildIndex(docs1)
	assert.NoError(t, err)
	h2, err := r.BuildIndex(docs2)
	assert.NoError(t, err)
	assert.Equal(t, h1, h2)
}

func TestBuildIndex_EmptyDocs(t *testing.T) {
	r := newTestRanker(t)
	_, err := r.BuildIndex(nil)
	assert.ErrorIs(t, err, ErrEmptyDocuments)

	_, err = r.BuildIndex([]Document{})
	assert.ErrorIs(t, err, ErrEmptyDocuments)
}

// ===== RankByIndex Tests (Cases 6-9) =====

// 注意：BM25 IDF = log((N+1-df+0.5)/(df+0.5))
// 当 df = (N+1)/2 时 IDF=0，所以测试数据需要确保查询词出现在少于半数文档中

func TestRankByIndex_Basic(t *testing.T) {
	r := newTestRanker(t)
	// "水果"只出现在 doc1，df=1, N=5, IDF>0
	docs := []Document{
		{ID: "1", Content: "苹果是一种常见的水果"},
		{ID: "2", Content: "今天天气很好适合出去玩"},
		{ID: "3", Content: "晚上要早点睡觉休息"},
		{ID: "4", Content: "北京的交通非常拥堵"},
		{ID: "5", Content: "学习编程需要大量练习"},
	}
	handle, err := r.BuildIndex(docs)
	require.NoError(t, err)

	results, err := r.RankByIndex(handle, "水果", 3)
	require.NoError(t, err)
	require.NotEmpty(t, results)
	// doc1 包含"水果"，应该排第一
	assert.Equal(t, "1", results[0].ID)
	assert.True(t, results[0].Score > 0)
}

func TestRankByIndex_TopN(t *testing.T) {
	r := newTestRanker(t)
	// "学习"只出现在 doc1 和 doc2，df=2, N=5, IDF>0
	docs := []Document{
		{ID: "1", Content: "学习数学需要认真思考"},
		{ID: "2", Content: "学习英语要多听多说"},
		{ID: "3", Content: "运动有助于身体健康"},
		{ID: "4", Content: "阅读可以开阔视野"},
		{ID: "5", Content: "旅行可以增长见识"},
	}
	handle, err := r.BuildIndex(docs)
	require.NoError(t, err)

	results, err := r.RankByIndex(handle, "学习", 1)
	require.NoError(t, err)
	// 虽然有2个文档含"学习"，topN=1 只返回1个
	assert.Equal(t, 1, len(results))
}

func TestRankByIndex_InvalidHandle(t *testing.T) {
	r := newTestRanker(t)
	_, err := r.RankByIndex(IndexHandle("nonexistent"), "query", 5)
	assert.ErrorIs(t, err, ErrIndexNotFound)
}

func TestRankByIndex_EmptyQuery(t *testing.T) {
	r := newTestRanker(t)
	docs := []Document{{ID: "1", Content: "测试文档内容"}}
	handle, err := r.BuildIndex(docs)
	require.NoError(t, err)

	results, err := r.RankByIndex(handle, "", 5)
	assert.NoError(t, err)
	assert.Nil(t, results)

	results, err = r.RankByIndex(handle, "test", 0)
	assert.NoError(t, err)
	assert.Nil(t, results)
}

// ===== RemoveIndex and TTL Tests (Cases 10-11) =====

func TestRemoveIndex(t *testing.T) {
	r := newTestRanker(t)
	docs := []Document{{ID: "1", Content: "测试文档内容"}}
	handle, err := r.BuildIndex(docs)
	require.NoError(t, err)

	r.RemoveIndex(handle)

	_, err = r.RankByIndex(handle, "测试", 5)
	assert.ErrorIs(t, err, ErrIndexNotFound)
}

func TestTTLExpiry(t *testing.T) {
	// 使用极短 TTL
	r := newTestRanker(t, 100*time.Millisecond)

	// "测试"只出现在 doc1，确保 IDF>0
	docs := []Document{
		{ID: "1", Content: "这是一个测试文档"},
		{ID: "2", Content: "今天天气很好"},
		{ID: "3", Content: "北京欢迎你"},
	}
	handle, err := r.BuildIndex(docs)
	require.NoError(t, err)

	// 立即查询应成功
	results, err := r.RankByIndex(handle, "测试", 5)
	require.NoError(t, err)
	require.NotEmpty(t, results)

	// 等待过期
	time.Sleep(150 * time.Millisecond)

	// 过期后查询应失败
	_, err = r.RankByIndex(handle, "测试", 5)
	assert.ErrorIs(t, err, ErrIndexNotFound)
}

// ===== Concurrency Tests (Cases 12-14) =====

func TestConcurrentRankByIndex_SameHandle(t *testing.T) {
	r := newTestRanker(t)
	// "水果"只出现在 doc1，IDF>0
	docs := []Document{
		{ID: "1", Content: "苹果是一种常见的水果"},
		{ID: "2", Content: "今天天气很好适合散步"},
		{ID: "3", Content: "晚上要早点睡觉休息"},
		{ID: "4", Content: "北京的交通非常拥堵"},
		{ID: "5", Content: "学习编程需要大量练习"},
	}
	handle, err := r.BuildIndex(docs)
	require.NoError(t, err)

	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results, err := r.RankByIndex(handle, "水果", 3)
			assert.NoError(t, err)
			assert.NotEmpty(t, results)
		}()
	}
	wg.Wait()
}

func TestConcurrentRankByIndex_MultiHandle(t *testing.T) {
	r := newTestRanker(t)
	// "水果" 只在 docs1 的 doc1 中出现
	docs1 := []Document{
		{ID: "1", Content: "苹果是一种常见的水果"},
		{ID: "2", Content: "今天天气很好适合散步"},
		{ID: "3", Content: "晚上要早点睡觉休息"},
	}
	// "天气" 只在 docs2 的 doc3 中出现
	docs2 := []Document{
		{ID: "3", Content: "今天天气很好适合运动"},
		{ID: "4", Content: "学习编程需要大量练习"},
		{ID: "5", Content: "北京的交通非常拥堵"},
	}
	h1, err := r.BuildIndex(docs1)
	require.NoError(t, err)
	h2, err := r.BuildIndex(docs2)
	require.NoError(t, err)

	var wg sync.WaitGroup
	for range 10 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			results, err := r.RankByIndex(h1, "水果", 2)
			assert.NoError(t, err)
			assert.NotEmpty(t, results)
		}()
		go func() {
			defer wg.Done()
			results, err := r.RankByIndex(h2, "天气", 2)
			assert.NoError(t, err)
			assert.NotEmpty(t, results)
		}()
	}
	wg.Wait()
}

func TestConcurrentMixedOps(t *testing.T) {
	r := newTestRanker(t)

	var wg sync.WaitGroup
	for i := range 20 {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			docs := []Document{
				{ID: fmt.Sprintf("doc-%d", idx), Content: "并发测试文档内容"},
				{ID: fmt.Sprintf("other-%d", idx), Content: "另一个完全不同的文档"},
				{ID: fmt.Sprintf("third-%d", idx), Content: "第三个也不一样的文档"},
			}
			handle, err := r.BuildIndex(docs)
			if err != nil {
				return
			}
			// 查询
			r.RankByIndex(handle, "测试", 1)
			// 部分删除
			if idx%3 == 0 {
				r.RemoveIndex(handle)
			}
		}(i)
	}
	wg.Wait()
}
