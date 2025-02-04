package decimal

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestTransferDecimal(t *testing.T) {
	res := TransferDecimal(float64(88), 1)
	assert.Equal(t, 8.8, res)

	res = TransferDecimal(float64(88), 2)
	assert.Equal(t, 0.88, res)

	res = TransferDecimal(float64(88), 3)
	assert.Equal(t, 0.08, res)

	res = TransferDecimal(float64(6950), 3)
	assert.Equal(t, 6.95, res)

	res = TransferDecimal(float64(12345), 3)
	assert.Equal(t, 12.34, res)

	res = TransferDecimal(float64(152668107691), 6)
	assert.Equal(t, 152668.10, res)
}

func TestValueFormat(t *testing.T) {
	convey.Convey("TestValueFormat", t, func() {
		convey.Convey("< 1000", func() {
			assert.Equal(t, "888", ValueFormat(888))
		})
		convey.Convey("< 1000000", func() {
			assert.Equal(t, "888.88K", ValueFormat(888888))
		})
		convey.Convey("< 1000000000", func() {
			assert.Equal(t, "888.88M", ValueFormat(888888888))
		})
		convey.Convey("< 1000000000000", func() {
			assert.Equal(t, "888.88B", ValueFormat(888888888888))
		})
		convey.Convey("< 1000000000000000", func() {
			assert.Equal(t, "888.88T", ValueFormat(888888888888888))
		})
		convey.Convey("< 1000000000000000000", func() {
			assert.Equal(t, "888.88Q", ValueFormat(888888888888888888))
		})
		convey.Convey("default", func() {
			assert.Equal(t, "888.88S", ValueFormat(888888888888888888888))
		})
	})
}
