# hashicorp/raft example

This repository contains an example of using [Hashicorp Raft implementation v1.1](https://github.com/hashicorp/raft)
together with [BadgerDB](https://github.com/dgraph-io/badger).

Build the application with

```bash
 go build -o example cmd/raft-example/main.go
```

This provides you with the `example` binary. Assuming a *nix system, you should be able
to form a working cluster by running the following three commands in three different terminals: 

```bash
./example -id=server_0 -datadir=/tmp/raft/server_0/data -raftdir=/tmp/raft/server_0/raft -http="127.0.0.1:9000" -raft="127.0.0.1:10000"
./example -id=server_1 -datadir=/tmp/raft/server_1/data -raftdir=/tmp/raft/server_1/raft -http="127.0.0.1:9001" -raft="127.0.0.1:10001" -join="127.0.0.1:9000"
./example -id=server_2 -datadir=/tmp/raft/server_2/data -raftdir=/tmp/raft/server_2/raft -http="127.0.0.1:9002" -raft="127.0.0.1:10002" -join="127.0.0.1:9000"
```

Here, `-datadir=...` specifies BadgerDB's data directory (provided to `badger.Open(...)`),
whereas `-raftdir=...` specifies the base directory of Raft's configuration, log and snapshot storage. Log and Config
directories are used by `raftbadger.NewBadgerStore(...)`; the snapshot directory is unrelated to BadgerDB and created by
[Raft.NewFileSnapshotStore()](https://pkg.go.dev/github.com/hashicorp/raft@v1.1.2?tab=doc#NewFileSnapshotStore).

## Project Layout

This project follows the _Standard Go Project Layout_ described [here](https://github.com/golang-standards/project-layout)
and [here](https://github.com/WeConnect/go-project-layout).

## Tests and benchmarks

To run all tests, execute

```bash
go test ./...
```

To run the benchmarks, execute

```bash
go test ./... -run=XXX -bench=.
```