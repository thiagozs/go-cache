package redis

import (
	"testing"

	redis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/thiagozs/go-cache/kind"
)

func TestOptions(t *testing.T) {
	opts, err := newRedisParams()
	assert.NoError(t, err)

	opts.SetOptions(OptDatabase(1))
	assert.Equal(t, 1, opts.GetDatabase())

	opts.SetOptions(OptHost("localhost"))
	assert.Equal(t, "localhost", opts.GetHost())

	opts.SetOptions(OptPassword("password"))
	assert.Equal(t, "password", opts.GetPassword())

	opts.SetOptions(OptUser("user"))
	assert.Equal(t, "user", opts.GetUser())

	opts.SetOptions(OptPort(6379))
	assert.Equal(t, 6379, opts.GetPort())

	opts.SetOptions(OptTTL(60))
	assert.Equal(t, 60, opts.GetTTL())

	logger := zerolog.New(nil)
	opts.SetOptions(OptLogger(logger))
	assert.Equal(t, logger, opts.GetLogger())

	opts.SetOptions(OptDriver(kind.REDIS))
	assert.Equal(t, kind.REDIS, opts.GetDriver())

	client := redis.NewClient(&redis.Options{})
	opts.SetOptions(OptRedis(client))
	assert.Equal(t, client, opts.GetRedis())

	opts.SetOptions(OptLogDebug(true))
	assert.Equal(t, true, opts.GetLogDebug())

	opts.SetOptions(OptLogDisable(true))
	assert.Equal(t, true, opts.GetLogDisable())

	redisOpts := opts.GetRedisOptions()
	assert.Equal(t, "localhost", redisOpts.Addr)
	assert.Equal(t, "password", redisOpts.Password)
	assert.Equal(t, 0, redisOpts.DB)

	redisOptsWithUser := opts.GetRedisOptionsWithUser()
	assert.Equal(t, "user", redisOptsWithUser.Username)

	redisOptsWithDB := opts.GetRedisOptionsWithDB()
	assert.Equal(t, 0, redisOptsWithDB.DB)
}
