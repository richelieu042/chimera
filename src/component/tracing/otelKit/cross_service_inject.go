package otelKit

import (
	"context"
	"net/http"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// InjectIntoRequest 使用 baggage 写入 trace id 和 span id（实际上是写到 Header 中）.
/*
PS:
(1) 需要先 set up;
(2) 适用场景: 跨服务（跨应用通讯） + 发送端.
*/
func InjectIntoRequest(r *http.Request, spanCtx context.Context, span trace.Span) error {
	carrier := propagation.HeaderCarrier(r.Header)
	return injectWithCarrier(carrier, spanCtx, span)
}

// InjectIntoMap
/*
PS:
(1) 需要先 set up;
(2) 适用场景: 跨服务（跨应用通讯） + 发送端.
*/
func InjectIntoMap(m map[string]string, spanCtx context.Context, span trace.Span) error {
	carrier := propagation.MapCarrier(m)
	return injectWithCarrier(carrier, spanCtx, span)
}

func injectWithCarrier(carrier propagation.TextMapCarrier, spanCtx context.Context, span trace.Span) (err error) {
	defer func() {
		if err != nil {
			err = errorKit.Wrapf(err, "fail to inject baggage")
		}
	}()

	traceMember, err := baggage.NewMember(KeyTraceId, span.SpanContext().TraceID().String())
	if err != nil {
		return
	}
	spanMember, err := baggage.NewMember(KeySpanId, span.SpanContext().SpanID().String())
	if err != nil {
		return
	}
	b, err := baggage.New(traceMember, spanMember)
	if err != nil {
		return
	}
	baggageCtx := baggage.ContextWithBaggage(spanCtx, b)

	//propagation.Baggage{}.Inject(baggageCtx, carrier)
	otel.GetTextMapPropagator().Inject(baggageCtx, carrier)

	return
}
