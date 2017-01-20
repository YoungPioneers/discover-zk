// Copyright 2016
// Author: huangjunwei@youmi.net

package zookeeper

import (
	"testing"
)

func TestClient(t *testing.T) {
	zkClient, err := NewClient("/discovery_test")
	if nil != err {
		t.Errorf("zk client init failed. err: %s", err.Error())
	}
	defer zkClient.Close()

	err = zkClient.Register("identifier", []byte("{'weight':10, 'work':true}"))
	if nil != err {
		t.Errorf("register nodes failed. err: %s", err.Error())
	}

	nodes, err := zkClient.Nodes()
	if nil != err {
		t.Errorf("get nodes failed. err: %s", err.Error())
	}

	t.Logf("nodes: %+v", nodes)

	err = zkClient.Update([]byte("{'weight':10, 'work':false}"))
	if nil != err {
		t.Errorf("update node failed. err: %s", err.Error())
	}
}
