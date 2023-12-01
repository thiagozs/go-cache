package gocache

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/kind"
	"github.com/thiagozs/go-xutils"
)

type GocacheLayer struct {
	params *MemoryParams
	utils  *xutils.XUtils
}

func NewMemory(opts ...Options) (*GocacheLayer, error) {
	params, err := newMemoryParams(opts...)
	if err != nil {
		return nil, err
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

	c := cache.New(params.GetTExpiration(), params.GetTCleanUpInt())

	params.SetLogger(log)
	params.SetCache(c)

	return &GocacheLayer{
		params: params,
		utils:  xutils.New(),
	}, nil
}

func (m *GocacheLayer) GetVal(key string) (string, error) {
	iface, found := m.params.GetCache().Get(key)
	if !found {
		err := fmt.Errorf("not found")
		m.params.log.Debug().Str("method", "cache.Get").Err(err).Msg("GetVal")
		return "", err
	}
	str := reflect.ValueOf(iface).String()
	m.params.log.Debug().Str("method", "get").
		Str("key", key).
		Str("value", str).
		Msg("GetVal")
	return str, nil
}

func (m *GocacheLayer) DeleteKey(key string) (string, error) {
	m.params.log.Debug().Str("method", "delete").
		Str("key", key).
		Msg("DeleteKey")
	m.params.GetCache().Delete(key)
	return "", nil
}

func (m *GocacheLayer) WriteKeyVal(key string, val string) error {
	m.params.log.Debug().Str("method", "write").
		Str("key", key).
		Str("value", val).
		Msg("WriteKeyVal")
	m.params.GetCache().Set(key, val, cache.NoExpiration)
	return nil
}

func (m *GocacheLayer) WriteKeyValTTL(key string, val string, insec int) error {
	if insec == 0 {
		return fmt.Errorf("ttl_seconds is zero")
	}

	ttlc := time.Duration(insec) * time.Second

	m.params.log.Debug().Int("ttl_seconds", int(ttlc)).Msg("WriteKeyValTTL")

	m.params.log.Debug().Str("method", "write").
		Str("key", key).
		Str("value", val).
		Int64("ttl", int64(ttlc)).
		Msg("WriteKeyValTTL")
	m.params.GetCache().Set(key, val, ttlc)
	return nil
}

func (m *GocacheLayer) WriteKeyValAsJSON(key string, val any) error {
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		m.params.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSON")
		return err
	}
	return m.WriteKeyVal(key, string(valueAsJSON))
}

func (m *GocacheLayer) WriteKeyValAsJSONTTL(key string, val any, insec int) error {
	if insec == 0 {
		return fmt.Errorf("ttl_seconds is zero")
	}

	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		m.params.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSONTTL")
		return err
	}

	return m.WriteKeyValTTL(key, string(valueAsJSON), insec)
}

func (m *GocacheLayer) GetDriver() kind.Driver {
	return m.params.GetDriver()
}

func (m *GocacheLayer) Incr(key string) (int64, error) {
	val, err := m.GetVal(key)
	if err != nil {
		err := fmt.Errorf("not found")
		m.params.log.Debug().Str("method", "cache.Get").Err(err).Msg("Decr")
		return 0, err
	}

	v, _ := m.utils.Convs().ToInt64(val)

	m.params.log.Debug().Str("method", "get").
		Str("key", key).
		Int64("value", v).
		Msg("Incr")

	v = v + 1

	if err := m.WriteKeyVal(key,
		fmt.Sprintf("%d", v)); err != nil {
		m.params.log.Debug().Str("method", "WriteKeyVal").
			Err(err).Msg("Incr")
		return 0, err
	}

	return v, nil
}

func (m *GocacheLayer) Decr(key string) (int64, error) {
	val, err := m.GetVal(key)
	if err != nil {
		err := fmt.Errorf("not found")
		m.params.log.Debug().Str("method", "cache.Get").Err(err).Msg("Decr")
		return 0, err
	}

	v, _ := m.utils.Convs().ToInt64(val)

	m.params.log.Debug().Str("method", "get").
		Str("key", key).
		Int64("value", v).
		Msg("Decr")

	v = v - 1

	if err := m.WriteKeyVal(key,
		fmt.Sprintf("%d", v)); err != nil {
		m.params.log.Debug().Str("method", "WriteKeyVal").
			Err(err).Msg("Decr")
		return 0, err
	}

	return v, nil
}
