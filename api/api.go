package api

import (
	"amica/api/handlers"
	"amica/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Api struct {
	config Config
	r      *gin.Engine
}
type Config struct {
	Addr string
}

func NewApi(c Config) *Api {
	r := gin.Default()
	s := &Api{
		config: c,
		r:      r,
	}
	s.UseRoutes()
	return s
}

func (a *Api) UseRoutes() {
	apiV1 := a.r.Group("/v1")
	cors := func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
	a.r.Use(cors)
	apiV1.Use(cors)

	apiV1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	apiV1.POST("/api/register", handlers.Register)
	apiV1.POST("/api/login", handlers.Login)
	apiV1.POST("/api/logout", handlers.Logout)

	apiV1.GET("/api/user", middleware.AuthMiddleware, handlers.User)

	// Запросы для учителя
	apiV1.GET("/api/courses", middleware.AuthMiddleware, handlers.ListCourses)
	apiV1.POST("/api/courses/create", middleware.AuthMiddleware, handlers.CreateCourse)
	apiV1.GET("/api/courses/:course", middleware.AuthMiddleware, handlers.GetCourse)
	apiV1.DELETE("/api/courses/:course", middleware.AuthMiddleware, handlers.DeleteCourse)
	apiV1.POST("/api/courses/buy/:course", middleware.AuthMiddleware, handlers.BuyCourse)

	// Запросы для отзывов
	apiV1.POST("/api/feedback/create/:course", middleware.AuthMiddleware, handlers.CreateFeedback)
	apiV1.GET("/api/feedback/:course", middleware.AuthMiddleware, handlers.ListFeedback)
	apiV1.DELETE("/api/feedback/delete/:feedback", middleware.AuthMiddleware, handlers.DeleteFeedback)
}
func (a *Api) Run() {
	if err := a.r.Run(a.config.Addr); err != nil {
		log.Fatal().Err(err).Msg("Сервер не запущен")
	}
}
