package string

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

func TestSubStr(t *testing.T) {
	str := "hello world"
	Convey("TestSubStr", t, func() {
		Convey("success", func() {
			l, r, err := SubStr(str, 6, 5)
			assert.Equal(t, l, int64(5))
			assert.Equal(t, r, "world")
			assert.Equal(t, err, nil)
		})
	})
}

func TestSplitString(t *testing.T) {
	content := "第一段。第二段?第三段？第四段.第五段!第六段？今天吃什么，随便吧。"

	res := SplitPunctuation(content, 10)
	assert.Equal(t, []string{
		"第一段。第二段?",
		"第三段？第四段.",
		"第五段!第六段？",
		"今天吃什么，随便吧。",
	}, res)

	res = ReverseSplitPunctuation(content, 10)
	assert.Equal(t, []string{
		"第一段。第二段?",
		"第三段？第四段.",
		"第五段!第六段？",
		"今天吃什么，随便吧。",
	}, res)
}
