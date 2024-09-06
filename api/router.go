package api

import (
	"api_gateway/api/handler"
	"api_gateway/api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title         Task Control API
// @version         1.0
// @description      Task Control API Documentation
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(h *handler.Handler, enforcer *casbin.Enforcer) *gin.Engine {
	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		RequestHeaders: "Authorization, Origin, Content-Type",
		Methods:        "POST, GET, PUT, DELETE, OPTION",
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.JWTMiddleware())
	r.Use(middleware.CasbinMiddleware(enforcer))

	user := r.Group("/user_service/")
	{
		user.GET("user/:id", h.GetUser)
		user.PUT("user/:id", h.UpdateUser)
		user.GET("users", h.GetAllUsers)
		user.DELETE("user/:id", h.DeleteUser)

	}

	task := r.Group("/task_service/")
	{
		task.POST("tasks", h.CreateTask)
		task.PUT("tasks/:id", h.UpdateTask)
		task.GET("tasks/:id", h.GetTask)
		task.GET("tasks", h.GetAllTasks)
		task.DELETE("tasks/:id", h.DeleteTask)
		task.POST("details", h.CreateDetail)
		task.PUT("details/:id", h.UpdateDetail)
		task.GET("details/:id", h.GetDetail)
		task.GET("details", h.GetAllDetail)
		task.DELETE("details/:id", h.DeleteDetail)
	}

	return r
}
