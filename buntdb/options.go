package buntdb

import (
	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/kind"
	"github.com/tidwall/buntdb"
)

type Options func(o *BuntDBParams) error

type BuntDBParams struct {
	db         *buntdb.DB
	file       string
	folder     string
	ttl        int
	log        zerolog.Logger
	driver     kind.Driver
	logDebug   bool
	logDisable bool
}

func newBuntDBParams(opts ...Options) (*BuntDBParams, error) {
	mts := &BuntDBParams{}
	for _, op := range opts {
		err := op(mts)
		if err != nil {
			return &BuntDBParams{}, err
		}
	}
	return mts, nil
}

func OptFile(file string) Options {
	return func(o *BuntDBParams) error {
		o.file = file
		return nil
	}
}

func OptFolder(folder string) Options {
	return func(o *BuntDBParams) error {
		o.folder = folder
		return nil
	}
}

func OptTTL(ttl int) Options {
	return func(o *BuntDBParams) error {
		o.ttl = ttl
		return nil
	}
}

func OptLog(log zerolog.Logger) Options {
	return func(o *BuntDBParams) error {
		o.log = log
		return nil
	}
}

func OptDriver(driver kind.Driver) Options {
	return func(o *BuntDBParams) error {
		o.driver = driver
		return nil
	}
}

func OptDB(db *buntdb.DB) Options {
	return func(o *BuntDBParams) error {
		o.db = db
		return nil
	}
}

func OptLogDebug(value bool) Options {
	return func(o *BuntDBParams) error {
		o.logDebug = value
		return nil
	}
}

func OptLogDisable(value bool) Options {
	return func(o *BuntDBParams) error {
		o.logDisable = value
		return nil
	}
}

// getters -----

func (o *BuntDBParams) GetFile() string {
	return o.file
}

func (o *BuntDBParams) GetFolder() string {
	return o.folder
}

func (o *BuntDBParams) GetTTL() int {
	return o.ttl
}

func (o *BuntDBParams) GetLogger() zerolog.Logger {
	return o.log
}

func (o *BuntDBParams) GetDriver() kind.Driver {
	return o.driver
}

func (o *BuntDBParams) GetDB() *buntdb.DB {
	return o.db
}

func (o *BuntDBParams) GetLogDebug() bool {
	return o.logDebug
}

func (o *BuntDBParams) GetLogDisable() bool {
	return o.logDisable
}

// setters -----

func (o *BuntDBParams) SetDB(db *buntdb.DB) {
	o.db = db
}

func (o *BuntDBParams) SetFile(file string) {
	o.file = file
}

func (o *BuntDBParams) SetFolder(folder string) {
	o.folder = folder
}

func (o *BuntDBParams) SetTTL(ttl int) {
	o.ttl = ttl
}

func (o *BuntDBParams) SetLogger(log zerolog.Logger) {
	o.log = log
}

func (o *BuntDBParams) SetDriver(driver kind.Driver) {
	o.driver = driver
}

func (o *BuntDBParams) SetLogDebug(value bool) {
	o.logDebug = value
}

func (o *BuntDBParams) SetLogDisable(value bool) {
	o.logDisable = value
}
