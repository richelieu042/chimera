package third

import (
	_ "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	_ "google.golang.org/grpc"
	_ "google.golang.org/protobuf/encoding/protojson"

	_ "golang.org/x/arch/x86/x86asm"
	_ "golang.org/x/crypto/cast5"
	_ "golang.org/x/exp/slog"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/mobile/bind"
	_ "golang.org/x/mod/modfile"
	_ "golang.org/x/net/proxy"
	_ "golang.org/x/oauth2/jws"
	_ "golang.org/x/sync/errgroup"
	_ "golang.org/x/sys/execabs"
	_ "golang.org/x/term"
	_ "golang.org/x/text/currency"
	_ "golang.org/x/time/rate"
	_ "golang.org/x/tools/blog"

	_ "filippo.io/edwards25519"        // 处理v1.1.0的漏洞
	_ "github.com/hashicorp/serf/serf" // 处理hashicorp系依赖的内部依赖问题
)
