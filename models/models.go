package models

import (
	"fmt"
	"github.com/BigerWANG/my-gin/pkg/setting"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreateOn int `json:"create_on"`
	ModifiedOn int `json:"modified_on"`
}

func init()  {
	// 初始化数据库连接
	var(
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil{
		log.Fatal(fmt.Sprintf("Fail to get database conf %v", err))
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
		))
	if err != nil{
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetConnMaxIdleTime(10)
	db.DB().SetMaxOpenConns(100)

}

func CloseDB()  {
	if err := db.Close(); err != nil{
		log.Fatal(err)
	}
}