package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	StudentID int
	CourseID  int
	Student   *User
	Course    *Course
	Rating    int
	Message   string
}
