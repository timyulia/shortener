package handler

import (
	"bytes"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	mock_serv "shortener/internal/handler/mock"

	"github.com/c2fo/testify/assert"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestHandler_getShort(t *testing.T) {
	type mockBehavior func(s *mock_serv.Mockservice, short, long string)

	tests := []struct {
		name                 string
		inputBody            string
		short                string
		long                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"URL": "https://www.google.com"}`,
			short:     "CL_rVxjFkR",
			long:      "https://www.google.com",
			mockBehavior: func(r *mock_serv.Mockservice, short, long string) {
				r.EXPECT().GetShortURL(long).Return(short, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"shortURL":"CL_rVxjFkR"}`,
		},
		{
			name:                 "bad request",
			inputBody:            `{"link"}`,
			short:                "",
			long:                 "",
			mockBehavior:         func(r *mock_serv.Mockservice, short, long string) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid character '}' after object key"}`,
		},
		{
			name:                 "not url",
			inputBody:            `{"URL": "not url"}`,
			short:                "",
			long:                 "",
			mockBehavior:         func(r *mock_serv.Mockservice, short, long string) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"this link is not a valid URL"}`,
		},
		{
			name:      "service error",
			inputBody: `{"URL": "https://www.google.com"}`,
			short:     "CL_rVxjFkR",
			long:      "https://www.google.com",
			mockBehavior: func(r *mock_serv.Mockservice, short, long string) {
				r.EXPECT().GetShortURL(long).Return("", errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)

			serv := mock_serv.NewMockservice(c)

			test.mockBehavior(serv, test.short, test.long)

			handler := Handler{serv}

			r := gin.New()
			r.POST("/", handler.getShort)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_getLong(t *testing.T) {
	type mockBehavior func(s *mock_serv.Mockservice, short, long string)

	tests := []struct {
		name                 string
		short                string
		long                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:  "ok",
			short: "CL_rVxjFkR",
			long:  "https://www.google.com",
			mockBehavior: func(r *mock_serv.Mockservice, short, long string) {
				r.EXPECT().GetLongURL(short).Return(long, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"longURL":"https://www.google.com"}`,
		},
		{
			name:  "no such short link",
			short: "noLinklink",
			long:  "",
			mockBehavior: func(r *mock_serv.Mockservice, short, long string) {
				r.EXPECT().GetLongURL(short).Return("", nil)
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"there is no such short link yet"}`,
		},
		{
			name:  "too short link",
			short: "tooShort",
			long:  "",
			mockBehavior: func(r *mock_serv.Mockservice, short, long string) {
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"the short link must consist of 10 characters"}`,
		},
		{
			name:  "service error",
			short: "CL_rVxjFkR",
			long:  "",
			mockBehavior: func(r *mock_serv.Mockservice, short, long string) {
				r.EXPECT().GetLongURL(short).Return("", errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)

			serv := mock_serv.NewMockservice(c)

			test.mockBehavior(serv, test.short, test.long)

			handler := Handler{serv}

			r := gin.New()
			r.GET("/:url", handler.getLong)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/%s", test.short),
				bytes.NewBufferString(""))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
