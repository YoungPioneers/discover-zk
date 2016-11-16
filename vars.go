// Copyright 2016
// Author: huangjunwei@youmi.net

package zookeeper

// Addrs .
func (zkClient *ZKClient) Addrs() []string {
	zkClient.lock.RLock()
	defer zkClient.lock.RUnlock()

	return zkClient.addrs
}
