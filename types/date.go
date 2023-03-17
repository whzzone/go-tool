package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Date 自定义日期
type Date time.Time

const dateFormat = "2006-01-02"

func (t *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(dateFormat, timeStr)
	*t = Date(t1)
	return err
}

func (t Date) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(dateFormat))
	return []byte(formatted), nil
}

func (t Date) Value() (driver.Value, error) {
	// Date 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format(dateFormat), nil
}

func (t *Date) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = Date(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *Date) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
