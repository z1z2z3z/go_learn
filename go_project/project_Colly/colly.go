/*
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	// "github.com/gocolly/colly/proxy"
)

func main()  {
	

	c := colly.NewCollector(
		colly.Async(true),
	)

	// 在访问的时候，使用随机的UserAgent，来模拟不同的浏览器访问
	extensions.RandomUserAgent(c)

	// 在访问的时候带上Referrer，意思就是这一次点击是从哪个页面产生的
	extensions.Referer(c)

	// 设置代理
	// rp, err := proxy.RoundRobinProxySwitcher()
	// c.SetProxy("")

	c.Limit(&colly.LimitRule{
		DomainGlob: "*.douban.*",
		Parallelism: 5,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting",r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnHTML(".hd", func(e *colly.HTMLElement) {
		fmt.Println(strings.Split(e.ChildAttr("a", "href"), "/")[4],
			strings.TrimSpace(e.DOM.Find("span.title").Eq(0).Text()))
    })

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.Visit("https://movie.douban.com/top250?start=0&filter=")
	c.Wait()
}