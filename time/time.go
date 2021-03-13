package time

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Date 等同于PHP的date函数
// Date("Y-m-d H:i:s", time.Now())
func Date(format string, ts ...time.Time) string {
	patterns := []string{
		// 年
		"Y", "2006", // 4 位数字完整表示的年份
		"y", "06", // 2 位数字表示的年份

		// 月
		"m", "01", // 数字表示的月份，有前导零
		"n", "1", // 数字表示的月份，没有前导零
		"M", "Jan", // 三个字母缩写表示的月份
		"F", "January", // 月份，完整的文本格式，例如 January 或者 March

		// 日
		"d", "02", // 月份中的第几天，有前导零的 2 位数字
		"j", "2", // 月份中的第几天，没有前导零

		"D", "Mon", // 星期几，文本表示，3 个字母
		"l", "Monday", // 星期几，完整的文本格式;L的小写字母

		// 时间
		"g", "3", // 小时，12 小时格式，没有前导零
		"G", "15", // 小时，24 小时格式，没有前导零
		"h", "03", // 小时，12 小时格式，有前导零
		"H", "15", // 小时，24 小时格式，有前导零

		"a", "pm", // 小写的上午和下午值
		"A", "PM", // 小写的上午和下午值

		"i", "04", // 有前导零的分钟数
		"s", "05", // 秒数，有前导零
	}
	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)

	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	if t.Unix() <= 0 {
		return ""
	}
	return t.Format(format)
}

// StrToTime 等同于PHP的strtotime函数
// StrToTime("2020-12-19 14:16:22")
func StrToTime(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, errors.New("value is null")
	}
	layouts := []string{
		"20060102",
		"20060102150405",
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return t, nil
		}
	}
	return t, errors.Wrap(err, "StrToTime fail:")
}

// StrToLocalTime 字符串转本地时间
func StrToLocalTime(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, errors.New("value is null")
	}
	zoneName, offset := time.Now().Zone()

	zoneValue := offset / 3600 * 100
	if zoneValue > 0 {
		value += fmt.Sprintf(" +%04d", zoneValue)
	} else {
		value += fmt.Sprintf(" -%04d", zoneValue)
	}

	if zoneName != "" {
		value += " " + zoneName
	}
	return StrToTime(value)
}

// DateFormat 格式化time.Time
// DateFormat("YYYY-MM-DD HH:mm:ss", time.Now())
func DateFormat(format string, t time.Time) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}

// GetDaysAgoZeroTime 以当天0点为基准，获取前后某天0点时间
// 昨天：GetDaysAgoZeroTime(-1)
// 今天：GetDaysAgoZeroTime(0)
// 明天：GetDaysAgoZeroTime(1)
func GetDaysAgoZeroTime(day int) time.Time {
	date := time.Now().AddDate(0, 0, day).Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", date, time.Local)
	return t
}

// TimeToHuman 根据时间戳获得人类可读时间
func TimeToHuman(ts int) string {
	var res = ""
	if ts == 0 {
		return res
	}

	t := int(time.Now().Unix()) - ts
	data := [7]map[string]interface{}{
		{"key": 31536000, "value": "年"},
		{"key": 2592000, "value": "个月"},
		{"key": 604800, "value": "星期"},
		{"key": 86400, "value": "天"},
		{"key": 3600, "value": "小时"},
		{"key": 60, "value": "分钟"},
		{"key": 1, "value": "秒"},
	}
	for _, v := range data {
		var c = t / v["key"].(int)
		if 0 != c {
			suffix := "前"
			if c < 0 {
				suffix = "后"
				c = -c
			}
			res = strconv.Itoa(c) + v["value"].(string) + suffix
			break
		}
	}

	return res
}

// StringTimestampToTime StringTimestampToTime("1600000000")
func StringTimestampToTime(str string) time.Time {
	intTmp, _ := strconv.Atoi(str)
	return time.Unix(int64(intTmp), 0)
}

// StringTimestampToDatetime StringTimestampToDatetime("1600000000")
func StringTimestampToDatetime(str string) string {
	intTmp, _ := strconv.Atoi(str)
	return Date("Y-m-d H:i:s", time.Unix(int64(intTmp), 0))
}

// GetCurrentDate 获取当前的时间 - 字符串
func GetCurrentDate() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// GetCurrentUnix 获取当前的时间 - Unix时间戳
func GetCurrentUnix() int64 {
	return time.Now().Unix()
}

// GetCurrentMilliUnix 获取当前的时间 - 毫秒级时间戳
func GetCurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// GetCurrentNanoUnix 获取当前的时间 - 纳秒级时间戳
func GetCurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}

// TruncateHour 小时向下取整
func TruncateHour(t time.Time) time.Time {
	return t.Truncate(1 * time.Hour)
}

// RoundHour 小时向上取整
func RoundHour(t time.Time) time.Time {
	return t.Round(1 * time.Hour)
}

// TruncateMinute 分钟向下取证
func TruncateMinute(t time.Time) time.Time {
	return t.Truncate(1 * time.Minute)
}

// RoundMinute 分钟向上取整
func RoundMinute(t time.Time) time.Time {
	return t.Round(1 * time.Minute)
}

// TruncateHourStr string小时向下取整
func TruncateHourStr(str string) (time.Time, error) {
	t, err := StrToTime(str)
	if err != nil {
		return t, err
	}
	return t.Truncate(1 * time.Hour), nil
}

// RoundHourStr string小时向上取整
func RoundHourStr(str string) (time.Time, error) {
	t, err := StrToTime(str)
	if err != nil {
		return t, err
	}
	return t.Round(1 * time.Hour), nil
}

// TruncateMinuteStr string分钟向下取整
func TruncateMinuteStr(str string) (time.Time, error) {
	t, err := StrToTime(str)
	if err != nil {
		return t, err
	}
	return t.Truncate(1 * time.Minute), nil
}

// RoundMinuteStr string分钟向上取整
func RoundMinuteStr(str string) (time.Time, error) {
	t, err := StrToTime(str)
	if err != nil {
		return t, err
	}
	return t.Round(1 * time.Minute), nil
}
