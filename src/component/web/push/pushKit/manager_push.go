package pushKit

import (
	"sync"

	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
)

// 解决问题：协程池提交失败会永久阻塞
func submitPush(wg *sync.WaitGroup, task func()) error {
	wg.Add(1)
	if err := pushPool.Submit(func() {
		defer wg.Done()
		task()
	}); err != nil {
		wg.Done()
		return err
	}
	return nil
}

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

			if submitErr := submitPush(&wg, func() {
				_ = channel.Push(data)
			}); submitErr != nil {
				err = errKit.Wrapf(submitErr, "fail to submit push task")
				break
			}
		}
		wg.Wait()
	})

	return err
}

func PushToBsid(data []byte, bsid string) (err error) {
	if err = CheckSetup(); err != nil {
		return
	}

	/* map读锁 */
	bsidMap.RLockFunc(func() {
		channel := bsidMap.Map[bsid]
		if channel == nil {
			err = errKit.Wrapf(NoSuitableChannelError, "fail to push to bsid(%s)", bsid)
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
			err = errKit.Wrapf(NoSuitableChannelError, "fail to push to user(%s)", user)
			return
		}

		/* set读锁 */
		userSet.RLockFunc(func() {
			var wg sync.WaitGroup
			userSet.Set.Each(func(channel Channel) bool {
				if sliceKit.Contains(exceptBsids, channel.GetBsid()) {
					return false // 不中断循环
				}

				if submitErr := submitPush(&wg, func() {
					_ = channel.Push(data)
				}); submitErr != nil {
					err = errKit.Wrapf(submitErr, "fail to submit push task for user(%s)", user)
					return true // 中断循环
				}
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
			err = errKit.Wrapf(NoSuitableChannelError, "fail to push to group(%s)", group)
			return
		}

		/* set读锁 */
		groupSet.RLockFunc(func() {
			var wg sync.WaitGroup
			groupSet.Set.Each(func(channel Channel) bool {
				if sliceKit.Contains(exceptBsids, channel.GetBsid()) {
					return false // 不中断循环
				}

				if submitErr := submitPush(&wg, func() {
					_ = channel.Push(data)
				}); submitErr != nil {
					err = errKit.Wrapf(submitErr, "fail to submit push task for group(%s)", group)
					return true // 中断循环
				}
				return false // 不中断循环
			})
			wg.Wait()
		})
	})
	return
}
