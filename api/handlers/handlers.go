package handlers

import (
	"amica/db"
	"amica/db/models"
	"amica/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

const SecretKey = "secret"

func Register(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var userExist int
	err := db.DB.Raw("SELECT COUNT(id) from users where name=?", data["name"]).Scan(&userExist).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if userExist > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "User already exist",
		})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Password: password,
		Role:     models.RoleClient,
	}
	db.DB.Create(&user)
	c.JSON(200, user)
}

func Login(c *gin.Context) {
	var data map[string]string
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var user models.User
	db.DB.Where("name = ?", data["name"]).First(&user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect password"})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "token": token, "user": user})
}

func User(c *gin.Context) {
	user, err := util.ExtractUserFromRequest(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Logout(c *gin.Context) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
