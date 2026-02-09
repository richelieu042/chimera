package pushKit

import (
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

func CloseAll(reason string) {
	reason = polyfillCloseReason(reason)

	/* 读锁 */
	idMap.RLockFunc(func() {
		for _, channel := range idMap.Map {
			_ = channel.Close(reason)
		}
	})
	return
}

// CloseById
/*
PS: 解绑（unbind）后续由 inner handler 处理.
*/
func CloseById(id string, reason string) (err error) {
	reason = polyfillCloseReason(reason)

	/* 读锁 */
	idMap.RLockFunc(func() {
		channel := idMap.Map[id]
		if channel == nil {
			return
		}
		err = channel.Close(reason)
	})
	return
}

// CloseByBsid
/*
PS: 解绑（unbind）后续由 inner handler 处理.
*/
func CloseByBsid(bsid string, reason string) (err error) {
	reason = polyfillCloseReason(reason)

	/* 读锁 */
	bsidMap.RLockFunc(func() {
		channel := bsidMap.Map[bsid]
		if channel == nil {
			return
		}
		err = channel.Close(reason)
		// 解绑（unbind）后续由 inner handler 处理
	})
	return
}

// CloseByUser
/*
PS: 解绑（unbind）后续由 inner handler 处理.
*/
func CloseByUser(user string, reason string) {
	reason = polyfillCloseReason(reason)

	/* 读锁 */
	userMap.RLockFunc(func() {
		userSet := userMap.Map[user]

		/* 读锁 */
		userSet.RLockFunc(func() {
			userSet.Set.Each(func(channel Channel) bool {
				_ = channel.Close(reason)
				return false
			})
		})
	})
}

// CloseByGroup
/*
PS: 解绑（unbind）后续由 inner handler 处理.
*/
func CloseByGroup(group string, reason string) {
	reason = polyfillCloseReason(reason)

	/* 读锁 */
	groupMap.RLockFunc(func() {
		groupSet := groupMap.Map[group]

		/* 读锁 */
		groupSet.RLockFunc(func() {
			groupSet.Set.Each(func(channel Channel) bool {
				_ = channel.Close(reason)
				return false
			})
		})
	})
}

func polyfillCloseReason(reason string) string {
	return strKit.EmptyToDefault(reason, "no reason")
}
