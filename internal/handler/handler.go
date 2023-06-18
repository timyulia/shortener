//go:generate go run github.com/golang/mock/mockgen -destination=./mock/handler_gen.go -source=handler.go -package=mock service

package handler

import (
	"context"
	"github.com/gin-gonic/gin"
)

type service interface {
	GetShortURL(ctx context.Context, long string) (string, error)
	GetLongURL(ctx context.Context, short string) (string, error)
}

type Handler struct {
	services service
}

func New(services service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	{
		router.POST("/", h.getShort)
		router.GET("/:url", h.getLong)
	}
	return router
}
