package slbKit

import "github.com/richelieu042/chimera/v3/src/core/errorKit"

var (
	HaveNotBeenStartedError = errorKit.Newf("have not been started")

	AlreadyStartedError = errorKit.Newf("already started")

	AlreadyDisposedError = errorKit.Newf("already disposed")

	NoAccessBackendError = errorKit.Newf("no access backend")

	NoBackendAddedError = errorKit.Newf("no backend added")
)
