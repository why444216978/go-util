package decimal

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestFormatPrice(t *testing.T) {
	r := FormatPrice(0.001)
	assert.Equal(t, "0.001", r)
	r = FormatPrice(0.0001)
	assert.Equal(t, "0.0₃1", r)
	r = FormatPrice(0.00010)
	assert.Equal(t, "0.0₃1", r)
	r = FormatPrice(0.00012345)
	assert.Equal(t, "0.0₃1234", r)
	r = FormatPrice(10.00000001234)
	assert.Equal(t, "10.0₇1234", r)
	r = FormatPrice(10.000000012345)
	assert.Equal(t, "10.0₇1234", r)
}

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

func TestDivFloorFloat64(t *testing.T) {
	res := DivFloorFloat64(1234, 2)
	assert.Equal(t, float64(1234), res)

	res = DivFloorFloat64(1234.5, 2)
	assert.Equal(t, 1234.5, res)

	res = DivFloorFloat64(1234.5, 2)
	assert.Equal(t, 1234.5, res)

	res = DivFloorFloat64(1234.5678, 2)
	assert.Equal(t, 1234.56, res)
}

func TestDivFloorString(t *testing.T) {
	r := DivFloorString(1234.5, 2)
	assert.Equal(t, "1234.50", r)

	r = DivFloorString(1234, 2)
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

func TestFormatVolume(t *testing.T) {
	convey.Convey("TestFormatVolume", t, func() {
		convey.Convey("abs <0.01", func() {
			assert.Equal(t, "< $-0.01", FormatVolume(-0.001))
		})
		convey.Convey("< 0", func() {
			assert.Equal(t, "$-0.10", FormatVolume(-0.1))
		})
		convey.Convey("== 0", func() {
			assert.Equal(t, "$0.00", FormatVolume(0))
		})
		convey.Convey("< 0.01", func() {
			assert.Equal(t, "< $0.01", FormatVolume(0.001))
		})
		convey.Convey("KMB", func() {
			assert.Equal(t, "$1.11K", FormatVolume(1111.11))
		})
	})
}

func TestFormatPercent(t *testing.T) {
	r := FormatPercent(-0.888, false)
	assert.Equal(t, "-0.88%", r)

	r = FormatPercent(0.001, false)
	assert.Equal(t, "< 0.01%", r)
	r = FormatPercent(0.001, true)
	assert.Equal(t, "< +0.01%", r)

	r = FormatPercent(0.888, false)
	assert.Equal(t, "0.88%", r)
	r = FormatPercent(0.888, true)
	assert.Equal(t, "+0.88%", r)

	r = FormatPercent(10.888, false)
	assert.Equal(t, "10.8%", r)
	r = FormatPercent(10.888, true)
	assert.Equal(t, "+10.8%", r)

	r = FormatPercent(1111.111, false)
	assert.Equal(t, "1.1K%", r)
	r = FormatPercent(1111.111, true)
	assert.Equal(t, "+1.1K%", r)
}
