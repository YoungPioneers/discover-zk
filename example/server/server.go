package main

import (
	"fmt"
	"scaling-zk"
	"time"
)

func main() {
	zkClient, err := zookeeper.NewClient("/example/nodes")
	if nil != err {
		fmt.Printf("zk client init failed. err: %s\n", err.Error())
		return
	}
	defer zkClient.Close()

	zkClient.Register("node1", []byte("value for node1"))

	for {
		// server runing
		fmt.Println("server is running")
		time.Sleep(10 * time.Second)
	}
}
