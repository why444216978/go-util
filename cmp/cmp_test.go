package cmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	assert.Equal(t, true, CompareLT(1, 2))
	assert.Equal(t, true, CompareLE(1, 2))
	assert.Equal(t, true, CompareEQ(1, 1))
	assert.Equal(t, true, CompareNE(1, 2))
	assert.Equal(t, true, CompareGT(2, 1))
	assert.Equal(t, true, CompareGE(1, 1))
	assert.Equal(t, true, CompareIN([]int{1, 2}, 1))
}
