package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kullaniciadi/url-shortener/config"
	"github.com/kullaniciadi/url-shortener/internal/handlers"
	"github.com/kullaniciadi/url-shortener/internal/models"
	"github.com/kullaniciadi/url-shortener/internal/services"
)

func main() {
	db := config.ConnectDB()
	rds := config.ConnectRedis()

	db.AutoMigrate(
		&models.URL{},
	)
	r := gin.Default()
	protected := r.Group("/")

	urlService := services.NewUrlService(db, rds)

	protected.POST("/shorten", func(c *gin.Context) {
		handlers.ShortenURL(urlService, c)
	})

	protected.GET("/:code", func(c *gin.Context) {
		handlers.GetURL(urlService, c)
	})
	protected.GET("/:code/stats", func(c *gin.Context) {
		handlers.GetStats(urlService, c)
	})
	r.Run(":" + os.Getenv("PORT"))
}
