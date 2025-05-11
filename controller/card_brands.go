package controller

import (
	"card-master/database"
	request "card-master/lib/middleware/request"
	models "card-master/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CardBrandsController struct {
	CardBrandsModel *gorm.DB
}

func CardBrandsControllerProvider() *CardBrandsController {
	// Initialize the CardBrandsController with the database connection
	return &CardBrandsController{
		CardBrandsModel: (&models.CardBrandEntity{}).LoadWithAssociations(database.DB),
	}
}

func (cc *CardBrandsController) GetBrands(c *fiber.Ctx) error {
	// Get all card brands from the database
	var brands []models.CardBrandEntity
	if err := cc.CardBrandsModel.Where(
		"deleted = false",
	).Find(&brands).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch card brands",
		})
	}
	
	// Return the list of cards as JSON
	return c.JSON(brands)
}

func (cc *CardBrandsController) GetBrandByID(c *fiber.Ctx) error {
	// Get the card ID from the URL parameters
	id := c.Params("id")

	// Find the card by ID
	var brand models.CardBrandEntity
	if err := cc.__GetBrandByID(&brand, id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Card brand not found",
		})
	}

	// Return the card as JSON
	return c.JSON(brand)
}

func (cc *CardBrandsController) CreateBrand(c *fiber.Ctx) error {
	// Parse the request body into a CardBrands struct
	var brand models.CardBrandsTable
	if err := request.Validate(c, &brand); err != nil {
		return err;
	}

	// Save the card to the database
	if err := cc.CardBrandsModel.Create(&brand).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create brand",
		})
	}

	// Return the created card as JSON
	return c.Status(201).JSON(brand)
}

func (cc *CardBrandsController) UpdateBrand(c *fiber.Ctx) error {
	// Get the card brands ID from the URL parameters
	id := c.Params("id")

	// Find the card brands by ID
	var brand models.CardBrandEntity
	if err := cc.__GetBrandByID(&brand, id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Card brand not found",
		})
	}

	// Parse the request body into a CardBrands struct
	if err := request.Validate(c, &brand); err != nil {
		return err;
	}

	// Update the brands in the database
	if err := cc.CardBrandsModel.Save(&brand).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update brand",
		})
	}

	// Return the updated brands as JSON
	return c.JSON(brand)
}

func (cc *CardBrandsController) DeleteBrand(c *fiber.Ctx) error {
	// Get the card brands ID from the URL parameters
	id := c.Params("id")

	// Find the card brands by ID
	var brand models.CardBrandEntity
	if err := cc.__GetBrandByID(&brand, id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Brand not found",
		})
	}
	// Delete the card brands from the database
	brand.Deleted = true
	// Set the Deleted field to true instead of deleting the record
	if err := cc.CardBrandsModel.Save(&brand).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete brand",
		})
	}
	// Return a success message
	return c.JSON(fiber.Map{
		"message": "Card brand deleted successfully",
	})
}

func (cc *CardBrandsController) __GetBrandByID(brands *models.CardBrandEntity, id string) (error) {
	// Find the card brands by ID
	if err := cc.CardBrandsModel.Where(
		"deleted = false AND id = ?", id,
	).First(&brands).Error; err != nil {
		return err
	}
	// If the brands is found, return nil
	return nil
}