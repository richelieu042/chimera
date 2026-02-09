package forwardKit

import (
	"net/http"
	"net/url"

	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
)

func NewSingleHostDirector(u *url.URL) (func(r *http.Request), error) {
	if err := interfaceKit.AssertNotNil(u, "u"); err != nil {
		return nil, err
	}

	rp, err := NewSingleHostReverseProxy(u)
	if err != nil {
		return nil, err
	}
	return rp.Director, nil
}
