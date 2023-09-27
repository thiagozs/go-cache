package options

import (
	"time"

	"github.com/rs/zerolog"
)

type Options func(o *OptionsCfg) error

type OptionsCfg struct {
	fileName    string
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
}

func OptFileName(filename string) Options {
	return func(o *OptionsCfg) error {
		o.fileName = filename
		return nil
	}
}

func OptFolder(folder string) Options {
	return func(o *OptionsCfg) error {
		o.folder = folder
		return nil
	}
}

func OptPath(path string) Options {
	return func(o *OptionsCfg) error {
		o.path = path
		return nil
	}
}

func OptTTL(sec int) Options {
	return func(o *OptionsCfg) error {
		o.ttl = sec
		return nil
	}
}

func OptLogDebug(value bool) Options {
	return func(o *OptionsCfg) error {
		o.logDebug = value
		return nil
	}
}

func OptLogDisable(value bool) Options {
	return func(o *OptionsCfg) error {
		o.logDisable = value
		return nil
	}
}

func OptHost(host string) Options {
	return func(o *OptionsCfg) error {
		o.host = host
		return nil
	}
}

func OptPassword(password string) Options {
	return func(o *OptionsCfg) error {
		o.password = password
		return nil
	}
}

func OptUser(user string) Options {
	return func(o *OptionsCfg) error {
		o.user = user
		return nil
	}
}

func OptPort(port int) Options {
	return func(o *OptionsCfg) error {
		o.port = port
		return nil
	}
}

func OptLogger(log zerolog.Logger) Options {
	return func(o *OptionsCfg) error {
		o.log = log
		return nil
	}
}

func OptTimeExpiration(value time.Duration) Options {
	return func(o *OptionsCfg) error {
		o.tExpiration = value
		return nil
	}
}

func OptTimeCleanUpInt(value time.Duration) Options {
	return func(o *OptionsCfg) error {
		o.tCleanUpInt = value
		return nil
	}
}

func OptDatabase(value string) Options {
	return func(o *OptionsCfg) error {
		o.database = value
		return nil
	}
}

func OptVersion(value string) Options {
	return func(o *OptionsCfg) error {
		o.version = value
		return nil
	}
}

// ------------- getters

func (o *OptionsCfg) GetFileName() string {
	return o.fileName
}

func (o *OptionsCfg) GetFolder() string {
	return o.folder
}

func (o *OptionsCfg) GetPath() string {
	return o.path
}

func (o *OptionsCfg) GetTTL() int {
	return o.ttl
}

func (o *OptionsCfg) GetLogDebug() bool {
	return o.logDebug
}

func (o *OptionsCfg) GetLogDisable() bool {
	return o.logDisable
}

func (o *OptionsCfg) GetHost() string {
	return o.host
}

func (o *OptionsCfg) GetPassword() string {
	return o.password
}

func (o *OptionsCfg) GetUser() string {
	return o.user
}

func (o *OptionsCfg) GetPort() int {
	return o.port
}

func (o *OptionsCfg) GetLog() zerolog.Logger {
	return o.log
}

func (o *OptionsCfg) GetTExpiration() time.Duration {
	return o.tExpiration
}

func (o *OptionsCfg) GetTCleanUpInt() time.Duration {
	return o.tCleanUpInt
}

func (o *OptionsCfg) GetDatabase() string {
	return o.database
}

func (o *OptionsCfg) GetVersion() string {
	return o.version
}
