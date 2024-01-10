package tms

import "time"

// 将类似2023-03-29T14:15:20+08:00这样的字符串时间类型转为time类型，方便时间的运算

func FormatStrTimeToTime(strTime string) time.Time {
	parseTime, _ := time.Parse(time.RFC3339, strTime)
	return parseTime
}

func Format(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
func FormatYMD(t time.Time) string {
	return t.Format("2006-01-02")
}
func FormatByMill(t int64) string {
	return time.UnixMilli(t).Format("2006-01-02 15:04:05")
}

func ParseTime(str string) int64 {
	parse, _ := time.Parse("2006-01-02 15:04:05", str)
	return parse.UnixMilli()
}

func FormatNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func FormatLatestTime() string {
	max := time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
	return max.Format("2006-01-02 15:04:05")
}

func GetTodayTS() int64 {
	now := time.Now()
	zeroTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return zeroTime.Unix()
}

func GetNowTimeFloat64() float64 {
	return float64(time.Now().UnixNano())
}

func GetTimeFloat64(t time.Time) float64 {
	return float64(t.UnixNano())
}
