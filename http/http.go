package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/why444216978/go-util/validate"
)

type Response struct {
	HTTPCode int
	Response string
}

// Send 发送请求
func Send(ctx context.Context, method, url string, header map[string]string, body io.Reader, timeout time.Duration) (ret Response, err error) {
	var req *http.Request

	client := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   timeout,
	}

	//构建req
	req, err = http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return
	}

	//设置请求header
	for k, v := range header {
		req.Header.Add(k, v)
	}

	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	ret.HTTPCode = resp.StatusCode
	if resp.StatusCode != http.StatusOK {
		err = errors.Errorf("http code is %d", resp.StatusCode)
		return
	}

	if b != nil {
		ret.Response = string(b)
	}

	return
}

// ExtractBoy 解析请求body并回写
func ExtractBody(req http.Request) string {
	var buf bytes.Buffer
	buf.ReadFrom(req.Body)
	req.Body = ioutil.NopCloser(&buf)

	return buf.String()
}

// ParseAndValidateBody 解析并校验http.Request.Body
func ParseAndValidateBody(req *http.Request, target interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(target); err != nil {
		return err
	}
	if err := validate.Validate(target); err != nil {
		return err
	}

	return nil
}
