package user

import (
	userService "awesomeProject3/test-health/internal/service/user"
	"context"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ctx     context.Context
	service userService.Service
}

func NewHanlder(ctx context.Context, service userService.Service) *Handler {
	return &Handler{
		ctx:     ctx,
		service: service,
	}
}

//POST 		http://localhost:8080/user 			- создать пользователя
//GET 		http://localhost:8080/user 			- получить пользователей 	get
//GET 		http://localhost:8080/user(id) 		- получить пользователя  	get-by-id
//DELETE 	http://localhost:8080/user/delete 	- удалить пользователя		delete
//PATCH 	http://localhost:8080/user/update 	- обновить пользователя 	update

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	user := router.Group("/")
	{
		user.POST("/", h.create)
		user.GET("/", h.get)
		user.GET("/id", h.getId)
		user.DELETE("/", h.delete)
		user.PATCH("/", h.patch)
	}

	return router
}
