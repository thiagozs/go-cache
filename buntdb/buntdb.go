package buntdb

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	kind "github.com/thiagozs/go-cache/kind"
	"github.com/thiagozs/go-utils/files"
	"github.com/thiagozs/go-xutils"
	"github.com/tidwall/buntdb"
)

type BuntDBLayer struct {
	params *BuntDBParams
	utils  *xutils.XUtils
}

func NewBuntDB(opts ...Options) (*BuntDBLayer, error) {
	params, err := newBuntDBParams(opts...)
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

	pathFile := fmt.Sprintf("%s/%s", params.GetFolder(), params.GetFile())

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

	params.SetLogger(log)
	params.SetDB(db)

	return &BuntDBLayer{
		params: params,
		utils:  xutils.New(),
	}, nil
}

func (d *BuntDBLayer) GetVal(key string) (string, error) {
	var value string
	err := d.params.GetDB().View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			d.params.log.Debug().Err(err).Msg("GetVal")
			return err
		}
		value = val
		return nil
	})
	d.params.log.Debug().Str("method", "get").
		Str("key", key).
		Str("value", value).
		Msg("GetVal")
	return value, err
}

func (d *BuntDBLayer) DeleteKey(key string) (string, error) {
	var value string
	err := d.params.GetDB().Update(func(tx *buntdb.Tx) error {
		val, err := tx.Delete(key)
		if err != nil {
			d.params.log.Debug().Err(err).Msg("DeleteKey")
			return err
		}
		value = val
		return nil
	})
	d.params.log.Debug().Str("method", "delete").
		Str("key", key).
		Str("value", value).
		Msg("DeleteKey")
	return value, err
}

func (d *BuntDBLayer) WriteKeyVal(key string, val string) error {
	err := d.params.GetDB().Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, val, nil)
		d.params.log.Debug().Err(err).Msg("WriteKeyVal")
		return err
	})
	if err != nil {
		return err
	}
	d.params.log.Debug().Str("method", "write").
		Str("key", key).
		Str("value", val).
		Msg("WriteKeyVal")
	return nil
}

func (d *BuntDBLayer) WriteKeyValTTL(key string, val string, insec int) error {
	if insec == 0 {
		return fmt.Errorf("ttl_seconds is zero")
	}
	err := d.params.GetDB().Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, val, &buntdb.SetOptions{Expires: true, TTL: time.Second * time.Duration(insec)})
		d.params.log.Debug().Str("method", "update").Err(err).Msg("WriteKeyValTTL")
		return err
	})
	if err != nil {
		d.params.log.Debug().Err(err).Msg("WriteKeyValTTL")
		return err
	}

	d.params.log.Debug().Str("method", "cache.Set").
		Str("key", key).
		Str("value", val).
		Int64("ttl", int64(insec)).
		Msg("WriteKeyValTTL")
	return nil
}

func (d *BuntDBLayer) WriteKeyValAsJSON(key string, val any) error {
	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		d.params.log.Debug().Str("method", "write").Err(err).Msg("WriteKeyValAsJSON")
		return err
	}
	return d.WriteKeyVal(key, string(valueAsJSON))
}

func (d *BuntDBLayer) WriteKeyValAsJSONTTL(key string, val any, insec int) error {
	if insec == 0 {
		return fmt.Errorf("ttl_seconds is zero")
	}
	d.params.log.Debug().Int("ttl_seconds", insec).Msg("WriteKeyValTTL")

	valueAsJSON, err := json.Marshal(val)
	if err != nil {
		d.params.log.Debug().Str("method", "json.Marshal").Err(err).Msg("WriteKeyValAsJSONTTL")
		return err
	}
	return d.WriteKeyValTTL(key, string(valueAsJSON), insec)
}

func (d *BuntDBLayer) GetDriver() kind.Driver {
	return d.params.GetDriver()
}

func (d *BuntDBLayer) Incr(key string) (int64, error) {
	val, err := d.GetVal(key)
	if err != nil {
		d.params.log.Debug().Err(err).Msg("Incr")
		return 0, err
	}

	d.params.log.Debug().Str("method", "incr").
		Str("key", key).
		Str("value", val).
		Msg("Incr")

	value, err := d.utils.Convs().ToInt64(val)
	if err != nil {
		d.params.log.Debug().Err(err).Msg("Incr")
		return 0, err
	}

	value = value + 1

	err = d.WriteKeyVal(key, fmt.Sprintf("%d", value))
	if err != nil {
		d.params.log.Debug().Err(err).Msg("Incr")
		return 0, err
	}

	d.params.log.Debug().Str("method", "incr").
		Str("key", key).
		Int64("value", value).
		Msg("Incr")

	return value, nil
}

func (d *BuntDBLayer) Decr(key string) (int64, error) {
	val, err := d.GetVal(key)
	if err != nil {
		d.params.log.Debug().Err(err).Msg("Decr")
		return 0, err
	}

	d.params.log.Debug().Str("method", "decr").
		Str("key", key).
		Str("value", val).
		Msg("Decr")

	value, err := d.utils.Convs().ToInt64(val)
	if err != nil {
		d.params.log.Debug().Err(err).Msg("Decr")
		return 0, err
	}

	value = value - 1

	err = d.WriteKeyVal(key, fmt.Sprintf("%d", value))
	if err != nil {
		d.params.log.Debug().Err(err).Msg("Decr")
		return 0, err
	}

	d.params.log.Debug().Str("method", "decr").
		Str("key", key).
		Int64("value", value).
		Msg("Decr")

	return value, nil
}
