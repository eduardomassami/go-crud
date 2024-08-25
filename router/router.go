package router

import (
	"net/http"

	"github.com/eduardomassami/go-crud/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(employeeController *controller.EmployeeController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/test")
	// employeeRouter := router.Group("/test")
	router.GET("", employeeController.FindAll)
	router.GET("/:id", employeeController.FindById)
	router.POST("", employeeController.Create)
	router.PATCH("/:id", employeeController.Update)
	router.DELETE("/:id", employeeController.Delete)

	return service
}
