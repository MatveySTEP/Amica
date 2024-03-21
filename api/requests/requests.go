package requests

type CreateCourseRequest struct {
	Name     string  `json:"name"`
	Desc     string  `json:"desc"`
	Price    float32 `json:"price"`
	Duration string  `json:"duration"`
}
