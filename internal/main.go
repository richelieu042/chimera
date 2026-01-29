package main

import (
	"fmt"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
)

type Config struct {
	ProductName   string        // 商品名称关键词
	CheckInterval time.Duration // 检查间隔
	MaxRetries    int           // 最大重试次数
	Headless      bool          // 是否无头模式
	SlowMo        int           // 慢速模式（毫秒）
	WaitTimeout   time.Duration // 等待超时时间
}

type SamsClubBot struct {
	config  Config
	browser playwright.Browser
	context playwright.BrowserContext
	page    playwright.Page
	pw      *playwright.Playwright
}

func NewSamsClubBot(config Config) *SamsClubBot {
	return &SamsClubBot{
		config: config,
	}
}

func (bot *SamsClubBot) Initialize() error {
	log.Println("初始化 Playwright...")

	pw, err := playwright.Run()
	if err != nil {
		return fmt.Errorf("无法启动 Playwright: %v", err)
	}
	bot.pw = pw

	log.Println("启动浏览器...")
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(bot.config.Headless),
		SlowMo:   playwright.Float(float64(bot.config.SlowMo)),
	})
	if err != nil {
		return fmt.Errorf("无法启动浏览器: %v", err)
	}
	bot.browser = browser

	// 创建浏览器上下文
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		Locale:     playwright.String("zh-CN"),
		TimezoneId: playwright.String("Asia/Shanghai"),
		Viewport:   &playwright.Size{Width: 1920, Height: 1080},
		UserAgent:  playwright.String("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	})
	if err != nil {
		return fmt.Errorf("无法创建浏览器上下文: %v", err)
	}
	bot.context = context

	// 创建新页面
	page, err := context.NewPage()
	if err != nil {
		return fmt.Errorf("无法创建页面: %v", err)
	}
	bot.page = page

	// 设置默认超时
	page.SetDefaultTimeout(float64(bot.config.WaitTimeout.Milliseconds()))

	log.Println("初始化完成")
	return nil
}

func (bot *SamsClubBot) Close() {
	if bot.browser != nil {
		bot.browser.Close()
	}
	if bot.pw != nil {
		bot.pw.Stop()
	}
}

func (bot *SamsClubBot) Login() error {
	log.Println("导航到山姆会员店首页...")
	if _, err := bot.page.Goto("https://www.samsclub.cn/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		return fmt.Errorf("无法打开网站: %v", err)
	}

	log.Println("请在浏览器中手动完成登录...")
	log.Println("等待登录完成（检测到用户信息后将自动继续）...")

	// 等待登录成功的标志（用户头像或用户名元素）
	// 这里需要根据实际页面结构调整选择器
	_, err := bot.page.WaitForSelector(".user-info, .login-user, [class*='user'], [class*='member']", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(120000), // 等待2分钟
		State:   playwright.WaitForSelectorStateVisible,
	})

	if err != nil {
		log.Println("未检测到登录成功标志，将继续尝试...")
		// 给用户额外的时间
		time.Sleep(5 * time.Second)
	} else {
		log.Println("✓ 登录成功！")
	}

	return nil
}

func (bot *SamsClubBot) SearchProduct() error {
	log.Printf("搜索商品: %s\n", bot.config.ProductName)

	// 查找搜索框（需要根据实际页面调整选择器）
	searchSelectors := []string{
		"input[type='search']",
		"input[placeholder*='搜索']",
		"input.search-input",
		"#search-input",
		".search-box input",
	}

	var searchBox playwright.Locator
	for _, selector := range searchSelectors {
		searchBox = bot.page.Locator(selector)
		count, _ := searchBox.Count()
		if count > 0 {
			log.Printf("找到搜索框: %s\n", selector)
			break
		}
	}

	// 输入搜索关键词
	if err := searchBox.Fill(bot.config.ProductName); err != nil {
		return fmt.Errorf("无法输入搜索关键词: %v", err)
	}

	// 提交搜索（按回车或点击搜索按钮）
	if err := searchBox.Press("Enter"); err != nil {
		log.Println("按回车失败，尝试点击搜索按钮...")

		searchButtonSelectors := []string{
			"button[type='submit']",
			".search-button",
			"button.search-btn",
		}

		for _, selector := range searchButtonSelectors {
			btn := bot.page.Locator(selector)
			count, _ := btn.Count()
			if count > 0 {
				if err := btn.Click(); err == nil {
					break
				}
			}
		}
	}

	// 等待搜索结果加载
	time.Sleep(2 * time.Second)
	log.Println("✓ 搜索完成")

	return nil
}

func (bot *SamsClubBot) CheckStockAndAddToCart() (bool, error) {
	log.Println("检查商品库存...")

	// 等待商品列表加载
	bot.page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateNetworkidle,
	})

	// 查找商品卡片（需要根据实际页面结构调整）
	productSelectors := []string{
		".product-item",
		".goods-item",
		"[class*='product']",
		".item-card",
	}

	var products playwright.Locator
	for _, selector := range productSelectors {
		products = bot.page.Locator(selector)
		count, _ := products.Count()
		if count > 0 {
			log.Printf("找到 %d 个商品\n", count)
			break
		}
	}

	// 检查第一个商品（或遍历所有商品找到目标）
	firstProduct := products.First()

	// 检查是否有货（查找"加入购物车"按钮）
	addToCartSelectors := []string{
		"button:has-text('加入购物车')",
		"button:has-text('立即购买')",
		".add-cart",
		"[class*='add-cart']",
	}

	var addButton playwright.Locator
	for _, selector := range addToCartSelectors {
		addButton = firstProduct.Locator(selector)
		count, _ := addButton.Count()
		if count > 0 {
			// 检查按钮是否可用
			isDisabled, _ := addButton.IsDisabled()
			if !isDisabled {
				log.Println("✓ 发现有货商品！")

				// 点击加入购物车
				log.Println("正在加入购物车...")
				if err := addButton.Click(); err != nil {
					return false, fmt.Errorf("无法点击加入购物车: %v", err)
				}

				time.Sleep(1 * time.Second)
				log.Println("✓ 已加入购物车")
				return true, nil
			}
		}
	}

	log.Println("× 商品暂时无货")
	return false, nil
}

func (bot *SamsClubBot) GoToCheckout() error {
	log.Println("前往购物车结算...")

	// 查找购物车图标/按钮
	cartSelectors := []string{
		"a[href*='cart']",
		".cart-icon",
		"[class*='cart']",
	}

	for _, selector := range cartSelectors {
		cart := bot.page.Locator(selector)
		count, _ := cart.Count()
		if count > 0 {
			if err := cart.Click(); err == nil {
				break
			}
		}
	}

	time.Sleep(2 * time.Second)

	// 查找结算按钮
	checkoutSelectors := []string{
		"button:has-text('去结算')",
		"button:has-text('结算')",
		".checkout-btn",
	}

	for _, selector := range checkoutSelectors {
		checkout := bot.page.Locator(selector)
		count, _ := checkout.Count()
		if count > 0 {
			if err := checkout.Click(); err == nil {
				log.Println("✓ 已进入结算页面")
				log.Println("⚠️  请手动确认订单信息并完成支付！")
				return nil
			}
		}
	}

	log.Println("已跳转到购物车，请手动进行结算")
	return nil
}

func (bot *SamsClubBot) MonitorAndBuy() error {
	retries := 0

	for retries < bot.config.MaxRetries {
		log.Printf("\n===== 第 %d 次尝试 =====\n", retries+1)

		// 刷新页面
		if retries > 0 {
			log.Println("刷新页面...")
			bot.page.Reload()
			time.Sleep(2 * time.Second)
		}

		// 检查库存并加入购物车
		success, err := bot.CheckStockAndAddToCart()
		if err != nil {
			log.Printf("错误: %v\n", err)
		}

		if success {
			// 成功加入购物车，前往结算
			if err := bot.GoToCheckout(); err != nil {
				log.Printf("结算失败: %v\n", err)
			}

			log.Println("\n✓ 抢购流程完成！请在浏览器中完成支付。")
			log.Println("浏览器将保持打开状态，按 Ctrl+C 退出程序...")

			// 保持浏览器打开
			select {}
		}

		retries++

		if retries < bot.config.MaxRetries {
			log.Printf("等待 %v 后重试...\n", bot.config.CheckInterval)
			time.Sleep(bot.config.CheckInterval)
		}
	}

	log.Printf("\n已达到最大重试次数 (%d)，程序结束。\n", bot.config.MaxRetries)
	return nil
}

func (bot *SamsClubBot) Run() error {
	// 初始化
	if err := bot.Initialize(); err != nil {
		return err
	}
	defer bot.Close()

	// 登录
	if err := bot.Login(); err != nil {
		return err
	}

	// 搜索商品
	if err := bot.SearchProduct(); err != nil {
		return err
	}

	// 监控并购买
	return bot.MonitorAndBuy()
}

func main() {
	// 配置参数
	config := Config{
		ProductName:   "爱他美奶粉",          // 修改为您要抢购的商品名称
		CheckInterval: 3 * time.Second,  // 检查间隔
		MaxRetries:    100,              // 最大重试次数
		Headless:      false,            // 显示浏览器窗口
		SlowMo:        100,              // 操作延迟（毫秒）
		WaitTimeout:   30 * time.Second, // 等待超时
	}

	log.Println("=================================")
	log.Println("   山姆会员店自动化辅助脚本")
	log.Println("=================================")
	log.Printf("商品名称: %s\n", config.ProductName)
	log.Printf("检查间隔: %v\n", config.CheckInterval)
	log.Printf("最大重试: %d 次\n", config.MaxRetries)
	log.Println("=================================\n")

	bot := NewSamsClubBot(config)

	if err := bot.Run(); err != nil {
		log.Fatalf("程序错误: %v", err)
	}
}
