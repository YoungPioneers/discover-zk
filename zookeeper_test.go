// Copyright 2016
// Author: huangjunwei@youmi.net

package zookeeper

import (
	"testing"
)

func TestClient(t *testing.T) {
	zkClient, err := NewClient("/discovery_test", "127.0.0.1:2181")
	if nil != err {
		t.Errorf("zk client init failed. err: %s", err.Error())
	}
	defer zkClient.Close()

	// 检查节点
	exists, err := zkClient.Exists()
	if nil != err {
		t.Errorf("existence check failed. err: %s", err.Error())
	}
	if exists {
		t.Errorf("node identifier should not exist")
	}

	// 注册节点
	err = zkClient.Register("identifier", []byte("{'weight':10, 'work':true}"))
	if nil != err {
		t.Errorf("register nodes failed. err: %s", err.Error())
	}

	// 检查节点
	exists, err = zkClient.Exists()
	if nil != err {
		t.Errorf("existence check failed. err: %s", err.Error())
	}
	if !exists {
		t.Errorf("node identifier should exist")
	}

	// 获取节点
	nodes, err := zkClient.Nodes()
	if nil != err {
		t.Errorf("get nodes failed. err: %s", err.Error())
	}

	t.Logf("nodes: %+v", nodes)

	// 更新节点
	err = zkClient.Update([]byte("{'weight':10, 'work':false}"))
	if nil != err {
		t.Errorf("update node failed. err: %s", err.Error())
	}
}
