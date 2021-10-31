package file

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
)

// WriteWithIo 使用io.WriteString()函数进行数据的写入，不存在则创建
func WriteWithIo(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	if content != "" {
		_, err := io.WriteString(file, content)
		if err != nil {
			return err
		}
		fmt.Println("Successful appending to the file with os.OpenFile and io.WriteString.", content)
	}

	return nil
}

// ReadLimit 读取指定字节
func ReadLimit(str string, len int64) string {
	reader := strings.NewReader(str)
	limitReader := &io.LimitedReader{R: reader, N: len}

	var res string
	for limitReader.N > 0 {
		tmp := make([]byte, 1)
		limitReader.Read(tmp)
		res += string(tmp)
	}
	return res
}

// ReadFile 读取整个文件
func ReadFile(dir string) (string, error) {
	data, err := ioutil.ReadFile(dir)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadFileLine 按行读取字典文件
func ReadFileLine(dir string) ([]string, error) {
	file, err := os.OpenFile(dir, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Stat()
	if err != nil {
		return nil, err
	}

	buf := bufio.NewReader(file)
	res := make([]string, 0)
	for {
		line, _, err := buf.ReadLine()
		s := string(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		res = append(res, s)
	}
	return res, nil
}

// ReadJsonFile 读取json文件
func ReadJsonFile(dir string) (string, error) {
	jsonFile, err := os.Open(dir)
	if err != nil {
		return "", err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return string(byteValue), nil
}

// GetFileInfo 获得文件Info
func GetFileInfo(file *os.File) os.FileInfo {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("file stat error:", err)
	}
	return fileInfo
}

// GetFileMode 获得文件权限Mode
func GetFileMode(file *os.File) os.FileMode {
	fileInfo := GetFileInfo(file)
	return fileInfo.Mode()
}

// GetFileStat 获得文件Stat
func GetFileStat(file *os.File) *syscall.Stat_t {
	fileInfo := GetFileInfo(file)
	sysInterface := fileInfo.Sys()
	sys := sysInterface.(*syscall.Stat_t)
	//fmt.Println(sys.Atimespec)
	return sys
}

// Chown 更改文件所有者
func Chown(file *os.File, uid, gid int) error {
	if uid == 0 {
		uid = os.Getuid()
	}
	if gid == 0 {
		gid = os.Getgid()
	}

	return file.Chown(uid, gid)
}

// Chmod 更改文件权限
func Chmod(file *os.File, mode int) error {
	return file.Chmod(os.FileMode(mode))
}

// Open 打开文件
func Open(dir string) (*os.File, error) {
	return os.Open(dir)
}

// Create 创建文件
func Create(dir string) (*os.File, error) {
	return os.Create(dir)
}

// CleanFile 清楚文件内容
func CleanFile(filePath string) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("")
	return err
}

// DownloadFileToBase64 通过url下载文件并转为base64
func DownloadFileToBase64(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// 读取获取的[]byte数据
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	data := base64.StdEncoding.EncodeToString(bytes)

	return data, nil
}

// Base64ToFile base64写入文件
func Base64ToFile(data []byte, file string) (err error) {
	decodeData, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return
	}

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.Write(decodeData)

	return
}
