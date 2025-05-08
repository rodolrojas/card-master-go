package database

import (
	models "card-master/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB() error {
	var err error
	// Connect to the database
	DB, err = gorm.Open(sqlite.Open("cards.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	if err = DB.AutoMigrate(&models.CardsTable{}, &models.CardSeriesTable{}, &models.CardBrandsTable{}); err != nil {
		return err
	}

	return nil
}