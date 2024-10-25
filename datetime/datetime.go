package datetime

import (
	"math"
	"time"
)

// NowDate 获取当前日期
func NowDate() string {
	return time.Now().Format(time.DateOnly)
}

// NowTime 获取当前时间
func NowTime() string {
	return time.Now().Format(time.TimeOnly)
}

// NowDateTime 获取当前日期时间
func NowDateTime() string {
	return time.Now().Format(time.DateTime)
}

// AddSecond 增加秒
func AddSecond(t time.Time, second int) time.Time {
	return t.Add(time.Duration(second) * time.Second)
}

// AddMinute 增加分钟
func AddMinute(t time.Time, minute int) time.Time {
	return t.Add(time.Duration(minute) * time.Minute)
}

// AddHour 增加小时
func AddHour(t time.Time, hour int) time.Time {
	return t.Add(time.Hour * time.Duration(hour))
}

// AddDay 增加天数
func AddDay(t time.Time, day int) time.Time {
	return t.AddDate(0, 0, day)
}

// AddMonth 增加月份
func AddMonth(t time.Time, month int) time.Time {
	return t.AddDate(0, month, 0)
}

// AddYear 增加年份
func AddYear(t time.Time, year int) time.Time {
	return t.AddDate(year, 0, 0)
}

// Yesterday 昨天
func Yesterday() time.Time {
	return AddDay(time.Now(), -1)
}

// YesterdayFormat 获取昨天日期时间
func YesterdayFormat() string {
	return Yesterday().Format(time.DateTime)
}

// Tomorrow 明天
func Tomorrow() time.Time {
	return AddDay(time.Now(), 1)
}

func TomorrowFormat() string {
	return Tomorrow().Format(time.DateTime)
}

// BeginOfDay 获取当天开始时间
func BeginOfDay() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// EndOfDay 获取当天结束时间
func EndOfDay() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, time.Local)
}

// CalcDaysBetweenDates 计算两个日期相差天数，layout为时间格式，常见时间格式见format.go
func CalcDaysBetweenDates(d1, d2 string, layout string) (error, int64) {
	t1, err := time.Parse(layout, d1)
	if err != nil {
		return err, 0
	}
	t2, err := time.Parse(layout, d2)
	if err != nil {
		return err, 0
	}
	return nil, int64(math.Abs(t1.Sub(t2).Hours() / 24))
}
