package otelKit

import "github.com/richelieu042/chimera/v3/src/core/error/errKit"

var (
	NotSetupError = errKit.Newf("haven’t been set up correctly")

	NotOtelRequestError = errKit.Newf("not otel request")
)
