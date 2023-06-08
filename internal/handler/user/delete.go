package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) delete(c *gin.Context) {
	if err := h.service.Delete(h.ctx); err != nil {
		c.JSON(500, fmt.Sprintf("deletion user error: %s", err.Error()))
	}
	c.JSON(200, "user deleted")
}
