package controller

import (
	"card-master/database"
	request "card-master/lib/middleware/request"
	models "card-master/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type CardController struct {}

func (cc *CardController) GetCards(c *fiber.Ctx) error {
	// Get all cards from the database
	var cards []models.CardEntity
	if err := database.DB.Where(
		"deleted = false",
	).Find(&cards).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch cards",
		})
	}
	
	// Return the list of cards as JSON
	return c.JSON(cards)
}

func (cc *CardController) GetCardByID(c *fiber.Ctx) error {
	// Get the card ID from the URL parameters
	id := c.Params("id")

	// Find the card by ID
	var card models.CardEntity
	if err := GetCardByID(&card, id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Card not found",
		})
	}

	// Return the card as JSON
	return c.JSON(card)
}

func (cc *CardController) CreateCard(c *fiber.Ctx) error {
	// Parse the request body into a Card struct
	var card models.CardsTable
	if err := request.Validate(c, &card); err != nil {
		fmt.Println("Validation error:", err)
		return err;
	}
	
	// Save the card to the database
	if err := database.DB.Create(&card).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create card",
		})
	}

	// Return the created card as JSON
	return c.Status(201).JSON(card)
}

func (cc *CardController) UpdateCard(c *fiber.Ctx) error {
	// Get the card ID from the URL parameters
	id := c.Params("id")

	// Find the card by ID
	var card models.CardEntity
	if err := GetCardByID(&card, id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Card not found",
		})
	}

	// Parse the request body into a Card struct
	if err := request.Validate(c, &card); err != nil {
		return err;
	}

	// Update the card in the database
	if err := database.DB.Save(&card).Error; err != nil {
		// Handle the error if the update fails
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update card",
		})
	}

	// Return the updated card as JSON
	return c.JSON(card)
}

func (cc *CardController) DeleteCard(c *fiber.Ctx) error {
	// Get the card ID from the URL parameters
	id := c.Params("id")

	// Find the card by ID
	var card models.CardEntity
	if err := GetCardByID(&card, id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Card not found",
		})
	}
	// Delete the card from the database
	card.Deleted = true
	// Set the Deleted field to true instead of deleting the record
	if err := database.DB.Save(&card).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete card",
		})
	}
	// Return a success message
	return c.JSON(fiber.Map{
		"message": "Card deleted successfully",
	})
}



func GetCardByID(card *models.CardEntity, id string) (error) {
	// Find the card by ID
	if err := database.DB.Where(
		"deleted = false AND id = ?", id,
	).First(&card).Error; err != nil {
		return err
	}
	// Check if the card was found
	return nil
}