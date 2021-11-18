# Go-Cache - simple way to keep data

Simple cache implementation **key**, **value** system.
Core use for this projet are **BuntDB**

## How to use the cache

Very simple.

```golang
package main

import (
	"github.com/thiagozs/go-cache/v1/cache"
)

func main() {
	println("Hello, world!")

	cache, err := cache.New("./db", "db.db", 200, false)
	if err != nil {
		println(err)
	}
	if err := cache.WriteKeyVal("key", "value"); err != nil {
		println(err)
		return
	}

	rr, err := cache.GetVal("key")
	if err != nil {
		println(err)
		return
	}
	println(rr)
}

```

## Versioning and license

Our version numbers follow the [semantic versioning specification](http://semver.org/). You can see the available versions by checking the [tags on this repository](https://github.com/thiagozs/go-cache/tags). For more details about our license model, please take a look at the [LICENSE](LICENSE) file.

**2021**, thiagozs.