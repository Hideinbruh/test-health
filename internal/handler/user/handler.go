package user

import (
	userService "awesomeProject3/internal/service/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service userService.Service
}

func NewHandler(service userService.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	user := router.Group("/user")
	{
		user.POST("/create", h.create)
		user.GET("/get", h.get)
	}
	return router
}
