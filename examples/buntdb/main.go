package main

import (
	"fmt"

	"github.com/thiagozs/go-cache/v1/cache"
	"github.com/thiagozs/go-cache/v1/cache/drivers"
	"github.com/thiagozs/go-cache/v1/cache/options"
)

func main() {
	fmt.Println("Hello World!")

	opts := []options.Options{
		options.OptFolder("./tmp/cache"),
		options.OptFileName("cache.db"),
		options.OptTTL(3000),
		options.OptLogDebug(false),
		options.OptLogDisable(false),
	}

	cache, err := cache.New(drivers.BUNTDB, opts...)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cache.WriteKeyVal("key1", "value1")
	cache.WriteKeyVal("key2", "value2")

	v1, err := cache.GetVal("key1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v2, err := cache.GetVal("key2")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)

}
