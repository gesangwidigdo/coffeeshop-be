package controllers

import (
	"net/http"

	"github.com/gesangwidigdo/coffeeshop-be/dto"
	"github.com/gesangwidigdo/coffeeshop-be/initializers"
	"github.com/gesangwidigdo/coffeeshop-be/models"
	"github.com/gesangwidigdo/coffeeshop-be/utils"
	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var employeeInput dto.EmployeeDTO

	// bind data
	if bindErr := utils.BindData(&employeeInput, c); !bindErr {
		return
	}

	var position models.Position
	if err := initializers.DB.First(&position, employeeInput.Position_id).Error; err != nil {
		utils.ReturnResponse(http.StatusBadRequest, "invalid position id", "error", err.Error(), c)
		return
	}

	var outlet models.Outlet
	if err := initializers.DB.First(&outlet, employeeInput.Outlet_id).Error; err != nil {
		utils.ReturnResponse(http.StatusBadRequest, "invalid outlet id", "error", err.Error(), c)
		return
	}

	// create employee
	employee := models.Employee{
		Employee_name:    employeeInput.Employee_name,
		Telephone_number: employeeInput.Telephone_number,
		Username:         employeeInput.Username,
		Password:         employeeInput.Password,
		Address:          employeeInput.Address,
		Gender:           employeeInput.Gender,
		Position_id:      employeeInput.Position_id,
		Outlet_id:        employeeInput.Outlet_id,
	}

	result := initializers.DB.Create(&employee)

	if result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to create data", "", nil, c)
		return
	}

	var createdEmployee models.Employee
	if err := initializers.DB.Preload("Position").Preload("Outlet").First(&createdEmployee, employee.ID).Error; err != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "error", err.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", createdEmployee, c)
}

func GetAllEmployee(c *gin.Context) {
	var employees []models.Employee
	result := initializers.DB.Model(&models.Employee{}).
		Preload("Position").
		Preload("Outlet").
		Find(&employees)

	if result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "Failed to get data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", employees, c)
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	var employees []models.Employee
	result := initializers.DB.Model(&models.Employee{}).
		Preload("Position").
		Preload("Outlet").
		First(&employees, id)

	if result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "Failed to get data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", employees, c)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employeeInput dto.EmployeeDTO

	if bindErr := utils.BindData(&employeeInput, c); !bindErr {
		return
	}

	var employee models.Employee
	findResult := initializers.DB.First(&employee, id)

	if findResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "", nil, c)
		return
	}

	// Update data
	newData := models.Employee{
		Employee_name:    employeeInput.Employee_name,
		Telephone_number: employee.Telephone_number,
		Username:         employeeInput.Username,
		Password:         employeeInput.Password,
		Address:          employeeInput.Address,
		Gender:           employeeInput.Gender,
		Position_id:      employeeInput.Position_id,
		Outlet_id:        employeeInput.Outlet_id,
	}

	updateResult := initializers.DB.Model(&employee).Updates(newData)

	if updateResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update data", "", nil, c)
		return
	}

	var updatedEmployee models.Employee
	updatedEmployeeRes := initializers.DB.Preload("Position").Preload("Outlet").First(&updatedEmployee, id)

	if updatedEmployeeRes.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", updatedEmployee, c)
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee

	deleteResult := initializers.DB.Delete(&employee, id)

	if deleteResult.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed delete data", "", nil, c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "", nil, c)
}
