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

type ProjectListResult struct {
	Project
	ProductName string `json:"product_name"` // 产品名

}

// 项目表, 项目属于产品
type Project struct {
	ProductID   int       `gorm:"type:integer(11)" json:"product_id"`
	ProjectName string    `gorm:"type:varchar(11)" json:"project_name"` // 项目名
	Description string    `gorm:"type:varchar(128)" json:"description"` // 项目描述
	Versions    []Version `gorm:"foreignkey:ProjectID;references:ID" json:"versions"`
	base.BaseModel
}

func (Project) TableName() string {
	return "s_project"
}

func InsertProject(project *Project) (err error) {
	err = orm.MyDB.Create(&project).Error
	return
}

func SelectProject(limit int, page int) (paginator *base.Paginator, err error) {
	var project []*ProjectListResult
	db := orm.MyDB.Table("s_project")
	db.Select(" s_project.* , s_product.product_name ").Joins(" left join s_product  on   s_project.product_id  = s_product.id ")
	paginator, err = base.Paging(db, limit, page, &project)
	return
}

func UpdateProject(project *Project) (err error) {
	err = orm.MyDB.Updates(project).Error
	return
}

func DeleteProject(project *Project) (err error) {
	err = orm.MyDB.Delete(project).Error
	return
}
