package buntdb

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/v1/cache/drivers/kind"
	"github.com/thiagozs/go-cache/v1/cache/options"
	"github.com/thiagozs/go-utils/files"
	"github.com/tidwall/buntdb"
)

type BuntDBLayer struct {
	db     *buntdb.DB
	file   string
	folder string
	ttl    int
	log    zerolog.Logger
	driver kind.Driver
}

func NewBuntDB(driver kind.Driver, opts ...options.Options) (*BuntDBLayer, error) {
	mts := &options.OptionsCfg{}
	for _, op := range opts {
		err := op(mts)
		if err != nil {
			return &BuntDBLayer{}, err
		}
	}
	return newInstance(driver, mts)
}

func newInstance(driver kind.Driver, opt *options.OptionsCfg) (*BuntDBLayer, error) {

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

	pathFile := fmt.Sprintf("%s/%s", opt.GetFolder(), opt.GetFileName())

	log.Info().Str("path_file", pathFile).Msg("file database")
	if !files.FileExists(pathFile) {
		if err := files.MkdirAll(path.Dir(pathFile)); err != nil {
			log.Info().Err(err).Msg("fail create a directory")
			return nil, err
		}
	}

	db, err := buntdb.Open(pathFile)
	if err != nil {
		log.Info().Err(err).Msg("could not open data file path")
		return nil, err
	}
	return &BuntDBLayer{
		db:     db,
		folder: opt.GetFolder(),
		file:   opt.GetFileName(),
		ttl:    opt.GetTTL(),
		log:    log,
		driver: driver,
	}, nil
}

func (d *BuntDBLayer) GetVal(key string) (string, error) {
	var value string
	err := d.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			d.log.Debug().Err(err).Msg("GetVal")
			return err
		}
		value = val
		return nil
	})
	d.log.Debug().Str("method", "get").
		Str("key", key).
		Str("value", value).
		Msg("GetVal")
	return value, err
}

func (d *BuntDBLayer) DeleteKey(key string) (string, error) {
	var value string
	err := d.db.Update(func(tx *buntdb.Tx) error {
		val, err := tx.Delete(key)
		if err != nil {
			d.log.Debug().Err(err).Msg("DeleteKey")
			return err
		}
		value = val
		return nil
	})
	d.log.Debug().Str("method", "delete").
		Str("key", key).
		Str("value", value).
		Msg("DeleteKey")
	return value, err
}

func (d *BuntDBLayer) WriteKeyVal(key string, val string) error {
	err := d.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, val, nil)
		d.log.Debug().Err(err).Msg("WriteKeyVal")
		return err
	})
	if err != nil {
		return err
	}
	d.log.Debug().Str("method", "write").
		Str("key", key).
		Str("value", val).
		Msg("WriteKeyVal")
	return nil
}

func (d *BuntDBLayer) WriteKeyValTTL(key string, val string, ttlSeconds int) error {
	if ttlSeconds == 0 {
		d.log.Debug().Int("ttl_seconds", d.ttl).Msg("WriteKeyValTTL")
		ttlSeconds = d.ttl
	}
	err := d.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, val, &buntdb.SetOptions{Expires: true, TTL: time.Second * time.Duration(ttlSeconds)})
		d.log.Debug().Str("method", "update").Err(err).Msg("WriteKeyValTTL")
		return err
	})
	if err != nil {
		d.log.Debug().Err(err).Msg("WriteKeyValTTL")
		return err
	}

	d.log.Debug().Str("method", "cache.Set").
		Str("key", key).
		Str("value", val).
		Int64("ttl", int64(ttlSeconds)).
		Msg("WriteKeyValTTL")
	return nil
}

func (d *BuntDBLayer) WriteKeyValAsJSON(key string, val interface{}) error {
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		d.log.Debug().Str("method", "write").Err(err).Msg("WriteKeyValAsJSON")
		return err
	}
	return d.WriteKeyVal(key, string(valueAsJSON))
}

func (d *BuntDBLayer) WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error {
	if ttlSeconds == 0 {
		d.log.Debug().Int("ttl_seconds", d.ttl).Msg("WriteKeyValTTL")
		ttlSeconds = d.ttl
	}

	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		d.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSONTTL")
		return err
	}
	return d.WriteKeyValTTL(key, string(valueAsJSON), ttlSeconds)
}

func (d *BuntDBLayer) GetDriver() kind.Driver {
	return d.driver
}
