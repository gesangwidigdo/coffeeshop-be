package controllers

import (
	"net/http"

	"github.com/gesangwidigdo/coffeeshop-be/dto"
	"github.com/gesangwidigdo/coffeeshop-be/initializers"
	"github.com/gesangwidigdo/coffeeshop-be/models"
	"github.com/gesangwidigdo/coffeeshop-be/utils"
	"github.com/gin-gonic/gin"
)

func CreateIngredient(c *gin.Context) {
	var ingredientInput dto.IngredientDTO

	if bindErr := utils.BindData(&ingredientInput, c); !bindErr {
		return
	}

	// Create ingredient
	ingredient := models.Ingredient{
		Ingredient_name: ingredientInput.Ingredient_name,
		Stock:           ingredientInput.Stock,
	}

	if result := initializers.DB.Create(&ingredient); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to create data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusCreated, "ok", "data", ingredient, c)
}

func GetAllIngredient(c *gin.Context) {
	var ingredients []models.Ingredient

	if result := initializers.DB.Find(&ingredients); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", ingredients, c)
}

func GetIngredientByID(c *gin.Context) {
	id := c.Param("id")

	var ingredient models.Ingredient
	if result := initializers.DB.First(&ingredient, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", ingredient, c)
}

func UpdateIngredient(c *gin.Context) {
	var ingredientInput dto.IngredientDTO

	if bindErr := utils.BindData(&ingredientInput, c); !bindErr {
		return
	}
	
	id := c.Param("id")

	var ingredient models.Ingredient
	if findRes := initializers.DB.First(&ingredient, id); findRes.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "error", findRes.Error.Error(), c)
		return
	}

	// update data
	newData := models.Ingredient{
		Ingredient_name: ingredientInput.Ingredient_name,
		Stock:           ingredientInput.Stock,
	}

	if updateResult := initializers.DB.Model(&ingredient).Updates(newData); updateResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "error", updateResult.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", newData, c)
}

func DeleteIngredient(c *gin.Context) {
	id := c.Param("id")

	var ingredient models.Ingredient

	if deleteResult := initializers.DB.Delete(&ingredient, id); deleteResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed delete data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "", nil, c)
}

func AddStock(c *gin.Context) {
	var stockInput dto.StockDTO

	if bindErr := utils.BindData(&stockInput, c); !bindErr {
		return
	}
	
	id := c.Param("id")

	var ingredient models.Ingredient
	if findRes := initializers.DB.First(&ingredient, id); findRes.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "error", findRes.Error.Error(), c)
		return
	}

	// update data
	newData := models.Ingredient{
		Stock:           ingredient.Stock + stockInput.Stock,
	}

	if updateResult := initializers.DB.Model(&ingredient).Updates(newData); updateResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "error", updateResult.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", newData, c)
}
