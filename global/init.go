/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/2/27 10:08 上午
 **/
package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"srilanka/global/orm"
	"srilanka/tools"
)

func DBSetup() {
	//// if MyDB == nil {
	//dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", tools.Cfg.DBUser, tools.Cfg.DBPwd, tools.Cfg.DBHost, tools.Cfg.DBPort, tools.Cfg.DBName)
	//fmt.Println(dsn)
	//orm.MyDB, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//// }
	var err error
	if tools.Cfg.DBType == "mysql" {
		dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", tools.Cfg.DBUser, tools.Cfg.DBPwd, tools.Cfg.DBHost, tools.Cfg.DBPort, tools.Cfg.DBName)
		orm.MyDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if tools.Cfg.DBType == "sqlite" {
		dsn := tools.Cfg.SQLiteDBName
		orm.MyDB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	} else {
		fmt.Println("数据库类型设置错误，请检查")
	}
	if err != nil {
		fmt.Println("初始化数据库连接失败: ", err.Error())
	}
}
