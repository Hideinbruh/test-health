package user

import (
	"fmt"

	"github.com/Hideinbruh/test-health/internal/converter"
	userModel "github.com/Hideinbruh/test-health/internal/model/user"

	"github.com/gin-gonic/gin"
)

func (h *Handler) create(c *gin.Context) {
	var user userModel.UserRequest
	if err := c.BindJSON(&user); err != nil {
		c.JSON(500, "invalid request")
		return
	}

	userId, err := h.service.Create(converter.ToCreateService(user))
	if err != nil {
		c.JSON(400, "user not created")
	}

	c.JSON(200, fmt.Sprintf("user %d created", userId))
}
