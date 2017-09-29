Introduction
=======

discover-zk借助zookeeper实现服务发现。辅助在无proxy架构下实现自动扩容

Quick-start
-----------------

```
# 服务端
package main

import (
	"fmt"
	"github.com/YoungPioneers/discover-zk"
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

```

```
# 客户端
package main

import (
	"fmt"
	"scaling-zk"
)

func main() {
	zkClient, err := zookeeper.NewClient("/example/nodes")
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
```
