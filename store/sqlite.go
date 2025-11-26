package store

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(dir string) (*SQLiteStore, error) {
	dbPath := dir + "/kv.db"
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	// Create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS kv (key BLOB PRIMARY KEY, value BLOB)`)
	if err != nil {
		db.Close()
		return nil, err
	}

	// Optimize for performance (optional but recommended for benchmarks)
	// WAL mode usually provides better concurrency and performance
	_, err = db.Exec(`PRAGMA journal_mode = WAL; PRAGMA synchronous = NORMAL;`)
	if err != nil {
		db.Close()
		return nil, err
	}

	return &SQLiteStore{db: db}, nil
}

func (s *SQLiteStore) Put(key, value []byte) error {
	_, err := s.db.Exec(`INSERT OR REPLACE INTO kv (key, value) VALUES (?, ?)`, key, value)
	return err
}

func (s *SQLiteStore) Get(key []byte) ([]byte, error) {
	var value []byte
	err := s.db.QueryRow(`SELECT value FROM kv WHERE key = ?`, key).Scan(&value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (s *SQLiteStore) Close() error {
	return s.db.Close()
}
