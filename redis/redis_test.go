package redis

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/alicebob/miniredis/v2"

	redis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/thiagozs/go-cache/kind"
)

func TestRedisLayer(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()

	fmt.Println("Miniredis started at: ", mr.Addr())

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	fmt.Println("Connecting to Redis at: ", mr.Addr())

	mrdis := strings.Split(mr.Addr(), ":")
	port, _ := strconv.Atoi(mrdis[1])

	redisLayer, err := NewRedis(OptRedis(client), OptPort(port), OptDriver(kind.REDIS))
	assert.NoError(t, err)

	t.Run("Ping", func(t *testing.T) {
		pong, err := redisLayer.Ping()
		assert.NoError(t, err)
		assert.Equal(t, "PONG", pong)
	})

	t.Run("WriteKeyVal", func(t *testing.T) {
		err := redisLayer.WriteKeyVal("key", "value")
		assert.NoError(t, err)
	})

	t.Run("GetVal", func(t *testing.T) {
		val, err := redisLayer.GetVal("key")
		assert.NoError(t, err)
		assert.Equal(t, "value", val)
	})

	t.Run("DeleteKey", func(t *testing.T) {
		val, err := redisLayer.DeleteKey("key")
		assert.NoError(t, err)
		assert.Equal(t, "1", val)
	})

	t.Run("WriteKeyValTTL", func(t *testing.T) {
		err := redisLayer.WriteKeyValTTL("key", "value", 10)
		assert.NoError(t, err)
	})

	t.Run("WriteKeyValAsJSON", func(t *testing.T) {
		data := map[string]string{"name": "Alice", "age": "30"}
		err := redisLayer.WriteKeyValAsJSON("key", data)
		assert.NoError(t, err)
	})

	t.Run("WriteKeyValAsJSONTTL", func(t *testing.T) {
		data := map[string]string{"name": "Bob", "age": "35"}
		err := redisLayer.WriteKeyValAsJSONTTL("key", data, 10)
		assert.NoError(t, err)
	})

	t.Run("GetDriver", func(t *testing.T) {
		assert.Equal(t, kind.REDIS, redisLayer.GetDriver())
	})
}
