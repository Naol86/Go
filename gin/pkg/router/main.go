package router

import (
	"myproject/pkg/database"
	"myproject/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the Gin router with routes and handlers
func SetupRouter() *gin.Engine {
    database.ConnectDB()
    r := gin.Default()
    
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
    })

    r.GET("/users", models.GetUsers)

    return r
}
