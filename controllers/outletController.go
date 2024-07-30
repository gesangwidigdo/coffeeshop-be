package controllers

import (
	"net/http"

	"github.com/gesangwidigdo/coffeeshop-be/dto"
	"github.com/gesangwidigdo/coffeeshop-be/initializers"
	"github.com/gesangwidigdo/coffeeshop-be/models"
	"github.com/gesangwidigdo/coffeeshop-be/utils"
	"github.com/gin-gonic/gin"
)

func CreateOutlet(c *gin.Context) {
	var outletInput dto.OutletDTO

	// Get Data from req body
	if bindErr := utils.BindData(&outletInput, c); !bindErr {
		return
	}

	// Create outlet
	outlet := models.Outlet{
		Outlet_name: outletInput.Outlet_name,
		Address:     outletInput.Address,
		City:        outletInput.City,
	}

	result := initializers.DB.Omit("Employees").Create(&outlet)

	if result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "Failed to create data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusCreated, "ok", "data", outlet, c)
}

func GetAllOutlet(c *gin.Context) {
	var outlet []models.Outlet
	result := initializers.DB.Find(&outlet)

	if result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "Failed to get data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", outlet, c)
}

func GetOutletByID(c *gin.Context) {
	// Get ID
	id := c.Param("id")

	// Get outlet by id
	var outlet models.Outlet
	result := initializers.DB.First(&outlet, id)

	if result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "Failed to get data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", outlet, c)
}

func UpdateOutlet(c *gin.Context) {
	var outletInput dto.OutletDTO

	// Get outlet id
	id := c.Param("id")

	// get data from req body
	if bindErr := utils.BindData(&outletInput, c); !bindErr {
		return
	}

	var outlet models.Outlet
	findResult := initializers.DB.First(&outlet, id)

	if findResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "Failed to Find Data", "", nil, c)
		return
	}

	// Update data
	newData := models.Outlet{
		Outlet_name: outletInput.Outlet_name,
		Address:     outletInput.Address,
		City:        outlet.City,
	}

	updateResult := initializers.DB.Model(&outlet).Updates(newData)

	if updateResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "Failed to update data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", newData, c)
}

func DeleteOutlet(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Find outlet by id
	var outlet models.Outlet

	deleteResult := initializers.DB.Delete(&outlet, id)

	if deleteResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed delete data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "", nil, c)
}
