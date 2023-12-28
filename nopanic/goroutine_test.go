package nopanic

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGo(t *testing.T) {
	assert.Nil(t, Go(context.Background(), func() error {
		panic("my panic")
		return nil
	}))
	assert.NotNil(t, Go(context.Background(), func() error {
		return errors.New("my error")
	}))
}

func TestGoVoid(t *testing.T) {
	GoVoid(context.Background(), func() {
		panic("my panic")
	})
}
