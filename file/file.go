package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

//ReadFileLine  按行读取文件
func ReadFileLine(dir string) (map[int]string, error) {
	file, err := os.OpenFile(dir, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	/* stat, err := file.Stat()
	util_err.Must(err)
	size := stat.Size */

	buf := bufio.NewReader(file)
	res := make(map[int]string)
	i := 0
	for {
		line, _, err := buf.ReadLine()
		context := string(line)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		res[i] = context
		i++
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
