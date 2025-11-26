package store

import (
	"github.com/lotusdblabs/lotusdb/v2"
)

type LotusDBStore struct {
	db *lotusdb.DB
}

func NewLotusDBStore(dir string) (*LotusDBStore, error) {
	options := lotusdb.DefaultOptions
	options.DirPath = dir
	db, err := lotusdb.Open(options)
	if err != nil {
		return nil, err
	}
	return &LotusDBStore{db: db}, nil
}

func (l *LotusDBStore) Put(key, value []byte) error {
	return l.db.Put(key, value)
}

func (l *LotusDBStore) Get(key []byte) ([]byte, error) {
	return l.db.Get(key)
}

func (l *LotusDBStore) Close() error {
	return l.db.Close()
}
