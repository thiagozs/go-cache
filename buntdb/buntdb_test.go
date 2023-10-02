package buntdb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thiagozs/go-cache/kind"
)

func TestBuntDBLayer(t *testing.T) {

	t.Run("NewBuntDB", func(t *testing.T) {
		_, err := NewBuntDB(OptFolder("/tmp/test"), OptFile("test.db"))
		assert.NoError(t, err)
	})

	buntDB, err := NewBuntDB(OptFolder("/tmp/test"), OptFile("test.db"))
	assert.NoError(t, err)

	t.Run("WriteKeyVal", func(t *testing.T) {
		err := buntDB.WriteKeyVal("key", "value")
		assert.NoError(t, err)
	})

	t.Run("GetVal", func(t *testing.T) {
		val, err := buntDB.GetVal("key")
		assert.NoError(t, err)
		assert.Equal(t, "value", val)
	})

	t.Run("DeleteKey", func(t *testing.T) {
		_, err := buntDB.DeleteKey("key")
		assert.NoError(t, err)

		_, err = buntDB.GetVal("key")
		assert.Error(t, err)
	})

	t.Run("WriteKeyValTTL", func(t *testing.T) {
		err := buntDB.WriteKeyValTTL("keyTTL", "value", 1)
		assert.NoError(t, err)

		time.Sleep(2 * time.Second)
		_, err = buntDB.GetVal("keyTTL")
		assert.Error(t, err)
	})

	t.Run("WriteKeyValAsJSON", func(t *testing.T) {
		data := map[string]string{"name": "Alice", "age": "30"}
		err := buntDB.WriteKeyValAsJSON("keyJSON", data)
		assert.NoError(t, err)

		val, err := buntDB.GetVal("keyJSON")
		assert.NoError(t, err)
		assert.JSONEq(t, `{"name":"Alice","age":"30"}`, val)
	})

	t.Run("WriteKeyValAsJSONTTL", func(t *testing.T) {
		data := map[string]string{"name": "Bob", "age": "35"}
		err := buntDB.WriteKeyValAsJSONTTL("keyJSONTTL", data, 1)
		assert.NoError(t, err)

		time.Sleep(2 * time.Second)
		_, err = buntDB.GetVal("keyJSONTTL")
		assert.Error(t, err)
	})

	t.Run("GetDriver", func(t *testing.T) {
		driver := buntDB.GetDriver()
		assert.Equal(t, kind.BUNTDB, driver)
	})
}
