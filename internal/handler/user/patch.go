package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) patch(c *gin.Context) {
	if err := h.service.Patch(h.ctx); err != nil {
		c.JSON(500, fmt.Sprintf("updating user error: %s", err.Error()))
	}
	c.JSON(200, "user updated")
}
