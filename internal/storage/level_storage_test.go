package storage

import (
	"encoding/gob"
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestLevelStorage_GetAndPut(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "temp")
	stor, err := NewLevelStorage(filepath.Join(tmpDir, "test-level-storage"), nil)
	if !assert.Equal(t, nil, err, "error open storage") {
		t.Fatalf("error open storage: %v", err)
	}

	defer stor.Destroy()
	defer stor.Close()

	type ValueType struct {
		Number int
		Name   string
	}

	gob.Register(ValueType{})

	key := "key1"
	value := ValueType{32, "value1"}

	err = stor.Put(key, &value)
	if err != nil {
		t.Fatal(err)
	}

	var output ValueType
	err = stor.Get(key, &output)
	if err != nil {
		t.Fatal(err)
	}

	err = stor.Get("key2", &output)
	if !assert.True(t, errors.Is(err, ErrNotFound), "should trigger ErrNotFound") {
		t.Fatal("result not correct")
	}

	ret, err := stor.Has("key1")

	if !assert.Equal(t, true, ret, "should have this key") {
		t.Fatal("result not correct")
	}

}
