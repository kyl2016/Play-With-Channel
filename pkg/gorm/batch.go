package main

import "github.com/jinzhu/gorm"

func BatchCreate(db *gorm.DB, models []interface{}) error {
	tx := db.Begin()
	for _, model := range models {
		tx.Create(model)
	}
	return tx.Commit().Error
}
