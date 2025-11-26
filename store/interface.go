package store

type KVStore interface {
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Close() error
}
