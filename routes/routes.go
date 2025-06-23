package routes

import (
	"hospitalapp/controllers"
	"hospitalapp/middleware"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	api := router.Group("/api")
	{
		api.GET("/patients", controllers.GetPatients)
		api.GET("/patients/:id", controllers.GetpatientById)

		api.POST("/patients",middleware.JWTAuthMiddleware("receptionist"), controllers.CreatePatient)
		api.PUT("/patients/:id",middleware.JWTAuthMiddleware("receptionist", "doctor"), controllers.UpdatePatient)
		api.DELETE("/patients/:id",middleware.JWTAuthMiddleware("receptionist"), controllers.DeletePatient)
	}
}