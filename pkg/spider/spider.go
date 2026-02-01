package spider

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func SpiderUrl() *rod.Browser {
	// 启动浏览器（调试时可设为false查看过程）
	url := launcher.New().
		Headless(false). // 先设为false调试，确认能抓取到后再改回true
		Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36").
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	return browser
}
