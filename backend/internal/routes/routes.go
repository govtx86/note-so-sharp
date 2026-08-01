package routes

import (
	"net/http"
	"os"
	"nss/internal/app"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *app.Application) http.Handler {
	r := gin.Default()
	frontendURL := os.Getenv("FRONTEND_URL")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	return r
}
