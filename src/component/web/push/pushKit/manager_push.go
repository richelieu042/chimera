package pushKit

import (
	"sync"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
)

func PushToAll(data []byte, exceptBsids []string) (err error) {
	if err = CheckSetup(); err != nil {
		return err
	}

	/* map读锁 */
	idMap.RLockFunc(func() {
		var wg sync.WaitGroup
		for _, channel := range idMap.Map {
			if sliceKit.Contains(exceptBsids, channel.GetBsid()) {
				continue
			}

			// for range + goroutine，必须使用 "同名变量覆盖v:=v"
			c := channel
			wg.Add(1)
			_ = pushPool.Submit(func() {
				defer wg.Done()
				_ = c.Push(data)
			})
		}
		wg.Wait()
	})

	return nil
}

func PushToBsid(data []byte, bsid string) (err error) {
	if err = CheckSetup(); err != nil {
		return
	}

	/* map读锁 */
	bsidMap.RLockFunc(func() {
		channel := bsidMap.Map[bsid]
		if channel == nil {
			err = errorKit.Wrapf(NoSuitableChannelError, "fail to push to bsid(%s)", bsid)
			return
		}
		err = channel.Push(data)
	})
	return
}

func PushToUser(data []byte, user string, exceptBsids []string) (err error) {
	if err = CheckSetup(); err != nil {
		return
	}

	/* map读锁 */
	userMap.RLockFunc(func() {
		userSet := userMap.Map[user]
		if userSet == nil {
			err = errorKit.Wrapf(NoSuitableChannelError, "fail to push to user(%s)", user)
			return
		}

		/* set读锁 */
		userSet.RLockFunc(func() {
			var wg sync.WaitGroup
			userSet.Set.Each(func(channel Channel) bool {
				if sliceKit.Contains(exceptBsids, channel.GetBsid()) {
					return false // 不中断循环
				}

				// 由于使用 Set.Each() 进行遍历，此处无需使用 "同名变量覆盖v:=v"
				wg.Add(1)
				_ = pushPool.Submit(func() {
					defer wg.Done()
					_ = channel.Push(data)
				})
				return false // 不中断循环
			})
			wg.Wait()
		})
	})
	return
}

func PushToGroup(data []byte, group string, exceptBsids []string) (err error) {
	if err = CheckSetup(); err != nil {
		return
	}

	/* map读锁 */
	groupMap.RLockFunc(func() {
		groupSet := groupMap.Map[group]
		if groupSet == nil {
			err = errorKit.Wrapf(NoSuitableChannelError, "fail to push to group(%s)", group)
			return
		}

		/* set读锁 */
		groupSet.RLockFunc(func() {
			var wg sync.WaitGroup
			groupSet.Set.Each(func(channel Channel) bool {
				if sliceKit.Contains(exceptBsids, channel.GetBsid()) {
					return false // 不中断循环
				}

				// 由于使用 Set.Each() 进行遍历，此处无需使用 "同名变量覆盖v:=v"
				wg.Add(1)
				_ = pushPool.Submit(func() {
					defer wg.Done()
					_ = channel.Push(data)
				})
				return false // 不中断循环
			})
			wg.Wait()
		})
	})
	return
}
