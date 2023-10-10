package xss

import (
	"github.com/microcosm-cc/bluemonday"
)

func FilterOnlyText(s string) string {
	return bluemonday.StrictPolicy().Sanitize(s)
}

func FilterSaveTag(s string) string {
	return bluemonday.UGCPolicy().Sanitize(s)
}
