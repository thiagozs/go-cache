default: help

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done


.PHONY: install-dep
install-dep: # Install dependencies
	# Install dependencies
	# --------------------
	# install ifacemaker...
	@go install github.com/vburenin/ifacemaker@latest
	# install mockgen...
	@go install github.com/golang/mock/mockgen@latest
	# Install deps complete +++

.PHONY: interfaces
interfaces: # Generate interfaces
	# Generate interfaces
	# -------------------
	# BuntDBLayer interfaces...
	@ifacemaker -f internal/buntdb/buntdb.go -s BuntDBLayer -i BuntDBLayerRepo -p buntdb -o internal/buntdb/repository.go
	# GocacheLayer interfaces...
	@ifacemaker -f internal/gocache/gocache.go -s GocacheLayer -i GocacheLayerRepo -p gocache -o internal/gocache/repository.go
	# RedisLayer interfaces...
	@ifacemaker -f internal/redis/redis.go -s RedisLayer -i RedisLayerRepo -p redis -o internal/redis/repository.go
	# Cache interfaces...
	@ifacemaker -f cache.go -s Cache -i CacheRepo -p cache -o repository.go

mocks: # Generate mocks
	# Generate mocks
	# --------------
	# BuntDBLayer mocks...
	@mockgen -source=internal/buntdb/repository.go -destination=internal/buntdb/buntdb_mock.go -package=buntdb
	# MemoryLayer mocks...
	@mockgen -source=internal/gocache/repository.go -destination=internal/gocache/gocache_mock.go -package=gocache
	# RedisLayer mocks...
	@mockgen -source=internal/redis/repository.go -destination=internal/redis/redis_mock.go -package=redis
	# Cache mocks...
	@mockgen -source=repository.go -destination=cache_mock.go -package=cache
