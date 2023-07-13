package db

import (
	"fmt"
	"github.com/setcy/spider/src/post"
	"github.com/setcy/spider/src/subscribe"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&post.Post{}, &subscribe.Subscribe{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}

	fmt.Println("Database and table created successfully.")

	return db
}
