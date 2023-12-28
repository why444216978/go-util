package nopanic

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {
	convey.Convey("TestGroup", t, func() {
		convey.Convey("panic", func() {
			wg := New(context.Background())
			wg.Go(func() error {
				return nil
			})
			wg.Go(func() error {
				panic("my panic")
				return nil
			})
			assert.Nil(t, wg.Wait())
		})
		convey.Convey("AllowFail", func() {
			wg := New(context.Background(), AllowFail())
			wg.Go(func() error {
				return errors.New("my error")
			})
			time.Sleep(time.Second)
			num := 0
			wg.Go(func() error {
				num = 1
				return nil
			})
			assert.NotNil(t, wg.Wait())
			assert.Equal(t, 1, num)
		})
		convey.Convey("notAllowFail", func() {
			wg := New(context.Background())
			wg.Go(func() error {
				return errors.New("my error")
			})
			time.Sleep(time.Second)
			num := 0
			wg.Go(func() error {
				num = 1
				return nil
			})
			assert.NotNil(t, wg.Wait())
			assert.Equal(t, 0, num)
		})
		convey.Convey("SetConcurrent", func() {
			start := time.Now()
			wg := New(context.Background(), SetConcurrent(1))
			itemTime := time.Millisecond * 50
			for i := 1; i <= 2; i++ {
				wg.Go(func() error {
					time.Sleep(itemTime)
					return nil
				})
			}
			assert.Nil(t, wg.Wait())
			assert.Equal(t, true, time.Since(start) > itemTime)
		})
	})
}
