package decimal

import (
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

var unicodeNum = map[string]string{
	"0": "\u2080",
	"1": "\u2081",
	"2": "\u2082",
	"3": "\u2083",
	"4": "\u2084",
	"5": "\u2085",
	"6": "\u2086",
	"7": "\u2087",
	"8": "\u2088",
	"9": "\u2089",
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func FormatPrice(v float64) string {
	vs := cast.ToString(v)
	arr := strings.Split(vs, ".")
	if len(arr) != 2 {
		return vs
	}

	if !strings.HasPrefix(arr[1], "000") {
		return strings.TrimRight(FormatFloorString(v, 4), "0")
	}

	r := []byte{}
	stopZero := false
	zeroCount := 0
	for _, v := range []byte(arr[1]) {
		if v != 0x30 {
			stopZero = true
		}

		if !stopZero {
			zeroCount++
		} else {
			r = append(r, v)
		}
	}

	z := ""
	for _, v := range cast.ToString(zeroCount) {
		z += unicodeNum[string(v)]
	}

	result := fmt.Sprintf("%s.0%s%s", arr[0], z, r[:min(4, len(r))])

	return strings.TrimRight(result, "0")
}

func FormatCountBuy(v int64) string {
	if v == 0 {
		return "0"
	}
	return fmt.Sprintf("+%s", FormatKMB(float64(v), 2))
}

func FormatCountSell(v int64) string {
	if v == 0 {
		return "0"
	}
	return fmt.Sprintf("-%s", FormatKMB(float64(v), 2))
}

func FormatCountBuyVolume(v float64) string {
	if v == 0 {
		return "$0.00"
	}
	return fmt.Sprintf("-$%s", FormatKMB(v, 2))
}

func FormatCountSellVolume(v float64) string {
	if v == 0 {
		return "$0.00"
	}
	return fmt.Sprintf("+$%s", FormatKMB(v, 2))
}

func FormatDecimal(v float64, n int32) float64 {
	return TransferDecimalFloat64(v, n)
}

func FormatDecimalVolume(v float64, n int32) string {
	return FormatVolume(TransferDecimalFloat64(v, n))
}

func TransferDecimalString(v float64, n int32) string {
	return TransferDecimal(v, n).String()
}

func TransferDecimalFloat64(v float64, n int32) float64 {
	d, _ := TransferDecimal(v, n).Float64()
	return d
}

func TransferDecimal(v float64, n int32) decimal.Decimal {
	if n <= 0 {
		return decimal.NewFromFloat(v)
	}

	if n <= 2 {
		return decimal.NewFromFloat(v).Div(decimal.NewFromFloat(math.Pow10(int(n))))
	}

	d := decimal.NewFromFloat(v).Div(decimal.NewFromFloat(math.Pow10(int(n) - 2)))
	d = d.Floor()
	return d.Div(decimal.NewFromFloat(100.0))
}

func FormatFloorFloat64(v float64, n int) float64 {
	return cast.ToFloat64(FormatFloorString(v, n))
}

func FormatFloorString(v float64, n int) string {
	// 1234.5678 => 1234.56

	arr := strings.Split(cast.ToString(v), ".")

	f := ""
	if len(arr) == 1 {
		for i := 1; i <= n; i++ {
			f += "0"
		}
	} else {
		l := len(arr[1])
		for i := 0; i < n; i++ {
			if i < l {
				f += string(arr[1][i])
			} else {
				f += "0"
			}
		}
	}

	return fmt.Sprintf("%s.%s", arr[0], f)
}

func FormatKMB(v float64, n int) string {
	if v == 0 {
		return "0.00"
	}

	if v < 1000 {
		return FormatFloorString(v, 2)
	}

	if v < 1000000 {
		return fmt.Sprintf("%sK", FormatFloorString(TransferDecimalFloat64(v, 3), n))
	}

	if v < 1000000000 {
		return fmt.Sprintf("%sM", FormatFloorString(TransferDecimalFloat64(v, 6), n))
	}

	if v < 1000000000000 {
		return fmt.Sprintf("%sB", FormatFloorString(TransferDecimalFloat64(v, 9), n))
	}

	if v < 1000000000000000 {
		return fmt.Sprintf("%sT", FormatFloorString(TransferDecimalFloat64(v, 12), n))
	}

	if v < 1000000000000000000 {
		return fmt.Sprintf("%sQ", FormatFloorString(TransferDecimalFloat64(v, 15), n))
	}

	return fmt.Sprintf("%sS", FormatFloorString(TransferDecimalFloat64(v, 18), n))
}

func FormatVolume(v float64) string {
	if v < 0 && math.Abs(v) < 0.01 {
		return "< $-0.01"
	}

	if v < 0 {
		return fmt.Sprintf("$%s", FormatFloorString(v, 2))
	}

	if v == 0 {
		return "$0.00"
	}

	if v < 0.01 {
		return "< $0.01"
	}

	if v < 10 {
		return fmt.Sprintf("$%s", FormatFloorString(v, 2))
	}

	if v >= 10 && v < 1000 {
		return fmt.Sprintf("$%s", FormatFloorString(v, 1))
	}

	return fmt.Sprintf("$%s", FormatKMB(v, 1))
}

func FormatAmount(v float64) string {
	if v < 0 {
		return cast.ToString(v)
	}

	if v == 0 {
		return "0.00"
	}

	if v < 0.01 {
		return "< 0.01"
	}

	if v < 10 {
		return FormatFloorString(v, 2)
	}

	if v >= 10 && v < 1000 {
		return FormatFloorString(v, 1)
	}

	return FormatKMB(v, 1)
}

func FormatPercent(v float64, withSymbol bool) string {
	symbol := ""
	if withSymbol {
		symbol = "+"
	}

	v, _ = decimal.NewFromFloat(v).Mul(decimal.NewFromInt(100)).Float64()

	if v < 0 && math.Abs(v) < 0.01 {
		return "< -0.01%"
	}

	if v < 0 {
		return fmt.Sprintf("%s", FormatFloorString(v, 2)) + "%"
	}

	if v < 0.01 {
		return fmt.Sprintf("< %s0.01", symbol) + "%"
	}

	if v < 10 {
		return fmt.Sprintf("%s%s", symbol, FormatFloorString(v, 2)) + "%"
	}

	if v >= 10 && v < 1000 {
		return fmt.Sprintf("%s%s", symbol, FormatFloorString(v, 1)) + "%"
	}

	return fmt.Sprintf("%s%s", symbol, FormatKMB(v, 1)) + "%"
}
