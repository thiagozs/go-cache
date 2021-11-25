package drivers

import (
	"fmt"

	"github.com/thiagozs/go-cache/v1/cache/drivers/buntdblayer"
	"github.com/thiagozs/go-cache/v1/cache/drivers/redislayer"
)

type Driver int

const (
	BUNTDB Driver = iota
	REDIS
)

func (d Driver) String() string {
	return []string{"buntdb", "redis"}[d]
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

func NewDriver(driver Driver, opts ...Options) (DriverPort, error) {
	switch driver {
	case BUNTDB:
		return buntdblayer.NewBuntDB(opts...), nil
	case REDIS:
		return redislayer.NewRedis(opts...), nil
	}
	return &Drivers{}, fmt.Errorf("unknown driver type: %s", driver.String())
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
