package service

import (
	"testing"

	"github.com/c2fo/testify/assert"
)

func TestGenerateShortURL(t *testing.T) {
	tests := []struct {
		name  string
		long  string
		short string
	}{
		{
			"with forbidden",
			"https://www.youtube.com/",
			"ELZgNmWcOJ",
		},
		{
			"ok",
			"https://www.google.com",
			"CL_rVxjFkR",
		},
	}
	for _, test := range tests {
		short := generateShortURL(test.long)
		assert.Equal(t, test.short, short)
	}
}
