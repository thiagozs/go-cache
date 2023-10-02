// Code generated by ifacemaker; DO NOT EDIT.

package buntdb

import (
	kind "github.com/thiagozs/go-cache/kind"
)

// BuntDBLayerRepo ...
type BuntDBLayerRepo interface {
	GetVal(key string) (string, error)
	DeleteKey(key string) (string, error)
	WriteKeyVal(key string, val string) error
	WriteKeyValTTL(key string, val string, insec int) error
	WriteKeyValAsJSON(key string, val any) error
	WriteKeyValAsJSONTTL(key string, val any, insec int) error
	GetDriver() kind.Driver
}