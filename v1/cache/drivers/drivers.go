package drivers

import (
	buntdblayer "github.com/thiagozs/go-cache/v1/cache/drivers/buntdb"
	gocachelayer "github.com/thiagozs/go-cache/v1/cache/drivers/gocache"
	"github.com/thiagozs/go-cache/v1/cache/drivers/kind"
	redislayer "github.com/thiagozs/go-cache/v1/cache/drivers/redis"
	"github.com/thiagozs/go-cache/v1/cache/options"
)

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
	GetDriver() kind.Driver
}

func NewDriver(driver kind.Driver, opts ...options.Options) (DriverPort, error) {
	var db DriverPort
	var err error
	switch driver {
	case kind.BUNTDB:
		db, err = buntdblayer.NewBuntDB(driver, opts...)
	case kind.REDIS:
		db, err = redislayer.NewRedis(driver, opts...)
	case kind.GOCACHE:
		db, err = gocachelayer.NewMemory(driver, opts...)
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

func (d *Drivers) GetDriver() kind.Driver {
	return d.db.GetDriver()
}
