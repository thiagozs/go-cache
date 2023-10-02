package buntdb

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/thiagozs/go-cache/kind"
	"github.com/tidwall/buntdb"
)

func TestBuntDBOptions(t *testing.T) {
	opts, err := newBuntDBParams()
	assert.NoError(t, err)

	opts.SetFile("test.db")
	assert.Equal(t, "test.db", opts.GetFile())

	opts.SetFolder("/tmp")
	assert.Equal(t, "/tmp", opts.GetFolder())

	opts.SetTTL(60)
	assert.Equal(t, 60, opts.GetTTL())

	logger := zerolog.New(nil)
	opts.SetLogger(logger)
	assert.Equal(t, logger, opts.GetLogger())

	opts.SetDriver(kind.BUNTDB)
	assert.Equal(t, kind.BUNTDB, opts.GetDriver())

	db, err := buntdb.Open(":memory:")
	assert.NoError(t, err)
	opts.SetDB(db)
	assert.Equal(t, db, opts.GetDB())

	opts.SetLogDebug(true)
	assert.Equal(t, true, opts.GetLogDebug())

	opts.SetLogDisable(true)
	assert.Equal(t, true, opts.GetLogDisable())
}
