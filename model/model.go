//数据库交互,数据库逻辑层
package model

import (
	"blog/pkg/setting"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB //供调用
	userName string
	passWord string
)

//链接数据库
func init() {
	//
	userName := setting.Cfg.Section("database").Key("USER").String()
	passWord := setting.Cfg.Section("database").Key("PASSWORD").String()
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local", userName, passWord)
	//dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	//自定义配置
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库链接失败", err)
	}
	//获取底层DB来设置最大连接数
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

}
