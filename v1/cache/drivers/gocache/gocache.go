package memorylayer

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/v1/cache/options"
)

type MemoryLayerRepo interface {
	WriteKeyVal(key string, val string) error
	WriteKeyValTTL(key string, val string, ttlSeconds int) error
	DeleteKey(key string) (string, error)
	WriteKeyValAsJSON(key string, val interface{}) error
	WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error
	GetVal(key string) (string, error)
}

type memorylayer struct {
	tExpiration time.Duration
	tCleanupInt time.Duration
	log         zerolog.Logger
	cache       *cache.Cache
}

func NewMemory(opts ...options.Options) (MemoryLayerRepo, error) {
	mts := &options.OptionsCfg{}
	for _, op := range opts {
		err := op(mts)
		if err != nil {
			return nil, err
		}
	}
	return newInstance(mts)
}

func newInstance(opt *options.OptionsCfg) (MemoryLayerRepo, error) {

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

	c := cache.New(opt.GetTExpiration(), opt.GetTCleanUpInt())

	return &memorylayer{
		log:         log,
		cache:       c,
		tExpiration: opt.GetTExpiration(),
		tCleanupInt: opt.GetTCleanUpInt(),
	}, nil
}

func (m *memorylayer) GetVal(key string) (string, error) {
	iface, found := m.cache.Get(key)
	if !found {
		err := fmt.Errorf("not found")
		m.log.Debug().Str("method", "cache.Get").Err(err).Msg("GetVal")
		return "", err
	}
	str := reflect.ValueOf(iface).String()
	m.log.Debug().Str("method", "get").
		Str("key", key).
		Str("value", str).
		Msg("GetVal")
	return str, nil
}

func (m *memorylayer) DeleteKey(key string) (string, error) {
	m.log.Debug().Str("method", "delete").
		Str("key", key).
		Msg("DeleteKey")
	m.cache.Delete(key)
	return "", nil
}

func (m *memorylayer) WriteKeyVal(key string, val string) error {
	m.log.Debug().Str("method", "write").
		Str("key", key).
		Str("value", val).
		Msg("WriteKeyVal")
	m.cache.Set(key, val, cache.NoExpiration)
	return nil
}

func (m *memorylayer) WriteKeyValTTL(key string, val string, ttlSeconds int) error {
	ttlc := time.Duration(60) * time.Second
	if ttlSeconds > 0 {
		m.log.Debug().Int("ttl_seconds", int(ttlc)).Msg("WriteKeyValTTL")
		ttlc = time.Duration(ttlSeconds) * time.Second
	}
	m.log.Debug().Str("method", "write").
		Str("key", key).
		Str("value", val).
		Int64("ttl", int64(ttlc)).
		Msg("WriteKeyValTTL")
	m.cache.Set(key, val, ttlc)
	return nil
}

func (m *memorylayer) WriteKeyValAsJSON(key string, val interface{}) error {
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		m.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSON")
		return err
	}
	return m.WriteKeyVal(key, string(valueAsJSON))
}

func (m *memorylayer) WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error {
	ttlc := 60
	if ttlSeconds > 0 {
		m.log.Debug().Int("ttl_seconds", ttlc).Msg("WriteKeyValAsJSONTTL")
		ttlc = ttlSeconds
	}
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		m.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSONTTL")
		return err
	}

	return m.WriteKeyValTTL(key, string(valueAsJSON), ttlSeconds)
}
