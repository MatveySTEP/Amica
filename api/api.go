package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"project/handles"
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
	middleware := func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
	a.r.Use(middleware)
	apiV1.Use(middleware)
	apiV1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	apiV1.POST("/api/register", handles.Register)
	apiV1.GET("/api/user", handles.User)
	apiV1.POST("/api/login", handles.Login)
	apiV1.POST("/api/logout", handles.Logout)
}
func (a *Api) Run() {
	if err := a.r.Run(a.config.Addr); err != nil {
		log.Fatal().Err(err).Msg("Сервер не запущен")
	}
}
