package routes

import "github.com/gin-gonic/gin"

func AllRoute(r *gin.Engine) {
	// Outlet Route
	outletRoutes := r.Group("/outlet")
	OutletRoute(outletRoutes)

	// Position Route
	positionRoutes := r.Group("/position")
	PositionRoute(positionRoutes)

	employeeRoutes := r.Group("/employee")
	EmployeeRoute(employeeRoutes)
}
