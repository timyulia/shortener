package service

import (
	"testing"

	"github.com/c2fo/testify/assert"
)

func TestGenerateShortURL(t *testing.T) {
	t.Run("with forbidden", func(t *testing.T) {
		t.Parallel()
		short := generateShortURL("https://www.youtube.com/")
		assert.Equal(t, "ELZgNmWcOJ", short)
	})
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		short := generateShortURL("https://www.google.com")
		assert.Equal(t, "CL_rVxjFkR", short)
	})
}
