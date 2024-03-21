package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	TeacherID int
	Teacher   *User
	Name      string
	Desc      string
	Price     float32
	Duration  string
}
