package requests

type CreateCourseRequest struct {
	Name     string  `json:"name"`
	Desc     string  `json:"desc"`
	Price    float32 `json:"price"`
	Duration string  `json:"duration"`
}

type CreateFeedbackRequest struct {
	Rating  int    `json:"rating"`
	Message string `json:"message"`
}
