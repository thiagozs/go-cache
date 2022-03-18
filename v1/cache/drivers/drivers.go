package drivers

import (
	buntdblayer "github.com/thiagozs/go-cache/v1/cache/drivers/buntdb"
	gocachelayer "github.com/thiagozs/go-cache/v1/cache/drivers/gocache"
	redislayer "github.com/thiagozs/go-cache/v1/cache/drivers/redis"
	"github.com/thiagozs/go-cache/v1/cache/options"
)

type Driver int

const (
	BUNTDB Driver = iota
	REDIS
	GOCACHE
)

func (d Driver) String() string {
	return []string{"buntdb", "redis", "gocache"}[d]
}

type Drivers struct {
	db DriverPort
}

type DriverPort interface {
	WriteKeyVal(key string, val string) error
	WriteKeyValTTL(key string, val string, ttlSeconds int) error
	DeleteKey(key string) (string, error)
	WriteKeyValAsJSON(key string, val interface{}) error
	WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error
	GetVal(key string) (string, error)
}

func NewDriver(driver Driver, opts ...options.Options) (DriverPort, error) {
	var db DriverPort
	var err error
	switch driver {
	case BUNTDB:
		db, err = buntdblayer.NewBuntDB(opts...)
	case REDIS:
		db, err = redislayer.NewRedis(opts...)
	case GOCACHE:
		db, err = gocachelayer.NewMemory(opts...)
	}
	if err != nil {
		return nil, err
	}
	return &Drivers{db: db}, nil
}

func (d *Drivers) WriteKeyVal(key string, val string) error {
	return d.db.WriteKeyVal(key, val)
}

func (d *Drivers) WriteKeyValTTL(key string, val string, ttlSeconds int) error {
	return d.db.WriteKeyValTTL(key, val, ttlSeconds)
}

func (d *Drivers) DeleteKey(key string) (string, error) {
	return d.db.DeleteKey(key)
}

func (d *Drivers) WriteKeyValAsJSON(key string, val interface{}) error {
	return d.db.WriteKeyValAsJSON(key, val)
}

func (d *Drivers) WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error {
	return d.db.WriteKeyValAsJSONTTL(key, val, ttlSeconds)
}

func (d *Drivers) GetVal(key string) (string, error) {
	return d.db.GetVal(key)
}
