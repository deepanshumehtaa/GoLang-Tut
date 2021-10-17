// The purpose of this file is to send the variable db to the whole file
// Basically a db cursor file, but as an app

package config

// package for MySql
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// function to talk with db
func Connect() {
	data, err := gorm.Open("mysql", "deepanshu:my_sql_pass/<table_name?>charset=utf8&parseTime&")
	if err != nil {
		panic(err) // if there is an error raise the panic
	}
	db = data
}

func GetDB() *gorm.DB {
	return db
}
