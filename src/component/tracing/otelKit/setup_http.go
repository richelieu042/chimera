package otelKit

import (
	"time"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
)

func MustSetUpWithHttp(httpEndpoint, serviceName string, attributeMap map[string]string, opts ...otlptracehttp.Option) {
	err := SetUpWithHttp(httpEndpoint, serviceName, attributeMap, opts...)
	if err != nil {
		console.Fatalf("Fail to set up, error: %+v", err)
	}
}

func SetUpWithHttp(httpEndpoint, serviceName string, attributeMap map[string]string, opts ...otlptracehttp.Option) (err error) {
	defer func() {
		if err == nil {
			setupFlag.Store(true)
		}
	}()

	if err = strKit.AssertNotEmpty(serviceName, "serviceName"); err != nil {
		return
	}

	/* TracerProvider */
	tp, err := NewHttpTracerProvider(httpEndpoint, serviceName, attributeMap, opts...)
	if err != nil {
		return
	}
	otel.SetTracerProvider(tp)

	/* TextMapPropagator */
	textMapPropagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	otel.SetTextMapPropagator(textMapPropagator)

	logrus.RegisterExitHandler(func() {
		ShutdownTracerProvider(tp, time.Second*3)
	})
	return
}
