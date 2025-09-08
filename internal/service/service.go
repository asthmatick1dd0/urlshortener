package service

import (
	"fmt"
	"hash/crc32"

	"github.com/asthmatick1dd0/urlshortener/internal/repository"
)

type URLService struct {
	repo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) *URLService {
	return &URLService{
		repo: repo,
	}
}

func (s *URLService) Shorten(url string) (string, error) {
	hash := crc32.ChecksumIEEE([]byte(url))
	shortened := fmt.Sprintf("%d", hash)
	if err := s.repo.Save(shortened, url); err != nil {
		return "", err
	}
	return shortened, nil
}

func (s *URLService) GetOriginalURL(code string) (string, error) {
	return s.repo.Get(code)
}
