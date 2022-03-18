package main

import (
	"fmt"
	"time"

	"github.com/thiagozs/go-cache/v1/cache"
	"github.com/thiagozs/go-cache/v1/cache/drivers/kind"
	"github.com/thiagozs/go-cache/v1/cache/options"
)

func main() {
	opts := []options.Options{
		options.OptTTL(3000),
		options.OptLogDebug(true),
		options.OptLogDisable(false),
		options.OptTimeExpiration(time.Second * 60),
		options.OptTimeCleanUpInt(time.Second * 120),
	}

	cache, err := cache.New(kind.GOCACHE, opts...)
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
