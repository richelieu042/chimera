package redisKit

import (
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/validateKit"
)

type (
	Config struct {
		UserName string `json:"userName" yaml:"userName"`
		// Password
		/*
		  如果Redis有密码，此处配置必须配置;
		  如果Redis无密码，此处配置配不配置皆可（配置了也不会报错）.
		*/
		Password  string `json:"password" yaml:"password"`
		KeyPrefix string `json:"keyPrefix" yaml:"keyPrefix"`

		Mode Mode `json:"mode" yaml:"mode" validate:"oneof=single sentinel cluster"`

		Single *SingleConfig `json:"single" yaml:"single" validate:"required_if=Mode single"`
		//MasterSlave *MasterSlaveConfig `json:"masterSlave" yaml:"masterSlave" validate:"required_if=Mode masterSlave"`
		Sentinel *SentinelConfig `json:"sentinel" yaml:"sentinel" validate:"required_if=Mode sentinel"`
		Cluster  *ClusterConfig  `json:"cluster" yaml:"cluster" validate:"required_if=Mode cluster"`
	}

	SingleConfig struct {
		// Addr address(host:port)
		Addr string `json:"addr" yaml:"addr" validate:"hostname_port"`

		// DB Database to be selected after connecting to the server.
		DB int `json:"db" yaml:"db" validate:"gte=0"`
	}

	//MasterSlaveConfig struct {
	//}

	SentinelConfig struct {
		// MasterName The master name.
		MasterName string `json:"masterName" yaml:"masterName"`

		// Addrs A seed list of host:port addresses of sentinel nodes.
		Addrs []string `json:"addrs" yaml:"addrs" validate:"required,gte=2,unique,dive,hostname_port"`

		DB int `json:"db" yaml:"db" validate:"gte=0"`
	}

	ClusterConfig struct {
		// Addrs
		/*
			A seed list of host:port addresses of cluster nodes.
			可以是: 所有的 master 的地址，
			也可以是: 所有的 master + slave 的地址（推荐）.
		*/
		Addrs []string `json:"addrs" yaml:"addrs" validate:"required,gte=2,unique,dive,hostname_port"`

		UseReplicasForReadOperations bool `json:"useReplicasForReadOperations" yaml:"useReplicasForReadOperations"`
	}
)

// Simplify 简化配置.
func (config *Config) Simplify() {
	if config == nil {
		return
	}

	switch config.Mode {
	case ModeSingle:
		config.Sentinel = nil
		config.Cluster = nil
	case ModeSentinel:
		config.Single = nil
		config.Cluster = nil
	case ModeCluster:
		config.Single = nil
		config.Sentinel = nil
	case ModeMasterSlave:
		fallthrough
	default:
		// do nothing
	}
}

// Validate
/*
@param config 可能为nil
*/
func (config *Config) Validate() error {
	if err := interfaceKit.AssertNotNil(config, "config"); err != nil {
		return err
	}

	/* 先简化，再验证（以免通不过验证） */
	config.Simplify()

	if err := validateKit.Struct(config); err != nil {
		return errorKit.Wrapf(err, "fail to verify")
	}
	return nil
}
