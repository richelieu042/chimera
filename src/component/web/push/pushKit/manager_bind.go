package pushKit

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/core/setKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

func BindId(channel Channel, id string) {
	if strKit.IsEmpty(id) {
		return
	}

	/* 写锁 */
	idMap.LockFunc(func() {
		if old, ok := idMap.Map[id]; ok {
			_ = old.Close(fmt.Sprintf("id(%s) is replaced by new channel", id))
		}
		idMap.Map[id] = channel
		// 此处无需更新 channel 的信息（id）
	})
}

func BindBsid(channel Channel, bsid string) {
	// 防止: 传参有误 || 重复绑定
	if strKit.IsEmpty(bsid) || channel.GetBsid() == bsid {
		return
	}
	// 先解绑 bsid（有的话）
	UnbindBsid(channel)

	/* 写锁 */
	bsidMap.LockFunc(func() {
		// 处理特殊情况: 已经有别的 channel 绑定了此 bsid
		old := bsidMap.Map[bsid]
		if old != nil {
			_ = old.Close(fmt.Sprintf("bsid(%s) is replaced by new channel", bsid))
			delete(bsidMap.Map, bsid)
		}

		bsidMap.Map[bsid] = channel
		// 更新 channel 的信息
		channel.SetBsid(bsid)
	})
}

func BindUser(channel Channel, user string) {
	// 防止: 传参有误 || 重复绑定
	if strKit.IsEmpty(user) || channel.GetUser() == user {
		return
	}
	// 先解绑 user（有的话）
	UnbindUser(channel)

	/* map写锁 */
	userMap.LockFunc(func() {
		userSet := userMap.Map[user]
		if userSet == nil {
			userSet = setKit.NewSetWithLock[Channel]()
			userMap.Map[user] = userSet
		}

		/* set写锁 */
		userSet.LockFunc(func() {
			_ = userSet.Set.Add(channel)
			// 更新 channel 的信息
			channel.SetUser(user)
		})
	})
}

func BindGroup(channel Channel, group string) {
	// 防止: 传参有误 || 重复绑定
	if strKit.IsEmpty(group) || channel.GetGroup() == group {
		return
	}
	// 先解绑 group（有的话）
	UnbindGroup(channel)

	/* map写锁 */
	groupMap.LockFunc(func() {
		groupSet := groupMap.Map[group]
		if groupSet == nil {
			groupSet = setKit.NewSetWithLock[Channel]()
			groupMap.Map[group] = groupSet
		}

		/* set写锁 */
		groupSet.LockFunc(func() {
			_ = groupSet.Set.Add(channel)
			// 更新 channel 的信息
			channel.SetGroup(group)
		})
	})
}
