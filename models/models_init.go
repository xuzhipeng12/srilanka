/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/1/9 8:07 下午
 **/
package models

import (
	"srilanka/global/orm"
	"srilanka/models/app"
	"srilanka/models/system"
	//"srilanka/tools"
)

func InitDbModels() error {
	//var err error
	//var db *gorm.DB
	//if tools.Cfg.DBType == "mysql" {
	//	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", tools.Cfg.DBUser, tools.Cfg.DBPwd, tools.Cfg.DBHost, tools.Cfg.DBPort, tools.Cfg.DBName)
	//	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//} else if tools.Cfg.DBType == "sqlite" {
	//	dsn := tools.Cfg.SQLiteDBName
	//	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	//} else {
	//	return nil
	//}
	//if err != nil {
	//	return err
	//}
	// 初始化表
	return orm.MyDB.AutoMigrate(
		new(app.Product),
		new(app.Project),
		new(app.SwitchData),
		new(app.SwitchTemplate),
		new(app.Version),
		new(system.CasbinRule),
		new(system.Role),
		new(system.User),
	)
}
