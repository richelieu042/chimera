package forwardKit

import (
	"net/http"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/urlKit"
	"github.com/richelieu042/chimera/v3/src/validateKit"
)

// NewDirector
/*
@param targetHost hostname || hostname:port
*/
func NewDirector(targetHost string, options ...DirectorOption) (director func(req *http.Request), err error) {
	if err = validateKit.Var(targetHost, "hostname|ipv4|hostname_port"); err != nil {
		err = errorKit.Newf("invalid targetHost(%s)", targetHost)
		return
	}

	opts := loadOptions(options...)
	director = func(req *http.Request) {
		req.URL.Scheme = opts.scheme
		req.URL.Host = targetHost
		if opts.requestUrlPath != nil {
			req.URL.Path = *opts.requestUrlPath
		}

		// 可能会修改 r.URL.RawQuery
		if opts.overrideQueryParams != nil {
			urlKit.OverrideRawQuery(req.URL, opts.overrideQueryParams)
		} else if opts.extraQueryParams != nil {
			urlKit.AddToRawQuery(req.URL, opts.extraQueryParams)
		}
	}
	return
}
