package otelKit

import (
	_ "github.com/richelieu042/chimera/v3/src/log/logrusInitKit"

	"testing"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc"
)

func TestMustSetUp(t *testing.T) {
	logrus.Info("---")
	MustSetUpWithGrpc("", "service-test", nil, otlptracegrpc.WithInsecure(), otlptracegrpc.WithDialOption(grpc.WithBlock()))
	logrus.Info("---")
}
