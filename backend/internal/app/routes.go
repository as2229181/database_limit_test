package app

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func (app * App)registerRoutes(router *gin.Engine) {
	router.GET("api/v1/health", app.handleHealthCheck) 
}

func (app * App)handleHealthCheck(c *gin.Context) {
	dbStatus := "ok"
	if err := app.pool.Ping(c.Request.Context()); err != nil{
		dbStatus = "error"
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"db": dbStatus,
	})
}