package kind

type Driver int

const (
	BUNTDB Driver = iota
	REDIS
	GOCACHE
)

func (d Driver) String() string {
	return []string{"buntdb", "redis", "gocache"}[d]
}
