package cookieKit

import (
	"net/http"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/urlKit"
)

type (
	options struct {
		// Path 可以为""（此时将采用默认值"/"）
		Path string

		// Domain 可以为""（此时为对应url的hostname，可能是ip）
		Domain string

		// MaxAge
		/*
			(1) > 0:	表示此cookie从创建到过期所能存在的时间，以秒为单位
			(2) == 0:	存储在浏览器进程中，进程退出cookie就会消失
			(3)	< 0:	删除此cookie（一般用-1）
		*/
		MaxAge int

		// Secure
		/*
			true:	浏览器只会在https、SSL等安全协议中传输此Cookie，不会在不安全的HTTP协议中传输此Cookie
			false:	此Cookie在所有协议中传输
		*/
		Secure bool

		// HttpOnly
		/*
			true:	不能被js检测到（读||写），发送请求时依旧会携带此Cookie（允许的话）。
			false:	能被js检测到（读||写），发送请求时依旧会携带此Cookie（允许的话）
		*/
		HttpOnly bool

		// SameSite
		/*
			http.SameSiteDefaultMode:	默认值
			http.SameSiteStrictMode:	Scrict最为严格，完全禁止第三方Cookie，跨站点时，任何情况下都不会发送Cookie。
			http.SameSiteLaxMode:		Lax规则稍稍放宽，大多数情况也是不发送第三方 Cookie，但是导航到目标网址的 Get 请求除外。
			http.SameSiteNoneMode:		网站可以选择显式关闭SameSite属性，将其设为None。不过，前提是必须同时设置Secure属性（Cookie 只能通过 HTTPS 协议发送），否则无效。
		*/
		SameSite http.SameSite
	}

	CookieOption func(*options)
)

func WithPath(path string) CookieOption {
	return func(opts *options) {
		opts.Path = path
	}
}

func WithDomain(domain string) CookieOption {
	return func(opts *options) {
		opts.Domain = domain
	}
}

func WithMaxAge(maxAge int) CookieOption {
	return func(opts *options) {
		opts.MaxAge = maxAge
	}
}

func WithSecure(secure bool) CookieOption {
	return func(opts *options) {
		opts.Secure = secure
	}
}

func WithHttpOnly(httpOnly bool) CookieOption {
	return func(opts *options) {
		opts.HttpOnly = httpOnly
	}
}

func WithSameSite(sameSite http.SameSite) CookieOption {
	return func(opts *options) {
		opts.SameSite = sameSite
	}
}

func loadOptions(cookieOptions ...CookieOption) *options {
	opts := &options{
		Path:     "/",
		Domain:   "",
		MaxAge:   0,
		Secure:   false,
		HttpOnly: false,
		SameSite: http.SameSiteDefaultMode,
	}
	for _, option := range cookieOptions {
		option(opts)
	}

	// 参考: gin.Context.SetCookie()
	opts.Path = strKit.BlankToDefault(opts.Path, "/")

	return opts
}

// NewCookie
/*
e.g. 跨域情况下，iframe页面（子页面）设置cookie失败的一种解决方法
满足以下条件:
	(1) WithSameSite(http.SameSiteNoneMode)
	(2) WithSecure(true)
	(3) https协议
*/
func NewCookie(name, value string, cookieOptions ...CookieOption) *http.Cookie {
	opts := loadOptions(cookieOptions...)

	cookie := &http.Cookie{
		Name:  name,
		Value: urlKit.EncodeURIComponent(value),

		Path:     opts.Path,
		Domain:   opts.Domain,
		MaxAge:   opts.MaxAge,
		Secure:   opts.Secure,
		HttpOnly: opts.HttpOnly,
		SameSite: opts.SameSite,
	}
	SetCookieMaxAge(cookie, opts.MaxAge)
	return cookie
}
