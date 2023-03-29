/*
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
	// "github.com/gocolly/colly/extensions"
	// "github.com/gocolly/colly/proxy"
)

// func main() {

// 	c := colly.NewCollector(
// 		colly.AllowURLRevisit(),
// 		colly.Async(true),
// 	)

// 	// 在访问的时候，使用随机的UserAgent，来模拟不同的浏览器访问
// 	extensions.RandomUserAgent(c)

// 	// 在访问的时候带上Referrer，意思就是这一次点击是从哪个页面产生的
// 	extensions.Referer(c)

// 	// 设置代理
// 	// rp, err := proxy.RoundRobinProxySwitcher()
// 	// c.SetProxy("")

// 	c.Limit(&colly.LimitRule{
// 		DomainGlob:  "*.douban.*",
// 		Parallelism: 5,
// 		RandomDelay: 1 * time.Second,
// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 	})

// 	c.OnError(func(r *colly.Response, err error) {
// 		fmt.Println("Something went wrong:", err)
// 	})

// 	c.OnHTML(".hd", func(e *colly.HTMLElement) {
// 		fmt.Println(strings.Split(e.ChildAttr("a", "href"), "/")[4],
// 			strings.TrimSpace(e.DOM.Find("span.title").Eq(0).Text()))
// 	})

// 	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
// 		e.Request.Visit(e.Attr("href"))
// 	})

// 	c.Visit("https://movie.douban.com/top250?start=0&filter=")
// 	c.Wait()
// }

func main1() {
	c := colly.NewCollector(
		colly.AllowedDomains("emojipedia.org"),
	)

	// Callback for when a scraped page contains an article element
	c.OnHTML("article", func(e *colly.HTMLElement) {
		isEmojiPage := false

		// Extract meta tags from the document
		metaTags := e.DOM.ParentsUntil("~").Find("meta")
		metaTags.Each(func(_ int, s *goquery.Selection) {
			// Search for og:type meta tags
			property, _ := s.Attr("property")
			if strings.EqualFold(property, "og:type") {
				content, _ := s.Attr("content")

				// Emoji pages have "article" as their og:type
				isEmojiPage = strings.EqualFold(content, "article")
			}
		})

		if isEmojiPage {
			// Find the emoji page title
			fmt.Println("Emoji: ", e.DOM.Find("h1").Text())
			// Grab all the text from the emoji's description
			fmt.Println(
				"Description: ",
				e.DOM.Find(".description").Find("p").Text())
		}
	})

	// Callback for links on scraped pages
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Extract the linked URL from the anchor tag
		link := e.Attr("href")
		fmt.Println("URL",link)
		// Have our crawler visit the linked URL
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://emojipedia.org")
}