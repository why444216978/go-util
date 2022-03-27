package dir

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// GetCurrentDirectory 获得当前绝对路径
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) // 返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) // 将\替换成/
}

// LeftAddPathPos 检测并补全路径左边的反斜杠
func LeftAddPathPos(path string) string {
	if path[:0] != "/" {
		path = "/" + path
	}
	return path
}

// RightAddPathPos 检测并补全路径右边的反斜杠
func RightAddPathPos(path string) string {
	if path[len(path)-1:len(path)] != "/" {
		path = path + "/"
	}
	return path
}

// FileNameByDate 根据当天日期和给定dir返回log文件名路径
func FileNameByDate(dir string) string {
	fileName := time.Now().Format("2006-01-02")
	dir = RightAddPathPos(dir)
	return dir + fileName + ".log"
}

// CreateDir 不存在则创建目录
func CreateDir(folderPath string) {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, 0777) // 0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
}

// CreateDateDir 根据当前日期，不存在则创建目录
func CreateDateDir(path string, prex string) string {
	folderName := time.Now().Format("20060102")
	if prex != "" {
		folderName = prex + folderName
	}
	folderPath := filepath.Join(path, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, 0777) // 0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

// GetDateDir 根据路径获得按照"年月日"生成的子路径
func GetDateDir(path string) string {
	return path + time.Now().Format("20660102")
}

// CreateHourLogFile 根据当前小时创建目录和日志文件
func CreateHourLogFile(path string, prex string) string {
	folderName := time.Now().Format("2006010215")
	if prex != "" {
		folderName = prex + folderName
	}
	folderPath := filepath.Join(path, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, 0777) // 0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

// ReadDirAll 读取目录
// example ReadDirAll("/Users/why/Desktop/go/test", 0)
func ReadDirAll(path string, curHier int) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	ret := []string{}
	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				ret = append(ret, "|\t")
			}
			ret = append(ret, info.Name()+"\\")
			ReadDirAll(path+"/"+info.Name(), curHier+1)
		} else {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				ret = append(ret, "|\t")
			}
			ret = append(ret, info.Name())
		}
	}
	return ret, nil
}

type FileInfo struct {
	Path      string
	Base      string
	BaseNoExt string
	Ext       string
	ExtNoSpot string
}

// GetPathInfo 获得路径信息
func GetPathInfo(f string) (info FileInfo, err error) {
	info.Path = f
	f = path.Base(filepath.ToSlash(f))
	info.Base = f
	ext := path.Ext(f)
	info.Ext = ext
	if ext == "" {
		err = errors.New("ext error")
		return
	}
	info.ExtNoSpot = ext[1:]
	info.BaseNoExt = strings.TrimSuffix(f, ext)

	return
}
