package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

// DateTime 自定义时间
type DateTime time.Time

const timeFormat = "2006-01-02 15:04:05"

func (t *DateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(timeFormat, timeStr)
	*t = DateTime(t1)
	return err
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(timeFormat))
	return []byte(formatted), nil
}

func (t DateTime) Value() (driver.Value, error) {
	// DateTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format(timeFormat), nil
}

func (t *DateTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = DateTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *DateTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
