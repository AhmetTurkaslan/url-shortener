package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kullaniciadi/url-shortener/internal/services"
)

type URLInput struct {
	LongURL string `json:"long_url" binding:"required"`
}

func ShortenURL(urlService *services.UrlService, c *gin.Context) {
	var inputURL URLInput
	if err := c.ShouldBindJSON(&inputURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := urlService.ShortenURL(inputURL.LongURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "kısa link oluştu"})
}

func GetURL(urlService *services.UrlService, c *gin.Context) {
	code := c.Param("code")
	longURL, err := urlService.GetURL(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, longURL)
}

func GetStats(urlService *services.UrlService, c *gin.Context) {
	code := c.Param("code")
	url, err := urlService.GetStats(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         url.ID,
		"code":       url.Code,
		"long_url":   url.LongURL,
		"clicks":     url.Clicks,
		"created_at": url.CreatedAt,
	})

}
