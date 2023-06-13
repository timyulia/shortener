package cache

import (
	"testing"

	"github.com/c2fo/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s := New()
		assert.NotNil(t, s)
	})
}

func TestInMemory_AddLinksPair(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		cache := New()
		err := cache.AddLinksPair("short", "long")

		assert.NoError(t, err)
		assert.Equal(t, "long", cache.pair["short"])
	})
}

func TestInMemory_GetLongURL(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		cache := New()
		_ = cache.AddLinksPair("short", "long")
		value, err := cache.GetLongURL("short")

		assert.NoError(t, err)
		assert.Equal(t, "long", value)
	})
}
