package main

import (
	"fmt"
	"github.com/YoungPioneers/discover-zk"
)

func main() {
	zkClient, err := zookeeper.NewClient("/example/nodes", "127.0.0.1:2181")
	if nil != err {
		fmt.Printf("zk client init failed. err: %s\n", err.Error())
		return
	}
	defer zkClient.Close()

	nodes, err := zkClient.Nodes()
	if nil != err {
		fmt.Printf("get nodes failed. err: %s\n", err.Error())
		return

	}

	fmt.Printf("nodes: %+v\n", nodes)

	snapshots, errors := zkClient.Mirror()
	for {
		select {
		case snapshot := <-snapshots:
			fmt.Printf("snapshot: %+v\n", snapshot)
		case err := <-errors:
			fmt.Printf("zk error. err: %s\n", err.Error())
		}
	}
}
