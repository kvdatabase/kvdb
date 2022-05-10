![rosedb_ico.png](https://i.loli.net/2021/04/28/gIL2FXZcOesPmyD.png)

![](https://img.shields.io/github/license/kvdatabase/kvdb)&nbsp;[![Go Report Card](https://goreportcard.com/badge/github.com/kvdatabase/kvdb)&nbsp;](https://goreportcard.com/report/github.com/kvdatabase/kvdb)![GitHub top language](https://img.shields.io/github/languages/top/kvdatabase/kvdb)&nbsp;[![GitHub stars](https://img.shields.io/github/stars/kvdatabase/kvdb)&nbsp;](https://github.com/kvdatabase/kvdb/stargazers)[![codecov](https://codecov.io/gh/kvdatabase/kvdb/branch/main/graph/badge.svg?token=YZUB9QT6XF)](https://codecov.io/gh/kvdatabase/kvdb) [![CodeFactor](https://www.codefactor.io/repository/github/kvdatabase/kvdb/badge)](https://www.codefactor.io/repository/github/kvdatabase/kvdb) [![Go Reference](https://pkg.go.dev/badge/github.com/kvdatabase/kvdb.svg)](https://pkg.go.dev/github.com/kvdatabase/kvdb) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#database) 

[English](https://github.com/kvdatabase/kvdb#kvdb) | [简体中文](https://github.com/kvdatabase/kvdb/blob/main/README-CN.md)

### kvdb is not production-ready, we are doing more test now.

kvdb is a fast, stable and embedded key-value (k-v) storage engine based on `bitcask`. Its on-disk files are organized as WAL(Write Ahead Log) in LSM trees, optimizing for write throughput. 

kvdb supports a variety of data structures such as `string`, `list`, `hash`, `set` and `sorted set`.

kvdb is written in pure `Go`, easy to understand and use.

Our vision is to build an efficient k-v storage engine for Go-based applications. Feel free to give us any advices, and also please give us a star ✨，that helps a lot, thanks!

 [![Stargazers over time](https://starchart.cc/kvdatabase/kvdb.svg)](https://starchart.cc/kvdatabase/kvdb)      

## Features

* Supports all common data structures :  `string`, `list`, `hash`, `set`, `zset`.
* Easy to embed (`import "github.com/kvdatabase/kvdb"`).
* Low latency and high throughput (see Benchmark).
* Has built-in parallel execution of data modification on many provided data structures.
* Comes with kvdb-cli for command line access, that is compatible with redis-cli.
* Supports TTL-based key eviction.
* Supports prefix scan and range scan for string keys.
* Support simple transaction, ACID features.
* Merge operation can be stopped manually.

## Usage

### Cli example

Navigate to kvdb/cmd/server and run `main.go`

![Xnip2021-04-14_14-33-11.png](https://i.loli.net/2021/04/14/EsMFv48YB3P9j7k.png)

Open a new terminal, navigate to kvdb/cmd/cli, and run `main.go`：

![Xnip2021-04-14_14-35-50.png](https://i.loli.net/2021/04/14/9uh1ElVF3C4D6dM.png)

Or you can just use `redis-cli` or any other redis client：

![2021-05-14 上午11.19.24.png](https://i.loli.net/2021/05/14/eYkMyTzl5CXUN83.png)

### Embedding example

Import kvdb in the application:

```go
import "github.com/kvdatabase/kvdb"
```

Open a connection to the database:

```go
package main

import (
	"github.com/kvdatabase/kvdb"
	"log"
)

func main() {
	config := kvdb.DefaultConfig()
	db, err := kvdb.Open(config)
	
	if err != nil {
		log.Fatal(err)
	}
	
  // don`t forget to close!
	defer db.Close()
	
	//...
}
```

### Deploy and run on Docker

```shell
docker build -t="kvdb:v1.2.9" .
docker run --name=kvdb -itd -p 5200:5200 kvdb:v1.2.9
docker exec -it kvdb sh

$ kvdb-cli
127.0.0.1:5200>set hello kvdb
OK
127.0.0.1:5200>get hello
kvdb
```

## Command

### String

* Set
* SetNx
* Get
* GetSet
* Append
* StrLen
* StrExists
* StrRem
* PrefixScan
* RangeScan
* Expire
* Persist
* TTL

### List

* LPush
* RPush
* LPop
* RPop
* LIndex
* LRem
* LInsert
* LSet
* LTrim
* LRange
* LLen
* LKeyExists
* LValExists

### Hash

* HSet
* HSetNx
* HGet
* HGetAll
* HDel
* HExists
* HLen
* HKeys
* HVals

### Set

* SAdd
* SPop
* SIsMember
* SRandMember
* SRem
* SMove
* SCard
* SMembers
* SUnion
* SDiff

### Zset

* ZAdd
* ZScore
* ZCard
* ZRank
* ZRevRank
* ZIncrBy
* ZRange
* ZRangeWithScores
* ZRevRange
* ZRevRangeWithScores
* ZRem
* ZGetByRank
* ZRevGetByRank
* ZScoreRange
* ZRevScoreRange

## TODO

see [here](https://github.com/kvdatabase/kvdb/projects/1).

## Benchmark

### Environment

* Go version：1.14.4
* OS: macOS Catalina 10.15.7
* CPU: 2.6GHz 6-Core Intel Core i7
* Memory: 16 GB 2667 MHz DDR4
* Code: [kvdb-bench](https://github.com/kvdatabase/kvdb-bench)
* Other databases for comparison:
  * Badger
  * GoLevelDB
  * Pudge

### Result

**1,000,000 iterations**

```
go test -bench=. -benchtime=1000000x
badger 2021/05/16 21:59:53 INFO: All 0 tables opened in 0s
badger 2021/05/16 21:59:53 INFO: Discard stats nextEmptySlot: 0
badger 2021/05/16 21:59:53 INFO: Set nextTxnTs to 0
goos: darwin
goarch: amd64
pkg: kvdb-bench
BenchmarkPutValue_BadgerDB-12                    1000000             11518 ns/op            2110 B/op         46 allocs/op
BenchmarkGetValue_BadgerDB-12                    1000000              3547 ns/op            1172 B/op         20 allocs/op
BenchmarkPutValue_GoLevelDB-12                   1000000              4659 ns/op             352 B/op          9 allocs/op
BenchmarkGetValue_GoLevelDB-12                   1000000              2838 ns/op             814 B/op         13 allocs/op
BenchmarkPutValue_Pudge-12                       1000000              8512 ns/op             791 B/op         22 allocs/op
BenchmarkGetValue_Pudge-12                       1000000              1253 ns/op             200 B/op          6 allocs/op
BenchmarkPutValue_RoseDB_KeyValRam-12            1000000              4371 ns/op             566 B/op         11 allocs/op
BenchmarkGetValue_RoseDB_KeyValRam-12            1000000               481 ns/op              56 B/op          3 allocs/op
BenchmarkPutValue_RoseDB_KeyOnlyRam-12           1000000              4255 ns/op             566 B/op         11 allocs/op
BenchmarkGetValue_RoseDB_KeyOnlyRam-12           1000000              2986 ns/op             312 B/op          8 allocs/op
PASS
ok      kvdb-bench    46.388s
```

**2,500,000 iterations**

```
go test -bench=. -benchtime=2500000x
badger 2021/05/16 22:06:08 INFO: All 0 tables opened in 0s
badger 2021/05/16 22:06:08 INFO: Discard stats nextEmptySlot: 0
badger 2021/05/16 22:06:08 INFO: Set nextTxnTs to 0
goos: darwin
goarch: amd64
pkg: kvdb-bench
BenchmarkPutValue_BadgerDB-12                    2500000             11660 ns/op            2150 B/op         46 allocs/op
BenchmarkGetValue_BadgerDB-12                    2500000              4180 ns/op            1222 B/op         21 allocs/op
BenchmarkPutValue_GoLevelDB-12                   2500000              4637 ns/op             336 B/op          9 allocs/op
BenchmarkGetValue_GoLevelDB-12                   2500000              2942 ns/op             817 B/op         14 allocs/op
BenchmarkPutValue_Pudge-12                       2500000              9238 ns/op             763 B/op         22 allocs/op
BenchmarkGetValue_Pudge-12                       2500000              1275 ns/op             200 B/op          6 allocs/op
BenchmarkPutValue_RoseDB_KeyValRam-12            2500000              4474 ns/op             566 B/op         11 allocs/op
BenchmarkGetValue_RoseDB_KeyValRam-12            2500000               525 ns/op              56 B/op          3 allocs/op
BenchmarkPutValue_RoseDB_KeyOnlyRam-12           2500000              4294 ns/op             566 B/op         11 allocs/op
BenchmarkGetValue_RoseDB_KeyOnlyRam-12           2500000              3038 ns/op             312 B/op          8 allocs/op
PASS
ok      kvdb-bench    119.529s
```

### Conclusion

**Badger**

Stable read and write performance. Write: 11000+ ns/op. Read: 4000+ ns/op.

**GoLevelDB**

Write performance is almost 2.5x that of Badger. Read is almost 3000 ns/op, a little faster than that of Badger.

**Pudge**

Write performance is between that of GoLevelDB and Badger, almost 8500 ns/op. Read is very fast and stable, almost twice as fast as that of GoLevelDB.

**RoseDB**

Write performance is stable, almost the same as that of GoLevelDB, 2.5x that of Badger. 

In KeyValueRamMode, since the values are all in memory, it is the fastest of all.

In KeyOnlyRamMode, it is almost the same as GoLevelDB.

## Contributing

If you are intrested in contributing to kvdb, please see here: [CONTRIBUTING](https://github.com/kvdatabase/kvdb/blob/main/CONTRIBUTING.md)

## Contact me

If you have any questions, you can contact me by email: roseduan520@gmail.com

## License

kvdb is licensed under the term of the [MIT License](https://github.com/kvdatabase/kvdb/blob/main/LICENSE)

