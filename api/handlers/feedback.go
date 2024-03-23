package handlers

import (
	"amica/api/requests"
	"amica/db"
	"amica/db/models"
	"amica/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateFeedback(c *gin.Context) {
	var r requests.CreateFeedbackRequest
	err := c.BindJSON(&r)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := util.ExtractUserFromRequest(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	courseID := c.Param("course")
	if courseID == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	courseId, err := strconv.Atoi(courseID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	//нужно еще написать проверку, что отзыв один
	feedback := &models.Feedback{
		Model:     gorm.Model{},
		StudentID: int(user.Id),
		CourseID:  courseId,
		Student:   nil,
		Course:    nil,
		Rating:    r.Rating,
		Message:   r.Message,
	}

	err = db.DB.Create(feedback).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, feedback)
}

func ListFeedback(c *gin.Context) {
	courseID := c.Param("course")
	if courseID == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	courseId, err := strconv.Atoi(courseID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	var feedbackList []models.Feedback
	db.DB.Where("course_id=?", courseId).Find(&feedbackList)
	c.JSON(http.StatusOK, feedbackList)
}

func DeleteFeedback(c *gin.Context) {
	feedbackID := c.Param("feedback")
	if feedbackID == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	feedbackId, err := strconv.Atoi(feedbackID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
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
	var feedback *models.Feedback
	//тут отъебнуло
	err = db.DB.Where("id = ?", feedbackId).First(&feedback).Error
	if err != nil || feedback == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	db.DB.Delete(feedback)

}
