package redislayer

import (
	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/v1/cache/drivers"
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

func NewRedis(opts ...drivers.Options) (RedisLayerRepo, error) {
	mts := &drivers.OptionsCfg{}
	for _, op := range opts {
		err := op(mts)
		if err != nil {
			return nil, err
		}
	}
	return newInstance(mts.GetFolder(), mts.GetFileName(),
		mts.GetTTL(), mts.GetLogDebug(), mts.GetLogDisable())
}

func newInstance(folder, file string, ttl int,
	logDebug bool, logDisable bool) (RedisLayerRepo, error) {
	return &redisblayer{}, nil
}

func (d *redisblayer) GetVal(key string) (string, error) {
	var value string

	return value, nil
}

func (d *redisblayer) DeleteKey(key string) (string, error) {
	var value string

	return value, nil
}

func (d *redisblayer) WriteKeyVal(key string, val string) error {

	return nil
}

func (d *redisblayer) WriteKeyValTTL(key string, val string, ttlSeconds int) error {

	return nil
}

func (d *redisblayer) WriteKeyValAsJSON(key string, val interface{}) error {

	return d.WriteKeyVal(key, string(""))
}

func (d *redisblayer) WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error {

	return d.WriteKeyValTTL(key, string(""), ttlSeconds)
}
