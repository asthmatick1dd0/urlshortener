package handler

import (
	"github.com/asthmatick1dd0/urlshortener/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.URLService
}

func NewHandler(service *service.URLService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/", h.index) // добавлено
	router.POST("/shorten", h.ShortenURL)
	router.GET("/:code", h.redirect)
}

func (h *Handler) ShortenURL(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	shortened, err := h.service.Shorten(req.URL)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not shorten URL"})
		return
	}
	c.JSON(200, gin.H{"shortened": shortened})
}

func (h *Handler) redirect(c *gin.Context) {
	// Implementation for redirecting to the original URL
	code := c.Param("code")
	originalURL, err := h.service.GetOriginalURL(code)
	if err != nil {
		c.JSON(404, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(302, originalURL)
}

func (h *Handler) index(c *gin.Context) {
	c.File("web/static/index.html")
}
