package middleware

import (
	"amica/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	_, err := util.ExtractUserFromRequest(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
