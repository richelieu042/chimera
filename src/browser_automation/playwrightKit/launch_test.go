package playwrightKit

import (
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/richelieu042/chimera/v3/src/consts"
	"github.com/richelieu042/chimera/v3/src/core/pathKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
)

func TestLaunchBrowser(t *testing.T) {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}

		console.Infof("wd: [%s].", wd)
	}

	url := "https://www.moulem.com/"

	pw, browser, err := LaunchBrowser(BrowserNameChromium, "_playwright-deps", true, &playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		panic(err)
	}
	defer pw.Stop()
	defer browser.Close() // defer语句的执行顺序是从上往下：先关 browser，再关 pw

	/* (1) 创建浏览器上下文 */
	bctx, err := browser.NewContext(playwright.BrowserNewContextOptions{
		UserAgent: playwright.String(
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
				"AppleWebKit/537.36 (KHTML, like Gecko) " +
				"Chrome/131.0.0.0 Safari/537.36",
		),

		ExtraHttpHeaders: map[string]string{
			"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
			// 👇 建议加上这个，让 Accept 头也更真实
			"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9," +
				"image/avif,image/webp,image/apng,*/*;q=0.8",
		},

		Viewport: &playwright.Size{
			Width:  1440,
			Height: 800,
		},
		// 👇 建议加上 Screen，和 Viewport 保持一致
		Screen: &playwright.Size{
			Width:  1440,
			Height: 800,
		},

		Locale:            playwright.String("zh-CN"),
		TimezoneId:        playwright.String("Asia/Shanghai"),
		DeviceScaleFactor: playwright.Float(1.0), // 👈 建议显式设置
	})

	/* (2) 创建新页面（即浏览器的一个tab） */
	page, err := bctx.NewPage()
	if err != nil {
		panic(err)
	}
	defer page.Close()

	/* (3) 设置默认超时 */
	page.SetDefaultTimeout(30_000) // 单位: ms

	if _, err = page.Goto(url); err != nil {
		panic(err)
	}

	console.Info("sleep starts")
	time.Sleep(time.Second * 3)
	console.Info("sleep ends")

	// 输入文本
	loc := page.Locator("input#search")
	count, err := loc.Count()
	if err != nil {
		panic(err)
	}
	console.Infof("count: %d", count)
	if err := loc.Fill("hello world!"); err != nil {
		panic(err)
	}
	// 点击按钮
	btnLoc := page.Locator("input#searchBtn")
	btnCount, err := btnLoc.Count()
	if err != nil {
		panic(err)
	}
	console.Infof("btnCount: %d", btnCount)
	if err := btnLoc.Click(); err != nil {
		panic(err)
	}

	select {}
}
