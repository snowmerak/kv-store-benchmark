# KV Store Benchmark

This project benchmarks the performance of various Key-Value stores available in Go.

## Supported Stores

- **[BadgerDB v4](https://github.com/dgraph-io/badger)**: Fast key-value DB in Go.
- **[NutsDB](https://github.com/nutsdb/nutsdb)**: Simple, fast, embeddable, persistent key/value store.
- **[Pebble](https://github.com/cockroachdb/pebble)**: RocksDB compatible key-value store.
- **[SQLite](https://gitlab.com/cznic/sqlite)**: Cgo-free, port of SQLite.
- **[Bbolt](https://github.com/etcd-io/bbolt)**: An embedded key/value database for Go.

## Benchmark Scenario

- **Payload Size**: 64KB fixed size random data.
- **Operations**:
  - `Put`: Write random keys with 64KB values.
  - `Get`: Read keys.

## How to Run

```bash
go test -bench=. -benchmem -timeout 20m
```

## Benchmark Results

### Environment 1: macOS (Apple Silicon)
> **Specs**: macOS, Apple M3 Pro, Go 1.25.4

| Store | Operation | Time/Op | Description |
| :--- | :--- | :--- | :--- |
| **BadgerDB** | Put | **~92 µs** | Fastest Write |
| | Get | ~23 µs | Fast Read |
| **NutsDB** | Put | ~5,058 µs | Slow Write |
| | Get | **~4 µs** | **Fastest Read** |
| **Pebble** | Put | ~5,987 µs | Slow Write |
| | Get | ~21 µs | Fast Read |
| **SQLite** | Put | ~190 µs | Fast Write |
| | Get | ~30 µs | Balanced Read |
| **Bbolt** | Put | ~9,958 µs | Slowest Write |
| | Get | ~4 µs | **Fastest Read** |

### Environment 2: Windows (AMD)
> **Specs**: Windows 11, AMD Ryzen 7 8845HS w/ Radeon 780M Graphics, Go 1.25.4

| Store | Operation | Time/Op | Description |
| :--- | :--- | :--- | :--- |
| **BadgerDB** | Put | **~90 µs** | Fastest Write |
| | Get | ~53 µs | Balanced Read/Write |
| **NutsDB** | Put | ~479 µs | |
| | Get | **~6.6 µs** | **Fastest Read** |
| **Pebble** | Put | ~518 µs | |
| | Get | ~38 µs | Fast Read |
| **SQLite** | Put | ~515 µs | |
| | Get | ~129 µs | Slower Read (SQL overhead) |
| **Bbolt** | Put | ~1,356 µs | Slowest Write (Safety focused) |
| | Get | ~15 µs | Very Fast Read |

### Summary

- **Read Heavy**: **NutsDB** and **Bbolt** consistently provide the fastest read performance (~4-15 µs) across both platforms.
- **Write Heavy**: **BadgerDB** is the consistent winner for write operations (~90 µs). **SQLite** also showed strong write performance on macOS.
- **Balanced**: **BadgerDB** offers the best balance of read/write performance on both systems.
- **Note**: Write performance for NutsDB, Pebble, and Bbolt was significantly slower on macOS compared to Windows, likely due to differences in file system sync (fsync) behavior.
