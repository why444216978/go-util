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

	util_err "github.com/why444216978/go-util/error"
)

// WriteWithIo 使用io.WriteString()函数进行数据的写入，不存在则创建
func WriteWithIo(filePath, content string) error {
	file := OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
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
func ReadFile(dir string) string {
	data, err := ioutil.ReadFile(dir)
	if err != nil {
		panic(err)
		return ""
	}
	return string(data)
}

//ReadFileLine  按行读取文件
func ReadFileLine(dir string) map[int]string {
	file, err := os.OpenFile(dir, os.O_RDWR, 0666)
	util_err.Must(err)
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
	return res
}

// ReadJsonFile 读取json文件
func ReadJsonFile(dir string) string {
	jsonFile, err := os.Open(dir)
	util_err.Must(err)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return string(byteValue)
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
func Chown(file *os.File, uid, gid int) {
	if uid == 0 {
		uid = os.Getuid()
	}
	if gid == 0 {
		gid = os.Getgid()
	}

	err := file.Chown(uid, gid)
	if err != nil {
		panic(err)
	}
}

// Chmod 更改文件权限
func Chmod(file *os.File, mode int) {
	err := file.Chmod(os.FileMode(mode))
	if err != nil {
		panic(err)
	}
}

// Open 打开文件
func Open(dir string) *os.File {
	file, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	return file
}

// Create 创建文件
func Create(dir string) *os.File {
	file, err := os.Create(dir)
	if err != nil {
		panic(err)
	}
	return file
}

// OpenFile 根据flag打开文件
func OpenFile(name string, flag int, perm os.FileMode) *os.File {
	file, err := os.OpenFile(name, flag, perm)
	if err != nil {
		panic(err)
	}
	return file
}
