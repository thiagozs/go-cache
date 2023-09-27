package main

import (
	"fmt"

	"github.com/thiagozs/go-cache/v1/cache"
	"github.com/thiagozs/go-cache/v1/cache/drivers/kind"
	"github.com/thiagozs/go-cache/v1/cache/options"
)

func main() {
	opts := []options.Options{
		options.OptTTL(3000),
		options.OptLogDebug(false),
		options.OptLogDisable(false),
		options.OptHost("localhost"),
		options.OptPort(6379),
	}

	cache, err := cache.New(kind.REDIS, opts...)
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

}
