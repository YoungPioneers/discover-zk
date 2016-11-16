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

// DefaultAddrsTXT 获取默认TXT地址
func DefaultAddrsTXT() string {
	return defaultAddrsTXT
}

// SetDefaultAddrsTXT 设置TXT地址
func SetDefaultAddrsTXT(txt string) {
	defaultAddrsTXT = txt
}

// DefaultAddrs 默认ZK地址,通过TXT记录查询获取
func DefaultAddrs() (nodes []string, err error) {
	nodes, err = net.LookupTXT(defaultAddrsTXT)
	if nil != err {
		return nil, err
	}

	return nodes, nil
}

// slient 为禁用zk lib的输出
type silent struct {
}

// Printf slient 为禁用zk lib的输出
func (silent *silent) Printf(string, ...interface{}) {
	return
}
