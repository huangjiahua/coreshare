package storage

const (
	ErrorIfExist      OpenOption = "Error If Exist"
	DefaultOpenOption OpenOption = "Default"
)

const (
	ErrNotFound KVStorageError = "Not Found"
)

type KVStorage interface {
	Put(key string, value interface{}) (err error)
	Get(key string, value interface{}) (err error)
	Has(key string) (ret bool, err error)
	Delete(key string) (err error)
	Close() (err error)
}

type OpenOption string

func (opt *OpenOption) ErrorIfExist() bool {
	return opt != nil && *opt == ErrorIfExist
}

type KVStorageError string

func (e KVStorageError) Error() string {
	return string(e)
}
