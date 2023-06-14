package handler

import (
	"testing"

	"shortener/internal/handler/mock"

	"github.com/c2fo/testify/assert"
	"github.com/golang/mock/gomock"
)

func TestHandler_InitRoutes(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		serviceMock := mock.NewMockservice(ctrl)
		h := New(serviceMock)
		ginEngine := h.InitRoutes()

		assert.NotNil(t, h)
		assert.NotNil(t, ginEngine)
	})
}
