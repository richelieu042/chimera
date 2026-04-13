package ip2RegionKit

import (
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
)

var NotSetupError = errorKit.Newf("haven’t been set up correctly")

// 缓存整个xdb数据的情况下，searcher对象可以安全用于并发
var (
	ipv4Searcher *xdb.Searcher

	ipv6Searcher *xdb.Searcher
)

func MustSetUp(ipv4XdbPath, ipv6XdbPath string) {
	err := SetUp(ipv4XdbPath, ipv6XdbPath)
	if err != nil {
		console.Fatalf("Fail to set up, error: %+v", err)
	}
}

func SetUp(ipv4XdbPath, ipv6XdbPath string) (err error) {
	defer func() {
		if err != nil {
			ipv4Searcher = nil
			ipv6Searcher = nil
		}
	}()

	ipv4Searcher, err = loadXdbFile(xdb.IPv4, ipv4XdbPath)
	if err != nil {
		return
	}
	ipv6Searcher, err = loadXdbFile(xdb.IPv6, ipv6XdbPath)
	return
}

// loadXdbFile
/*
@param path xdb文件的路径
*/
func loadXdbFile(version *xdb.Version, xdbPath string) (*xdb.Searcher, error) {
	if err := interfaceKit.AssertNotNil(version, "version"); err != nil {
		return nil, err
	}
	if err := fileKit.AssertExistAndIsFile(xdbPath); err != nil {
		return nil, err
	}

	// 缓存整个xdb数据
	cBuff, err := xdb.LoadContentFromFile(xdbPath)
	if err != nil {
		return nil, err
	}
	return xdb.NewWithBuffer(version, cBuff)
}
