package demo2

import (
	"fmt"
	"github.com/Cao-yu-xin1/SpiderCrawl/pkg/spider"
	"time"
)

func ImageCrawler(pageUrl string, elementSelector string) {
	browser := spider.SpiderUrl()
	defer browser.MustClose()

	// 访问1688页面（移除URL中的空格）
	page := browser.MustPage(pageUrl)

	// 等待页面基本加载完成
	page.MustWaitLoad()

	// 关键：等待特定元素出现（修正选择器并增加等待时间）
	err := page.Timeout(10 * time.Second).MustElement(elementSelector).WaitVisible()
	if err != nil {
		fmt.Println("等待元素超时:", err)
		return
	}

	// 简短暂停，确保所有图片加载完成
	time.Sleep(2 * time.Second)

	// 获取所有匹配的元素（修正CSS选择器）
	elements := page.MustElements(elementSelector)

	fmt.Printf("找到 %d 个商品图片:\n", len(elements))

	// 遍历并提取src属性
	for i, el := range elements {
		// 确保元素可见
		if el.MustVisible() {
			src, err := el.Attribute("src")
			if err == nil && src != nil {
				fmt.Printf("%d. %s\n", i+1, *src)
			}
		}
	}
}

//func main() {
//	ImageCrawler("https://air.1688.com/kapp/channel-fe/cps-4c-pc/sytm?type=1&offerIds=660390230106,574965204819,949033739317",
//		`.offer-item__img`)
//}
