package raftKit

import (
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

var (
	// NewInmemStore
	/*
		PS:
		(1) 实现了接口: raft.LogStore、raft.StableStore
	*/
	NewInmemStore func() *raft.InmemStore = raft.NewInmemStore
)

// NewBoltStore
/*
	PS:
	(1) hashicorp提供了一个 raft-boltdb 来实现底层存储，它是一个嵌入式的数据库，能够 持久化 存储数据，我们直接用它来实现 LogStore 和 StableStore.
	(2) 实现了接口: raft.LogStore、raft.StableStore

	@param filePath string 存储路径（建议以".bolt"为文件后缀）
*/
func NewBoltStore(filePath string) (*raftboltdb.BoltStore, error) {
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}

	return raftboltdb.NewBoltStore(filePath)
}
