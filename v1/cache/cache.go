package cache

import (
	"github.com/thiagozs/go-cache/v1/cache/drivers"
	"github.com/thiagozs/go-cache/v1/cache/drivers/kind"
	"github.com/thiagozs/go-cache/v1/cache/options"
)

type CachePort interface {
	WriteKeyVal(key string, val string) error
	WriteKeyValTTL(key string, val string, ttlSeconds int) error
	DeleteKey(key string) (string, error)
	WriteKeyValAsJSON(key string, val interface{}) error
	WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error
	GetVal(key string) (string, error)
	GetDriver() kind.Driver
}
type cache struct {
	db CachePort
}

func New(driver kind.Driver, opts ...options.Options) (CachePort, error) {

	port, err := drivers.NewDriver(driver, opts...)
	if err != nil {
		return nil, err
	}

	return &cache{
		db: port,
	}, nil
}

func (c *cache) GetVal(key string) (string, error) {
	return c.db.GetVal(key)
}

func (c *cache) DeleteKey(key string) (string, error) {
	return c.db.DeleteKey(key)
}

func (c *cache) WriteKeyVal(key string, val string) error {
	return c.db.WriteKeyVal(key, val)
}

func (c *cache) WriteKeyValTTL(key string, val string, ttlSeconds int) error {
	return c.db.WriteKeyValTTL(key, val, ttlSeconds)
}

func (c *cache) WriteKeyValAsJSON(key string, val interface{}) error {
	return c.db.WriteKeyValAsJSON(key, val)
}

func (c *cache) WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error {
	return c.db.WriteKeyValAsJSONTTL(key, val, ttlSeconds)
}

func (c *cache) GetDriver() kind.Driver {
	return c.db.GetDriver()
}
