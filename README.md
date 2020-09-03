# hashicorp/raft example

This repository contains an example of using [Hashicorp Raft implementation v1.1](https://github.com/hashicorp/raft)
together with [BadgerDB](https://github.com/dgraph-io/badger).

Build the application with

```bash
go build
```

This should provide you with the `hashicorp-raft-example` binary. Assuming a *nix system, you should be able
to form a working cluster by running the following three commands in three different terminals: 

```bash
./hashicorp-raft-example -id=server_0 -datadir=/tmp/raft/server_0/data -raftdir=/tmp/raft/server_0/raft -http="127.0.0.1:9000" -raft="127.0.0.1:10000"
./hashicorp-raft-example -id=server_1 -datadir=/tmp/raft/server_1/data -raftdir=/tmp/raft/server_1/raft -http="127.0.0.1:9001" -raft="127.0.0.1:10001" -join="127.0.0.1:9000"
./hashicorp-raft-example -id=server_2 -datadir=/tmp/raft/server_2/data -raftdir=/tmp/raft/server_2/raft -http="127.0.0.1:9002" -raft="127.0.0.1:10002" -join="127.0.0.1:9000"
```

Here, `-datadir=...` specifies BadgerDB's data directory (provided to `badger.Open(...)`),
whereas `-raftdir=...` specifies the base directory of Raft's configuration, log and snapshot storage. Log and Config
directories are used by `raftbadger.NewBadgerStore(...)`; the snapshot directory is unrelated to BadgerDB and created by
[Raft.NewFileSnapshotStore()](https://pkg.go.dev/github.com/hashicorp/raft@v1.1.2?tab=doc#NewFileSnapshotStore).
