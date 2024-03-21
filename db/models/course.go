package models

import "gorm.io/gorm"

const (
	CourseDurationWeek     = "week"
	CourseDurationMonth    = "month"
	CourseDurationHalfYear = "half_year"
	CourseDurationYear     = "year"
)

var (
	CourseDurations = map[string]struct{}{
		"week":      {},
		"month":     {},
		"half_year": {},
		"year":      {},
	}
)

type Course struct {
	gorm.Model
	TeacherID      int
	Teacher        *User
	Name           string
	Desc           string
	Price          float32
	Duration       string
	UsersPurchased []*User `json:"users_purchased" gorm:"many2many:purchased_courses;"`
}
