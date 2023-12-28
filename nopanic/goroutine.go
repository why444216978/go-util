package nopanic

import (
	"context"

	"github.com/why444216978/go-util/runtime"
)

func Go(ctx context.Context, fn func() error) error {
	defer func() {
		if re := recover(); re != nil {
			handleCallback(ctx, runtime.WrapStackError(re))
		}
	}()
	return fn()
}

func GoVoid(ctx context.Context, fn func()) {
	defer func() {
		if re := recover(); re != nil {
			handleCallback(ctx, runtime.WrapStackError(re))
		}
	}()
	fn()
}
