package otelKit

import (
	"context"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

func TestNewGrpcTracerProvider(t *testing.T) {
	tp, err := NewGrpcTracerProvider("", "SERVICE_NAME_GRPC", map[string]string{
		"hello": "world",
	}, otlptracegrpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	otel.SetTracerProvider(tp)

	funA := func(spanCtx context.Context, wg *sync.WaitGroup) {
		defer func() {
			logrus.Info("funA ends.")
		}()
		defer wg.Done()

		time.Sleep(time.Second * 3)

		tracer := otel.Tracer("my-tracer-a")
		_, span := tracer.Start(spanCtx, "funA")
		defer span.End()
		span.SetAttributes(attribute.KeyValue{
			Key:   "funA",
			Value: attribute.StringValue("funA"),
		})

		// span持续2s
		time.Sleep(time.Second * 2)
	}
	funB := func(spanCtx context.Context) {
		defer func() {
			logrus.Info("funB ends.")
		}()

		tracer := otel.Tracer("my-tracer-b")
		_, span := tracer.Start(spanCtx, "funB")
		defer span.End()
		span.SetAttributes(attribute.KeyValue{
			Key:   "funB",
			Value: attribute.StringValue("funB"),
		})

		// span持续1s
		time.Sleep(time.Second)
	}

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		logrus.Info("Start.")
		defer func() {
			logrus.Info("end.")
		}()

		tracer := otel.Tracer("TRACER-NAME")
		spanCtx, span := tracer.Start(ctx, "SPAN_NAME")
		defer span.End()

		wg := new(sync.WaitGroup)
		wg.Add(1)
		go funA(spanCtx, wg)
		funB(spanCtx)
		wg.Wait()

		ctx.String(http.StatusOK, "hello")
	})
	if err := engine.Run(":8888"); err != nil {
		logrus.Fatal(engine)
	}
}
