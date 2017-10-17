package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=gorm dbname=channels sslmode=disable password=mypassword")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
}
