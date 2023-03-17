package types

import (
	"database/sql/driver"
	"fmt"
)

// BitBool 自定义布尔值,与数据库bit(0)转换
type BitBool bool

func (bitBool *BitBool) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 && v[0] == 1 {
			*bitBool = true
		} else {
			*bitBool = false
		}
		return nil
	default:
		return fmt.Errorf("unsupported type for BitBool: %T", src)
	}
}

func (bitBool BitBool) Value() (driver.Value, error) {
	if bitBool {
		return []byte{1}, nil
	}
	return []byte{0}, nil
}
