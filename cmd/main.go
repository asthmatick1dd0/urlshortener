package main

import (
	"github.com/asthmatick1dd0/urlshortener/internal/handler"
	"github.com/asthmatick1dd0/urlshortener/internal/repository"
	"github.com/asthmatick1dd0/urlshortener/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	URLRepository := repository.NewInMemoryURLRepository()
	urlService := service.NewURLService(URLRepository)
	handler := handler.NewHandler(urlService)
	handler.RegisterRoutes(router)
	router.Run(":5000")
}
