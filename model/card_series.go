package models

import (
	"gorm.io/gorm"
)

type CardSeriesTable struct {
	gorm.Model
	Title       	string 		`json:"title" validate:"required,min=3,max=100"`
	Year			string    	`json:"year" validate:"required,min=4,max=5"`
	CardBrandsID 	uint   		`json:"card_brands_id" validate:"required"`
	Deleted			bool   		`json:"deleted" gorm:"default:false"`
	
}

func (cc CardSeriesTable) TableName() string {
	return "card_series"
}

type CardSeriesEntity struct {
	CardSeriesTable
	CardBrand		CardBrandsTable 	`json:"card_brands" gorm:"foreignKey:CardBrandsID;references:ID"`
	Card			[]CardsTable 		`json:"card" gorm:"foreignKey:CardSeriesID;references:ID"`
}	

func (cc CardSeriesEntity) LoadWithAssociations(db *gorm.DB) *gorm.DB {
	return db.Preload("CardBrand").Preload("Card")
}