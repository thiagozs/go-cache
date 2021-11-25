# Go Cache 

Simple cache implementation **key**, **value** system.

The cache you can choose between drivers for different storage.

### Implementation
* [x] **BuntDB** 
* [x] **Redis**

## How to use the cache

Very simple.

```golang
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

```

A example of a cache implementation. You can find it in the `examples` folder.	

## Versioning and license

Our version numbers follow the [semantic versioning specification](http://semver.org/). You can see the available versions by checking the [tags on this repository](https://github.com/thiagozs/go-cache/tags). For more details about our license model, please take a look at the [LICENSE](LICENSE) file.

**2021**, thiagozs.