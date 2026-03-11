package reqKit

import (
	"io"
	"log"

	"github.com/imroc/req/v3"
)

var (
	NewLogger func(output io.Writer, prefix string, flag int) req.Logger = req.NewLogger

	NewLoggerFromStandardLogger func(l *log.Logger) req.Logger = req.NewLoggerFromStandardLogger
)
