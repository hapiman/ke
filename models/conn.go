package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var KePush *gorm.DB

// ConnPush 连接push库
func ConnKe() *gorm.DB {
	if KePush != nil {
		return KePush
	}
	sSqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "root", "123456", "localhost", "3306", "ke")
	db, err := gorm.Open("mysql", sSqlConnStr)
	if err != nil {
		panic(err)
	}
	// AutoMigrate
	// db.AutoMigrate(&model.MsgBenefitExpire{},
	// 	&model.MsgBenefitReceive{},
	// 	&model.MsgPaymentReceive{},
	// 	&model.MsgForecast{},
	// 	&model.MsgUser{}
	// )

	// 不要取复数
	db.SingularTable(true)
	// 查看执行日志
	db.LogMode(false)

	KePush = db
	return db
}
