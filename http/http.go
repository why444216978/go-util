package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Send 发送http请求
func Send(method, url string, header map[string]string, body string) (ret map[string]interface{}, err error) {
	var req *http.Request
	ret = make(map[string]interface{}, 3)

	client := &http.Client{}

	reader := bytes.NewReader([]byte(body))

	//构建req
	req, err = http.NewRequest(method, url, reader)
	if err != nil {
		err = errors.Wrap(err, "util structure http err:")
		return
	}

	//设置请求header
	for k, v := range header {
		req.Header.Add(k, v)
	}

	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "util http send err:")
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "util http read body err：")
		return
	}

	ret["http_code"] = resp.StatusCode

	if resp.StatusCode != http.StatusOK {
		str := fmt.Sprintf("http code is %d", resp.StatusCode)
		err = errors.New(str)
		return
	}

	if b != nil {
		ret["response"] = string(b)
	}

	return
}

// PostForm 发送POST表单请求
func PostForm(ctx context.Context, postUrl string, data map[string]interface{}) (string, error) {
	body := url.Values{}
	for k, v := range data {
		body.Add(k, fmt.Sprintf("%v", v))
	}

	res, err := http.PostForm(postUrl, body)
	if err != nil {
		err = errors.Wrap(err, "util post from fail")
		return "", err
	}
	defer res.Body.Close()

	resp, _ := ioutil.ReadAll(res.Body)

	return string(resp), nil
}
