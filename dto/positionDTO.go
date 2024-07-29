package dto

type PositionDTO struct {
	Position_name string `json:"position_name" binding:"required"`
}