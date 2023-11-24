package xss

import (
	"github.com/microcosm-cc/bluemonday"
)

var (
	defaultStrict = bluemonday.StrictPolicy()
	defaultUGC    = func() *bluemonday.Policy {
		p := bluemonday.UGCPolicy()
		p.RequireNoFollowOnLinks(false)
		p.AllowAttrs("style").Globally()
		p.AllowAttrs("data-origin").OnElements("img")
		p.AllowAttrs("data-size").OnElements("img")
		p.AllowAttrs("target").OnElements("a")
		p.AllowAttrs("rel").OnElements("a")
		return p
	}()
)

type options struct {
	policy *bluemonday.Policy
}

type Option func(*options)

func WithPolicy(p *bluemonday.Policy) Option {
	return func(o *options) { o.policy = p }
}

func FilterOnlyText(s string, opts ...Option) string {
	o := handleOptions(opts...)
	if o.policy != nil {
		return o.policy.Sanitize(s)
	}

	return defaultStrict.Sanitize(s)
}

func FilterSaveTag(s string, opts ...Option) string {
	o := handleOptions(opts...)
	if o.policy != nil {
		return o.policy.Sanitize(s)
	}

	return defaultUGC.Sanitize(s)
}

func handleOptions(opts ...Option) *options {
	option := &options{}
	for _, o := range opts {
		o(option)
	}
	return option
}
