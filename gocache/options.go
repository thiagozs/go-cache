package gocache

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/kind"
)

type Options func(o *MemoryParams) error

type MemoryParams struct {
	logDebug    bool
	logDisable  bool
	driver      kind.Driver
	log         zerolog.Logger
	tExpiration time.Duration
	tCleanUpInt time.Duration
	cache       *cache.Cache
}

func newMemoryParams(opts ...Options) (*MemoryParams, error) {
	mts := &MemoryParams{}
	for _, op := range opts {
		err := op(mts)
		if err != nil {
			return &MemoryParams{}, err
		}
	}
	return mts, nil
}

func OptTExpiration(t time.Duration) Options {
	return func(o *MemoryParams) error {
		o.tExpiration = t
		return nil
	}
}

func OptTCleanUpInt(t time.Duration) Options {
	return func(o *MemoryParams) error {
		o.tCleanUpInt = t
		return nil
	}
}

func OptLogDebug(value bool) Options {
	return func(o *MemoryParams) error {
		o.logDebug = value
		return nil
	}
}

func OptLogDisable(value bool) Options {
	return func(o *MemoryParams) error {
		o.logDisable = value
		return nil
	}
}

func OptLog(log zerolog.Logger) Options {
	return func(o *MemoryParams) error {
		o.log = log
		return nil
	}
}

func OptCache(cache *cache.Cache) Options {
	return func(o *MemoryParams) error {
		o.cache = cache
		return nil
	}
}

func OptDriver(driver kind.Driver) Options {
	return func(o *MemoryParams) error {
		o.driver = driver
		return nil
	}
}

func OptTTL(t time.Duration) Options {
	return func(o *MemoryParams) error {
		o.tExpiration = t
		return nil
	}
}

// getters -----

func (m *MemoryParams) GetLogDebug() bool {
	return m.logDebug
}

func (m *MemoryParams) GetLogDisable() bool {
	return m.logDisable
}

func (m *MemoryParams) GetDriver() kind.Driver {
	return m.driver
}

func (m *MemoryParams) GetLog() zerolog.Logger {
	return m.log
}

func (m *MemoryParams) GetTExpiration() time.Duration {
	return m.tExpiration
}

func (m *MemoryParams) GetTCleanUpInt() time.Duration {
	return m.tCleanUpInt
}

func (m *MemoryParams) GetCache() *cache.Cache {
	return m.cache
}

// setters -----

func (m *MemoryParams) SetLogDebug(value bool) {
	m.logDebug = value
}

func (m *MemoryParams) SetLogDisable(value bool) {
	m.logDisable = value
}

func (m *MemoryParams) SetDriver(driver kind.Driver) {
	m.driver = driver
}

func (m *MemoryParams) SetLogger(log zerolog.Logger) {

	m.log = log
}

func (m *MemoryParams) SetTExpiration(t time.Duration) {
	m.tExpiration = t
}

func (m *MemoryParams) SetTCleanUpInt(t time.Duration) {
	m.tCleanUpInt = t
}

func (m *MemoryParams) SetCache(cache *cache.Cache) {
	m.cache = cache
}
