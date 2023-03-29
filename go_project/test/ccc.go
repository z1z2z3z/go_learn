/*
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func TranslateEn2Ch(text string) (string, error) {

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:10809")
	}

	// proxy,_ := url.Parse("http://127.0.0.1:10809")

	transport := http.Transport{
		Proxy: proxy,
	}

	client := http.Client{
		Transport: &transport,
		Timeout:   10 * time.Second,
	}

	fmt.Println("url转码", url.QueryEscape(text))
	urlS := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=zh-cn&tl=en&dt=t&q=%s", url.QueryEscape(text))
	resp, err := client.Get(urlS)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	reg,_:=regexp.Compile("您的IP地址是：\\[.+?\\]");
	sproxy:=reg.FindString(string(bs))
	fmt.Println("代理地址",sproxy)

	//返回的json反序列化比较麻烦, 直接字符串拆解
	ss := string(bs)
	// fmt.Println("语言",ss)
	ss = strings.ReplaceAll(ss, "[", "")
	ss = strings.ReplaceAll(ss, "]", "")
	ss = strings.ReplaceAll(ss, "null,", "")
	ss = strings.Trim(ss, `"`)
	ps := strings.Split(ss, `","`)
	return ps[0], nil
}
func main() {
	str, err := TranslateEn2Ch("www.topgoer.com是个不错的go语言中文文档")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
