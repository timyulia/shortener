package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	mock_repo "shortener/internal/repository/mocks"

	"testing"

	"github.com/c2fo/testify/assert"
)

func TestService_GetShortURL(t *testing.T) {
	type mockBehavior func(r *mock_repo.MockRepository, short, long string)
	tests := []struct {
		name         string
		short        string
		res          string
		long         string
		mockBehavior mockBehavior
		err          string
	}{
		{
			"ok",
			"CL_rVxjFkR",
			"CL_rVxjFkR",
			"https://www.google.com",
			func(r *mock_repo.MockRepository, short, long string) {
				r.EXPECT().GetLongURL(short).Return(long, nil)
			},
			"",
		},
		{
			"error",
			"CL_rVxjFkR",
			"",
			"https://www.google.com",
			func(r *mock_repo.MockRepository, short, long string) {
				r.EXPECT().GetLongURL(short).Return("", errors.New("something went wrong"))
			},
			"something went wrong",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repo.NewMockRepository(c)
			test.mockBehavior(repo, test.short, test.long)
			serv := New(repo)
			short, err := serv.GetShortURL(test.long)
			if test.err == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.err)
			}
			assert.Equal(t, test.res, short)
		})
	}
}

//func TestService_GetLongURL(t *testing.T) {
//	type mockBehavior func(r *mock_repo.MockRepository, short, long string)
//	tests := []struct {
//		name         string
//		short        string
//		long         string
//		mockBehavior mockBehavior
//	}{
//		{
//			"ok",
//			"CL_rVxjFkR",
//			"https://www.google.com",
//			func(r *mock_repo.MockRepository, short, long string) {
//				r.EXPECT().GetLongURL(short).Return(long, nil)
//			},
//		},
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			c := gomock.NewController(t)
//			defer c.Finish()
//
//			repo := mock_repo.NewMockRepository(c)
//			test.mockBehavior(repo, test.short, test.long)
//			serv := New(repo)
//			short, err := serv.GetShortURL(test.long)
//			assert.NoError(t, err)
//			assert.Equal(t, test.short, short)
//		})
//	}
//}
