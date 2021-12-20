package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type XTime struct {
	time.Time
}

func (t *XTime) UnmarshalJSON(data []byte) error {
	if len(data) == 2 {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = XTime{t1}
	return err
}

func (t XTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`""`), nil
	}
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
	ID        uint `gorm:"primary_key" json:"ID"`
	CreatedAt XTime
	UpdatedAt XTime
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type XDate struct {
	time.Time
}

func (t *XDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02", timeStr)
	*t = XDate{t1}
	return err
}

func (t XDate) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`""`), nil
	}
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(output), nil
}

func (t XDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *XDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
