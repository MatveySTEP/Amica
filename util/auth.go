package util

import (
	"amica/db"
	"amica/db/models"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func ExtractUserFromRequest(c *gin.Context) (*models.User, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("ошибка получения пользователя")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return nil, errors.New("ошибка получения пользователя")
	}
	if headerParts[0] != "Bearer" {
		return nil, errors.New("ошибка получения пользователя")
	}

	claims, err := ParseToken(headerParts[1])
	if err != nil {
		return nil, errors.New("ошибка получения пользователя")
	}

	var user models.User
	err = db.DB.Where("id = ?", claims["iss"]).First(&user).Error
	if err != nil || user.Id == 0 {
		return nil, errors.New("ошибка получения пользователя")
	}

	return &user, nil
}
