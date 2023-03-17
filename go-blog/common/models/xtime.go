package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type XTime struct {
	time.Time
}

// MarshalJSON 重写MarshalJSON方法 自定义时间格式
func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// Value 实现Value方法 写入数据库时会调用方法将自定义的时间类型写入数据库
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not comvert %v to timestamp", v)
}
