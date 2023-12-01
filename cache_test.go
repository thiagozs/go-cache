package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagozs/go-cache/buntdb"
	"github.com/thiagozs/go-cache/gocache"
	"github.com/thiagozs/go-cache/kind"
	"github.com/thiagozs/go-cache/redis"

	"github.com/golang/mock/gomock"
)

func TestNewDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRedis := redis.NewMockRedisLayerRepo(ctrl)
	mockGocache := gocache.NewMockGocacheLayerRepo(ctrl)
	mockBuntdb := buntdb.NewMockBuntDBLayerRepo(ctrl)

	tests := []struct {
		name        string
		driver      kind.Driver
		expectedErr error
	}{
		{
			name:        "Success Redis",
			driver:      kind.REDIS,
			expectedErr: nil,
		},
		{
			name:        "Success GoCache",
			driver:      kind.GOCACHE,
			expectedErr: nil,
		},
		{
			name:        "Success BuntDB",
			driver:      kind.BUNTDB,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		opts := []Options{
			OptLogDebug(false),
			OptLogDisable(false),
		}

		t.Run(tt.name, func(t *testing.T) {
			switch tt.driver {
			case kind.REDIS:
				mockRedis.EXPECT().GetDriver().Return(kind.REDIS).AnyTimes()

				opts = append(opts, OptTTL(10))
				opts = append(opts, OptCache(mockRedis))

			case kind.GOCACHE:
				mockGocache.EXPECT().GetDriver().Return(kind.GOCACHE).AnyTimes()

				opts = append(opts, OptTTL(10))
				opts = append(opts, OptCache(mockGocache))

			case kind.BUNTDB:
				mockBuntdb.EXPECT().GetDriver().Return(kind.BUNTDB).AnyTimes()

				opts = append(opts, OptFile("test.txt"))
				opts = append(opts, OptFolder("/tmp/test"))
				opts = append(opts, OptTTL(10))
				opts = append(opts, OptCache(mockBuntdb))
			}

			_, err := New(opts...)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCacheGetVal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		expected string
		mock     any
	}{
		{
			name:     "Redis",
			expected: "value1",
			mock:     redis.NewMockRedisLayerRepo(ctrl),
		},
		{
			name:     "GoCache",
			expected: "value2",
			mock:     gocache.NewMockGocacheLayerRepo(ctrl),
		},
		{
			name:     "BuntDB",
			expected: "value3",
			mock:     buntdb.NewMockBuntDBLayerRepo(ctrl),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.mock.(type) {
			case *redis.MockRedisLayerRepo:
				v.EXPECT().GetVal("key").Return(tt.expected, nil).Times(1)
			case *gocache.MockGocacheLayerRepo:
				v.EXPECT().GetVal("key").Return(tt.expected, nil).Times(1)
			case *buntdb.MockBuntDBLayerRepo:
				v.EXPECT().GetVal("key").Return(tt.expected, nil).Times(1)
			}

			cache, err := New(OptCache(tt.mock.(CacheRepo)))
			assert.NoError(t, err)

			value, err := cache.GetVal("key")
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestCacheWriteKeyVal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		mock     any
		key      string
		value    string
		expected error
	}{
		{
			name:     "Redis",
			mock:     redis.NewMockRedisLayerRepo(ctrl),
			key:      "key1",
			value:    "value1",
			expected: nil,
		},
		{
			name:     "GoCache",
			mock:     gocache.NewMockGocacheLayerRepo(ctrl),
			key:      "key2",
			value:    "value2",
			expected: nil,
		},
		{
			name:     "BuntDB",
			mock:     buntdb.NewMockBuntDBLayerRepo(ctrl),
			key:      "key3",
			value:    "value3",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.mock.(type) {
			case *redis.MockRedisLayerRepo:
				v.EXPECT().WriteKeyVal(tt.key, tt.value).Return(tt.expected).Times(1)
			case *gocache.MockGocacheLayerRepo:
				v.EXPECT().WriteKeyVal(tt.key, tt.value).Return(tt.expected).Times(1)
			case *buntdb.MockBuntDBLayerRepo:
				v.EXPECT().WriteKeyVal(tt.key, tt.value).Return(tt.expected).Times(1)
			default:
				t.Fatalf("unexpected type %T", v)
			}

			cacheRepo, ok := tt.mock.(CacheRepo)
			if !ok {
				t.Fatalf("could not assert %T to CacheRepo", tt.mock)
			}

			cache, err := New(OptCache(cacheRepo))
			assert.NoError(t, err)

			err = cache.WriteKeyVal(tt.key, tt.value)
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestCacheWriteKeyValTTL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		mock     interface{}
		key      string
		value    string
		ttl      int
		expected error
	}{
		{
			name:     "Redis",
			mock:     redis.NewMockRedisLayerRepo(ctrl),
			key:      "key1",
			value:    "value1",
			ttl:      10,
			expected: nil,
		},
		{
			name:     "GoCache",
			mock:     gocache.NewMockGocacheLayerRepo(ctrl),
			key:      "key2",
			value:    "value2",
			ttl:      20,
			expected: nil,
		},
		{
			name:     "BuntDB",
			mock:     buntdb.NewMockBuntDBLayerRepo(ctrl),
			key:      "key3",
			value:    "value3",
			ttl:      30,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.mock.(type) {
			case *redis.MockRedisLayerRepo:
				v.EXPECT().WriteKeyValTTL(tt.key, tt.value, tt.ttl).Return(tt.expected).Times(1)
			case *gocache.MockGocacheLayerRepo:
				v.EXPECT().WriteKeyValTTL(tt.key, tt.value, tt.ttl).Return(tt.expected).Times(1)
			case *buntdb.MockBuntDBLayerRepo:
				v.EXPECT().WriteKeyValTTL(tt.key, tt.value, tt.ttl).Return(tt.expected).Times(1)
			default:
				t.Fatalf("unexpected type %T", v)
			}

			cacheRepo, ok := tt.mock.(CacheRepo)
			if !ok {
				t.Fatalf("could not assert %T to CacheRepo", tt.mock)
			}

			cache, err := New(OptCache(cacheRepo))
			assert.NoError(t, err)

			err = cache.WriteKeyValTTL(tt.key, tt.value, tt.ttl)
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestCacheDeleteKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		mock     any
		key      string
		expected error
	}{
		{
			name:     "Redis",
			mock:     redis.NewMockRedisLayerRepo(ctrl),
			key:      "key1",
			expected: nil,
		},
		{
			name:     "GoCache",
			mock:     gocache.NewMockGocacheLayerRepo(ctrl),
			key:      "key2",
			expected: nil,
		},
		{
			name:     "BuntDB",
			mock:     buntdb.NewMockBuntDBLayerRepo(ctrl),
			key:      "key3",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.mock.(type) {
			case *redis.MockRedisLayerRepo:
				v.EXPECT().DeleteKey(tt.key).Return("1", tt.expected).Times(1)
			case *gocache.MockGocacheLayerRepo:
				v.EXPECT().DeleteKey(tt.key).Return("1", tt.expected).Times(1)
			case *buntdb.MockBuntDBLayerRepo:
				v.EXPECT().DeleteKey(tt.key).Return("1", tt.expected).Times(1)
			}

			cache, err := New(OptCache(tt.mock.(CacheRepo)))
			assert.NoError(t, err)

			_, err = cache.DeleteKey(tt.key)
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestCacheWriteKeyValAsJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type MyStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	tests := []struct {
		name     string
		mock     any
		key      string
		value    interface{}
		expected error
	}{
		{
			name:     "Redis",
			mock:     redis.NewMockRedisLayerRepo(ctrl),
			key:      "key1",
			value:    &MyStruct{Name: "Alice", Value: 42},
			expected: nil,
		},
		{
			name:     "GoCache",
			mock:     gocache.NewMockGocacheLayerRepo(ctrl),
			key:      "key2",
			value:    &MyStruct{Name: "Bob", Value: 43},
			expected: nil,
		},
		{
			name:     "BuntDB",
			mock:     buntdb.NewMockBuntDBLayerRepo(ctrl),
			key:      "key3",
			value:    &MyStruct{Name: "Charlie", Value: 44},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.mock.(type) {
			case *redis.MockRedisLayerRepo:
				v.EXPECT().WriteKeyValAsJSON(tt.key, tt.value).Return(tt.expected).Times(1)
			case *gocache.MockGocacheLayerRepo:
				v.EXPECT().WriteKeyValAsJSON(tt.key, tt.value).Return(tt.expected).Times(1)
			case *buntdb.MockBuntDBLayerRepo:
				v.EXPECT().WriteKeyValAsJSON(tt.key, tt.value).Return(tt.expected).Times(1)
			}

			cache, err := New(OptCache(tt.mock.(CacheRepo)))
			assert.NoError(t, err)

			err = cache.WriteKeyValAsJSON(tt.key, tt.value)
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestCacheWriteKeyValAsJSONTTL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type MyStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	tests := []struct {
		name     string
		mock     any
		key      string
		value    interface{}
		ttl      int
		expected error
	}{
		{
			name:     "Redis",
			mock:     redis.NewMockRedisLayerRepo(ctrl),
			key:      "key1",
			value:    &MyStruct{Name: "Alice", Value: 42},
			ttl:      60,
			expected: nil,
		},
		{
			name:     "GoCache",
			mock:     gocache.NewMockGocacheLayerRepo(ctrl),
			key:      "key2",
			value:    &MyStruct{Name: "Bob", Value: 43},
			ttl:      120,
			expected: nil,
		},
		{
			name:     "BuntDB",
			mock:     buntdb.NewMockBuntDBLayerRepo(ctrl),
			key:      "key3",
			value:    &MyStruct{Name: "Charlie", Value: 44},
			ttl:      180,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.mock.(type) {
			case *redis.MockRedisLayerRepo:
				v.EXPECT().WriteKeyValAsJSONTTL(tt.key, tt.value, tt.ttl).Return(tt.expected).Times(1)
			case *gocache.MockGocacheLayerRepo:
				v.EXPECT().WriteKeyValAsJSONTTL(tt.key, tt.value, tt.ttl).Return(tt.expected).Times(1)
			case *buntdb.MockBuntDBLayerRepo:
				v.EXPECT().WriteKeyValAsJSONTTL(tt.key, tt.value, tt.ttl).Return(tt.expected).Times(1)
			}

			cache, err := New(OptCache(tt.mock.(CacheRepo)))
			assert.NoError(t, err)

			err = cache.WriteKeyValAsJSONTTL(tt.key, tt.value, tt.ttl)
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestCacheIncr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		mock     any
		key      string
		expected int64
	}{
		{
			name:     "Redis",
			mock:     redis.NewMockRedisLayerRepo(ctrl),
			key:      "key1",
			expected: 1,
		},
		{
			name:     "GoCache",
			mock:     gocache.NewMockGocacheLayerRepo(ctrl),
			key:      "key2",
			expected: 2,
		},
		{
			name:     "BuntDB",
			mock:     buntdb.NewMockBuntDBLayerRepo(ctrl),
			key:      "key3",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.mock.(type) {
			case *redis.MockRedisLayerRepo:
				v.EXPECT().Incr(tt.key).Return(tt.expected, nil).Times(1)
			case *gocache.MockGocacheLayerRepo:
				v.EXPECT().Incr(tt.key).Return(tt.expected, nil).Times(1)
			case *buntdb.MockBuntDBLayerRepo:
				v.EXPECT().Incr(tt.key).Return(tt.expected, nil).Times(1)
			}

			cache, err := New(OptCache(tt.mock.(CacheRepo)))
			assert.NoError(t, err)

			value, err := cache.Incr(tt.key)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestCacheDecr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		mock     any
		key      string
		expected int64
	}{
		{
			name:     "Redis",
			mock:     redis.NewMockRedisLayerRepo(ctrl),
			key:      "key1",
			expected: 1,
		},
		{
			name:     "GoCache",
			mock:     gocache.NewMockGocacheLayerRepo(ctrl),
			key:      "key2",
			expected: 2,
		},
		{
			name:     "BuntDB",
			mock:     buntdb.NewMockBuntDBLayerRepo(ctrl),
			key:      "key3",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.mock.(type) {
			case *redis.MockRedisLayerRepo:
				v.EXPECT().Decr(tt.key).Return(tt.expected, nil).Times(1)
			case *gocache.MockGocacheLayerRepo:
				v.EXPECT().Decr(tt.key).Return(tt.expected, nil).Times(1)
			case *buntdb.MockBuntDBLayerRepo:
				v.EXPECT().Decr(tt.key).Return(tt.expected, nil).Times(1)
			}

			cache, err := New(OptCache(tt.mock.(CacheRepo)))
			assert.NoError(t, err)

			value, err := cache.Decr(tt.key)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}
