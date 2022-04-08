/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/1/9 11:45 下午
 **/

package base

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type BaseModel struct {
	ID        int       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
	CreatedAt JSONTime  `gorm:"column:create_at" json:"create_at"`
	UpdatedAt JSONTime  `gorm:"column:update_at" json:"update_at"`
	DeletedAt *JSONTime `gorm:"column:delete_at" sql:"index" json:"-"`

	CreateBy string `gorm:"column:create_by" json:"create_by"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
}

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("无法转换 %v 的时间格式", v)
}
