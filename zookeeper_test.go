// Copyright 2016
// Author: huangjunwei@youmi.net

package zookeeper

import (
	"testing"
)

func TestClient(t *testing.T) {
	zkClient, err := NewClient("/tagocean/fetchers")
	if nil != err {
		t.Errorf("zk client init failed. err: %s", err.Error())
	}

	err = zkClient.Register("127.0.0.1:18014", []byte("{'weight':10, 'work':true}"))
	if nil != err {
		t.Errorf("register nodes failed. err: %s", err.Error())
	}

	nodes, err := zkClient.Nodes()
	if nil != err {
		t.Errorf("get nodes failed. err: %s", err.Error())
	}

	t.Logf("nodes: %+v", nodes)
}
