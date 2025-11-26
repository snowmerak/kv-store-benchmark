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

> **Environment**: Windows 11, AMD Ryzen 7 8845HS w/ Radeon 780M Graphics, Go 1.25.4

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

- **Read Heavy**: Use **NutsDB** or **Bbolt**.
- **Write Heavy / Balanced**: Use **BadgerDB**.
- **SQL Support**: Use **SQLite**.
