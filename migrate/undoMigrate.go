package main

import (
	"github.com/gesangwidigdo/coffeeshop-be/initializers"
	"github.com/gesangwidigdo/coffeeshop-be/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.Migrator().DropTable(&models.Outlet{})
}
