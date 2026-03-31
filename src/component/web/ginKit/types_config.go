package ginKit

type (
	Config struct {
		Mode string `json:"mode" yaml:"mode" validate:"omitempty,oneof=debug release test"`
		// HostName 一般情况下，用""就够了
		HostName string `json:"hostName" yaml:"hostName" validate:"omitempty,hostname|ipv4|ipv6"`
		// Port
		/*
			0（默认）: 不使用 http port
		*/
		Port         int  `json:"port" yaml:"port" validate:"port|eq=0,necsfield=SSL.Port"`
		DisableColor bool `json:"disableColor" yaml:"disableColor"`
		Pprof        bool `json:"pprof" yaml:"pprof"`

		SSL        SslConfig        `json:"ssl" yaml:"ssl"`
		Middleware MiddlewareConfig `json:"middleware" yaml:"middleware"`
	}

	SslConfig struct {
		// Port
		/*
			0（默认）: 不使用 https port
		*/
		Port     int    `json:"port" yaml:"port" validate:"port|eq=0"`
		CertFile string `json:"certFile" yaml:"certFile" validate:"file_unless=Port 0"`
		KeyFile  string `json:"keyFile" yaml:"keyFile" validate:"file_unless=Port 0"`
	}

	MiddlewareConfig struct {
		BodyLimit int64 `json:"bodyLimit" yaml:"bodyLimit"`
		//Gzip      int   `json:"gzip" yaml:"gzip" validate:"min=-1,max=9"`
		//XFrameOptions string `json:"xFrameOptions" yaml:"xFrameOptions" validate:"omitempty,lowercase,oneof=deny sameorigin|startswith=allow-from "`

		Gzip *GzipConfig `json:"gzip" yaml:"gzip"`

		Cors CorsConfig `json:"cors" yaml:"cors"`

		RateLimiter *RateLimiterConfig `json:"rateLimiter" yaml:"rateLimiter"`

		ResponseHeaders map[string]string `json:"responseHeaders" yaml:"responseHeaders" mapstructure:"responseHeaders"`
	}

	// CorsConfig cors（跨源资源共享）的配置
	CorsConfig struct {
		Access bool `json:"access" yaml:"access"`

		// AllowCredentials
		/*
			true: 	设置响应头 "Access-Control-Allow-Credentials" 为 "true"
			false:	do nothing
		*/
		AllowCredentials bool `json:"allowCredentials" yaml:"allowCredentials"`

		Origins []string `json:"origins" yaml:"origins" validate:"unique,dive,required"`
	}

	// RateLimiterConfig 限流器（令牌桶算法）的配置
	RateLimiterConfig struct {
		R int `json:"r" yaml:"r" validate:"gt=0"`
		B int `json:"b" yaml:"b" validate:"gt=0,gtecsfield=R"`
	}

	GzipConfig struct {
		Level int `json:"level" yaml:"level" validate:"min=-1,max=9"`

		//MinLength int `json:"minLength" yaml:"minLength" validate:"gt=0"`
		MinLength int `json:"minLength" yaml:"minLength"`

		ExcludedExtensions   []string `json:"excludedExtensions" yaml:"excludedExtensions"`
		ExcludedPaths        []string `json:"excludedPaths" yaml:"excludedPaths"`
		ExcludedPathsRegexps []string `json:"excludedPathsRegexps" yaml:"excludedPathsRegexps"`
	}
)
