package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitialMigration() {
	db, err := gorm.Open("postgres", "host=localhost port=54320 user=db_user dbname=db_name password=db_password sslmode=disable")
	if err != nil {
		fmt.Println(err.Error)
		panic("Db connection failed.")
	}
	defer db.Close()

	db.AutoMigrate(&Beer{})
	db.AutoMigrate(&BeerDatasetElement{})
}
