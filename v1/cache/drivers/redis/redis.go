package redislayer

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/v1/cache/options"
)

type RedisLayerRepo interface {
	WriteKeyVal(key string, val string) error
	WriteKeyValTTL(key string, val string, ttlSeconds int) error
	DeleteKey(key string) (string, error)
	WriteKeyValAsJSON(key string, val interface{}) error
	WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error
	GetVal(key string) (string, error)
}

type redisblayer struct {
	host     string
	password string
	user     string
	port     int
	ttl      int
	log      zerolog.Logger
}

func NewRedis(opts ...options.Options) (RedisLayerRepo, error) {
	mts := &options.OptionsCfg{}
	for _, op := range opts {
		err := op(mts)
		if err != nil {
			return nil, err
		}
	}
	return newInstance(mts)
}

func newInstance(opt *options.OptionsCfg) (RedisLayerRepo, error) {

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

	return &redisblayer{
		log:      log,
		host:     opt.GetHost(),
		password: opt.GetPassword(),
		user:     opt.GetUser(),
		port:     opt.GetPort(),
		ttl:      opt.GetTTL(),
	}, nil
}

func (d *redisblayer) GetVal(key string) (string, error) {
	var value string
	// TODO: implement
	return value, nil
}

func (d *redisblayer) DeleteKey(key string) (string, error) {
	var value string
	// TODO: implement

	return value, nil
}

func (d *redisblayer) WriteKeyVal(key string, val string) error {
	// TODO: implement

	return nil
}

func (d *redisblayer) WriteKeyValTTL(key string, val string, ttlSeconds int) error {
	// TODO: implement

	return nil
}

func (d *redisblayer) WriteKeyValAsJSON(key string, val interface{}) error {
	// TODO: implement

	return d.WriteKeyVal(key, string(""))
}

func (d *redisblayer) WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error {
	// TODO: implement

	return d.WriteKeyValTTL(key, string(""), ttlSeconds)
}
