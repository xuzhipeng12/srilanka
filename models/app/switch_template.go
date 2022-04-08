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
)

// 开关模板表
type SwitchTemplate struct {
	TemplateName string         `gorm:"type:varchar(128)" json:"template_name"`
	Template     datatypes.JSON `gorm:"type:json" json:"template"`
	Description  string         `gorm:"type:varchar(128)" json:"description"`
	Versions     []Version      `gorm:"foreignkey:SwitchTemplateId;references:ID" json:"versions"`
	base.BaseModel
}

func (SwitchTemplate) TableName() string {
	return "s_switch_template"
}

func InsertTemplate(template *SwitchTemplate) (err error) {
	err = orm.MyDB.Create(&template).Error
	return
}

func SelectTemplate(limit int, page int, name string) (paginator *base.Paginator, err error) {
	var switchTemplate []*SwitchTemplate
	db := orm.MyDB.Table("s_switch_template")
	if name != "" {
		db = db.Where("template_name like ?", "%"+name+"%")
	}
	paginator, err = base.Paging(db, limit, page, &switchTemplate)
	return
}

func UpdateTemplate(template *SwitchTemplate) (err error) {
	err = orm.MyDB.Updates(template).Error
	return
}

func DeleteTemplate(template *SwitchTemplate) (err error) {
	err = orm.MyDB.Delete(template).Error
	return
}

func SelectTemplateById(template *SwitchTemplate) (err error) {
	err = orm.MyDB.Find(template, template.ID).Error
	return
}
