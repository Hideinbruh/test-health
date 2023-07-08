package user

import (
	"awesomeProject3/internal/converter"
	userModel "awesomeProject3/internal/model/user"
	"awesomeProject3/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) create(c *gin.Context) {
	var user userModel.HandlerUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(500, "invalid request")
		return
	}

	userId, err := h.service.Create(converter.ToCreateService(user))
	if err != nil {
		logger.Logger(err)
		c.JSON(400, "user not created")
	}
	c.JSON(200, fmt.Sprintf("user %d created", userId))
}
