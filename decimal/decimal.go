package decimal

import (
	"fmt"
	"math"
	"strconv"

	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

func DecimalFormat(v float64, n int) string {
	return ValueFormat(TransferDecimal(v, n))
}

func TransferDecimalString(v float64, n int) string {
	return cast.ToString(TransferDecimal(v, n))
}

func TransferDecimal(v float64, n int) float64 {
	if n <= 0 {
		return v
	}

	if n <= 2 {
		r, _ := decimal.NewFromFloat(v).Div(decimal.NewFromFloat(math.Pow10(n))).Float64()
		return r
	}

	r, _ := decimal.NewFromFloat(v).Div(decimal.NewFromFloat(math.Pow10(n - 2))).Float64()
	r = math.Floor(r)
	r, _ = decimal.NewFromFloat(r).Div(decimal.NewFromFloat(100.0)).Float64()
	return r
}

func DivFloorFloat64(v float64, n int) float64 {
	return divFloor(v, n)
}

func DivFloorString(v float64, n int) string {
	return strconv.FormatFloat(divFloor(v, n), 'f', n, 64)
}

func divFloor(v float64, n int) float64 {
	// 1234.5678 => 123456.78
	r, _ := decimal.NewFromFloat(v).Mul(decimal.NewFromFloat(math.Pow10(n))).Float64()
	// 123456.78 => 123456
	r = math.Floor(r)
	// 123456 => 1234.56
	d, _ := decimal.NewFromFloat(r).Div(decimal.NewFromFloat(math.Pow10(n))).Float64()
	return d
}

func ValueFormat(v float64) string {
	if v < 1000 {
		return DivFloorString(v, 2)
	}

	if v < 1000000 {
		return fmt.Sprintf("%.2fK", TransferDecimal(v, 3))
	}

	if v < 1000000000 {
		return fmt.Sprintf("%.2fM", TransferDecimal(v, 6))
	}

	if v < 1000000000000 {
		return fmt.Sprintf("%.2fB", TransferDecimal(v, 9))
	}

	if v < 1000000000000000 {
		return fmt.Sprintf("%.2fT", TransferDecimal(v, 12))
	}

	if v < 1000000000000000000 {
		return fmt.Sprintf("%.2fQ", TransferDecimal(v, 15))
	}

	return fmt.Sprintf("%.2fS", TransferDecimal(v, 18))
}
