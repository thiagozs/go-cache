package main

import (
	"fmt"
	"time"

	"github.com/thiagozs/go-cache"
	"github.com/thiagozs/go-cache/kind"
)

func main() {
	opts := []cache.Options{
		cache.OptTTL(3000),
		cache.OptLogDebug(true),
		cache.OptLogDisable(false),
		cache.OptTimeExpiration(time.Second * 60),
		cache.OptTimeCleanUpInt(time.Second * 120),
		cache.OptDriverKind(kind.GOCACHE),
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
