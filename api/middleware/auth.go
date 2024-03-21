package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/util"
)

func AuthMiddleware(c *gin.Context) {
	_, err := util.ExtractUserFromRequest(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
