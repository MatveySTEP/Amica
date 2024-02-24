package handles

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"project/models"
)

func Register(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Password: password,
	}
	c.JSON(200, user)
}
