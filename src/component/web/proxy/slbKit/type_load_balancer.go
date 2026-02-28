package slbKit

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu042/chimera/v3/src/component/web/proxy/forwardKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/cronKit"
	"github.com/richelieu042/chimera/v3/src/log/logKit"
	"github.com/robfig/cron/v3"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

// LoadBalancer 负载均衡器.
/*
!!!: 创建实例后，需要先调用 Start.
*/
type LoadBalancer struct {
	gmutex.RWMutex

	logger *zap.Logger

	backends []*Backend

	// current 当前的下标
	/*
		PS: 虽然值会不停地变大，但以一般理性而论，达不到最大值的（int64的最大值约92233720368亿）.
	*/
	current *atomic.Int64

	//// retry
	//retry int16
	//// retryInterval
	//retryInterval time.Duration

	//// attempt
	//attempt int16

	// c 用于定期健康检查
	c *cron.Cron

	status Status
}

func (lb *LoadBalancer) AddBackend(be *Backend) (err error) {
	if be == nil {
		return
	}
	be.logger = lb.logger

	/* 写锁 */
	lb.LockFunc(func() {
		switch lb.status {
		case StatusInitialized, StatusStarted:
			// 允许添加后端服务
			lb.backends = append(lb.backends, be)
		case StatusDisposed:
			err = AlreadyDisposedError
		default:
			err = errorKit.Newf("invalid status: %s", lb.status)
		}
	})
	return
}

// nextIndex 返回下一个下标.
/*
PS:
(1) 此方法无需加锁;
(2) 需要自行对返回值先进行 其余操作(%) 才能继续使用.
*/
func (lb *LoadBalancer) nextIndex() int64 {
	return lb.current.Inc()
}

// casIndex
/*
PS: 此方法无需加锁.
*/
func (lb *LoadBalancer) casIndex(old, new int64) {
	_ = lb.current.CompareAndSwap(old, new)
}

// healthCheck 对目前所有的后端服务，进行 1次 健康检查.
func (lb *LoadBalancer) healthCheck() {
	/* 读锁 */
	lb.RLock()
	defer lb.RUnlock()

	lb.logger.Info("Health check starts.")
	// 只有 StatusStarted 情况下，才会进行健康检查
	if lb.status != StatusStarted {
		lb.logger.Error("Health check is interrupted.", zap.String("status", string(lb.status)))
		return
	}
	defer lb.logger.Info("Health check ends.")

	var wg sync.WaitGroup
	for _, backend := range lb.backends {
		wg.Add(1)
		go func() {
			defer wg.Done()
			backend.HealthCheck()
		}()
	}
	wg.Wait()
}

// Start 手动启动.
func (lb *LoadBalancer) Start() error {
	/* 写锁 */
	lb.Lock()
	defer lb.Unlock()

	switch lb.status {
	case StatusInitialized:
		lb.status = StatusStarted
	case StatusStarted:
		return AlreadyStartedError
	case StatusDisposed:
		return AlreadyDisposedError
	default:
		return errorKit.Newf("invalid status: %s", lb.status)
	}

	/* 以 10s 为周期，对所有后端服务进行健康检查 */
	spec := fmt.Sprintf("@every %s", HealthCheckInterval.String())
	c, _, err := cronKit.NewCronWithTask(spec, func() {
		lb.healthCheck()
	})
	if err != nil {
		return err
	}
	c.Start() // 不阻塞
	lb.c = c

	lb.logger.Info("Started.")

	return nil
}

// Dispose 手动中止.
func (lb *LoadBalancer) Dispose() {
	/* 写锁 */
	lb.Lock()
	defer lb.Unlock()

	if lb.status != StatusDisposed {
		cronKit.StopCron(lb.c)

		lb.backends = nil
		lb.c = nil
		lb.status = StatusDisposed

		lb.logger.Warn("Disposed.")
	}
}

func (lb *LoadBalancer) HandleRequest(w http.ResponseWriter, r *http.Request) (err error) {
	details := &bytes.Buffer{}
	detailLogger := logKit.NewLogger(details, "", log.Ltime|log.Lmicroseconds|log.Lshortfile)

	/* 输出请求的信息，以便后续定位 */
	//detailLogger.Printf("%s %s\n", r.Method, r.URL.String())
	clientIp := httpKit.GetClientIP(r)
	path := r.URL.Path
	raw := r.URL.RawQuery
	if raw != "" {
		path = path + "?" + raw
	}
	detailLogger.Printf("client ip: %s, method: %s, path: %s", clientIp, r.Method, path)

	defer func() {
		if err != nil {
			lb.logger.Sugar().Errorf("Fail to handle request, error: %s\ndetails:\n%s", err.Error(), details.String())
		} else {
			lb.logger.Sugar().Infof("Succeed to handle request, details:\n%s", details.String())
		}
	}()

	switch lb.status {
	case StatusInitialized:
		return HaveNotBeenStartedError
	case StatusStarted:
		// do nothing
	case StatusDisposed:
		return AlreadyDisposedError
	default:
		return errorKit.Newf("invalid status: %s", lb.status)
	}

	if err := r.Context().Err(); err != nil {
		// 请求已经被取消
		return err
	}

	/* 读锁 */
	lb.RLock()
	defer lb.RUnlock()

	length := int64(len(lb.backends))
	if length == 0 {
		return NoBackendAddedError
	}
	start := lb.nextIndex()
	limit := start + length
	for i := start; i < limit; i++ {
		idx := i % length
		be := lb.backends[idx]
		if !be.IsAlive() {
			/* (1) 找到的服务不可用，继续找 */
			detailLogger.Printf("(%d/%d, %s) Not alive, continue...", idx+1, length, be.String())
			continue
		}
		err := be.HandleRequest(w, r)
		if err != nil {
			if forwardKit.IsInterruptedError(err) {
				/* (2) 请求被中断了 */
				detailLogger.Printf("(%d/%d, %s) Request is interrupted, end.", idx+1, length, be.String())
				return err
			}
			/* (3) 当前找到的后端服务有问题，继续找 */
			be.Disable("fail to forward request, error: %s", err.Error())
			detailLogger.Printf("(%d/%d, %s) Fail to proxy with error(%s), continue...", idx+1, length, be.String(), err.Error())
			continue
		}
		/* (4) 代理请求成功 */
		/*
			下面三行代码是为了解决bug: 有三个后端节点（8000、8001、8002），依次Add.假如8001挂了，会导致8002压力增加（相对于8000来说）.
			当满足条件时（即某次代理，最前面的后端服务不可用的情况下），尝试手动更新index.
		*/
		if i != start {
			lb.casIndex(start, i)
		}
		detailLogger.Printf("(%d/%d, %s) Succeed to proxy, end.", idx+1, length, be.String())
		return nil
	}
	/* (5) 无可用后端服务 */
	return NoAccessBackendError
}
