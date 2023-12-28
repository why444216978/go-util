package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	msg := "my error"
	se := WrapStackError(msg)
	assert.Equal(t, msg, se.Error())
	_ = se.Stack()

	se = WrapStackErrorSkip(msg, 1)
	assert.Equal(t, msg, se.Error())
}
