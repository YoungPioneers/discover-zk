// Copyright 2016
// Author: huangjunwei@youmi.net

package zookeeper

import (
	"errors"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const ()

var (
	// ErrClosedInstance 实例已关闭
	ErrClosedInstance = errors.New("Closed instance")
	// ErrNodesNeeded 未指定节点
	ErrNodesNeeded = errors.New("Nodes needed")
)

// ZKClient zookeeper client
type ZKClient struct {
	path  string
	addrs []string
	conn  *zk.Conn

	closed bool
	lock   *sync.RWMutex
}

// NewClient 初始化, path为server&&client共用路径,addrs为zookeepers ip
func NewClient(path string, addrs ...string) (client *ZKClient, err error) {
	if len(addrs) < 1 {
		// 通过TXT记录获取默认节点列表
		addrs, err = DefaultAddrs()
		if nil != err {
			return nil, err
		}
	}

	client = new(ZKClient)
	client.path = path
	client.closed = false
	client.lock = new(sync.RWMutex)

	err = client.connect(addrs)
	if nil != err {
		return nil, err
	}

	return
}

// connect 链接节点
func (zkClient *ZKClient) connect(addrs []string) (err error) {
	if len(addrs) < 1 {
		return ErrNodesNeeded
	}

	// 链接节点
	zkClient.conn, _, err = zk.Connect(addrs, time.Second)
	zkClient.conn.SetLogger(new(silent))
	if nil != err {
		return err
	}
	zkClient.addrs = addrs

	return nil
}

// ensurePath path不存在时创建
func (zkClient *ZKClient) ensurePath(path string) (err error) {
	slash := fmt.Sprintf("%c", filepath.Separator)
	tmpPath := slash
	directories := strings.Split(path, slash)

	for _, directory := range directories {
		tmpPath = filepath.Join(tmpPath, directory)

		exists, _, err := zkClient.conn.Exists(tmpPath)
		if nil != err {
			return err
		}

		if !exists {
			flags := int32(0)
			acl := zk.WorldACL(zk.PermAll)
			_, err = zkClient.conn.Create(tmpPath, nil, flags, acl)
		}
	}

	return err
}

// Register 节点注册, value为自定义信息(权重，开关等)
func (zkClient *ZKClient) Register(name string, value []byte) (err error) {
	zkClient.lock.RLock()
	defer zkClient.lock.RUnlock()

	if zkClient.closed {
		return ErrClosedInstance
	}

	// 确保使用的目录存在
	err = zkClient.ensurePath(zkClient.path)
	if nil != err {
		return err
	}

	// 创建ephemeral znodes 作临时节点
	flags := int32(zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)

	dest := filepath.Join(zkClient.path, name)

	_, err = zkClient.conn.Create(dest, value, flags, acl)
	if nil != err {
		return err
	}

	return nil
}

// Nodes 返回节点列表
func (zkClient *ZKClient) Nodes() (nodeValues map[string][]byte, err error) {
	zkClient.lock.RLock()
	defer zkClient.lock.RUnlock()

	if zkClient.closed {
		return nil, ErrClosedInstance
	}

	nodes, _, err := zkClient.conn.Children(zkClient.path)
	if nil != err {
		return nil, err
	}

	nodeValues, err = zkClient.getNodeValues(nodes)
	if nil != err {
		return nil, err
	}

	return nodeValues, nil
}

// Mirror 指定路径的snapshots chan, 每个snapshot为 map[string][]byte{"ip:port":"value","ip:port":"value"}
func (zkClient *ZKClient) Mirror() (snapshots chan map[string][]byte, errors chan error) {
	// TODO 同时返回启用开关、权重等信息
	zkClient.lock.RLock()
	defer zkClient.lock.RUnlock()

	snapshots = make(chan map[string][]byte)
	errors = make(chan error)
	go func() {
		for {
			// 变化后，获取所有节点名
			nodes, _, events, err := zkClient.conn.ChildrenW(zkClient.path)
			if nil != err {
				errors <- err
				continue
			}

			// 获取节点的值
			snapshot, err := zkClient.getNodeValues(nodes)
			if nil != err {
				errors <- err
				continue
			}

			snapshots <- snapshot
			event := <-events
			if nil != event.Err {
				errors <- err
				return
			}
		}
	}()

	return snapshots, errors
}

// getNodeValues 获取节点信息
func (zkClient *ZKClient) getNodeValues(nodes []string) (nodeValues map[string][]byte, err error) {
	nodeValues = make(map[string][]byte)
	for _, node := range nodes {
		path := filepath.Join(zkClient.path, node)
		bytes, _, err := zkClient.conn.Get(path)
		if nil != err {
			return nil, err
		}
		nodeValues[node] = bytes
	}
	return nodeValues, nil
}

// Close .
func (zkClient *ZKClient) Close() (err error) {
	zkClient.lock.Lock()
	defer zkClient.lock.Unlock()

	if zkClient.closed {
		return ErrClosedInstance
	}

	zkClient.conn.Close()
	zkClient.closed = true

	return nil
}
