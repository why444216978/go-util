package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/why444216978/go-util/conversion"

	"github.com/pkg/errors"
)

func HttpSend(method, url string, header, queryData map[string]string, body map[string]interface{}) (ret map[string]interface{}, err error) {
	var req *http.Request
	ret = make(map[string]interface{}, 3)

	client := &http.Client{}

	if queryData != nil {
		url += "?"
		querySlice := make([]string, 0, len(queryData))
		for k, v := range queryData {
			querySlice = append(querySlice, k+"="+v)
		}
		url += strings.Join(querySlice, "&")
	}

	//请求数据
	byteDates, err := json.Marshal(body)
	if err != nil {
		err = errors.Wrap(err, "request data to json err")
		return
	}
	reader := bytes.NewReader(byteDates)

	//构建req
	req, err = http.NewRequest(method, url, reader)
	if err != nil {
		err = errors.Wrap(err, "structure http err")
		return
	}

	//设置请求header
	for k, v := range header {
		//req.Header.Add("content-type", "application/json")
		req.Header.Add(k, v)
	}

	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "send http err：")
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "read body err：")
		return
	}

	ret["http_code"] = resp.StatusCode
	ret["response"] = make(map[string]interface{})

	if resp.StatusCode != http.StatusOK {
		ret["msg"] = "http code:" + strconv.Itoa(resp.StatusCode)
	}

	if b != nil {
		bStr := string(b)
		bMap, err := conversion.JsonToMap(bStr)
		if err != nil {
			ret["response"] = string(b)
		} else {
			ret["response"] = bMap
		}
	}

	return
}
