package handlers

import (
	"amica/api/requests"
	"amica/db"
	"amica/db/models"
	"amica/util"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func ListCourses(c *gin.Context) {
	user, err := util.ExtractUserFromRequest(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = db.DB.Preload("Courses").
		Preload("PurchasedCourses").
		Where("id = ?", user.Id).
		First(&user).Error
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var courses []*models.Course

	switch user.Role {
	case models.RoleClient:
		courses = user.PurchasedCourses
	case models.RoleTeacher:
		courses = user.Courses
	}

	c.JSON(200, courses)
}

func GetCourse(c *gin.Context) {
	courseID := c.Param("course")
	if courseID == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var course *models.Course
	err := db.DB.Where("id = ?", courseID).First(&course).Error
	if err != nil || course == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(200, course)
}

func CreateCourse(c *gin.Context) {
	var r requests.CreateCourseRequest
	err := c.BindJSON(&r)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if _, ok := models.CourseDurations[r.Duration]; !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := util.ExtractUserFromRequest(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if user.Role != models.RoleTeacher {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	course := &models.Course{
		TeacherID: int(user.Id),
		Name:      r.Name,
		Desc:      r.Desc,
		Price:     r.Price,
		Duration:  r.Duration,
	}
	err = db.DB.Create(course).Error
	if err != nil {
		log.Error().Err(err).Msg("error creating course")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, course)
}

func DeleteCourse(c *gin.Context) {
	courseID := c.Param("course")
	if courseID == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var course *models.Course
	err := db.DB.Where("id = ?", courseID).First(&course).Error
	if err != nil || course == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	user, err := util.ExtractUserFromRequest(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if course.TeacherID != int(user.Id) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	db.DB.Delete(course)
}

// запросы для ученика

func BuyCourse(c *gin.Context) {
	user, err := util.ExtractUserFromRequest(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	courseID := c.Param("course")
	if courseID == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var course *models.Course
	err = db.DB.Where("id = ?", courseID).First(&course).Error
	if err != nil || course == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	err = db.DB.Model(user).Association("PurchasedCourses").Append(course)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
