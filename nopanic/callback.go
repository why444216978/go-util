package nopanic

import (
	"context"

	"github.com/why444216978/go-util/runtime"
)

type RecoverData struct {
	StackError *runtime.StackError
	Ctx        context.Context
}

var recoverCallback Slice[func(re RecoverData)]

func RegisterRecoverCallback(fn func(re RecoverData)) {
	recoverCallback.Add(fn)
}

func handleCallback(ctx context.Context, se *runtime.StackError, ext ...any) {
	if se == nil {
		return
	}

	fns := recoverCallback.All()
	if len(fns) == 0 {
		return
	}
	data := RecoverData{
		StackError: se,
		Ctx:        ctx,
	}
	for i := 0; i < len(fns); i++ {
		fns[i](data)
	}
}
