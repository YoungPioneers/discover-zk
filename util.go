// Copyright 2016
// Author: huangjunwei@youmi.net

package zookeeper

import (
	"net"
)

var (
	// defaultNodesTXT 默认使用内网dns的TXT记录获取zk节点列表
	defaultNodesTXT = "zookeepers.y.cn"
)

func DefaultNodesTXT() string {
	return defaultNodesTXT
}

func SetDefaultNodesTXT(txt string) {
	defaultNodesTXT = txt
}

func defaultNodes() (nodes []string, err error) {
	nodes, err = net.LookupTXT(defaultNodesTXT)
	if nil != err {
		return nil, err
	}

	return nodes, nil
}
