/**
 * @Author xuzhipeng
 * @Description
 * @Date 2021/9/26 6:21 下午
 **/
package app

import (
	"srilanka/global/orm"
	"srilanka/models/base"
)

// 产品表
type Product struct {
	ProductName string    `gorm:"type:varchar(128)" json:"product_name"` // 产品名
	Description string    `gorm:"type:varchar(1024)" json:"description"`  // 产品描述
	Projects    []Project `gorm:"foreignkey:ProductID;references:ID" json:"projects"`
	base.BaseModel
}

func (Product) TableName() string {
	return "s_product"
}

func InsertProduct(product *Product) (err error) {
	err = orm.MyDB.Create(&product).Error
	return
}

func SelectProduct(limit int, page int) (paginator *base.Paginator, err error) {
	var products []*Product
	paginator, err = base.Paging(orm.MyDB.Table("s_product"), limit, page, &products)
	return
}
func SelectAllProductDepth() (products *[]Product, err error) {
	err = orm.MyDB.Preload("Projects").Preload("Projects.Versions").Find(&products).Error
	return
}

func UpdateProduct(product *Product) (err error) {
	err = orm.MyDB.Updates(product).Error
	return
}

func DeleteProduct(product *Product) (err error) {
	err = orm.MyDB.Delete(product).Error
	return
}
