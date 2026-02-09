package pushKit

import (
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// UnbindId
/*
PS: 仅是解绑，不会关闭channel.
*/
func UnbindId(channel Channel) {
	id := channel.GetId()

	/* 写锁 */
	idMap.LockFunc(func() {
		// Richelieu: 此处额外判断一下，因为 inner listener 触发事件（此处主要是OnClose）是异步的.
		target := idMap.Map[id]
		if channel.Equals(target) {
			delete(idMap.Map, id)
		}
	})
}

// UnbindBsid
/*
PS: 仅是解绑，不会关闭channel.
*/
func UnbindBsid(channel Channel) {
	bsid := channel.GetBsid()
	if strKit.IsEmpty(bsid) {
		return
	}
	defer channel.ClearBsid()

	/* 写锁 */
	bsidMap.LockFunc(func() {
		// Richelieu: 此处额外判断一下，因为 inner listener 触发事件（此处主要是OnClose）是异步的.
		target := bsidMap.Map[bsid]
		if channel.Equals(target) {
			delete(bsidMap.Map, bsid)
		}
	})
}

// UnbindUser
/*
PS:
(1) 仅是解绑，不会关闭channel;
(2) 解绑成功后，如果set为空，应该移除掉.
*/
func UnbindUser(channel Channel) {
	user := channel.GetUser()
	if strKit.IsEmpty(user) {
		return
	}
	defer channel.ClearUser()

	/* map写锁 */
	userMap.LockFunc(func() {
		userSet, ok := userMap.Map[user]
		if !ok {
			return
		}
		if userSet == nil {
			delete(userMap.Map, user)
			return
		}
		/* set写锁 */
		userSet.LockFunc(func() {
			userSet.Set.Remove(channel)
			if userSet.Set.Cardinality() == 0 {
				delete(userMap.Map, user)
			}
		})
	})
}

// UnbindGroup
/*
PS:
(1) 仅是解绑，不会关闭channel;
(2) 解绑成功后，如果set为空，应该移除掉.
*/
func UnbindGroup(channel Channel) {
	group := channel.GetGroup()
	if strKit.IsEmpty(group) {
		return
	}
	defer channel.ClearGroup()

	/* map写锁 */
	groupMap.LockFunc(func() {
		groupSet, ok := groupMap.Map[group]
		if !ok {
			return
		}
		if groupSet == nil {
			delete(groupMap.Map, group)
			return
		}
		/* set写锁 */
		groupSet.LockFunc(func() {
			groupSet.Set.Remove(channel)
			if groupSet.Set.Cardinality() == 0 {
				delete(groupMap.Map, group)
			}
		})
	})
}
