package cache

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/kind"
)

type Options func(d *CacheParams) error

type CacheParams struct {
	file        string
	folder      string
	path        string
	ttl         int
	logDebug    bool
	logDisable  bool
	host        string
	password    string
	user        string
	port        int
	log         zerolog.Logger
	tExpiration time.Duration
	tCleanUpInt time.Duration
	database    string
	version     string
	cache       CacheRepo
	driver      kind.Driver
}

func newCacheParams(opts ...Options) (*CacheParams, error) {
	mts := &CacheParams{}
	for _, op := range opts {
		err := op(mts)
		if err != nil {
			return &CacheParams{}, err
		}
	}
	return mts, nil
}

func OptFile(file string) Options {
	return func(d *CacheParams) error {
		d.file = file
		return nil
	}
}

func OptFolder(folder string) Options {
	return func(d *CacheParams) error {
		d.folder = folder
		return nil
	}
}

func OptPath(path string) Options {
	return func(d *CacheParams) error {
		d.path = path
		return nil
	}
}

func OptTTL(sec int) Options {
	return func(d *CacheParams) error {
		d.ttl = sec
		return nil
	}
}

func OptLogDebug(value bool) Options {
	return func(d *CacheParams) error {
		d.logDebug = value
		return nil
	}
}

func OptLogDisable(value bool) Options {
	return func(d *CacheParams) error {
		d.logDisable = value
		return nil
	}
}

func OptHost(host string) Options {
	return func(d *CacheParams) error {
		d.host = host
		return nil
	}
}

func OptPassword(password string) Options {
	return func(d *CacheParams) error {
		d.password = password
		return nil
	}
}

func OptUser(user string) Options {
	return func(d *CacheParams) error {
		d.user = user
		return nil
	}
}

func OptPort(port int) Options {
	return func(d *CacheParams) error {
		d.port = port
		return nil
	}
}

func OptLogger(log zerolog.Logger) Options {
	return func(d *CacheParams) error {
		d.log = log
		return nil
	}
}

func OptTimeExpiration(value time.Duration) Options {
	return func(d *CacheParams) error {
		d.tExpiration = value
		return nil
	}
}

func OptTimeCleanUpInt(value time.Duration) Options {
	return func(d *CacheParams) error {
		d.tCleanUpInt = value
		return nil
	}
}

func OptDatabase(value string) Options {
	return func(d *CacheParams) error {
		d.database = value
		return nil
	}
}

func OptVersion(value string) Options {
	return func(d *CacheParams) error {
		d.version = value
		return nil
	}
}

func OptCache(value CacheRepo) Options {
	return func(d *CacheParams) error {
		d.cache = value
		return nil
	}
}

func OptDriverKind(value kind.Driver) Options {
	return func(d *CacheParams) error {
		d.driver = value
		return nil
	}
}

//  getters -----

func (o *CacheParams) GetFile() string {
	return o.file
}

func (o *CacheParams) GetFolder() string {
	return o.folder
}

func (o *CacheParams) GetPath() string {
	return o.path
}

func (o *CacheParams) GetTTL() int {
	return o.ttl
}

func (o *CacheParams) GetLogDebug() bool {
	return o.logDebug
}

func (o *CacheParams) GetLogDisable() bool {
	return o.logDisable
}

func (o *CacheParams) GetHost() string {
	return o.host
}

func (o *CacheParams) GetPassword() string {
	return o.password
}

func (o *CacheParams) GetUser() string {
	return o.user
}

func (o *CacheParams) GetPort() int {
	return o.port
}

func (o *CacheParams) GetLog() zerolog.Logger {
	return o.log
}

func (o *CacheParams) GetTExpiration() time.Duration {
	return o.tExpiration
}

func (o *CacheParams) GetTCleanUpInt() time.Duration {
	return o.tCleanUpInt
}

func (o *CacheParams) GetDatabase() string {
	return o.database
}

func (o *CacheParams) GetVersion() string {
	return o.version
}

func (o *CacheParams) GetDriver() kind.Driver {
	return o.driver
}

func (o *CacheParams) GetCache() CacheRepo {
	return o.cache
}

// setters -----

func (o *CacheParams) SetFile(value string) {
	o.file = value
}

func (o *CacheParams) SetFolder(value string) {
	o.folder = value
}

func (o *CacheParams) SetPath(value string) {
	o.path = value
}

func (o *CacheParams) SetTTL(value int) {
	o.ttl = value
}

func (o *CacheParams) SetLogDebug(value bool) {
	o.logDebug = value
}

func (o *CacheParams) SetLogDisable(value bool) {
	o.logDisable = value
}

func (o *CacheParams) SetHost(value string) {
	o.host = value
}

func (o *CacheParams) SetPassword(value string) {
	o.password = value
}

func (o *CacheParams) SetUser(value string) {
	o.user = value
}

func (o *CacheParams) SetPort(value int) {
	o.port = value
}

func (o *CacheParams) SetLogger(value zerolog.Logger) {
	o.log = value
}

func (o *CacheParams) SetTExpiration(value time.Duration) {
	o.tExpiration = value
}

func (o *CacheParams) SetTCleanUpInt(value time.Duration) {
	o.tCleanUpInt = value
}

func (o *CacheParams) SetDatabase(value string) {
	o.database = value
}

func (o *CacheParams) SetVersion(value string) {
	o.version = value
}

func (o *CacheParams) SetDriver(value kind.Driver) {
	o.driver = value
}

func (o *CacheParams) SetCache(value CacheRepo) {
	o.cache = value
}
