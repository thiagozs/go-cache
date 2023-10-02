package redis

import (
	redis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/thiagozs/go-cache/kind"
)

type Options func(o *OptionsParams) error

type OptionsParams struct {
	host       string
	password   string
	user       string
	database   int
	port       int
	ttl        int
	log        zerolog.Logger
	rdb        *redis.Client
	driver     kind.Driver
	logDebug   bool
	logDisable bool
}

func newRedisParams(opts ...Options) (*OptionsParams, error) {
	mts := &OptionsParams{}
	for _, op := range opts {
		err := op(mts)
		if err != nil {
			return &OptionsParams{}, err
		}
	}
	return mts, nil
}

func OptDatabase(database int) Options {
	return func(o *OptionsParams) error {
		o.database = database
		return nil
	}
}

func OptHost(host string) Options {
	return func(o *OptionsParams) error {
		o.host = host
		return nil
	}
}

func OptPassword(password string) Options {
	return func(o *OptionsParams) error {
		o.password = password
		return nil
	}
}

func OptUser(user string) Options {
	return func(o *OptionsParams) error {
		o.user = user
		return nil
	}
}

func OptPort(port int) Options {
	return func(o *OptionsParams) error {
		o.port = port
		return nil
	}
}

func OptLogger(log zerolog.Logger) Options {
	return func(o *OptionsParams) error {
		o.log = log
		return nil
	}
}

func OptTTL(sec int) Options {
	return func(o *OptionsParams) error {
		o.ttl = sec
		return nil
	}
}

func OptDriver(driver kind.Driver) Options {
	return func(o *OptionsParams) error {
		o.driver = driver
		return nil
	}
}

func OptRedis(rdb *redis.Client) Options {
	return func(o *OptionsParams) error {
		o.rdb = rdb
		return nil
	}
}

func OptLogDebug(value bool) Options {
	return func(o *OptionsParams) error {
		o.logDebug = value
		return nil
	}
}

func OptLogDisable(value bool) Options {
	return func(o *OptionsParams) error {
		o.logDisable = value
		return nil
	}
}

// getters -----

func (o *OptionsParams) GetDatabase() int {
	return o.database
}

func (o *OptionsParams) GetHost() string {
	return o.host
}

func (o *OptionsParams) GetPassword() string {
	return o.password
}

func (o *OptionsParams) GetUser() string {
	return o.user
}

func (o *OptionsParams) GetPort() int {
	return o.port
}

func (o *OptionsParams) GetTTL() int {
	return o.ttl
}

func (o *OptionsParams) GetLogger() zerolog.Logger {
	return o.log
}

func (o *OptionsParams) GetDriver() kind.Driver {
	return o.driver
}

func (o *OptionsParams) GetRedis() *redis.Client {
	return o.rdb
}

func (o *OptionsParams) GetLogDebug() bool {
	return o.logDebug
}

func (o *OptionsParams) GetLogDisable() bool {
	return o.logDisable
}

// setters -----

func (o *OptionsParams) SetDatabase(database int) {
	o.database = database
}

func (o *OptionsParams) SetRedis(rdb *redis.Client) {
	o.rdb = rdb
}

func (o *OptionsParams) SetLogger(log zerolog.Logger) {
	o.log = log
}

func (o *OptionsParams) SetTTL(sec int) {
	o.ttl = sec
}

func (o *OptionsParams) SetHost(host string) {
	o.host = host
}

func (o *OptionsParams) SetPassword(password string) {
	o.password = password
}

func (o *OptionsParams) SetUser(user string) {
	o.user = user
}

func (o *OptionsParams) SetPort(port int) {
	o.port = port
}

func (o *OptionsParams) SetDriver(driver kind.Driver) {
	o.driver = driver
}

func (o *OptionsParams) SetLogDebug(value bool) {
	o.logDebug = value
}

func (o *OptionsParams) SetLogDisable(value bool) {
	o.logDisable = value
}

func (o *OptionsParams) SetOptions(opts ...Options) error {
	for _, op := range opts {
		err := op(o)
		if err != nil {
			return err
		}
	}
	return nil
}

// options -----

func (o *OptionsParams) GetOptions() *OptionsParams {
	return o
}

func (o *OptionsParams) GetRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     o.GetHost(),
		Password: o.GetPassword(),
		DB:       0,
	}
}

func (o *OptionsParams) GetRedisOptionsWithUser() *redis.Options {
	return &redis.Options{
		Addr:     o.GetHost(),
		Password: o.GetPassword(),
		DB:       0,
		Username: o.GetUser(),
	}
}

func (o *OptionsParams) GetRedisOptionsWithDB() *redis.Options {
	return &redis.Options{
		Addr:     o.GetHost(),
		Password: o.GetPassword(),
		DB:       0,
	}
}
