package xss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const s = `<p>我的2022</p><img src='https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png' onload='alert("aaa")' onerror='alert("bbb")'>`

func TestFilterOnlyText(t *testing.T) {
	assert.Equal(t, "我的2022", FilterOnlyText(s))
}

func TestFilterSaveTag(t *testing.T) {
	assert.Equal(t, `<p>我的2022</p><img src="https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png">`, FilterSaveTag(s))
}
