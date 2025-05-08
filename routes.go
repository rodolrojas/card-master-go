package main

import (
	"card-master/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	cards := api.Group("/cards")
	CardController := controller.CardController{}
	cards.Get("/", CardController.GetCards)
	cards.Get("/:id", CardController.GetCardByID)
	cards.Post("/", CardController.CreateCard)
	cards.Put("/:id", CardController.UpdateCard)
	cards.Delete("/:id", CardController.DeleteCard)
	
	series := api.Group("/card_series")
	CardSeriesController := controller.CardSeriesController{}
	series.Get("/", CardSeriesController.GetSeries)
	series.Get("/:id", CardSeriesController.GetSeriesByID)
	series.Post("/", CardSeriesController.CreateSeries)
	series.Put("/:id", CardSeriesController.UpdateSeries)
	series.Delete("/:id", CardSeriesController.DeleteSeries)
	
	brands := api.Group("/card_brands")
	CardBrandsController := controller.CardBrandsController{}
	brands.Get("/", CardBrandsController.GetBrands)
	brands.Get("/:id", CardBrandsController.GetBrandByID)
	brands.Post("/", CardBrandsController.CreateBrand)
	brands.Put("/:id", CardBrandsController.UpdateBrand)
	brands.Delete("/:id", CardBrandsController.DeleteBrand)
	
}