package controllers

import (
	"net/http"

	"github.com/gesangwidigdo/coffeeshop-be/dto"
	"github.com/gesangwidigdo/coffeeshop-be/initializers"
	"github.com/gesangwidigdo/coffeeshop-be/models"
	"github.com/gesangwidigdo/coffeeshop-be/utils"
	"github.com/gin-gonic/gin"
)

func GetAllPosition(c *gin.Context) {
	var position []models.Position
	result := initializers.DB.Find(&position)

	if result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "Failed to get data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", position, c)
}

func GetPositionByID(c *gin.Context) {
	// Get id from param
	id := c.Param("id")

	// get outlet by id
	var position models.Position
	result := initializers.DB.First(&position, id)

	// check error
	if result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", position, c)
}

func CreatePosition(c *gin.Context) {
	var positionInput dto.PositionDTO

	// bind req body
	if bindErr := utils.BindData(&positionInput, c); !bindErr {
		return
	}

	// create position
	position := models.Position{
		Position_name: positionInput.Position_name,
	}

	result := initializers.DB.Omit("Employees").Create(&position)

	if result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to create data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusCreated, "ok", "data", position, c)
}

func UpdatePosition(c *gin.Context) {
	// get id from param
	id := c.Param("id")
	
	// find position by id
	var position models.Position
	findResult := initializers.DB.First(&position, id)

	// check error
	if findResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to find data", "", nil, c)
		return
	}

	// bind data
	var positionInput dto.PositionDTO
	if bindErr := utils.BindData(&positionInput, c); !bindErr {
		return
	}

	// create new position
	newPosition := models.Position{
		Position_name: positionInput.Position_name,
	}

	updateResult := initializers.DB.Model(&position).Updates(&newPosition)

	if updateResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", newPosition, c)
}

func DeletePosition(c *gin.Context) {
	// get id from param
	id := c.Param("id")

	// find data by id
	var position models.Position
	findResult := initializers.DB.First(&position, id)

	if findResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "", nil, c)
		return
	}

	// delete data
	deleteResult := initializers.DB.Delete(&position, id)

	if deleteResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to delete data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "", nil, c)
}