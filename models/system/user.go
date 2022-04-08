/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/2/26 4:08 下午
 **/
package system

import (
	"gorm.io/datatypes"
	"srilanka/global/orm"
	"srilanka/models/base"
)

// 系统用户表
type User struct {
	Name     string         `gorm:"type:varchar(128)" json:"name"`     // 昵称
	NickName string         `gorm:"type:varchar(128)" json:"nickName"` // 昵称
	Password string         `gorm:"type:varchar(255)" json:"password"` // 密码
	Avatar   string         `gorm:"type:varchar(255)" json:"avatar"`   //头像
	Email    string         `gorm:"type:varchar(128)" json:"email"`    //邮箱
	Status   int            `gorm:"type:int" json:"status"`            //用户状态
	Roles    datatypes.JSON `gorm:"type:json" json:"roles"`            //角色
	base.BaseModel
}

func (User) TableName() string {
	return "sys_user"
}
func InsertUser(user *User) (err error) {
	err = orm.MyDB.Create(&user).Error
	return
}

func SelectUserByPasswd(userName string, password string) (user *User, err error) {
	err = orm.MyDB.Where("name = ?  AND password = ? AND status = 1", userName, password).First(&user).Error
	return
}

func SelectlUser(limit int, page int, name string) (paginator *base.Paginator, err error) {
	var user []*User
	db := orm.MyDB.Table("sys_user").Where("status != ?", 0).Select("id", "name", "nick_name", "avatar", "email", "status", "roles")
	if name != "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	paginator, err = base.Paging(db, limit, page, &user)
	return
}

func UpdateUser(user *User) (err error) {
	err = orm.MyDB.Updates(user).Error
	return
}

func DeleteUser(user *User) (err error) {
	err = orm.MyDB.Model(&user).Updates(map[string]interface{}{"status": 0}).Error
	return
}
