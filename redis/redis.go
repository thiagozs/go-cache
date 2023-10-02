package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	redis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/kind"
)

type RedisLayer struct {
	params *OptionsParams
}

func NewRedis(opts ...Options) (*RedisLayer, error) {
	params, err := newRedisParams(opts...)
	if err != nil {
		return &RedisLayer{}, err
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()

	if params.GetLogDebug() {
		log.Info().Bool("debug", true).Msg("log debug")
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if params.GetLogDisable() {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	redisOpts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", params.GetHost(), params.GetPort()),
		Password: params.GetPassword(),
		DB:       params.GetDatabase(),
	}

	rdb := redis.NewClient(redisOpts)
	ctx := context.Background()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Error().Err(err).Msg("error ping")
		return nil, err
	}

	params.SetLogger(log)
	params.SetRedis(rdb)

	return &RedisLayer{params: params}, nil
}

func (r *RedisLayer) Ping() (string, error) {
	return r.params.GetRedis().Ping(context.Background()).Result()
}

func (r *RedisLayer) GetVal(key string) (string, error) {
	r.params.log.Debug().Str("method", "get").
		Str("key", key).
		Msg("GetVal")

	val, err := r.params.GetRedis().Get(context.Background(), key).Result()
	switch {
	case err == redis.Nil:
		return val, fmt.Errorf("key does not exist")
	case err != nil:
		return val, fmt.Errorf("Get failed %s", err)
	case val == "":
		return val, fmt.Errorf("value is empty")
	}

	return val, nil
}

func (r *RedisLayer) DeleteKey(key string) (string, error) {
	val, err := r.params.GetRedis().Del(context.Background(), key).Result()
	if err != nil {
		r.params.log.Debug().Err(err).Msg("DeleteKey")
		return "", err
	}
	r.params.log.Debug().Str("method", "delete").
		Str("key", key).
		Msg("DeleteKey")
	return fmt.Sprintf("%d", val), nil
}

func (r *RedisLayer) WriteKeyVal(key string, val string) error {
	r.params.log.Debug().Str("method", "cache.Set").
		Str("key", key).
		Str("value", val).
		Msg("WriteKeyVal")
	return r.params.GetRedis().Set(context.Background(), key, val, time.Duration(0)).Err()
}

func (r *RedisLayer) WriteKeyValTTL(key string, val string, insec int) error {
	if insec == 0 {
		return fmt.Errorf("ttl_seconds is zero")
	}
	r.params.log.Debug().Int("ttl_seconds", r.params.GetTTL()).Msg("WriteKeyValTTL")
	return r.params.GetRedis().Set(context.Background(), key, val, time.Duration(insec)*time.Second).Err()
}

func (r *RedisLayer) WriteKeyValAsJSON(key string, val any) error {
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		r.params.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSON")
		return err
	}
	return r.WriteKeyVal(key, string(valueAsJSON))
}

func (r *RedisLayer) WriteKeyValAsJSONTTL(key string, val any, insec int) error {
	if insec == 0 {
		return fmt.Errorf("ttl_seconds is zero")
	}
	r.params.log.Debug().Int("ttl_seconds", r.params.GetTTL()).Msg("WriteKeyValAsJSONTTL")

	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		r.params.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSONTTL")
		return err
	}

	return r.WriteKeyValTTL(key, string(valueAsJSON), insec)
}

func (r *RedisLayer) GetDriver() kind.Driver {
	return r.params.GetDriver()
}
