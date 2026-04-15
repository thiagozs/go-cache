package gocache

import (
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/thiagozs/go-cache/kind"
)

func TestMemoryOptions(t *testing.T) {
	opts, err := newMemoryParams()
	assert.NoError(t, err)

	opts.SetLogDebug(true)
	assert.Equal(t, true, opts.GetLogDebug())

	opts.SetLogDisable(true)
	assert.Equal(t, true, opts.GetLogDisable())

	opts.SetDriver(kind.GOCACHE)
	assert.Equal(t, kind.GOCACHE, opts.GetDriver())

	// Test setting and getting the logger option
	logger := zerolog.New(nil)
	opts.SetLogger(logger)
	assert.Equal(t, logger, opts.GetLog())

	opts.SetTExpiration(5 * time.Minute)
	assert.Equal(t, 5*time.Minute, opts.GetTExpiration())

	opts.SetTCleanUpInt(10 * time.Minute)
	assert.Equal(t, 10*time.Minute, opts.GetTCleanUpInt())

	c := cache.New(5*time.Minute, 10*time.Minute)
	opts.SetCache(c)
	assert.Equal(t, c, opts.GetCache())
}
