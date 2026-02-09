package centrifugoKit

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/grpcKit"
	"github.com/richelieu042/chimera/v3/src/idKit"
	"github.com/richelieu042/chimera/v3/src/micro/centrifugoKit/apiproto"
	"github.com/richelieu042/chimera/v3/src/netKit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

// NewGrpcClient centrifugo服务的grpc客户端，支持客户端负载均衡（slb）
/*
@param hosts		centrifugo服务的grpc地址列表 e.g.[]string{"127.0.0.1:10000", "127.0.0.1:10001"}
@param scheme		grpc客户端负载均衡(slb)使用的scheme
					(1) 可以为""，将自动生成
					(2) 其中不能有大写字母
					(3) 可以有: 小写字母、数字、"-"...
					(4) 长度貌似有限制
					(5) 不要以 数字 开头
@param grpcApiKey	对应centrifugo服务配置文件中的 "grpc_api_key"
*/
func NewGrpcClient(hosts []string, scheme string, grpcApiKey string) (*GrpcClient, error) {
	/* hosts */
	hosts, err := netKit.PolyfillHosts(hosts, 1)
	if err != nil {
		return nil, err
	}

	/* scheme */
	if strKit.IsEmpty(scheme) {
		scheme = fmt.Sprintf("chimera-centrifugo-%s", idKit.NewXid())
	} else {
		scheme = strKit.ToLower(scheme)
	}

	/* grpcApiKey */
	if err := strKit.AssertNotEmpty(grpcApiKey, "grpcApiKey"); err != nil {
		return nil, err
	}

	var target string
	if len(hosts) > 1 {
		/* slb */
		builder, err := grpcKit.NewResolverBuilder(scheme, hosts)
		if err != nil {
			return nil, err
		}
		resolver.Register(builder)

		// Richelieu: target中的"hello"随意，甚至可以去掉
		target = fmt.Sprintf("%s:///hello", scheme)
	} else {
		// 此种情况下（就一个host），不使用 scheme
		target = hosts[0]
	}

	/* new client */
	conn, err := grpc.NewClient(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithPerRPCCredentials(NewKeyAuth(grpcApiKey)),
	)
	if err != nil {
		return nil, err
	}
	client := apiproto.NewCentrifugoApiClient(conn)
	return &GrpcClient{
		CentrifugoApiClient: client,
		conn:                conn,
	}, nil
}
