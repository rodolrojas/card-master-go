package models

import (
	"gorm.io/gorm"
)

type CardsTable struct {
	gorm.Model
	Title       	string		`json:"title" validate:"required,min=3,max=100"`
	CardID			string		`json:"card_id" validate:"required,min=3,max=100"`
	CardSeriesID 	int		    `json:"card_series_id" validate:"required"`
	Value			float32		`json:"value" validate:"required,gte=0,lte=10000"`
	Deleted 		bool		`json:"deleted" gorm:"default:false"`	
}

func (cc CardsTable) TableName() string {
	return "cards"
}


type CardEntity struct {
	CardsTable
	CardSeries		CardSeriesEntity	`json:"card_series" gorm:"foreignKey:CardSeriesID;references:ID"`
}

func (cc CardEntity) LoadWithAssociations(db *gorm.DB) *gorm.DB {
	return db.Preload("CardSeries").Preload("CardSeries.CardBrand")
}