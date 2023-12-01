package cache

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/thiagozs/go-cache/kind"
)

func TestCacheOptions(t *testing.T) {
	opts, err := newCacheParams()
	assert.NoError(t, err)

	opts.SetFile("test.db")
	assert.Equal(t, "test.db", opts.GetFile())

	opts.SetFolder("/tmp")
	assert.Equal(t, "/tmp", opts.GetFolder())

	opts.SetPath("/tmp/test.db")
	assert.Equal(t, "/tmp/test.db", opts.GetPath())

	opts.SetTTL(60)
	assert.Equal(t, 60, opts.GetTTL())

	opts.SetLogDebug(true)
	assert.Equal(t, true, opts.GetLogDebug())

	opts.SetLogDisable(true)
	assert.Equal(t, true, opts.GetLogDisable())

	opts.SetHost("localhost")
	assert.Equal(t, "localhost", opts.GetHost())

	opts.SetPassword("password")
	assert.Equal(t, "password", opts.GetPassword())

	opts.SetUser("user")
	assert.Equal(t, "user", opts.GetUser())

	opts.SetPort(6379)
	assert.Equal(t, 6379, opts.GetPort())

	logger := zerolog.New(nil)
	opts.SetLogger(logger)
	assert.Equal(t, logger, opts.GetLog())

	opts.SetTExpiration(time.Minute)
	assert.Equal(t, time.Minute, opts.GetTExpiration())

	opts.SetTCleanUpInt(time.Hour)
	assert.Equal(t, time.Hour, opts.GetTCleanUpInt())

	opts.SetDatabase("db")
	assert.Equal(t, "db", opts.GetDatabase())

	opts.SetVersion("1.0.0")
	assert.Equal(t, "1.0.0", opts.GetVersion())

	opts.SetDriver(kind.REDIS)
	assert.Equal(t, kind.REDIS, opts.GetDriver())

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cachemock := NewMockCacheRepo(ctrl)
	opts.SetCache(cachemock)
	assert.Equal(t, cachemock, opts.GetCache())
}
