package raftbadger

import (
	"os"
	"testing"

	"github.com/hashicorp/raft/bench"
)

func BenchmarkBadgerStore_FirstIndex(b *testing.B) {
	store, path := testBadgerStore(b)
	defer os.Remove(path)
	defer store.Close()

	raftbench.FirstIndex(b, store)
}

func BenchmarkBadgerStore_LastIndex(b *testing.B) {
	store, path := testBadgerStore(b)
	defer os.Remove(path)
	defer store.Close()

	raftbench.LastIndex(b, store)
}

func BenchmarkBadgerStore_GetLog(b *testing.B) {
	store, path := testBadgerStore(b)
	defer os.Remove(path)
	defer store.Close()

	raftbench.GetLog(b, store)
}

func BenchmarkBadgerStore_StoreLog(b *testing.B) {
	store, path := testBadgerStore(b)
	defer os.Remove(path)
	defer store.Close()

	raftbench.StoreLog(b, store)
}

func BenchmarkBadgerStore_StoreLogs(b *testing.B) {
	store, path := testBadgerStore(b)
	defer os.Remove(path)
	defer store.Close()

	raftbench.StoreLogs(b, store)
}

func BenchmarkBadgerStore_DeleteRange(b *testing.B) {
	store, path := testBadgerStore(b)
	defer os.Remove(path)
	defer store.Close()

	raftbench.DeleteRange(b, store)
}

func BenchmarkBadgerStore_Set(b *testing.B) {
	store, path := testBadgerStore(b)
	defer os.Remove(path)
	defer store.Close()

	raftbench.Set(b, store)
}

func BenchmarkBadgerStore_Get(b *testing.B) {
	store, path := testBadgerStore(b)
	defer os.Remove(path)
	defer store.Close()

	raftbench.Get(b, store)
}

func BenchmarkBadgerStore_SetUint64(b *testing.B) {
	store, path := testBadgerStore(b)
	defer os.Remove(path)
	defer store.Close()

	raftbench.SetUint64(b, store)
}

func BenchmarkBadgerStore_GetUint64(b *testing.B) {
	store, path := testBadgerStore(b)
	defer os.Remove(path)
	defer store.Close()

	raftbench.GetUint64(b, store)
}
