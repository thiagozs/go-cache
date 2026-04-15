package cache

import (
	"encoding/json"
	"fmt"

	"github.com/thiagozs/go-cache/internal/buntdb"
	"github.com/thiagozs/go-cache/internal/gocache"
	"github.com/thiagozs/go-cache/internal/redis"
	"github.com/thiagozs/go-cache/kind"
)

type Cache struct {
	db CacheRepo
}

func New(opts ...Options) (*Cache, error) {
	var db CacheRepo
	var err error

	params, err := newCacheParams(opts...)
	if err != nil {
		return nil, err
	}

	if params.GetCache() != nil {
		return &Cache{db: params.GetCache()}, nil
	}

	switch params.GetDriver() {
	case kind.BUNTDB:
		opts := []buntdb.Options{
			buntdb.OptFile(params.GetFile()),
			buntdb.OptFolder(params.GetFolder()),
			buntdb.OptTTL(params.GetTTL()),
			buntdb.OptLogDebug(params.GetLogDebug()),
			buntdb.OptLogDisable(params.GetLogDisable()),
			buntdb.OptDriver(params.GetDriver()),
		}
		db, err = buntdb.NewBuntDB(opts...)
	case kind.REDIS:
		opts := []redis.Options{
			redis.OptHost(params.GetHost()),
			redis.OptPort(params.GetPort()),
			redis.OptPassword(params.GetPassword()),
			redis.OptUser(params.GetUser()),
			redis.OptTTL(params.GetTTL()),
			redis.OptLogDebug(params.GetLogDebug()),
			redis.OptLogDisable(params.GetLogDisable()),
			redis.OptDriver(params.GetDriver()),
		}
		db, err = redis.NewRedis(opts...)
	case kind.GOCACHE:
		opts := []gocache.Options{
			gocache.OptTExpiration(params.GetTExpiration()),
			gocache.OptTCleanUpInt(params.GetTCleanUpInt()),
			gocache.OptLogDebug(params.GetLogDebug()),
			gocache.OptLogDisable(params.GetLogDisable()),
			gocache.OptDriver(params.GetDriver()),
		}
		db, err = gocache.NewMemory(opts...)
	default:
		return nil, fmt.Errorf("unsupported cache driver: %v", params.GetDriver())
	}

	if err != nil {
		return nil, err
	}

	return &Cache{db: db}, nil
}

func (d *Cache) WriteKeyVal(key string, val string) error {
	return d.db.WriteKeyVal(key, val)
}

func (d *Cache) WriteKeyValTTL(key string, val string, ttlSeconds int) error {
	return d.db.WriteKeyValTTL(key, val, ttlSeconds)
}

func (d *Cache) DeleteKey(key string) (string, error) {
	return d.db.DeleteKey(key)
}

func (d *Cache) WriteKeyValAsJSON(key string, val interface{}) error {
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return d.db.WriteKeyVal(key, string(valueAsJSON))
}

func (d *Cache) WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error {
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return d.db.WriteKeyValTTL(key, string(valueAsJSON), ttlSeconds)
}

func (d *Cache) GetVal(key string) (string, error) {
	return d.db.GetVal(key)
}

func (d *Cache) GetDriver() kind.Driver {
	return d.db.GetDriver()
}

func (d *Cache) Incr(key string) (int64, error) {
	return d.db.Incr(key)
}

func (d *Cache) Decr(key string) (int64, error) {
	return d.db.Decr(key)
}
