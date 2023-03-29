package main

import (
    "encoding/json"
    "fmt"
    "github.com/gocolly/colly"
)

type Listing struct {
    Price string `json:"price"`
}

type NFT struct {
    Name    string  `json:"name"`
    Listing Listing `json:"listing"`
}

func main() {
    c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"),
        colly.AllowURLRevisit(),
	)

    var nfts []NFT

	c.OnRequest(func(r *colly.Request) {
        r.Headers.Set("Host", "www.gem.xyz")
        r.Headers.Set("Connection", "keep-alive")
        r.Headers.Set("Cache-Control", "max-age=0")
        r.Headers.Set("Upgrade-Insecure-Requests", "1")
        r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
        r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
        r.Headers.Set("Referer", "https://www.gem.xyz/")
        r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
        r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
		r.Headers.Set("x-api-key","rLnNH1tdrT09EQjGsjrSS7V3uGonfZLW")
        // r.Headers.Set("Cookie", "_ga=GA1.2.XXXXXXXXX.XXXXXXXXXX; _gid=GA1.2.XXXXXXXXX.XXXXXXXXXX; _gat_gtag_UA_XXXXXXXXX_1=1; __cfduid=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
    })

    c.OnHTML(".jsx-2959205572.token-card", func(e *colly.HTMLElement) {
        var nft NFT
        nft.Name = e.ChildText(".jsx-2959205572.name")
        json.Unmarshal([]byte(e.ChildAttr("script[type=\"application/ld+json\"]", "innerHTML")), &nft.Listing)
        nfts = append(nfts, nft)
    })

    c.Visit("https://www.gem.xyz/collection/boredapeyachtclub/")

    for _, nft := range nfts {
        fmt.Printf("%s - %s\n", nft.Name, nft.Listing.Price)
    }
}
