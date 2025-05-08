package controller

import (
	"card-master/database"
	request "card-master/lib/middleware/request"
	models "card-master/model"

	"github.com/gofiber/fiber/v2"
)

type CardSeriesController struct {}

func (cc *CardSeriesController) GetSeries(c *fiber.Ctx) error {
	// Get all card series from the database
	var series []models.CardSeriesEntity
	if err := database.DB.Where(
		"deleted = false",
	).Find(&series).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch card series",
		})
	}
	
	// Return the list of cards as JSON
	return c.JSON(series)
}

func (cc *CardSeriesController) GetSeriesByID(c *fiber.Ctx) error {
	// Get the card ID from the URL parameters
	id := c.Params("id")

	// Find the card by ID
	var series models.CardSeriesEntity
	if err := GetSeriesByID(&series, id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Card Series not found",
		})
	}

	// Return the card as JSON
	return c.JSON(series)
}

func (cc *CardSeriesController) CreateSeries(c *fiber.Ctx) error {
	// Parse the request body into a CardSeries struct
	var series models.CardSeriesTable
	if err := request.Validate(c, &series); err != nil {
		return err;
	}

	// Save the card to the database
	if err := database.DB.Create(&series).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create series",
		})
	}

	// Return the created card as JSON
	return c.Status(201).JSON(series)
}

func (cc *CardSeriesController) UpdateSeries(c *fiber.Ctx) error {
	// Get the card series ID from the URL parameters
	id := c.Params("id")

	// Find the card series by ID
	var series models.CardSeriesEntity
	if err := GetSeriesByID(&series, id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Card series not found",
		})
	}

	// Parse the request body into a CardSeries struct
	if err := request.Validate(c, &series); err != nil {
		return err;
	}

	// Update the series in the database
	if err := database.DB.Save(&series).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update card",
		})
	}

	// Return the updated series as JSON
	return c.JSON(series)
}

func (cc *CardSeriesController) DeleteSeries(c *fiber.Ctx) error {
	// Get the card series ID from the URL parameters
	id := c.Params("id")

	// Find the card series by ID
	var series models.CardSeriesEntity
	if err := GetSeriesByID(&series, id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Card not found",
		})
	}
	// Delete the card series from the database
	series.Deleted = true
	// Set the Deleted field to true instead of deleting the record
	if err := database.DB.Save(&series).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete series",
		})
	}
	// Return a success message
	return c.JSON(fiber.Map{
		"message": "Card Series deleted successfully",
	})
}

func GetSeriesByID(series *models.CardSeriesEntity, id string) (error) {
	// Find the card series by ID
	if err := database.DB.Where(
		"deleted = false AND id = ?", id,
	).First(&series).Error; err != nil {
		return err
	}
	// Return nil if found
	return nil
}