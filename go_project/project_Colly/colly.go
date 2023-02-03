/*
 * @Description: 请填写简介
 */
package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		// colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	// 在请求之前调用
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Vis", r.URL)
	})

	// 在请求中出现错误时调用
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// 响应接收到之后调用
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	// OnResponse 正确执行后，如果接收到的文本是HTML时执行
	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	// 在OnXML/OnHTML回调完成后调用
    c.OnScraped(func(r *colly.Response) {
        fmt.Println("Finished", r.Request.URL)
    })

	c.Visit("https://movie.douban.com/top250?start=0&filter=")
}
