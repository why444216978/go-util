package nopanic

import (
	"context"
	"sync"

	"github.com/panjf2000/ants/v2"
)

type options struct {
	concurrent int
	allowFail  bool
}

type optionFunc func(*options)

func SetConcurrent(concurrent int) optionFunc {
	return func(o *options) { o.concurrent = concurrent }
}

func AllowFail() optionFunc {
	return func(o *options) { o.allowFail = true }
}

type Group struct {
	*options

	ctx     context.Context
	wg      sync.WaitGroup
	errOnce sync.Once
	lock    sync.RWMutex
	pool    *ants.Pool
	err     error
}

func New(ctx context.Context, opts ...optionFunc) *Group {
	option := &options{}
	for _, o := range opts {
		o(option)
	}

	p, _ := ants.NewPool(option.concurrent)

	return &Group{
		options: option,
		ctx:     ctx,
		pool:    p,
	}
}

func (g *Group) Go(fn func() error) {
	var err error

	g.lock.RLock()
	err = g.err
	g.lock.RUnlock()

	if !g.allowFail && err != nil {
		return
	}

	g.wg.Add(1)
	g.pool.Submit(func() {
		defer g.wg.Done()
		err := Go(g.ctx, fn)
		if err != nil {
			g.setError(err)
		}
	})
}

func (g *Group) setError(err error) {
	g.errOnce.Do(func() {
		g.lock.Lock()
		g.err = err
		g.lock.Unlock()
	})
}

func (g *Group) Wait() error {
	g.wg.Wait()

	g.lock.RLock()
	err := g.err
	g.lock.RUnlock()

	return err
}
