package main

import (
	"context"
	"fmt"

	"github.com/pingcap/tidb/store/tikv"
)

func main() {
	run()
}

func run() error {
	driver := tikv.Driver{}
	store, err := driver.Open("tikv://localhost:2379")
	if err != nil {
		return err
	}

	txn, err := store.Begin()
	if err != nil {
		return err
	}

	txn.Set([]byte("k1"), []byte("v1"))
	txn.Commit(context.Background())

	txn2, err := store.Begin()
	if err != nil {
		return err
	}

	val, err := txn2.Get(context.Background(), []byte("k1"))
	if err != nil {
		return err
	}

	fmt.Printf("k1:\t %v\n", val)

	return nil
}
