package url

import (
	"net/url"
	"path"
	"strings"

	util_dir "github.com/why444216978/go-util/dir"
)

// MapMapToHttpQuery map转uri query
func MapToHttpQuery(m map[string]interface{}) string {
	var p = url.Values{}

	for k, v := range m {
		p.Add(k, fmt.Sprintf("%v", v))
	}
	return p.Encode()
}

// uRUriToFilePathByDate url的path转文件名
func UriToFilePathByDate(uriPath string, dir string) string {
	pathArr := strings.Split(uriPath, "/")
	fileName := strings.Join(pathArr, "-")
	writePath := util_dir.CreateDateDir(dir, "") //根据时间检测是否存在目录，不存在创建
	writePath = util_dir.RightAddPathPos(writePath)
	fileName = path.Join(writePath, fileName[1:len(fileName)]+".log")
	return fileName
}

// uRUriToFileName uri转log文件名
func UriToFileName(uri string) string {
	pathArr := strings.Split(uri, "/")
	fileName := strings.Join(pathArr, "-")
	fileName = fileName + ".log"
	return fileName
}

// LogByUrl url转log文件名
func LogByUrl(fullUrl string) (string, error) {
	u, err := url.Parse(fullUrl)
	if err != nil {
		return "", err
	}

	pathArr := strings.Split(u.Path, "/")
	fileName := strings.Join(pathArr, "-")
	writePath := "/tmp/logs/2020-01-12"
	fileName = path.Join(writePath, fileName[1:len(fileName)]+".log")

	return fileName, nil
}

// ParseUriQueryToMap uri的query转map
func ParseUriQueryToMap(query string) map[string]interface{} {
	queryMap := strings.Split(query, "&")
	res := make(map[string]interface{}, len(queryMap))
	if query == "" {
		return res
	}
	for _, item := range queryMap {
		itemMap := strings.Split(item, "=")
		res[itemMap[0]] = itemMap[1]
	}
	return res
}

// UriToFilePath uri转log路径
func UriToFilePath(uri string, dir string) string {
	pathArr := strings.Split(uri, "/")
	fileName := strings.Join(pathArr, "-")
	fileName = path.Join(dir, fileName[1:len(fileName)]+".log")
	if fileName[len(fileName)-1:len(fileName)] != "/" {
		fileName = fileName + "/"
	}
	return fileName
}
