package dal

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local"

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf(s,
		"mytalk",
		"mytalk",
		"localhost:3306",
		"mytalk",
		"utf8",
		"True")),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})

	if err != nil {
		log.Panic(err)
	}

	if err = DB.AutoMigrate(&User{}); err != nil {
		log.Panic(err)
	}
}
