package storage

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"os"
)

type LevelStorage struct {
	db   *leveldb.DB
	path string
}

func NewLevelStorage(path string, option *OpenOption) (*LevelStorage, error) {
	var dbOption opt.Options
	if option.ErrorIfExist() {
		dbOption.ErrorIfExist = true
	}

	db, err := leveldb.OpenFile(path, &dbOption)

	if err != nil {
		return nil, fmt.Errorf("unable to open storage: %w", err)
	}

	stor := LevelStorage{
		db,
		path,
	}

	return &stor, nil
}

func (l *LevelStorage) Put(key string, value interface{}) (err error) {
	data, err := encode(value)
	if err != nil {
		return fmt.Errorf("data encode error: %w", err)
	}

	err = l.db.Put([]byte(key), data, nil)
	if err != nil {
		return fmt.Errorf("storage put error: %w", err)
	}

	return nil
}

func (l *LevelStorage) Has(key string) (ret bool, err error) {
	ret, err = l.db.Has([]byte(key), nil)
	if err != nil {
		return false, fmt.Errorf("storage get error: %w", err)
	}
	return
}

func (l *LevelStorage) Delete(key string) (err error) {
	err = l.db.Delete([]byte(key), nil)

	if err != nil {
		return fmt.Errorf("storage delete error: %w", err)
	}

	return nil
}

func (l *LevelStorage) Get(key string, value interface{}) (err error) {
	data, err := l.db.Get([]byte(key), nil)
	if errors.Is(err, leveldb.ErrNotFound) {
		return ErrNotFound
	}

	if err != nil {
		return fmt.Errorf("storage get error: %w", err)
	}

	err = decode(data, value)

	if err != nil {
		return fmt.Errorf("data decode error: %w", err)
	}

	return nil
}

func (l *LevelStorage) Close() (err error) {
	err = l.db.Close()
	return
}

func (l *LevelStorage) Destroy() error {
	return os.RemoveAll(l.path)
}

func encode(obj interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(obj)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(data []byte, obj interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(obj)
	return err
}
