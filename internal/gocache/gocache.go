package gocache

import (
	"fmt"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/thiagozs/go-cache/internal/logging"
	"github.com/thiagozs/go-cache/kind"
	"github.com/thiagozs/go-xutils"
)

type GocacheLayer struct {
	mu     sync.Mutex
	params *MemoryParams
	utils  *xutils.XUtils
}

func NewMemory(opts ...Options) (*GocacheLayer, error) {
	params, err := newMemoryParams(opts...)
	if err != nil {
		return nil, err
	}

	log := logging.New(params.GetLogDebug(), params.GetLogDisable())

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
	str, ok := iface.(string)
	if !ok {
		return "", fmt.Errorf("value for key %q is not a string", key)
	}
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

	m.params.log.Debug().Str("method", "write").
		Str("key", key).
		Str("value", val).
		Int64("ttl", int64(ttlc)).
		Msg("WriteKeyValTTL")
	m.params.GetCache().Set(key, val, ttlc)
	return nil
}

func (m *GocacheLayer) GetDriver() kind.Driver {
	return m.params.GetDriver()
}

func (m *GocacheLayer) Incr(key string) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	val, err := m.GetVal(key)
	if err != nil {
		m.params.log.Debug().Str("method", "cache.Get").Err(err).Msg("Incr")
		return 0, err
	}

	v, err := m.utils.Convs().ToInt64(val)
	if err != nil {
		m.params.log.Debug().Err(err).Msg("Incr")
		return 0, err
	}

	v++

	if err := m.WriteKeyVal(key, fmt.Sprintf("%d", v)); err != nil {
		m.params.log.Debug().Str("method", "WriteKeyVal").Err(err).Msg("Incr")
		return 0, err
	}

	m.params.log.Debug().Str("method", "incr").
		Str("key", key).
		Int64("value", v).
		Msg("Incr")

	return v, nil
}

func (m *GocacheLayer) Decr(key string) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	val, err := m.GetVal(key)
	if err != nil {
		m.params.log.Debug().Str("method", "cache.Get").Err(err).Msg("Decr")
		return 0, err
	}

	v, err := m.utils.Convs().ToInt64(val)
	if err != nil {
		m.params.log.Debug().Err(err).Msg("Decr")
		return 0, err
	}

	v--

	if err := m.WriteKeyVal(key, fmt.Sprintf("%d", v)); err != nil {
		m.params.log.Debug().Str("method", "WriteKeyVal").Err(err).Msg("Decr")
		return 0, err
	}

	m.params.log.Debug().Str("method", "decr").
		Str("key", key).
		Int64("value", v).
		Msg("Decr")

	return v, nil
}
