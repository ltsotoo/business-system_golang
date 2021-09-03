package model

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type XTime struct {
	time.Time
}

func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

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
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt XTime
	UpdatedAt XTime
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
