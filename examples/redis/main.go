package main

import (
	"fmt"

	"github.com/thiagozs/go-cache"
	"github.com/thiagozs/go-cache/kind"
)

func main() {
	opts := []cache.Options{
		cache.OptTTL(3000),
		cache.OptLogDebug(false),
		cache.OptLogDisable(false),
		cache.OptHost("localhost"),
		cache.OptPort(6379),
		cache.OptDriverKind(kind.REDIS),
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
	}

	v2, err := cache.GetVal("key2")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)

	v3, err := cache.GetVal("key3")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("v3:", v3)

	type data struct {
		Name string
	}

	d := data{
		Name: "Thiago",
	}
	if err := cache.WriteKeyValAsJSONTTL("key4", d, 2000); err != nil {
		fmt.Println("Error:", err)
	}

	v, err := cache.GetVal("key4")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("v4:", v)

	cache.WriteKeyVal("key5", "1")

	v5, err := cache.Incr("key5")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("v5:", v5)

	v6, err := cache.Decr("key5")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("v6:", v6)

}
