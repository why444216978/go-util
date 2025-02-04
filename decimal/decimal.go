package decimal

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

func DecimalFormat(v float64, n int32) string {
	return ValueFormat(TransferDecimal(v, n))
}

func TransferDecimalString(v float64, n int32) string {
	return cast.ToString(TransferDecimal(v, n))
}

func TransferDecimal(v float64, n int32) float64 {
	if n <= 0 {
		return v
	}
	if n == 1 {
		r, _ := decimal.NewFromFloat(v).Div(decimal.NewFromFloat(10)).Float64()
		return r
	}
	if n == 2 {
		r, _ := decimal.NewFromFloat(v).Div(decimal.NewFromFloat(100)).Float64()
		return r
	}
	r, _ := decimal.NewFromFloat(v).Div(decimal.NewFromFloat(math.Pow10(int(n) - 2))).Float64()
	r = math.Floor(r)
	r, _ = decimal.NewFromFloat(r).Div(decimal.NewFromFloat(100.0)).Float64()
	return r
}

func ValueFormat(v float64) string {
	if v < 1000 {
		return cast.ToString(v)
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
