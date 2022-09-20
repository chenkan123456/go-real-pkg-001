package util

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	timeDateFormat      = "2006-01-02 15:04:05"
	timeDateShortFormat = "2006-01-02"
)

type FormatTime time.Time
type FormatShortTime FormatTime

func (t FormatTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(timeDateFormat))
	return []byte(formatted), nil
}

func (t FormatShortTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(timeDateShortFormat))
	return []byte(formatted), nil
}

func (t *FormatTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(timeDateFormat, timeStr)
	*t = FormatTime(t1)
	return err
}

func (t FormatTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format(timeDateFormat), nil
}

func (t *FormatTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = FormatTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t FormatShortTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format(timeDateShortFormat), nil
}

func CheckIsMinDate(date FormatShortTime) bool {
	var ti time.Time
	tTime := time.Time(date)
	return tTime == ti

}

func ReturnDateRange(s_date, e_date string) []string {

	var (
		year_list = [12]int{2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022}
		dates     = make([]string, 0)
		s         = 0
		e         = time.Now().Year()
		err       error
	)

	if s_date != "" {
		s, err = strconv.Atoi(s_date[0:4])
		if err != nil {
			s = year_list[0]
		}
	}

	if e_date != "" {
		e, err = strconv.Atoi(e_date[0:4])
		if err != nil {
			e = year_list[len(year_list)-1]
		}
	}

	for _, v := range year_list {
		if v >= s && v <= e {
			dates = append(dates, strconv.Itoa(v))
		}
	}

	return dates

}

func GetStartDate(date string) string {
	if date == "" {
		return time.Now().Format(timeDateShortFormat)
	} else {
		return date
	}

}
func GetEndDate(date string) string {
	if date == "" {
		return timeDateShortFormat
	} else {
		return date
	}

}

func ReturnEndDate(date string) string {
	if date != "" {
		date = strings.Replace(date, "/", "-", -1)
		//这边需要按照周进行处理，计算出date 周的最后一天
		t, err := time.Parse(timeDateShortFormat, date)
		if err != nil {
			return time.Now().Format(timeDateShortFormat)
		}
		d := ReturnNowWeekRange(t)
		return d[1]
	} else {

		return time.Now().Format(timeDateShortFormat)
	}

}

func ReturnStartDate(date string) string {
	if date != "" {
		date = strings.Replace(date, "/", "-", -1)
		//这边需要按照周进行处理，计算出date 周的第一天
		t, err := time.Parse(timeDateShortFormat, date)
		if err != nil {
			return time.Now().Format(timeDateShortFormat)
		}
		d := ReturnNowWeekRange(t)
		return d[0]
	} else {
		return timeDateShortFormat
	}

}

//本月第一天
func ReturnFirstDateOfMonth(date string) string {

	d := StringToShortDate(date)
	d = d.AddDate(0, 0, -d.Day()+1)
	f_date := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	return f_date.Format(timeDateShortFormat)

}

//本月最后一天
func ReturnLastDateOfMonth(date string) string {
	d := StringToShortDate(date)
	d = d.AddDate(0, 0, -d.Day()+1)
	f_date := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	l_date := f_date.AddDate(0, 1, -1)
	return l_date.Format(timeDateShortFormat)

}

func StringDateFormat(date string) string {
	if date == "" {
		return timeDateShortFormat
	} else {
		return date[0:10]
	}
}

func StringToShortDate(date string) time.Time {
	//获取平台项目信息  描述对比信息
	t, err := time.Parse(timeDateShortFormat, date)
	if err != nil {
		return time.Now()
	}
	return t
}

func ReturnNowWeekRange(date time.Time) []string {
	//now := time.Now().Add(-24 * 4 * time.Hour)
	//now := time.Now()
	offset := int(time.Monday - date.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStart := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)

	s_date := weekStart.Format(timeDateShortFormat)
	e_date := weekStart.Add(6 * 24 * time.Hour).Format(timeDateShortFormat)

	return []string{s_date, e_date}
}

func ReturnLastWeekRange(date time.Time) []string {
	n_date := ReturnNowWeekRange(date)
	l_date, _ := time.Parse(timeDateShortFormat, n_date[0])
	lw_date := ReturnNowWeekRange(l_date.AddDate(0, 0, -1))
	return lw_date
}

func ReturnLastMonthRange(date time.Time) []string {

	s_date := date.Add(-30 * 24 * time.Hour).Format(timeDateShortFormat)
	e_date := date.Format(timeDateShortFormat)
	return []string{s_date, e_date}
}

//获取当前日期的第一天和最后一天
func ReturnNowMonthkRange(d time.Time) []string {

	day := d.Day()

	if day >= 20 {
		d = d.AddDate(0, 0, -d.Day()+1)
	} else {
		//不满10天 取上个月数据
		d = d.AddDate(0, -1, -d.Day()+1)
	}

	f_date := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	l_date := f_date.AddDate(0, 1, -1)
	e_m := f_date.Format(timeDateShortFormat)
	l_m := l_date.Format(timeDateShortFormat)
	return []string{e_m, l_m}

}

//获取月
func ReturnDateMonthText(d string) string {

	date, err := time.Parse(timeDateShortFormat, d)
	if err != nil {
		return ""
	}
	m, err := strconv.Atoi(date.Format("01"))
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v月", strconv.Itoa(m))

}

func ReturnLastWeekRangeByYear(date time.Time, year int) []string {

	var (
		y_date_arry []string
		s_date      = date.Format(timeDateShortFormat)
		e_date      = date.AddDate(year, 0, 0)
	)

	if e_date.After(time.Now()) {
		e_date = time.Now()
	}
	y_date_slice := append(y_date_arry, s_date)

	n_date := ReturnNowWeekRange(date)
	l_date, _ := time.Parse(timeDateShortFormat, n_date[1])

	for l_date.Before(e_date) {
		l_date_str := l_date.Format(timeDateShortFormat)
		y_date_slice = append(y_date_slice, l_date_str)

		n_date = ReturnNowWeekRange(l_date.AddDate(0, 0, 1))
		l_date, _ = time.Parse(timeDateShortFormat, n_date[1])
	}

	return y_date_slice
}
