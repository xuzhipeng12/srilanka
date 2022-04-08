/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/2/26 4:08 下午
 **/
package system

import "srilanka/models/base"

// 系统角色表
type Role struct {
	RoleName string `json:"roleName" gorm:"type:varchar(128);"` // 角色名称
	Remark   string `json:"remark" gorm:"type:varchar(255);"`   //备注
	Admin    bool   `json:"admin" gorm:"type:char(1);"`
	Params   string `json:"params" gorm:"-"`
	base.BaseModel
}

func (Role) TableName() string {
	return "sys_role"
}
