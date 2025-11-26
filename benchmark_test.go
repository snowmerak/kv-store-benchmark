package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"testing"

	"github.com/snowmerak/kv-store-benchmark/store"
)

const (
	MaxDataSize = 64 * 1024 // 64KB
)

var testValues [][]byte

func init() {
	testValues = make([][]byte, 100)
	for i := range testValues {
		testValues[i] = make([]byte, MaxDataSize)
		rand.Read(testValues[i])
	}
}

func benchmarkPut(b *testing.B, s store.KVStore) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := []byte(fmt.Sprintf("key-%09d", i))
		value := testValues[i%len(testValues)]
		if err := s.Put(key, value); err != nil {
			b.Fatal(err)
		}
	}
}

func benchmarkGet(b *testing.B, s store.KVStore) {
	// Pre-populate with enough keys for the benchmark or a fixed set
	// Since b.N can be large, we can't pre-populate b.N items easily without taking time.
	// We'll populate a fixed number of items and cycle through them.
	const numItems = 10000
	for i := 0; i < numItems; i++ {
		key := []byte(fmt.Sprintf("key-%09d", i))
		value := testValues[i%len(testValues)]
		if err := s.Put(key, value); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := []byte(fmt.Sprintf("key-%09d", i%numItems))
		_, err := s.Get(key)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func runBenchmark(b *testing.B, name string, setup func(dir string) (store.KVStore, error)) {
	// Put Benchmark
	b.Run("Put", func(b *testing.B) {
		dir, err := os.MkdirTemp("", "kv-bench-put-"+name)
		if err != nil {
			b.Fatal(err)
		}
		defer os.RemoveAll(dir)

		s, err := setup(dir)
		if err != nil {
			b.Fatal(err)
		}
		defer s.Close()

		benchmarkPut(b, s)
	})

	// Get Benchmark
	b.Run("Get", func(b *testing.B) {
		dir, err := os.MkdirTemp("", "kv-bench-get-"+name)
		if err != nil {
			b.Fatal(err)
		}
		defer os.RemoveAll(dir)

		s, err := setup(dir)
		if err != nil {
			b.Fatal(err)
		}
		defer s.Close()

		benchmarkGet(b, s)
	})
}

func BenchmarkBadger(b *testing.B) {
	runBenchmark(b, "Badger", func(dir string) (store.KVStore, error) {
		return store.NewBadgerStore(dir)
	})
}

func BenchmarkNutsDB(b *testing.B) {
	runBenchmark(b, "NutsDB", func(dir string) (store.KVStore, error) {
		return store.NewNutsDBStore(dir)
	})
}

func BenchmarkPebble(b *testing.B) {
	runBenchmark(b, "Pebble", func(dir string) (store.KVStore, error) {
		return store.NewPebbleStore(dir)
	})
}

func BenchmarkSQLite(b *testing.B) {
	runBenchmark(b, "SQLite", func(dir string) (store.KVStore, error) {
		return store.NewSQLiteStore(dir)
	})
}

func BenchmarkBbolt(b *testing.B) {
	runBenchmark(b, "Bbolt", func(dir string) (store.KVStore, error) {
		return store.NewBboltStore(dir)
	})
}

func BenchmarkLotusDB(b *testing.B) {
	runBenchmark(b, "LotusDB", func(dir string) (store.KVStore, error) {
		return store.NewLotusDBStore(dir)
	})
}

func TestDummy(t *testing.T) {
}
