package ginKit

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/appKit"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/ioKit"
	"github.com/richelieu042/chimera/v3/src/core/signalKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/richelieu042/chimera/v3/src/netKit"
	"github.com/richelieu042/chimera/v3/src/time/timeKit"
	"github.com/richelieu042/chimera/v3/src/validateKit"
)

// serviceInfo e.g."Agent-127.0.0.1:12345"
var serviceInfo = ""

func MustSetUp(config *Config, businessLogic func(engine *gin.Engine) error, options ...GinOption) {
	err := SetUp(config, businessLogic, options...)
	if err != nil {
		console.Fatalf("Fail to setup, error: %s", err)
	}
}

// SetUp
/*
PS: 正常执行的情况下，此方法会阻塞调用的协程.

@param config			不能为nil（否则将返回error）
@param businessLogic 	业务逻辑，可以在其中进行 路由绑定 等操作...（可以为nil，但不推荐这么干）
*/
func SetUp(config *Config, businessLogic func(engine *gin.Engine) error, options ...GinOption) error {
	if err := validateKit.Struct(config); err != nil {
		return err
	}

	opts := loadOptions(options...)
	serviceInfo = opts.ServiceInfo

	// Gin的模式，后续也可以在 businessLogic 里面调整
	if strKit.IsEmpty(config.Mode) {
		// 默认: debug模式
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(config.Mode)
	}

	/*
		gin框架中如何让日志文字带颜色输出？
			https://mp.weixin.qq.com/s/eHtIC5egDoqx4LdAvcE5Qw

		PS: 如果 Gin.ForceConsoleColor() 和 Gin.DisableConsoleColor() 都不调用，那么默认是在终端中输出日志是带颜色的，输出到其他地方是不带颜色的.
	*/
	if config.DisableColor {
		// 禁用日志带颜色输出
		gin.DisableConsoleColor()
	} else {
		// 强制日志带颜色输出（无论是在终端还是其他输出设备）
		gin.ForceConsoleColor()
	}

	// 默认: os.Stdout
	gin.DefaultWriter = ioKit.LockedStdout
	// 默认: os.Stderr
	gin.DefaultErrorWriter = ioKit.LockedStderr

	engine := DefaultEngine()

	// pprof
	if config.Pprof {
		pprof.Register(engine, pprof.DefaultPrefix) // 等价于 pprof.Register(engine)
	}

	// middleware
	if err := attachMiddlewares(engine, config.Middleware, opts); err != nil {
		return err
	}

	/* favicon.ico */
	if opts.DefaultFavicon {
		DefaultFavicon(engine)
	}

	/* 404 */
	if opts.DefaultNoRouteHtml {
		if err := DefaultNoRouteHtml(engine); err != nil {
			return err
		}
	}

	/* 405（不设置的话，就会走到404） */
	engine.HandleMethodNotAllowed = opts.DefaultNoMethod
	if engine.HandleMethodNotAllowed {
		DefaultNoMethod(engine)
	}

	/* 业务逻辑 */
	if businessLogic != nil {
		if err := businessLogic(engine); err != nil {
			return errorKit.Wrapf(err, "Fail to execute businessLogic().")
		}
	}

	/* http */
	httpPort := config.Port
	var httpSrv *http.Server
	if httpPort != 0 {
		httpSrv = &http.Server{
			Addr:    netKit.JoinToHost(config.HostName, httpPort),
			Handler: engine.Handler(),
		}
		console.Infof("Listening and serving HTTP on address(%s)", httpSrv.Addr)

		go func() {
			err := httpSrv.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				console.Fatalf("Fail to start http server with addr(%s), error: %+v", httpSrv.Addr, err)
			}
		}()
	}
	/* https */
	sslConfig := config.SSL
	httpsPort := sslConfig.Port
	var httpsSrv *http.Server
	if httpsPort != 0 {
		httpsSrv = &http.Server{
			Addr:    netKit.JoinToHost(config.HostName, httpsPort),
			Handler: engine.Handler(),
		}
		console.Infof("Listening and serving HTTPS on [%s]", httpsSrv.Addr)

		go func() {
			err := httpsSrv.ListenAndServeTLS(sslConfig.CertFile, sslConfig.KeyFile)
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				console.Fatalf("Fail to start https server with addr(%s), error: %+v", httpsSrv.Addr, err)
			}
		}()
	}

	if httpPort == 0 && httpsPort == 0 {
		return errorKit.Newf("both httpPort and httpsPort are zero")
	}

	/*
		优雅地重启或停止
			https://gin-gonic.com/zh-cn/docs/examples/graceful-restart-or-stop/
			https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go
	*/
	appKit.RegisterExitHandler(func() {
		var wg sync.WaitGroup

		ctx, cancel := context.WithTimeout(context.TODO(), timeKit.Second*5)
		defer cancel()
		if httpSrv != nil {
			wg.Add(1)
			go func() {
				defer wg.Done()

				if err := httpSrv.Shutdown(ctx); err != nil {
					console.Errorf("Fail to shut down http server, error: %+v", err)
					return
				}
				console.Info("Manager to shut down http server.")
			}()
		}
		if httpsSrv != nil {
			wg.Add(1)
			go func() {
				defer wg.Done()

				if err := httpsSrv.Shutdown(ctx); err != nil {
					console.Errorf("Fail to shut down https server, error: %+v", err)
					return
				}
				console.Info("Manager to shut down https server.")
			}()
		}
		wg.Wait()
	})
	signalKit.MonitorExitSignals()
	return nil
}
