// Copyright 2016
// Author: huangjunwei@youmi.net

package zookeeper

// Addrs .
func (zkClient *ZKClient) Addrs() []string {
	zkClient.lock.RLock()
	defer zkClient.lock.RUnlock()

	return zkClient.addrs
}

// Name .
func (zkClient *ZKClient) Name() string {
	zkClient.lock.RLock()
	defer zkClient.lock.RUnlock()

	return zkClient.name
}

// getVersion .
func (zkClient *ZKClient) getVersion() int32 {
	zkClient.versionLock.RLock()
	defer zkClient.versionLock.RUnlock()

	return zkClient.version
}

// setVersion .
func (zkClient *ZKClient) setVersion(version int32) {
	zkClient.versionLock.Lock()
	defer zkClient.versionLock.Unlock()

	zkClient.version = version
}
