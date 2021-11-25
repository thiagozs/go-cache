package buntdblayer

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/v1/cache/drivers"
	"github.com/thiagozs/go-utils/files"
	"github.com/tidwall/buntdb"
)

type BuntDBLayerRepo interface {
	WriteKeyVal(key string, val string) error
	WriteKeyValTTL(key string, val string, ttlSeconds int) error
	DeleteKey(key string) (string, error)
	WriteKeyValAsJSON(key string, val interface{}) error
	WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error
	GetVal(key string) (string, error)
}

type buntdblayer struct {
	db     *buntdb.DB
	file   string
	folder string
	ttl    int
	log    zerolog.Logger
}

func NewBuntDB(opts ...drivers.Options) (BuntDBLayerRepo, error) {
	mts := &drivers.OptionsCfg{}
	for _, op := range opts {
		err := op(mts)
		if err != nil {
			return nil, err
		}
	}
	return newInstance(mts.GetFolder(), mts.GetFileName(),
		mts.GetTTL(), mts.GetLogDebug(), mts.GetLogDisable())
}

func newInstance(folder, file string, ttl int,
	logDebug bool, logDisable bool) (BuntDBLayerRepo, error) {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()

	if logDebug {
		log.Info().Bool("debug", true).Msg("log debug")
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if logDisable {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	pathFile := fmt.Sprintf("%s/%s", folder, file)

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
	return &buntdblayer{
		db:     db,
		folder: folder,
		file:   file,
		ttl:    ttl,
		log:    log,
	}, nil
}

func (d *buntdblayer) GetVal(key string) (string, error) {
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
	return value, err
}

func (d *buntdblayer) DeleteKey(key string) (string, error) {
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
	return value, err
}

func (d *buntdblayer) WriteKeyVal(key string, val string) error {
	err := d.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, val, nil)
		d.log.Debug().Err(err).Msg("WriteKeyVal")
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *buntdblayer) WriteKeyValTTL(key string, val string, ttlSeconds int) error {
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
	return nil
}

func (d *buntdblayer) WriteKeyValAsJSON(key string, val interface{}) error {
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		d.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSON")
		return err
	}
	return d.WriteKeyVal(key, string(valueAsJSON))
}

func (d *buntdblayer) WriteKeyValAsJSONTTL(key string, val interface{}, ttlSeconds int) error {
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		d.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSONTTL")
		return err
	}
	return d.WriteKeyValTTL(key, string(valueAsJSON), ttlSeconds)
}
