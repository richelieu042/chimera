package grpcKit

import (
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"google.golang.org/grpc/resolver"
)

// NewResolverBuilder 用于: grpc客户端负载均衡(slb).
/*
PS: 第2个返回值为nil的情况下，应该将第1个返回值作为传参调用 resolver.Register().

@param scheme 协议（不能有大写字符）
*/
func NewResolverBuilder(scheme string, hosts []string) (resolver.Builder, error) {
	if err := strKit.AssertNotEmpty(scheme, "scheme"); err != nil {
		return nil, err
	}
	hosts = sliceKit.PolyfillStringSlice(hosts)
	if err := sliceKit.AssertNotEmpty(hosts, "hosts"); err != nil {
		return nil, err
	}

	addrs := hostsToAddressSlice(hosts)
	return &manualResolverBuilder{
		scheme:    scheme,
		addresses: addrs,
	}, nil
}

type manualResolverBuilder struct {
	resolver.Builder

	// scheme returns the scheme supported by this resolver.  Scheme is defined
	// at https://github.com/grpc/grpc/blob/master/doc/naming.md.  The returned
	// string should not contain uppercase characters, as they will not match
	// the parsed target's scheme as defined in RFC 3986.
	scheme    string
	addresses []resolver.Address
}

func (builder *manualResolverBuilder) Scheme() string {
	return builder.scheme
}

func (builder *manualResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &manualResolver{
		target:    target,
		cc:        cc,
		addresses: builder.addresses,
	}
	if err := r.start(); err != nil {
		return nil, err
	}
	return r, nil
}

type manualResolver struct {
	resolver.Resolver

	target resolver.Target
	cc     resolver.ClientConn

	addresses []resolver.Address
}

func (r *manualResolver) start() error {
	return r.cc.UpdateState(resolver.State{
		Addresses: r.addresses,
	})
}

func (r *manualResolver) ResolveNow(o resolver.ResolveNowOptions) {
}

func (r *manualResolver) Close() {
}

func hostsToAddressSlice(hosts []string) []resolver.Address {
	return sliceKit.ConvertElementType(hosts, func(item string, index int) resolver.Address {
		return resolver.Address{
			Addr: item,
		}
	})
}
