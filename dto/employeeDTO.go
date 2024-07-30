package dto

type EmployeeDTO struct {
	Employee_name string `json:"employee_name" binding:"required,alpha"`
	Telephone_number string `json:"telephone_number" binding:"required,numeric,min=10"`
	Username string `json:"username" binding:"required,alpha"`
	Password string `json:"password" binding:"required,alpha"`
	Address string `json:"address" binding:"required"`
	Gender string `json:"gender" binding:"required,alpha"`
	Position_id uint `json:"position_id" binding:"required,numeric"`
	Outlet_id uint `json:"outlet_id" binding:"required,numeric"`
}