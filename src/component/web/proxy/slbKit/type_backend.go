package slbKit

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/richelieu042/chimera/v3/src/component/web/proxy/forwardKit"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/netKit"
	"go.uber.org/zap"
)

// Backend 后端节点.
type Backend struct {
	gmutex.Mutex

	logger *zap.Logger

	// alive 当前节点是否可用？
	alive        bool
	u            *url.URL
	reverseProxy *httputil.ReverseProxy
}

func NewBackend(urlStr string) (*Backend, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, errorKit.Wrapf(err, "invalid url: %s", urlStr)
	}
	rp, err := forwardKit.NewSingleHostReverseProxy(u)
	if err != nil {
		return nil, err
	}

	return &Backend{
		alive:        true,
		u:            u,
		reverseProxy: rp,
	}, nil
}

func (be *Backend) Enable(reason string, a ...any) {
	reason = fmt.Sprintf(reason, a...)

	/* 锁 */
	be.LockFunc(func() {
		if be.alive {
			return
		}
		be.alive = true
		be.logger.Info("Backend is [enable].", zap.String("backend", be.String()), zap.String("reason", reason))
	})
}

func (be *Backend) Disable(reason string, a ...any) {
	reason = fmt.Sprintf(reason, a...)

	/* 锁 */
	be.LockFunc(func() {
		if !be.alive {
			return
		}
		be.alive = false
		be.logger.Warn("Backend is [disable].", zap.String("backend", be.String()), zap.String("reason", reason))
	})
}

func (be *Backend) IsAlive() (alive bool) {
	/* 锁 */
	be.LockFunc(func() {
		alive = be.alive
	})
	return
}

// HealthCheck 健康检查（此方法会修改 alive 字段）.
/*
@return 后端服务是否可用？
*/
func (be *Backend) HealthCheck() {
	conn, err := netKit.DialTimeout("tcp", be.u.Host, HealthCheckTimeout)
	if err != nil {
		be.Disable("health check fails, error: %s", err.Error())
		return
	}
	_ = conn.Close()
	be.Enable("health check succeeds")
}

func (be *Backend) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	return forwardKit.ForwardByReverseProxy(w, r, be.reverseProxy)
}

func (be *Backend) String() string {
	if be == nil {
		return "null"
	}
	return be.u.String()
}
