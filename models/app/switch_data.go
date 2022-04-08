/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/1/9 6:22 下午
 **/
package app

import (
	"gorm.io/datatypes"
	"srilanka/global/orm"
	"srilanka/models/base"
	"strings"
)

// 开关信息数据
type SwitchData struct {
	Title       string         `gorm:"type:varchar(128)" json:"title"`
	Data        datatypes.JSON `gorm:"type:json" json:"data"`
	Description string         `gorm:"type:varchar(128)" json:"description"`
	ShareToken  string         `gorm:"type:varchar(512)" json:"share_token"`
	base.BaseModel
}

func (SwitchData) TableName() string {
	return "s_switch_data"
}

func InsertSwitchData(switchData *SwitchData) (err error) {
	err = orm.MyDB.Create(&switchData).Error
	return
}

func SelectSwitchData(limit int, page int, createBy string, roleStr string, title string) (paginator *base.Paginator, err error) {
	var switchData []*SwitchData
	db := orm.MyDB.Table("s_switch_data").Order("id desc")
	if !strings.Contains(roleStr, "admin"){
		db.Where("create_by = ?",  createBy)
	}
	if title != "" {
		db = db.Where("create_by = ?  and  title like ?", createBy, "%"+title+"%")
	}
	paginator, err = base.Paging(db, limit, page, &switchData)
	return
}

func DeleteSwitchData(switchData *SwitchData) (err error) {
	err = orm.MyDB.Delete(switchData).Error
	return
}

func SelectSwitchDataByID(switchData *SwitchData) (err error) {
	err = orm.MyDB.Find(switchData, switchData.ID).Error
	return
}
