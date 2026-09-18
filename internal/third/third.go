package third

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

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

func init() {
	///* golang.org/x/ */
	//{
	//	var _ x86asm.Arg
	//	var _ *cast5.Cipher
	//	var _ *bind.Generator
	//	var _ *modfile.Comment
	//	var _ jws.Signer
	//	var _ *errgroup.Group
	//	var _ *execabs.Error
	//	var _ *term.Terminal
	//	var _ *blog.Doc
	//	var _ slog.Handler
	//	var _ *tiff.Options
	//	var _ proxy.Dialer
	//	var _ *currency.Amount
	//	var _ *rate.Limit
	//}

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
