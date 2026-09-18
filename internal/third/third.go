package third

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"golang.org/x/arch/x86/x86asm"
	"golang.org/x/crypto/cast5"
	"golang.org/x/exp/slog"
	"golang.org/x/image/tiff"
	"golang.org/x/mobile/bind"
	"golang.org/x/mod/modfile"
	"golang.org/x/net/proxy"
	"golang.org/x/oauth2/jws"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sys/execabs"
	"golang.org/x/term"
	"golang.org/x/text/currency"
	"golang.org/x/time/rate"
	"golang.org/x/tools/blog"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/hashicorp/serf/serf"
)

func init() {
	/* golang.org/x/ */
	{
		var _ x86asm.Arg
		var _ *cast5.Cipher
		var _ *bind.Generator
		var _ *modfile.Comment
		var _ jws.Signer
		var _ *errgroup.Group
		var _ *execabs.Error
		var _ *term.Terminal
		var _ *blog.Doc
		var _ slog.Handler
		var _ *tiff.Options
		var _ proxy.Dialer
		var _ *currency.Amount
		var _ *rate.Limit
	}

	/* otel */
	{
		var _ = otlptrace.Version()
		var _ = otlptracehttp.NewClient
		var _ = otlptracegrpc.NewClient
	}

	/* grpc && protobuf */
	{
		var _ *grpc.ConnectParams
		var _ *protojson.UnmarshalOptions
	}

	///* github.com/ulikunitz/xz v0.5.10 是脆弱的 */
	//{
	//	var _ xz.Writer
	//}
}
