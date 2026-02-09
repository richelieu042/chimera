package pushKit

import (
	"github.com/richelieu042/chimera/v3/src/core/mapKit"
	"github.com/richelieu042/chimera/v3/src/core/setKit"
)

var (
	// idMap 即allMap
	/*
		key: id（一对一）
	*/
	idMap = mapKit.NewMapWithLock[string, Channel]()

	// bsidMap
	/*
		key: bsid（一对一）
	*/
	bsidMap = mapKit.NewMapWithLock[string, Channel]()

	// userMap
	/*
		key: user（一对多）
	*/
	userMap = mapKit.NewMapWithLock[string, *setKit.SetWithLock[Channel]]()

	// groupMap
	/*
		key: group（一对多）
	*/
	groupMap = mapKit.NewMapWithLock[string, *setKit.SetWithLock[Channel]]()
)

func GetCountOfIdMap() int {
	return idMap.Size()
}

func GetCountOfBsidMap() int {
	return bsidMap.Size()
}

func GetCountOfUserMap() (count int) {
	/* 读锁 */
	userMap.RLockFunc(func() {
		for _, userSet := range userMap.Map {
			count += userSet.Size()
		}
	})
	return count
}

func GetCountOfGroupMap() (count int) {
	/* 读锁 */
	groupMap.RLockFunc(func() {
		for _, groupSet := range groupMap.Map {
			count += groupSet.Size()
		}
	})
	return count
}
