package utility

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func DB() *gorm.DB {
	dbInfo := "host=localhost user=postgres dbname=postgres sslmode=disable password=root"
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
