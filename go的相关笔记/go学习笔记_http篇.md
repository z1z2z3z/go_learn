### Go标准库net/http创建客户端(发送get、post请求)

# 1. Get 请求(response body是不能多次读取的)

## 1.1 使用 net/http 包的快捷方法 GET

```go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    resp, err := http.Get("http://www.baidu.com")
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(body))
}
```

## 1.2 自定义客户端

```go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    client := &http.Client{}
    request, err := http.NewRequest("GET", "http://www.baidu.com", nil)
    if err != nil {
        fmt.Println(err)
    }

    resp, err := client.Do(request)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(body))
}
```

使用自定义 `HTTP` 客户端意味着可对请求设置报头、基本身份验证和 cookies 。鉴于使用快捷方法和自定义HTTP 客户端时， 发出请求所需代码的差别很小， 建议除非要完成的任务非常简单，否则都使用自定义HTTP 客户端。

# 2. POST 请求

```go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

func main() {
    data := strings.NewReader(`{"some": "json"}`)
    resp, err := http.Post("https://httpbin.org/post", "application/json", data)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(body))
}


//结果
{
  "args": {}, 
  "data": "{\"some\": \"json\"}", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Content-Length": "16", 
    "Content-Type": "application/json", 
    "Host": "httpbin.org", 
    "User-Agent": "Go-http-client/2.0", 
    "X-Amzn-Trace-Id": "Root=1-60575025-22341e95217463712e18068e"
  }, 
  "json": {
    "some": "json"
  }, 
  "origin": "192.168.0.110", 
  "url": "https://httpbin.org/post"
}
```

# 3. 调试 HTTP

net/http/httputil 也提供了能够让您轻松调试 HTTP 客户端和服务器的方法。这个包中的方法DumpRequestOut 和 DumpResponse 能够让您查看请求和响应。

可对前一个示例进行改进，以使用 net/http/httputil 包中的 DumpRequestOut 和 DumpResponse方法来支持日志功能。这些方法显示请求和响应的报头，还有返回的响应体.

```go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/httputil"
    "strings"
)

func main() {
    client := &http.Client{}
    data := strings.NewReader(`{"some": "json"}`)

    request, err := http.NewRequest("POST", "https://httpbin.org/post", data)
    request.Header.Add("Accept", "application/json") // 增加请求报文头
    /*
        通过使用 Accept 报头， 客户端告诉服务器它想要的是 application/json，而服务器返回数
        据时将 Content-Type 报头设置成了application/json。

    */
    if err != nil {
        fmt.Println(err)
    }

    debugReq, err := httputil.DumpRequestOut(request, true)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("debugReq is ", string(debugReq))

    resp, err := client.Do(request)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()

    debugResponse, err := httputil.DumpResponse(resp, true)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("debugResponse is ", string(debugResponse))

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(body))
}
```

# 4. 处理超时

使用默认的 `HTTP` 客户端时，没有对请求设置超时时间。这意味着如果服务器没有响应，则请求将无限期地等待或挂起。对于任何请求，都建议设置超时时间。这样如果请求在指定的时间内没有完成， 将返回错误。

```go
    client := &http.Client{
        Timeout: 1 * time.Second,
    }
```

上述配置要求客户端在 1s 内完成请求。但是因为服务器的响应速度不够快。完全有可能发生请求超时的情况。如下：

```go
Post https://httpbin.org/post: net/http:
 request canceled (Client.Timeout exceeded while awaiting headers)
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x40 pc=0x63491b]
```

# 5.请求后要加resp.body.close()原因

原因：[参考链接](https://juejin.cn/post/6987372070120194055)

**<mark>defer使用注意事项：</mark>**

* 如果有多个defer表达式，调用顺序类似于栈，越后面的defer表达式越先被调用。

* [defer真正的执行顺序](https://blog.51cto.com/u_15127535/4037146)

* **return xxx这一条语句并不是一条原子指令!**  (不可拆分:原子指令)

* return ××× 可以拆分为先赋值  后返回  defer就在这之间运行

**代码实例**

```go
package main

import "fmt"


func f1() (result int) {
    defer func ()  {
        result++
    }()
    return 0
}

func f2() (r int) {
    t := 5
    defer func ()  {
        t = t + 5
    }()
    return t
}


///当defer被声明时，其参数就会被实时解析
func f3() (r int) {

    ////不会改变要返回的那个r值
    defer func (r int)  {
        r = r + 5
        fmt.Println(r)
    }(r) ///  这里的值就已经定了就是默认值0
    return 1
}

func main()  {
    fmt.Println(f1())
    fmt.Println(f2())
    fmt.Println(f3())

}
```
