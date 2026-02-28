package otelKit

import (
	"context"
	"net/http"

	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// ExtractFromRequest
/*
PS:
(1) 需要先 set up
(2) 适用场景: 跨服务（跨应用通讯） + 接收端
(3) 返回的error != nil的情况下，可以判断是否等于 otelKit.NotOtelRequestError.
*/
func ExtractFromRequest(r *http.Request) (remoteSpanCtx context.Context, err error) {
	// Richelieu: 如果不在此处就行特殊情况判断，下面会在 trace.TraceIDFromHex(b.Member(KeyTraceId).Value()) 处返回error
	baggageStr := httpKit.GetHeader(r.Header, HeaderBaggage)
	if strKit.IsEmpty(baggageStr) {
		// 非链路追踪请求
		err = NotOtelRequestError
		return
	}

	defer func() {
		if err != nil {
			err = errorKit.Wrapf(err, "fail to extract from request")
		}
	}()

	//var propagator = propagation.TextMapPropagator(propagation.Baggage{})
	//propagatorCtx := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
	propagatorCtx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))

	b := baggage.FromContext(propagatorCtx)
	traceId, err := trace.TraceIDFromHex(b.Member(KeyTraceId).Value())
	if err != nil {
		return
	}
	spanId, err := trace.SpanIDFromHex(b.Member(KeySpanId).Value())
	if err != nil {
		return
	}
	var spanCtx trace.SpanContext = trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    traceId,
		SpanID:     spanId,
		TraceFlags: trace.FlagsSampled, // 这个没写，是不会记录的
		TraceState: trace.TraceState{},
		Remote:     true,
	})
	remoteSpanCtx = trace.ContextWithRemoteSpanContext(propagatorCtx, spanCtx)
	return
}
