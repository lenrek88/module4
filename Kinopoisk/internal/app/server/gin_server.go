package server

import "github.com/gin-gonic/gin"

func NewGin() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/v1/films", func(c *gin.Context) { c.JSON(200, []any{}) })
	r.GET("/v1/films/:id", func(c *gin.Context) { c.JSON(200, "film") })
	r.GET("/health/check", func(c *gin.Context) { c.String(200, "ok") })
	return r
}
