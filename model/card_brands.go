package models

import (
	"gorm.io/gorm"
)

type CardBrandsTable struct {
	gorm.Model
	Name     	string `json:"title" validate:"required,min=3,max=100"`
	Deleted	 	bool	`json:"deleted" gorm:"default:false"`
	
}

func (cc CardBrandsTable) TableName() string {
	return "card_brands"
}

type CardBrandEntity struct {
	CardBrandsTable
	CardSeries  []CardSeriesEntity `json:"card_series" gorm:"foreignKey:CardBrandsID"`
}

func (cc CardBrandEntity) LoadWithAssociations(db *gorm.DB) *gorm.DB {
	return db.Preload("CardSeries").Preload("CardSeries.Cards")
}