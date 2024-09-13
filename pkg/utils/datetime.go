package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	DateLayout     = "2006-01-02"
	DateTimeLayout = "2006-01-02 15:04:05"

	DateLayoutSlash     = "2006/01/02"
	DateTimeLayoutSlash = "2006/01/02 15:04:05"

	DateShortLayout      = "06-01-02"
	DateShortLayoutSlash = "06/01/02"

	MonthDateLayout      = "01-02"
	MonthDateLayoutSlash = "01/02"

	TimeLayout       = "15:04:05"
	TimeMinuteLayout = "15:04"
)

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(DateTimeLayout))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	// 判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
