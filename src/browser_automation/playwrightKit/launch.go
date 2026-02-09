package playwrightKit

import (
	"os"
	"strings"

	"github.com/playwright-community/playwright-go"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// LaunchBrowser
/*
@param browserName 		浏览器名称
@param driverDir 		用于存放浏览器和操作系统的依赖（初次执行会下载，会有点慢）
@param launchOptions 	可以为nil，将采用默认值(headless == true)
*/
func LaunchBrowser(browserName string, driverDir string, installFlag bool,
	launchOptions *playwright.BrowserTypeLaunchOptions) (pw *playwright.Playwright, browser playwright.Browser, err error) {

	defer func() {
		if err != nil {
			if browser != nil {
				_ = browser.Close()
				browser = nil
			}
			if pw != nil {
				_ = pw.Stop()
				pw = nil
			}
		}
	}()

	if err = fileKit.AssertNotExistOrIsDir(driverDir); err != nil {
		return
	}
	if err = fileKit.MkDirs(driverDir); err != nil {
		return
	}

	if browserName == "" {
		browserName = BrowserNameChromium
	}
	browserName = strings.ToLower(browserName)
	switch browserName {
	case BrowserNameChromium, BrowserNameFirefox, BrowserNameWebkit:
	default:
		err = errorKit.Newf("invalid browserName(%s)", browserName)
		return
	}

	runOptions := &playwright.RunOptions{
		DriverDirectory:     driverDir,
		SkipInstallBrowsers: false,
		Browsers:            []string{browserName},
		Verbose:             true,
		Stdout:              os.Stdout,
		Stderr:              os.Stderr,
	}

	if installFlag {
		err = playwright.Install(runOptions)
		if err != nil {
			err = errorKit.Wrapf(err, "fail to install dependencies")
			return
		}
	}
	pw, err = playwright.Run(runOptions)
	if err != nil {
		err = errorKit.Wrapf(err, "fail to run playwright")
		return
	}

	var tmp []playwright.BrowserTypeLaunchOptions = nil
	if launchOptions != nil {
		tmp = append(tmp, *launchOptions)
	}
	switch browserName {
	case BrowserNameChromium:
		browser, err = pw.Chromium.Launch(tmp...)
	case BrowserNameFirefox:
		browser, err = pw.Firefox.Launch(tmp...)
	case BrowserNameWebkit:
		browser, err = pw.WebKit.Launch(tmp...)
	default:
		// Richelieu: 理论上代码不会走到此处
		err = errorKit.Newf("invalid browserName(%s)", browserName)
		return
	}
	if err != nil {
		err = errorKit.Wrapf(err, "fail to launch browser(%s)", browserName)
	}
	return
}
