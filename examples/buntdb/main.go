package main

import (
	"fmt"

	"github.com/thiagozs/go-cache"
	"github.com/thiagozs/go-cache/kind"
)

func main() {
	opts := []cache.Options{
		cache.OptFolder("./tmp/cache"),
		cache.OptFile("cache.db"),
		cache.OptTTL(3000),
		cache.OptLogDebug(false),
		cache.OptLogDisable(false),
		cache.OptDriverKind(kind.BUNTDB),
	}

	cache, err := cache.New(opts...)
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
