package migration

import (
	"fmt"
	"log"

	"github.com/samsul-rijal/go-api/config/database"
	"github.com/samsul-rijal/go-api/model/entity"
)

func RunMigration()  {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}