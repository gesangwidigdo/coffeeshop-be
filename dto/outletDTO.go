package dto

type OutletDTO struct {
	Outlet_name string `json:"outlet_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	City        string `json:"city" binding:"required"`
}
