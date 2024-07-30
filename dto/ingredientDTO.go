package dto

type IngredientDTO struct {
	Ingredient_name string `json:"ingredient_name" binding:"required,alpha"`
	Stock           int    `json:"stock" binding:"numeric"`
}
