package playwrightKit

import (
	"github.com/playwright-community/playwright-go"
)

// GetBrowserContextPages
/*
@param ctx 不能为nil
@return 有可能len(pages) == 0
*/
func GetBrowserContextPages(ctx playwright.BrowserContext) (pages []playwright.Page) {
	pages = ctx.Pages()
	return
}

// GetBrowserPages
/*
@param browser 不能为nil
@return 有可能len(pages) == 0
*/
func GetBrowserPages(browser playwright.Browser) (pages []playwright.Page) {
	contexts := browser.Contexts()
	for _, context := range contexts {
		pages = append(pages, context.Pages()...)
	}
	return
}
