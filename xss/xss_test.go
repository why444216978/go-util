package xss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const s = `<p>我的2022</p><img src='https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png' onload='alert("aaa")' onerror='alert("bbb")'>`

func TestFilterOnlyText(t *testing.T) {
	assert.Equal(t, "我的2022", FilterOnlyText(s))
}

// nolint
var equalHTML = []string{
	`<p>正文</p>`,
	`<p><strong style="color: rgb(230, 0, 0); background-color: rgb(153, 51, 255);"><em><s><u>文字样式</u></s></em></strong><sub style="color: rgb(230, 0, 0); background-color: rgb(153, 51, 255);"><strong><em><s><u>1</u></s></em></strong></sub><sup style="color: rgb(230, 0, 0); background-color: rgb(153, 51, 255);"><strong><em><s><u>2</u></s></em></strong></sup></p>`,
	`<ol><li><sup style="color: rgb(230, 0, 0); background-color: rgb(153, 51, 255);"><strong><em><s><u>1</u></s></em></strong></sup></li><li><sup style="color: rgb(230, 0, 0); background-color: rgb(153, 51, 255);"><strong><em><s><u>2</u></s></em></strong></sup></li></ol>`,
	`<ul><li><sup style="color: rgb(230, 0, 0); background-color: rgb(153, 51, 255);"><strong><em><s><u>1</u></s></em></strong></sup></li><li><sup style="color: rgb(230, 0, 0); background-color: rgb(153, 51, 255);"><strong><em><s><u>2</u></s></em></strong></sup></li></ul>`,
	`<p><a href="https://drs.baidu.com/dcp/editor" rel="noopener noreferrer" target="_blank">百度一下</a></p>`,
	`<img src="https://health-community.cdn.bcebos.com/image/2023-11/23/14/2fnyMBi6dR09A/12ace5ac61f90253bdd3e2f298af3ece.png?x-bce-process=image/auto-orient,o_1/quality,Q_60" data-origin="https://health-community.cdn.bcebos.com/image/2023-11/23/14/2fnyMBi6dR09A/12ace5ac61f90253bdd3e2f298af3ece.png" data-size="151929">`,
	`<p><br></p>`,
}

// nolint
var filterHTML = []string{
	`<script>alert(111)</script>`,
	`<p>1<a href='http://www.baidu.com'>23</a></p>`,
	`<p>1<a href='www"onmouseover="alert(1)'>23</a></p>`,
	`<img onload='alert("aaa")' onerror=alert(111) src="https://health-community.bj.bcebos.com/image/2023-01/31/17/vtUJR28LTCHzAeeeB/100617641dca342ca1d8f7433b8d9736.jpeg" data-origin="https://health-community.bj.bcebos.com/image/2023-01/31/17/vtUJR28LTCHzAeeeB/100617641dca342ca1d8f7433b8d9736.jpeg" data-size="1010404">`,
}

func TestFilterSaveTag(t *testing.T) {
	for _, s := range equalHTML {
		assert.Equal(t, true, s == FilterSaveTag(s))
	}
	for _, s := range filterHTML {
		assert.Equal(t, false, s == FilterSaveTag(s))
	}

	assert.Equal(t, "我的2022", FilterSaveTag(s, WithPolicy(defaultStrict)))
}
