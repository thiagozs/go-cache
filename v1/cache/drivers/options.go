package drivers

type Options func(o *OptionsCfg) error

type OptionsCfg struct {
	fileName   string
	folder     string
	path       string
	ttl        int
	logDebug   bool
	logDisable bool
	host       string
	password   string
	user       string
	port       int
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
