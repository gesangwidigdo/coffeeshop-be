package dto

type StockDTO struct {
	Stock           int    `json:"stock" binding:"required,numeric"`
}