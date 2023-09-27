package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	redis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/v1/cache/drivers/kind"
	"github.com/thiagozs/go-cache/v1/cache/options"
)

type RedisLayer struct {
	host     string
	password string
	user     string
	port     int
	ttl      int
	log      zerolog.Logger
	rdb      *redis.Client
	driver   kind.Driver
}

func NewRedis(driver kind.Driver, opts ...options.Options) (*RedisLayer, error) {
	mts := &options.OptionsCfg{}
	for _, op := range opts {
		err := op(mts)
		if err != nil {
			return &RedisLayer{}, err
		}
	}
	return newInstance(driver, mts)
}

func newInstance(driver kind.Driver, opt *options.OptionsCfg) (*RedisLayer, error) {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()

	if opt.GetLogDebug() {
		log.Info().Bool("debug", true).Msg("log debug")
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if opt.GetLogDisable() {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	db, _ := strconv.Atoi(opt.GetDatabase())

	// redis ACL system need to be set on 6.0.0 higher
	// Default for redis lower than 6.0.0 is empty string
	redisOpts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", opt.GetHost(), opt.GetPort()),
		Password: opt.GetPassword(),
		DB:       db,
	}
	if opt.GetVersion() >= "6.0.0" {
		redisOpts.Username = opt.GetUser()
	}

	rdb := redis.NewClient(redisOpts)
	ctx := context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error().Err(err).Msg("error ping")
		return nil, err
	}

	return &RedisLayer{
		log:      log,
		host:     opt.GetHost(),
		password: opt.GetPassword(),
		user:     opt.GetUser(),
		port:     opt.GetPort(),
		ttl:      opt.GetTTL(),
		rdb:      rdb,
		driver:   driver,
	}, nil
}

func (r *RedisLayer) Ping() (string, error) {
	return r.rdb.Ping(context.Background()).Result()
}

func (d *RedisLayer) GetVal(key string) (string, error) {
	d.log.Debug().Str("method", "get").
		Str("key", key).
		Msg("GetVal")

	val, err := d.rdb.Get(context.Background(), key).Result()
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

func (d *RedisLayer) DeleteKey(key string) (string, error) {
	val, err := d.rdb.Del(context.Background(), key).Result()
	if err != nil {
		d.log.Debug().Err(err).Msg("DeleteKey")
		return "", err
	}
	d.log.Debug().Str("method", "delete").
		Str("key", key).
		Msg("DeleteKey")
	return fmt.Sprintf("%d", val), nil
}

func (d *RedisLayer) WriteKeyVal(key string, val string) error {
	d.log.Debug().Str("method", "cache.Set").
		Str("key", key).
		Str("value", val).
		Msg("WriteKeyVal")
	return d.rdb.Set(context.Background(), key, val, time.Duration(0)).Err()
}

func (d *RedisLayer) WriteKeyValTTL(key string, val string, ttlSeconds int) error {
	if ttlSeconds == 0 {
		d.log.Debug().Int("ttl_seconds", d.ttl).Msg("WriteKeyValTTL")
		ttlSeconds = d.ttl
	}
	return d.rdb.Set(context.Background(), key, val, time.Duration(ttlSeconds)).Err()
}

func (d *RedisLayer) WriteKeyValAsJSON(key string, val interface{}) error {
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		d.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSON")
		return err
	}
	return d.WriteKeyVal(key, string(valueAsJSON))
}

func (d *RedisLayer) WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error {
	if ttlSeconds == 0 {
		d.log.Debug().Int("ttl_seconds", d.ttl).Msg("WriteKeyValAsJSONTTL")
		ttlSeconds = d.ttl
	}
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		d.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSONTTL")
		return err
	}

	return d.WriteKeyValTTL(key, string(valueAsJSON), ttlSeconds)
}

func (d *RedisLayer) GetDriver() kind.Driver {
	return d.driver
}
