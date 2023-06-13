package handler

import (
	"shortener/internal/handler/mock"

	"testing"

	"github.com/golang/mock/gomock"
)

func TestHandler_getLong(t *testing.T) {
	t.Skip()
	ctrl := gomock.NewController(t)

	mockService := mock.NewMockservice(ctrl)
	mockService.EXPECT().GetLongURL(gomock.All()).Return()
}
