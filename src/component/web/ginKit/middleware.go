package ginKit

import (
	"fmt"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/core/mapKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/micro/rateLimitKit"
	"golang.org/x/time/rate"
)

//// UseMiddlewares
///*
//@param middlewares 其中的元素不能为nil!!!
//*/
//func UseMiddlewares(engine *gin.Engine, middlewares ...gin.HandlerFunc) (err error) {
//	if len(middlewares) == 0 {
//		return
//	}
//
//	sliceKit.Each(middlewares, func(middleware gin.HandlerFunc, index int) bool {
//		if middleware == nil {
//			err = errorKit.Sign("middlewares[%d] == nil", index)
//			return true
//		}
//		return false
//	})
//	if err != nil {
//		return
//	}
//
//	engine.Use(middlewares...)
//	return
//}

// attachMiddlewares 绑定一些常用的中间件.
func attachMiddlewares(engine *gin.Engine, config MiddlewareConfig, opts *ginOptions) error {
	// gzip
	/*
		PS:
		(1) 必须在 recoveryMiddleware 前面，为了万无一失还是放在 最前面 吧；
		(2) gzip会使得响应头中的 Content-Length 不生效.
	*/
	gc := config.Gzip
	if gc != nil && gc.Level != gzip.NoCompression {
		if sliceKit.IsEmpty(gc.ExcludedExtensions) {
			// 默认值
			gc.ExcludedExtensions = []string{".png", ".gif", ".jpeg", ".jpg", ".webp"}
		}

		gzipMiddleware := NewGzipMiddleware(gc.Level,
			gzip.WithExcludedExtensions(gc.ExcludedExtensions),
			gzip.WithExcludedPaths(gc.ExcludedPaths),
			gzip.WithExcludedPathsRegexs(gc.ExcludedPathsRegexps),
		)
		engine.Use(gzipMiddleware)
	}

	/* logger(necessary) && recovery(necessary) */
	engine.Use(gin.Logger(), opts.RecoveryMiddleware)

	/* cors(optional) */
	cors := config.Cors
	if cors.Access {
		// 配置cors
		origins := sliceKit.PolyfillStringSlice(cors.Origins)
		engine.Use(NewCorsMiddleware(origins, cors.AllowCredentials))
	}

	/* OPTIONS */
	if opts.DefaultOptionsMiddleware {
		engine.Use(NewOptionsMiddleware())
	}

	/* rate limiter（限流器） */
	rlConfig := config.RateLimiter
	if rlConfig != nil {
		forbiddenText := fmt.Sprintf("Exceed rate limit(r: %d, b: %d).", rlConfig.R, rlConfig.B)
		if strKit.IsNotEmpty(serviceInfo) {
			forbiddenText = fmt.Sprintf("[%s] %s", serviceInfo, forbiddenText)
		}

		middleware := rateLimitKit.NewGinMiddleware(rate.Limit(rlConfig.R), rlConfig.B, forbiddenText)
		engine.Use(middleware)
	}

	/* bodyLimit */
	// TODO: 因为http.MaxBytesReader()，如果涉及"请求转发（代理）"，转发方不要全局配置此属性，否则会导致: 有时成功，有时代理失败（error），有时http客户端失败
	if config.BodyLimit > 0 {
		middleware, err := NewSizeLimiterMiddleware(config.BodyLimit)
		if err != nil {
			return err
		}
		engine.Use(middleware)
	}

	/* response headers */
	if mapKit.IsNotEmpty(config.ResponseHeaders) {
		engine.Use(func(ctx *gin.Context) {
			for k, v := range config.ResponseHeaders {
				ctx.Header(k, v)
			}
			ctx.Next()
		})
	}

	/*
	   # 响应头的"X-Frame-Options"（全部小写！），可选值:
	   # (1) deny                页面不允许在frame中显示
	   # (2) sameorigin          页面允许在相同域名页面的frame中显示
	   # (3) allow-from $(uri)   页面可以在制定来源的frame中显示，e.g. allow-from https://example.com/
	   xFrameOptions:
	*/
	//// others
	//engine.Use(func(ctx *gin.Context) {
	//	if strKit.IsNotEmpty(config.XFrameOptions) {
	//		// e.g.不能被嵌入到任何iframe或frame中
	//		ctx.Header("X-Frame-Options", config.XFrameOptions)
	//	}
	//
	//	// 解决漏洞: 未启用Web浏览器XSS保护
	//	ctx.Header("X-XSS-Protection", "1;mode=block")
	//})

	return nil
}
