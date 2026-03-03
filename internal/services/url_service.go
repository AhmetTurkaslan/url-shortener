package services

import (
	"context"
	"errors"
	"time"

	"github.com/kullaniciadi/url-shortener/internal/models"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UrlService struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUrlService(db *gorm.DB, rdb *redis.Client) *UrlService {
	return &UrlService{db: db, rdb: rdb}
}

func (s *UrlService) ShortenURL(longURL string) error {

	code, err := gonanoid.New(8)
	if err != nil {
		return errors.New("kod üretilemedi")
	}

	url := models.URL{
		LongURL: longURL,
		Code:    code,
	}
	if err := s.db.Create(&url).Error; err != nil {
		return errors.New("kısa link oluşturulamadı")
	}
	return nil
}

func (s *UrlService) GetURL(code string) (string, error) {
	ctx := context.Background()
	result, err := s.rdb.Get(ctx, code).Result()
	if err == nil {
		s.db.Model(&models.URL{}).Where("code = ?", code).Update("clicks", gorm.Expr("clicks + 1"))
		return result, nil
	}
	var url models.URL
	if err := s.db.Where("code=?", code).First(&url).Error; err != nil {
		return "", errors.New("Link bulunamadı")
	}
	s.rdb.Set(ctx, code, url.LongURL, 24*time.Hour)
	s.db.Model(&models.URL{}).Where("code = ?", code).Update("clicks", gorm.Expr("clicks + 1"))
	return url.LongURL, nil
}

func (s *UrlService) GetStats(code string) (models.URL, error) {

	var url models.URL
	if err := s.db.Where("code = ?", code).First(&url).Error; err != nil {
		return models.URL{}, errors.New("Link veritabanında bulunamadı")
	}
	return url, nil
}
