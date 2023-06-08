package binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteList2Int(t *testing.T) {
	list := []uint8{}
	n := ByteList2Int(list)
	assert.Equal(t, 0, n)

	list = []uint8{1, 0, 0}
	n = ByteList2Int(list)
	assert.Equal(t, 1, n)
}

func TestIntToByteList(t *testing.T) {
	list := IntToByteList(22, 5)
	assert.Equal(t, []uint8{0, 1, 1, 0, 1}, list)
}
