package demo1

import (
	"fmt"
	"github.com/Cao-yu-xin1/SpiderCrawl/pkg/spider"
	"time"
)

func TitleCrawler(pageUrl string, elementSelector string) {
	browser := spider.SpiderUrl()
	defer browser.MustClose()

	// 访问1688页面
	page := browser.MustPage(pageUrl)

	// 等待页面基本加载完成
	page.MustWaitLoad()

	// 1688页面有动态加载，增加额外等待时间
	// 等待元素出现，最长等待10秒
	err := page.Timeout(10 * time.Second).MustElement(elementSelector).WaitVisible()
	if err != nil {
		fmt.Println("等待元素超时:", err)
		return
	}

	// 获取所有匹配的元素
	elements := page.MustElements(elementSelector)

	fmt.Printf("找到 %d 个商品标题:\n", len(elements))

	// 遍历并提取文本
	for i, el := range elements {
		// 确保元素在DOM中且可见
		if err := el.WaitVisible(); err == nil {
			text := el.MustText()
			fmt.Printf("%d. %s\n", i+1, text)
		}
	}

	// 如果需要更精确的选择（例如只获取可见的）
	fmt.Println("\n--- 只获取可见元素 ---")
	visibleElements := page.MustElements(`.offer-title.ellipsis`)
	for i, el := range visibleElements {
		// 检查元素是否可见且在视口中
		if el.MustVisible() {
			text := el.MustText()
			fmt.Printf("%d. %s\n", i+1, text)
		}
	}
}

//func main() {
//	TitleCrawler("https://air.1688.com/kapp/channel-fe/cps-4c-pc/sytm?type=1&offerIds=660390230106,574965204819,949033739317",
//		`.offer-title.ellipsis`)
//}
