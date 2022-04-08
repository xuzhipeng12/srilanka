/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/1/9 6:22 下午
 **/
package app

import (
	"srilanka/global/orm"
	"srilanka/models/base"
)

type VersionListResult struct {
	Version
	ProjectName string `json:"project_name"` // 项目名

}

// 项目版本表
type Version struct {
	ProjectID        int    `gorm:"type:integer(11)" json:"project_id"`
	SwitchTemplateId int    `gorm:"type:integer(11)" json:"switch_template_id"`
	VersionName      string `gorm:"type:varchar(11)" json:"version_name"` // 版本名
	Description      string `gorm:"type:varchar(128)" json:"description"` // 版本说明
	base.BaseModel
}

func (Version) TableName() string {
	return "s_version"
}

func InsertVersion(version *Version) (err error) {
	err = orm.MyDB.Create(&version).Error
	return
}

func SelectVersion(limit int, page int) (paginator *base.Paginator, err error) {
	var version []*VersionListResult
	db := orm.MyDB.Table("s_version").
		Select("s_version.*, s_project.project_name").
		Joins("left join s_project on s_version.project_id  = s_project.id ")
	paginator, err = base.Paging(db, limit, page, &version)
	return
}

func UpdateVersion(version *Version) (err error) {
	err = orm.MyDB.Updates(version).Error
	return
}

func DeleteVersion(version *Version) (err error) {
	err = orm.MyDB.Delete(version).Error
	return
}

func SelectVersionTemplateById(versionId int) (template *SwitchTemplate, err error) {
	err = orm.MyDB.Raw("select sst.* from s_switch_template sst , s_version sv  where sst.id = sv.switch_template_id  and sv.id= ?", versionId).Scan(&template).Error
	return
}
