package demo4

import (
	"fmt"
	"github.com/Cao-yu-xin1/SpiderCrawl/pkg/spider"
	"time"

	"github.com/go-rod/rod"
)

func Spider3(page *rod.Page) {
	browser := spider.SpiderUrl()
	//获取当前页信息(处理可能得页面跳转情况)
	var detailPage *rod.Page

	//检查是否有新窗口打开
	pages := browser.MustPages()
	if len(pages) > 1 {
		//如果开了新的标签页，则打开新的标签页
		detailPage = pages[len(pages)-1]
	} else {
		//否则使用当前页（本页跳转）
		detailPage = page
	}

	//确认详情页已经加载完毕
	detailPage.MustWaitLoad()

	//等待详情页关键元素出现（例如商品标题或者图片）
	detailPage.Timeout(10 * time.Second).MustElement(`[class*="title"], [class*="detail"], img`)
	fmt.Printf("已经进去详情页,URL:%s\n", detailPage.MustInfo().URL)

	//在详情页执行操作（实例：获取标题和截图）
	title, err := detailPage.Element(`h1, .title .[class*="title"]`)
	if err == nil {
		titleText, _ := title.Text()
		fmt.Printf("商品标题：%s\n", titleText)
	}
	// 截图保存详情页
	detailPage.MustScreenshot("detail_page.png")
	fmt.Println("详情页已截图保存为 detail_page.png")
	//可选：获取详情页的特定信息
	//例如获取所有图片
	images := detailPage.MustElements(`img`)
	fmt.Printf("详情页弓%d张图片\n", len(images))
	for i, image := range images {
		if src, err := image.Attribute("src"); err == nil && src != nil {
			fmt.Printf("图片 %d: %s\n", i+1, *src)
		} else {
			fmt.Println("未找到商品标题元素:", err)
			detailPage.MustScreenshot("title_debug.png") // 截图调试
		}
	}
}
