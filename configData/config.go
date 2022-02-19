package configData

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() (*gorm.DB, error) {
	d := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"root", "", "0.0.0.0", 3306, "todos")
	DB, err = gorm.Open(mysql.Open(d), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return DB, err
}
