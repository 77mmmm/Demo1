package mysql

import (
	"awesomeProject/model"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/dbsearch?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		err = fmt.Errorf("conn mysql error: %s", err)
		return nil, err
	}
	log.Println("连接成功")

	err = db.DB().Ping()
	if err != nil {
		err = fmt.Errorf("error:%s", err)
		fmt.Println(err.Error())
		return nil, err
	}
	db.AutoMigrate(&model.User{})
	db.SingularTable(true)
	return db, nil
}
