package otelKit

import "github.com/richelieu042/chimera/v3/src/core/error/errorKit"

var (
	NotSetupError = errorKit.Newf("haven’t been set up correctly")

	NotOtelRequestError = errorKit.Newf("not otel request")
)
