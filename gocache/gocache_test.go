package gocache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thiagozs/go-cache/kind"
)

func TestGocacheLayer(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{
			"NewMemory",
			func(t *testing.T) {
				_, err := NewMemory()
				assert.NoError(t, err)
			},
		},
		{
			"WriteKeyVal",
			func(t *testing.T) {
				gocache, err := NewMemory()
				assert.NoError(t, err)

				err = gocache.WriteKeyVal("key", "value")
				assert.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFunc)
	}
}

func TestGetVal(t *testing.T) {
	gocache, err := NewMemory()
	assert.NoError(t, err)

	gocache.WriteKeyVal("key", "value")
	val, err := gocache.GetVal("key")
	assert.NoError(t, err)
	assert.Equal(t, "value", val)
}

func TestDeleteKey(t *testing.T) {
	gocache, err := NewMemory()
	assert.NoError(t, err)

	gocache.WriteKeyVal("key", "value")
	_, err = gocache.DeleteKey("key")
	assert.NoError(t, err)

	_, err = gocache.GetVal("key")
	assert.Error(t, err)
}

func TestWriteKeyValTTL(t *testing.T) {
	gocache, err := NewMemory()
	assert.NoError(t, err)

	err = gocache.WriteKeyValTTL("key", "value", 1)
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)
	_, err = gocache.GetVal("key")
	assert.Error(t, err)
}

func TestWriteKeyValAsJSON(t *testing.T) {
	gocache, err := NewMemory()
	assert.NoError(t, err)

	data := map[string]string{"name": "Alice", "age": "30"}
	err = gocache.WriteKeyValAsJSON("key", data)
	assert.NoError(t, err)

	val, err := gocache.GetVal("key")
	assert.NoError(t, err)
	assert.JSONEq(t, `{"name":"Alice","age":"30"}`, val)
}

func TestWriteKeyValAsJSONTTL(t *testing.T) {
	gocache, err := NewMemory()
	assert.NoError(t, err)

	data := map[string]string{"name": "Bob", "age": "35"}
	err = gocache.WriteKeyValAsJSONTTL("key", data, 1)
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)
	_, err = gocache.GetVal("key")
	assert.Error(t, err)
}

func TestGetDriver(t *testing.T) {
	gocache, err := NewMemory(OptDriver(kind.GOCACHE))
	assert.NoError(t, err)

	driver := gocache.GetDriver()
	assert.Equal(t, kind.GOCACHE, driver)
}
