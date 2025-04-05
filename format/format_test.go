package format

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestTransferDecimalFloat64(t *testing.T) {
	res := TransferDecimalFloat64(float64(88), 1)
	assert.Equal(t, 8.8, res)

	res = TransferDecimalFloat64(float64(88), 2)
	assert.Equal(t, 0.88, res)

	res = TransferDecimalFloat64(float64(88), 3)
	assert.Equal(t, 0.08, res)

	res = TransferDecimalFloat64(float64(6950), 3)
	assert.Equal(t, 6.95, res)

	res = TransferDecimalFloat64(float64(12345), 3)
	assert.Equal(t, 12.34, res)

	res = TransferDecimalFloat64(float64(152668107691), 6)
	assert.Equal(t, 152668.10, res)
}

func TestFormatFloorFloat64(t *testing.T) {
	res := FormatFloorFloat64(1234, 2)
	assert.Equal(t, float64(1234), res)

	res = FormatFloorFloat64(1234.5, 2)
	assert.Equal(t, 1234.5, res)

	res = FormatFloorFloat64(1234.5, 2)
	assert.Equal(t, 1234.5, res)

	res = FormatFloorFloat64(1234.5678, 2)
	assert.Equal(t, 1234.56, res)
}

func TestFormatFloorString(t *testing.T) {
	r := FormatFloorString(1234.5, 2)
	assert.Equal(t, "1234.50", r)

	r = FormatFloorString(1234, 2)
	assert.Equal(t, "1234.00", r)
}

func TestFormatKMB(t *testing.T) {
	convey.Convey("TestValueFormat", t, func() {
		convey.Convey("< 1000", func() {
			assert.Equal(t, "888.00", FormatKMB(888, 2))
			assert.Equal(t, "888.12", FormatKMB(888.123456789, 2))
		})
		convey.Convey("< 1000000", func() {
			assert.Equal(t, "888.88K", FormatKMB(888888, 2))
		})
		convey.Convey("< 1000000000", func() {
			assert.Equal(t, "888.88M", FormatKMB(888888888, 2))
		})
		convey.Convey("< 1000000000000", func() {
			assert.Equal(t, "888.88B", FormatKMB(888888888888, 2))
		})
		convey.Convey("< 1000000000000000", func() {
			assert.Equal(t, "888.88T", FormatKMB(888888888888888, 2))
		})
		convey.Convey("< 1000000000000000000", func() {
			assert.Equal(t, "888.88Q", FormatKMB(888888888888888888, 2))
		})
		convey.Convey("default", func() {
			assert.Equal(t, "888.88S", FormatKMB(888888888888888888888, 2))
		})
	})
}

func TestFormatPercent(t *testing.T) {
	r := FormatPercent(-0.888, false)
	assert.Equal(t, "-88.80%", r)

	r = FormatPercent(0.00001, false)
	assert.Equal(t, "< 0.01%", r)
	r = FormatPercent(0.00001, true)
	assert.Equal(t, "< +0.01%", r)

	r = FormatPercent(0.00888, false)
	assert.Equal(t, "0.88%", r)
	r = FormatPercent(0.00888, true)
	assert.Equal(t, "+0.88%", r)

	r = FormatPercent(0.10888, false)
	assert.Equal(t, "10.88%", r)
	r = FormatPercent(0.10888, true)
	assert.Equal(t, "+10.88%", r)

	r = FormatPercent(11.11111, false)
	assert.Equal(t, "1.1K%", r)
	r = FormatPercent(11.11111, true)
	assert.Equal(t, "+1.1K%", r)
}
