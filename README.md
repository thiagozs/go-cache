# go-cache

A multi-driver cache library for Go with a unified key/value API. Switch between embedded, in-memory, or Redis storage by changing a single option — no code changes needed.

## Drivers

| Driver | Backend | Storage |
| ------ | ------- | ------- |
| `kind.BUNTDB` | [tidwall/buntdb](https://github.com/tidwall/buntdb) | File-based (embedded) |
| `kind.REDIS` | [redis/go-redis v9](https://github.com/redis/go-redis) | Redis server |
| `kind.GOCACHE` | [patrickmn/go-cache](https://github.com/patrickmn/go-cache) | In-memory |

## Installation

```bash
go get github.com/thiagozs/go-cache
```

## Quick start

```go
import (
    "github.com/thiagozs/go-cache"
    "github.com/thiagozs/go-cache/kind"
)
```

### BuntDB (file-based)

```go
c, err := cache.New(
    cache.OptDriver(kind.BUNTDB),
    cache.OptFolder("./data"),
    cache.OptFile("cache.db"),
)
```

### Redis

```go
c, err := cache.New(
    cache.OptDriver(kind.REDIS),
    cache.OptHost("localhost"),
    cache.OptPort(6379),
)
```

### GoCache (in-memory)

```go
c, err := cache.New(
    cache.OptDriver(kind.GOCACHE),
    cache.OptTimeExpiration(5 * time.Minute),
    cache.OptTimeCleanUpInt(10 * time.Minute),
)
```

## API

```go
// Write
c.WriteKeyVal(key, value string) error
c.WriteKeyValTTL(key, value string, ttlSeconds int) error
c.WriteKeyValAsJSON(key string, value any) error
c.WriteKeyValAsJSONTTL(key string, value any, ttlSeconds int) error

// Read / Delete
c.GetVal(key string) (string, error)
c.DeleteKey(key string) (string, error)

// Counters (atomic)
c.Incr(key string) (int64, error)
c.Decr(key string) (int64, error)

// Meta
c.GetDriver() kind.Driver
```

`WriteKeyValAsJSON` / `WriteKeyValAsJSONTTL` marshal any value to JSON before storing — the stored string is retrieved unchanged via `GetVal`.

`Incr` / `Decr` expect the stored value to be a numeric string (e.g. `"0"`). They are mutex-protected on BuntDB and GoCache; Redis uses its native `INCR`/`DECR` commands.

## Dependency injection

Pass `OptCache(repo)` to bypass driver construction and inject any `CacheRepo` implementation directly — useful for testing with mocks:

```go
ctrl := gomock.NewController(t)
mock := cache.NewMockCacheRepo(ctrl)

c, _ := cache.New(cache.OptCache(mock))
```

## Options reference

### Common (all drivers)

| Option | Description |
| ------ | ----------- |
| `OptDriver(kind.Driver)` | Selects the backend |
| `OptLogDebug(bool)` | Enable debug logging |
| `OptLogDisable(bool)` | Disable all logging |
| `OptCache(CacheRepo)` | Inject a custom implementation |

### BuntDB options

| Option | Description |
| ------ | ----------- |
| `OptFile(string)` | Database file name |
| `OptFolder(string)` | Directory for the database file |
| `OptTTL(int)` | Default TTL in seconds |

### Redis options

| Option | Description |
| ------ | ----------- |
| `OptHost(string)` | Server host |
| `OptPort(int)` | Server port |
| `OptPassword(string)` | Server password |
| `OptUser(string)` | Server username |
| `OptDatabase(int)` | Database index |
| `OptTTL(int)` | Default TTL in seconds |
| `OptRedis(*redis.Client)` | Inject an existing client |

### GoCache options

| Option | Description |
| ------ | ----------- |
| `OptTimeExpiration(time.Duration)` | Item expiration duration |
| `OptTimeCleanUpInt(time.Duration)` | Cleanup interval |

## Examples

Full working examples are in the [`examples/`](examples/) directory:

- [`examples/buntdb/`](examples/buntdb/main.go)
- [`examples/redis/`](examples/redis/main.go)
- [`examples/gocache/`](examples/gocache/main.go)

## Versioning and license

Version numbers follow [semantic versioning](http://semver.org/). See [tags](https://github.com/thiagozs/go-cache/tags) for available releases. License: [MIT](LICENSE).

**2026**, Thiago Zilli Sarmento. :heart:
