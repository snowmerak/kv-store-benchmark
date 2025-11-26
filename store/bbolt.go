package store

import (
	"go.etcd.io/bbolt"
)

type BboltStore struct {
	db *bbolt.DB
}

const bboltBucket = "benchmark"

func NewBboltStore(dir string) (*BboltStore, error) {
	dbPath := dir + "/kv.db"
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bboltBucket))
		return err
	})
	if err != nil {
		db.Close()
		return nil, err
	}

	return &BboltStore{db: db}, nil
}

func (b *BboltStore) Put(key, value []byte) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bboltBucket))
		return bucket.Put(key, value)
	})
}

func (b *BboltStore) Get(key []byte) ([]byte, error) {
	var val []byte
	err := b.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bboltBucket))
		v := bucket.Get(key)
		if v != nil {
			val = make([]byte, len(v))
			copy(val, v)
		}
		return nil
	})
	return val, err
}

func (b *BboltStore) Close() error {
	return b.db.Close()
}
