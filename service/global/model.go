package global

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type GsModel struct {
	ID        uint           `gorm:"primarykey;autoIncrement" json:"id"` // 主键ID
	CreatedAt *LocalTime     `json:"created_at" gorm:"type:datetime(3)"` // 创建时间
	UpdatedAt *LocalTime     `json:"updated_at" gorm:"type:datetime(3)"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                     // 删除时间
}

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// Scan @ignore waring
func (t *LocalTime) Scan(v interface{}) error {

	switch v := v.(type) {
	case string:
		ct, err := time.Parse("2006-01-02 15:04:05.999", v)
		*t = LocalTime(ct)
		return err
	default:
		return fmt.Errorf("unsupported type for CustomTime: %T", v)
	}
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
