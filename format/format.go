package format

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

func FormatUnicode(v float64) string {
	vs := cast.ToString(v)
	arr := strings.Split(vs, ".")
	if len(arr) != 2 {
		return fmt.Sprintf("$%s", vs)
	}

	if !strings.HasPrefix(arr[1], "000") {
		zeroList := []rune{}
		valList := []rune{}
		for _, s := range arr[1] {
			if len(valList) >= 4 {
				break
			}
			if len(valList) == 0 && s == 48 {
				zeroList = append(zeroList, s)
				continue
			}
			valList = append(valList, s)
		}
		return strings.TrimRight(fmt.Sprintf("$%s.%s%s", arr[0], string(zeroList), string(valList)), "0")
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

	result := fmt.Sprintf("$%s.0%s%s", arr[0], z, r[:min(4, len(r))])

	return strings.TrimRight(result, "0")
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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

	if val, _ := d.Float64(); val > 1 {
		d = d.Floor()
	}
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

	if f == "" {
		return arr[0]
	}
	return fmt.Sprintf("%s.%s", arr[0], f)
}

func FormatValueComma(v float64) string {
	// 将价格转换为字符串，保留2位小数
	str := fmt.Sprintf("%.2f", v)

	// 分割整数部分和小数部分
	parts := strings.Split(str, ".")
	integerPart := parts[0]
	decimalPart := parts[1]

	// 处理负数情况
	negative := false
	if integerPart[0] == '-' {
		negative = true
		integerPart = integerPart[1:]
	}

	// 从右向左每3位加逗号
	var formatted strings.Builder
	length := len(integerPart)

	for i := 0; i < length; i++ {
		// 不是第一个字符且是3的倍数位置时加逗号
		if (length-i)%3 == 0 && i != 0 {
			formatted.WriteString(",")
		}
		formatted.WriteByte(integerPart[i])
	}

	// 重新添加负号和小数部分
	result := formatted.String()
	if negative {
		result = "-" + result
	}
	return result + "." + decimalPart
}

func FormatPercent(v float64, withSymbol bool) string {
	symbol := ""
	if withSymbol {
		symbol = "+"
	}

	v, _ = decimal.NewFromFloat(v).Mul(decimal.NewFromInt(100)).Float64()

	if v < -1000 {
		return FormatKMB(v, 1) + "%"
	}

	if v < -0.01 {
		return fmt.Sprintf("%s%s", symbol, FormatFloorString(v, 2)) + "%"
	}

	if v > -0.01 && v < 0 {
		return "< -0.01%"
	}

	if v < 0.01 {
		return "< " + symbol + "0.01%"
	}

	if v < 1000 {
		return fmt.Sprintf("%s%s", symbol, FormatFloorString(v, 2)) + "%"
	}

	if v >= 1000 {
		return fmt.Sprintf("%s%s", symbol, FormatKMB(v, 1)) + "%"
	}

	return cast.ToString(v)
}

func FormatKMB(v float64, n int) string {
	if v <= 0 {
		return "0"
	}

	if v < 1000 {
		return FormatFloorString(v, n)
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
