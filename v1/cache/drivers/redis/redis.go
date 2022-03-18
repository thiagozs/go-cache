package redislayer

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/v1/cache/drivers/kind"
	"github.com/thiagozs/go-cache/v1/cache/options"
)

type RedisLayerRepo interface {
	WriteKeyVal(key string, val string) error
	WriteKeyValTTL(key string, val string, ttlSeconds int) error
	DeleteKey(key string) (string, error)
	WriteKeyValAsJSON(key string, val interface{}) error
	WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error
	GetVal(key string) (string, error)
	GetDriver() kind.Driver
}

type redisblayer struct {
	host     string
	password string
	user     string
	port     int
	ttl      int
	log      zerolog.Logger
	rdb      *redis.Client
	driver   kind.Driver
}

func NewRedis(driver kind.Driver, opts ...options.Options) (RedisLayerRepo, error) {
	mts := &options.OptionsCfg{}
	for _, op := range opts {
		err := op(mts)
		if err != nil {
			return nil, err
		}
	}
	return newInstance(driver, mts)
}

func newInstance(driver kind.Driver, opt *options.OptionsCfg) (RedisLayerRepo, error) {

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

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", opt.GetHost(), opt.GetPort()),
		Password: opt.GetPassword(), // no password set
		DB:       0,                 // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Error().Err(err).Msg("error ping")
		return nil, err
	}

	return &redisblayer{
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

func (r *redisblayer) Ping() (string, error) {
	return r.rdb.Ping().Result()
}

func (d *redisblayer) GetVal(key string) (string, error) {
	d.log.Debug().Str("method", "get").
		Str("key", key).
		Msg("GetVal")
	return d.rdb.Get(key).Result()
}

func (d *redisblayer) DeleteKey(key string) (string, error) {
	val, err := d.rdb.Del(key).Result()
	if err != nil {
		d.log.Debug().Err(err).Msg("DeleteKey")
		return "", err
	}
	d.log.Debug().Str("method", "delete").
		Str("key", key).
		Msg("DeleteKey")
	return fmt.Sprintf("%d", val), nil
}

func (d *redisblayer) WriteKeyVal(key string, val string) error {
	d.log.Debug().Str("method", "cache.Set").
		Str("key", key).
		Str("value", val).
		Msg("WriteKeyVal")
	return d.rdb.Set(key, val, time.Duration(0)).Err()
}

func (d *redisblayer) WriteKeyValTTL(key string, val string, ttlSeconds int) error {
	if ttlSeconds == 0 {
		d.log.Debug().Int("ttl_seconds", d.ttl).Msg("WriteKeyValTTL")
		ttlSeconds = d.ttl
	}
	return d.rdb.Set(key, val, time.Duration(ttlSeconds)).Err()
}

func (d *redisblayer) WriteKeyValAsJSON(key string, val interface{}) error {
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		d.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSON")
		return err
	}
	return d.WriteKeyVal(key, string(valueAsJSON))
}

func (d *redisblayer) WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error {
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

func (d *redisblayer) GetDriver() kind.Driver {
	return d.driver
}
