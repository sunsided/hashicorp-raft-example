# hashicorp/raft example

This repository contains an example of using [Hashicorp Raft v1.1](https://github.com/hashicorp/raft)
together with [BadgerDB](https://github.com/dgraph-io/badger).
It was forked from [aalda/hashicorp-raft-example](https://github.com/aalda/hashicorp-raft-example)
and updated for API changes, as well as project layout.

## Usage example 

Assuming the compiled `example` binary and a *nix system, you should be able
to form a working cluster by running the following three commands in three different terminals: 

```bash
./example -id=server_0 -http="127.0.0.1:9000" -raft="127.0.0.1:10000" -datadir=/tmp/raft/server_0/data -raftdir=/tmp/raft/server_0/raft
./example -id=server_1 -http="127.0.0.1:9001" -raft="127.0.0.1:10001" -join="127.0.0.1:9000" -datadir=/tmp/raft/server_1/data -raftdir=/tmp/raft/server_1/raft
./example -id=server_2 -http="127.0.0.1:9002" -raft="127.0.0.1:10002" -join="127.0.0.1:9000" -datadir=/tmp/raft/server_2/data -raftdir=/tmp/raft/server_2/raft
```

Here, `-datadir=...` specifies BadgerDB's data directory (provided to `badger.Open(...)`),
whereas `-raftdir=...` specifies the base directory of Raft's configuration, log and snapshot storage. Log and Config
directories are used by `raftbadger.NewBadgerStore(...)`; the snapshot directory is unrelated to BadgerDB and created by
[Raft.NewFileSnapshotStore()](https://pkg.go.dev/github.com/hashicorp/raft@v1.1.2?tab=doc#NewFileSnapshotStore).

To set keys, issue a `POST` to `/key` providing a JSON string map. The following example sets the keys 

- `answer` to value `42` and
- `foo` to value `bar`

```bash
curl --location --request POST 'localhost:9000/key/' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "answer": "41",
        "foo": "bar"
    }'
```

To get the values, issue a `GET` to `/key/{name}`, e.g.

```bash
curl --location --request GET 'localhost:9000/key/answer'
curl --location --request GET 'localhost:9000/key/foo'
``` 

To delete a key, issue a `DELETE` to `/key/{name}`, e.g.

```bash
curl --location --request DELETE 'localhost:9000/key/answer'
```

Note that while reading keys is possible on every node, creating and deleting keys is only possible on the leader
node. Attempting to mutate state on a follower will result in a 500 Internal Server Error; it's a bit crude,
but it gets the job done for an example.

## Project Layout

This project follows the _Standard Go Project Layout_ described [here](https://github.com/golang-standards/project-layout)
and [here](https://github.com/WeConnect/go-project-layout).

## Building

To get Go 1.15 on Ubuntu 20.04, run either:

```bash
sudo snap install --classic --channel=1.15/stable go
sudo snap refresh --classic --channel=1.15/stable go
```

### go build

Build the application with

```bash
go build -o example cmd/raft-example/main.go
```

### Dockerfile

To build a minimal image containing only the binary, run

```bash
docker build --target release -t raft-example .
```

or specify `--target release`. Note that this build is without compiler
optimizations and inlining in order to help debugging (with [Delve](https://github.com/go-delve/delve) in particular).

##### Docker-Compose

Using Docker Compose for testing (see [docker-compose.yaml](docker-compose.yaml)) is a bit fiddly, but it is possible
to get a configuration working by starting with

```bash
docker-compose up
``` 

in one terminal, then fiddling around with the other nodes in another terminal:

```bash
docker-compose stop node_1 node_2
docker-compose start node_1 node_2
```

To clean everything up, run

```
docker-compose stop
docker-compose rm -v
```

This may still leave the volume around, so use `docker volume ls` to spy for it. You can then delete it
using a command similar to the following:

```bash
docker volume rm hashicorp-raft-example_raft-data
```

#### Debugging with sources

To build a Docker image containing both the sources and the binary, then shell into it, run:

```bash
docker build --target dev-env -t raft-example .
docker run --rm -it --entrypoint sh raft-example
```

#### Debugging with Delve

To build for [Delve](https://github.com/go-delve/delve):

```bash
docker build --target debug -t raft-example .
```

You can then run the application e.g. like so (here, passing the `--help` command-line arguments).

```bash
docker run --rm -t -p 40000:40000 -t raft-example -- --help
```

From the host, connect using

```bash
dlv connect "localhost:40000"
```

To continue the application from there, run `continue` in the debugger.

## Tests and benchmarks

To run all tests, execute

```bash
go test ./...
```

To run the benchmarks, execute

```bash
go test ./... -run=XXX -bench=.
```