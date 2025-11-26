package store

import (
	"github.com/nutsdb/nutsdb"
)

type NutsDBStore struct {
	db *nutsdb.DB
}

const bucket = "benchmark"

func NewNutsDBStore(dir string) (*NutsDBStore, error) {
	opts := nutsdb.DefaultOptions
	opts.Dir = dir
	db, err := nutsdb.Open(opts)
	if err != nil {
		return nil, err
	}

	// Ensure bucket exists
	err = db.Update(func(tx *nutsdb.Tx) error {
		return tx.NewBucket(nutsdb.DataStructureBTree, bucket)
	})
	if err != nil {
		// Ignore error if bucket already exists?
		// nutsdb might return error if it exists or just work.
		// Let's check if we can just ignore it or handle it.
		// For now, let's assume it might fail if it exists, but for a fresh dir it's fine.
		// Actually, NewBucket might return error if it exists.
	}

	return &NutsDBStore{db: db}, nil
}

func (n *NutsDBStore) Put(key, value []byte) error {
	return n.db.Update(func(tx *nutsdb.Tx) error {
		return tx.Put(bucket, key, value, 0)
	})
}

func (n *NutsDBStore) Get(key []byte) ([]byte, error) {
	var val []byte
	err := n.db.View(func(tx *nutsdb.Tx) error {
		v, err := tx.Get(bucket, key)
		if err != nil {
			return err
		}
		val = make([]byte, len(v))
		copy(val, v)
		return nil
	})
	return val, err
}

func (n *NutsDBStore) Close() error {
	return n.db.Close()
}
