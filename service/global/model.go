package global

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type GAD_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt *LocalTime     `json:"created_at"`           // 创建时间
	UpdatedAt *LocalTime     `json:"updated_at"`           // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
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

type MyJson map[string]interface{}

func (j *MyJson) Scan(value interface{}) error {

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var mJson MyJson
	err := json.Unmarshal(bytes, &mJson)
	*j = mJson
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j MyJson) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	b, err := json.Marshal(j)
	return string(b), err
}
