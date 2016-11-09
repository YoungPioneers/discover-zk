// Copyright 2016
// Author: huangjunwei@youmi.net

package zookeeper

import (
	"net"
)

var (
	// defaultAddrsTXT 默认使用内网dns的TXT记录获取zk节点列表
	defaultAddrsTXT = "zookeepers.y.cn"
)

func DefaultAddrsTXT() string {
	return defaultAddrsTXT
}

func SetDefaultAddrsTXT(txt string) {
	defaultAddrsTXT = txt
}

func DefaultAddrs() (nodes []string, err error) {
	nodes, err = net.LookupTXT(defaultAddrsTXT)
	if nil != err {
		return nil, err
	}

	return nodes, nil
}

type silent struct {
}

func (silent *silent) Printf(string, ...interface{}) {
	return
}
