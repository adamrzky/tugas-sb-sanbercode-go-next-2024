package controllers

import (
	"final-project-golang-individu/config"
	"final-project-golang-individu/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResp struct {
	Error string `json:"error"`
}

// FoodInput represents the input for creating a food item
type FoodInput struct {
	Name        string  `json:"name" binding:"required"`
	Restaurant  uint    `json:"restaurant_id" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
}

// FoodResponse represents the response for a food item
type FoodResponse struct {
	ID          uint    `json:"id"`
	Restaurant  uint    `json:"restaurant_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// CreateFood godoc
// @Summary Create a new food item
// @Description Create a new food item in the restaurant
// @Tags food
// @Accept json
// @Produce json
// @Param food body FoodInput true "Food Input"
// @Success 201 {object} FoodResponse
// @Failure 400 {object} ErrorResp "Bad Request"
// @Router /foods [post]
// @SecurityDefinitionsApikey AuthKey
// @in header
// @name Authorization
func CreateFood(c *gin.Context) {
	var input FoodInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the RestaurantID from the input
	food := models.Food{
		Name:         input.Name,
		Description:  input.Description,
		Price:        input.Price,
		RestaurantID: input.Restaurant, // Memastikan RestaurantID diatur dengan benar
	}

	if err := config.DB.Create(&food).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := FoodResponse{
		ID:          food.ID,
		Restaurant:  food.RestaurantID,
		Name:        food.Name,
		Description: food.Description,
		Price:       food.Price,
	}

	c.JSON(http.StatusCreated, response)
}

// GetFoods godoc
// @Summary Get all food items
// @Description Retrieve all food items in the restaurant
// @Tags food
// @Produce json
// @Success 200 {array} FoodResponse
// @Router /foods [get]
func GetFoods(c *gin.Context) {
	var foods []models.Food
	if err := config.DB.Find(&foods).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []FoodResponse
	for _, food := range foods {
		response = append(response, FoodResponse{
			ID:          food.ID,
			Restaurant:  food.RestaurantID,
			Name:        food.Name,
			Description: food.Description,
			Price:       food.Price,
		})
	}

	c.JSON(http.StatusOK, response)
}
