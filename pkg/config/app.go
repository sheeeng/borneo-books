package config

import (
	"github.com/jinzhu/gorm"
	// Use the blank identifier as explicit package name to import a package solely for its side-effects (initialization).
	// https://golang.org/ref/spec#Import_declarations
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func Connect() {
	// Please define your user name and password for my sql.
	//d, err := gorm.Open("mysql", "root:root@/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
