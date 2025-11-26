package store

import (
	"github.com/cockroachdb/pebble"
)

type PebbleStore struct {
	db *pebble.DB
}

func NewPebbleStore(dir string) (*PebbleStore, error) {
	db, err := pebble.Open(dir, &pebble.Options{})
	if err != nil {
		return nil, err
	}
	return &PebbleStore{db: db}, nil
}

func (p *PebbleStore) Put(key, value []byte) error {
	return p.db.Set(key, value, pebble.Sync)
}

func (p *PebbleStore) Get(key []byte) ([]byte, error) {
	val, closer, err := p.db.Get(key)
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	// Make a copy because closer.Close() invalidates val
	v := make([]byte, len(val))
	copy(v, val)
	return v, nil
}

func (p *PebbleStore) Close() error {
	return p.db.Close()
}
