package nopanic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/why444216978/go-util/runtime"
)

func TestCallback(t *testing.T) {
	res := []RecoverData{}
	RegisterRecoverCallback(func(re RecoverData) {
		res = append(res, re)
	})

	handleCallback(context.Background(), runtime.WrapStackError("panic"))

	assert.Equal(t, 1, len(res))
}
