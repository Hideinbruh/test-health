package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getId(c *gin.Context) {
	if err := h.service.GetId(h.ctx); err != nil {
		c.JSON(500, fmt.Sprintf("getting user error: %s", err.Error()))
	}
	c.JSON(200, "getting user")
}
