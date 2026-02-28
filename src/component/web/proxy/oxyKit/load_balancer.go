package oxyKit

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/sirupsen/logrus"
	"github.com/vulcand/oxy/v2/buffer"
	"github.com/vulcand/oxy/v2/forward"
	"github.com/vulcand/oxy/v2/roundrobin"
	"github.com/vulcand/oxy/v2/utils"
)

var (
	defaultOxyReverseProxy = forward.New(true)
)

// NewLoadBalancerHandler
/*
Deprecated: 详见 _info.md

@param reverseProxy	可以为nil（将采用默认值）
@param servers 		e.g. []string{"http://localhost:8001", "http://localhost:8002"}
@param logrusLogger	可以为nil，但不建议这么干，因为会无输出
@param verbose		true: 详细的信息
*/
func NewLoadBalancerHandler(reverseProxy *httputil.ReverseProxy, servers []string, logrusLogger *logrus.Logger, verbose bool) (func(http.ResponseWriter, *http.Request), error) {
	if reverseProxy == nil {
		reverseProxy = defaultOxyReverseProxy
	}
	servers = sliceKit.PolyfillStringSlice(servers)
	if err := sliceKit.AssertNotEmpty(servers, "servers"); err != nil {
		return nil, err
	}

	var errHandler utils.ErrorHandler
	var logger utils.Logger

	if logrusLogger == nil {
		// 不输出
		errHandler = utils.DefaultHandler
		logger = &utils.NoopLogger{}
	} else {
		errHandler = &errorHandler{
			logger: logrusLogger,
		}
		logger = NewLogger(logrusLogger)
	}

	lb, err := roundrobin.New(reverseProxy,
		roundrobin.ErrorHandler(errHandler),
		roundrobin.Verbose(verbose),
		roundrobin.Logger(logger),
	)
	if err != nil {
		return nil, errorKit.Wrapf(err, "roundrobin.New() fails")
	}
	for _, server := range servers {
		u, err := url.Parse(server)
		if err != nil {
			return nil, errorKit.Wrapf(err, "server(%s) is invalid", server)
		}
		if err := lb.UpsertServer(u); err != nil {
			return nil, errorKit.Wrapf(err, "lb.UpsertServer() fails with server(%s)", server)
		}
	}
	/*
		Richelieu: 有几个服务，就最多重试几次.
		e.g. `IsNetworkError() && Attempts() < 3`: 最多试 3 次
	*/
	predicate := fmt.Sprintf("IsNetworkError() && Attempts() < %d", len(servers))
	// buf will read the request body and will replay the request again in case if forward returned status
	// corresponding to nework error (e.g. Gateway Timeout)
	buf, err := buffer.New(lb,
		buffer.Retry(predicate),
		buffer.ErrorHandler(errHandler),
		buffer.Verbose(verbose),
		buffer.Logger(logger),
	)
	if err != nil {
		return nil, errorKit.Wrapf(err, "buffer.New() fails")
	}
	return buf.ServeHTTP, nil
}
