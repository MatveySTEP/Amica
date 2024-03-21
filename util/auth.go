package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"project/db"
	"project/db/models"
	"strings"
)

func ExtractUserFromRequest(c *gin.Context) (*models.User, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("хуйня")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return nil, errors.New("хуйня")
	}
	if headerParts[0] != "Bearer" {
		return nil, errors.New("хуйня")
	}

	claims, err := ParseToken(headerParts[1])
	if err != nil {
		return nil, errors.New("хуйня")
	}

	var user models.User
	err = db.DB.Where("id = ?", claims["iss"]).First(&user).Error
	if err != nil || user.Id == 0 {
		return nil, errors.New("хуйня")
	}

	return &user, nil
}
