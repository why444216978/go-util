package md5

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
)

func MD5(str string) string {
	h := md5.New()
	r := bytes.NewBufferString(str)
	_, _ = io.Copy(h, r)
	return hex.EncodeToString(h.Sum(nil))
}
