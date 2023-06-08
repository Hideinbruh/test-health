package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) get(c *gin.Context) {
	if err := h.service.Get(h.ctx); err != nil {
		c.JSON(500, fmt.Sprintf("get user error: %s", err.Error()))
	}
	c.JSON(200, "here is the list of users")
}
