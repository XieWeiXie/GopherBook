
`web` 开发是软件开发领域一个重要的组成部分，`Go` 凭借着出色的语言特性，在各个领域内都有它的影子，比如微服务、中间件、区块链等。

同样凭借着出色的开源能力、完善的社区生态，`Go` 在`web` 领域也表现的非常不错。使用内置库`net/http` 可以快速的构建 `web` 服务，丰富的第三方 `web` 框架，诸如：`gin、echo、iris、beego` 等，使得整个生态更为完善。

本章使用内置 `net/http` 库和第三方 `web` 框架 `iris` 构建 `web` 服务，核心是掌握 `Restful` 风格的`API` 设计。整体章节的学习，读者可以对原生的构建 `web` 服务和使用第三方框架能很好的掌握。

- 使用原生 `net/http` 和 `html/template` 构建简易的 `web` 服务
- 从企业级服务出发，使用原生的 `net/http` 构建完善的 `web` 服务
- 从企业级服务出发，使用第三方 `iris` 构建完善的 `web` 服务



## 1. `net/http` 构建简易的 `web` 服务

主要内容:

- 如何启动 `web` 服务
- 如何使用 `template` 及其相应的应用
- 如何使用中间件
- 如何设计整个系统



## 1.1 如何启动 `web` 服务

在 `Go` 中短短的几行代码，便可以启动一个简易的 `web` 服务，源于 `Go` 在语言层面原生支持并发，所以其特别适用于构建高性能的 `web` 服务。

在内置库常用的操作章节中，已经介绍过相应的内置库 `net/http` , 启动 `web` 服务靠的就是内置库 `net/http` 的使用。

访问网页是通过浏览器渲染出前端页面，网页中有些内容是静态的，即不可变，有些内容是动态的，可以和服务器交互，进行数据层面的交互，这样就能动态的改变网页的内容。

真实的企业开发 `web` 服务的内容通常要和这几种职业的人打交道：

- 产品经理梳理真实的用户需求，对产品进行定位，竞品分析等，决定产品的走向
- 设计人员根据产品经理的需求和产品原型图进行设计
- 前后端分离的情况下，前端人员或者客户端（APP 或者 IOS）根据后端提供的接口和服务器交互，根据设计人员的设计稿进行产品的完成
- 后端人员完成的是最核心的任务，根据需求文档，完成相应的内容开发
- 测试人员对完成的内容进行相应的测试


这是一般的企业的开发流程，当然功能会不断的变更、甚至整个产品完全的改版。每一种职业的人都在开发过程中对产品负责。产品开发完成，把代码在真实的服务器上进行运行。这一步完成的是所谓的产品上线。


后端人员在这个过程中，需要接触哪些技术完成 `web` 开发？

- 熟悉一门编程语言：完成功能开发，后端人员选择对应的技术栈进行完成
- 关系型数据库：产品提供服务，内容的存储，一定是会和数据库打交道，比如：抖音，用户看到的是视频、图片、点赞、评论等，实际一定是存储在某种数据库内，这样才能完成数据持久化，当然数据库根据数据量、使用场景，选择也不一样，在某个阶段甚至需要对数据库进行再设计阶段，比如当前数据库使用场景变了、或者数据量变大了，上亿级别。这些因素都影响着数据库的选择。
- 非关系型：有些情况下关系型数据库同样可以不适合当前的使用场景，这个时候的数据存储，有可能选择非关系型，比如频繁读取的数据，需要当做缓存使用，这个时候有可能选择 `Redis`  等，比如对搜索要求很高，这个时候有可能选择 `Elasticsearch` 等
- 代码版本管理：当前企业中最流行的代码版本管理是 `Git`, 其他的代码管理技术几乎可以忽略不计，不学习即可。如果你在一家不是使用 `Git` 进行版本管理的，你可能需要考虑这家公司的技术是不是严重落后了
- 如何维护代码的质量？测试人员会对功能测试，但不可能所有的 `Bug` 都被测到，作为开发人员需要编写相应的测试对代码层面进行测试
- 如何让生产环境和本地，甚至其他开发人员的环境一致？当前企业最流行的技术是 `Docker` 容器，用于维护环境的一致性，极大的方便后续的持续部署，开发人员要拥抱容器技术
- 如何持续集成和持续部署？分支内合并新代码，自动触发相应的测试，确保测试通过才能继续代码合并动作，否则查看错误，排查问题。新代码一旦合入，触发构建新容器服务，容器内运行当前正确的代码，比如 `web` 服务，自动启动新容器，代替原容器，提供服务
- 线上真实的服务运行，使用的用户一旦增大，比如同一时刻，用户访问激增，当前服务器同一时刻可能不能应对所有的请求，这个时候就要使用 `nginx`、负载均衡、集群等服务，将网络请求分配给其他服务器，缓解访问请求


可以看到，整个流程中其实作为后端人员接触的技术非常的多，从编程语言到服务器部署、维护稳定等。

`web` 服务，提供的是用户访问资源、获取资源、更新资源和删除资源等行为。

可以看下真实的网络请求：`https://www.baidu.com`，借助浏览器的调试功能，推荐使用`Chrome` 浏览器。

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3ek5lplf4j213x0m40xw.jpg)

可以看到网络请求，交互了很多的内容，使用搜索引擎进行搜索，浏览器之所以会渲染出不同的内容，是和搜索引擎的服务器交互的结果。

这些内容其实都是`HTTP` (超文本传输协议：Hyper Text Transfer Protocol)，协议约定了网络传输过程中需要遵循的规则，是服务器传输超文本到本地浏览器的传送协议。

发起访问的一方，称之为客户端，提供服务的一方，称之为服务端。对浏览器访问的行为，浏览器可以称为客户端。

客户端发起网络请求遵循`HTTP` 协议，客户端发起一个HTTP请求到服务器的请求消息包括以下格式：

- 请求行：Request Line
- 请求头部: Header
- 空行
- 请求数据

```
> GET /persons HTTP/1.1
> Host: localhost:9999
> User-Agent: curl/7.54.0
> Accept: */*
> Content-Type: application/json
```

上文示例: 客户端的发起 `GET` 请求，访问的地址是: `localhost:9999/persons`，头部信息包含：`Host、User-Agent、Accept、Content-Type`


服务器响应信息：包含这些部分：

- 状态行: 协议版本、状态码
- 消息报文头
- 空行
- 正文

```
< HTTP/1.1 200 OK
< Content-Type: application/json
< Set-Cookie: expires=Get; Expires=Mon, 27 May 2019 04:13:43 GMT
< Date: Sun, 26 May 2019 04:13:43 GMT
< Content-Length: 390
<
[{"avatar":"http://images.org/123","origin_id":"08ff7c7ccbe23a21ff3856fe13b8cb8e","id":1,"created_at":"2019-05-26T12:13:43.134858+08:00","telephone":"1234567890","gender":1,"what_is_up":"Python"},{"avatar":"http://images.org/456","origin_id":"b9fdb24e2947f9290bd7474346e33f25","id":2,"created_at":"2019-05-26T12:13:43.13486+08:00","telephone":"987654321","gender":0,"what_is_up":"Golang"}]

```

上文示例: 服务端的响应是通过`HTTP/1.1` 版本的，状态码是`200`，响应头部信息包括：`Content-Type、Set-Cookie、Date、Content-Length` ，响应正文包括：json 字符串。


`HTTP` 协议标准提供了多种请求方法：主要包括：

- GET、POST、PUT、PATCH、DELETE、HEAD、OPTIONS、CONNECT、TRACE


各请方法对应的应用作用是：

- GET 获取指定资源
- POST 在服务器上创建新资源，需要传入请求参数
- PUT/PATCH 在服务器上更新资源，需要传入请求参数
- DELETE 在服务请上删除指定资源
- HEAD 只获取请求的头部信息
- OPTIONS 查看服务器的性能
- CONNECT 将连接改为管道方式的代理服务器
- TRACE 回显服务器收到的请求，主要用于测试或诊断


`HTTP` 协议标准提供了由三个十进制数字组成的状态码，常见的比如：

- 200 请求成功
- 307 重定向
- 404 请求资源有误
- 504 服务器连接有误

状态码也有很多，开发人员需要了解其分类：

- 1XX 接收到请求
- 2XX 请求正确
- 3XX 重定向
- 4XX 客户端错误
- 5XX 服务端错误



对开发人员，构建 `web` 服务的核心是设计：请求动作、路由、响应信息、状态码等信息。

在 `Go` 中启动一个 `web` 服务只需简单的几行：

```
package main

import (
    "net/http"
    "log"
)


func Hello(writer http.ResponseWriter, request *http.Request){
	writer.Write([]byte("Hello World"))
}

func main(){
    
    http.HandleFunc("/", Hello)
    log.Fatal(http.ListenAndServe(":9999", nil))
    
}

```
上述启动端口为`9999` 的本地服务，访问路由：`localhost:9999` 调用 `Hello` 函数，触发返回响应 `Hello World` 信息。


下面的使用内置的库，构建这一个简易的`web` 服务，提供简单的前端网页，从中学习，如何进行 `web` 开发。

主页：

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3eo0rh0sqj213x0m240l.jpg)

页面：

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3eou96sqej213v0m1wi1.jpg)

登录界面：

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3eo0rjtwlj213x0m3q4q.jpg)

页面：

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3eo0rgrztj213x0m3gns.jpg)



进度条：

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3eo0rou67j213v0m40un.jpg)




开发 `web` 服务，你要明确哪些是前端的职责，哪些是后端的职责，上面是一个非常简单的网页页面，并没有使用到很复杂的内容，页面的展示也是使用了 `Bootstrap` 前端框架，这样能够快速的帮助后端人员构建网页页面。

开发这样一个简单的`web` 服务，要遵循哪些步骤？

- 明确目的：最终目标是什么？需要开发多少页面，页面的内容包括哪些？
- 技术选型：既然是前端页面，开发人员是否熟悉前端，如何进行技术选型，原生还是框架？
- 代码组织：完成功能开发必要的项目组织



## 1.2 目标

开发一个类似博客系统的 `web` 服务，页面主要包括：

- 主页：导航栏：主页按钮、登录、登出界面、API 界面、进度条界面、正文：文章列表，可以跳转文章详情
- 文章详情：正文文章内容，侧边栏显示状态：用户数、分享数、收藏等
- 登录界面：用户名、密码输入框、提交按钮
- API 界面：表格显示一些歌曲、下文显示服务提供 API 详情
- 进度条：计算当年已过多少天，得出进度条



## 1.3 模板的使用

页面的展示很多内容其实都是前端处理，后端进行数据的处理，比如正文文章的列表、文章详情、API 内容、进度条数据信息等内容。

逐个完成页面的编写，在这之前需要了解下模板引擎：即在静态 HTML 内插入动态语言生成的数据，模板的作用可以复用很多静态代码。

```
package main

import (
    "net/http"
    "log"
    "http/template"
)


func HelloTemplate(writer http.ResponseWriter, request *http.Request) {
	t, err := template.New("hello").Parse(`
	<html>
		<head>
			<title>{{.Title}}</title>
		</head>
		<body>
			<h1>
				{{.Content}}
			</h1>
		</body>
	</html>
`)
	if err != nil {
		panic(fmt.Sprintf("template fail : %s", err.Error()))
	}
	text := struct {
		Title   string
		Content string
	}{
		Title:   "Hello Golang",
		Content: "Golang",
	}
	err = t.Execute(writer, text)
	if err != nil {
		panic(err)
	}
}


func main(){
    http.HandleFunc("/", HelloTemplate)
    log.Fatal(http.ListenAndServe(":9999", nil))
}

```

上文中的模板是：

```
<html>
	<head>
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>
			{{.Content}}
		</h1>
	</body>
</html>
```

模板内的内容语法都需要带上 `{{}}`。上文中 `{{.Title}}` 和 `{{ .Content }}` 表示模板中的动态数据，`.` 表示当前上下文的变量，比如上文的 `text`, `text` 包含属性 `Title` 和 `Content` 。模板引擎可以根据传入的变量，将动态数据填充进去。


模板引擎还支持哪些操作？

- 遍历： `{{range .}} ... {{end}}`、`{{with .}} ... {{end}}`
- 分支：`{{ if . }} ... {{ else }} ... {{end}}`
- 管道：`{{ .| function }}`, 管道后接处理函数，对当前变量进行操作
- 模板继承：`{{ define "name"}} ... {{end}}` 定义模板，`{{ template }}`
- 模版中可以直接调用当前变量的方法

```
<html>
	<head>
		<title>{{.Title}}</title>
	</head>
	<body>
		{{if .Content}}
			{{range .Content}}
			<p>{{ . | handle}}</p>
			{{end}}
		{{ else}}
		<button>No Result</button>
		{{end}}
	</body>
</html>

```

模板内判断传入的变量的 `Content` 是否有值，如果有值进行遍历，否则，输出`No Result`.

```

func textHandle(value string) string {
	return fmt.Sprintf(value + " Yes")
}

func HelloTemplate(writer http.ResponseWriter, request *http.Request) {
	t := template.New("hello")
	t = t.Funcs(template.FuncMap{"handle": textHandle})
	t, err := t.Parse(`
	<html>
		<head>
			<title>{{.Title}}</title>
		</head>
		<body>
			{{if .Content}}
				{{range .Content}}
				<p>{{ . | handle}}</p>
				{{end}}
			{{ else}}
			<button>No Result</button>
			{{end}}
		</body>
	</html>
`)
	if err != nil {
		panic(fmt.Sprintf("template fail : %s", err.Error()))
	}
	text := struct {
		Title   string
		Content []string
	}{
		Title:   "Hello Golang",
		Content: []string{"Go", "Python", "Java", "JavaScript"},
	}
	err = t.Execute(writer, text)
	if err != nil {
		panic(err)
	}
}

func main() {
    http.HandleFunc("/template", HelloTemplate)
	log.Fatal(http.ListenAndServe(":9999", nil))
}


```


管道信息的处理，需在加载模板之前载入，即：`t.Funcs(template.FuncMap{"handle":textHandle})`。`handle` 表示模板中管道内使用的值，实际的处理函数是 `textHandle`。

实际的结果：

```

<html>
        <head>
                <title>Hello Golang</title>
        </head>
        <body>


                        <p>Go Yes</p>

                        <p>Python Yes</p>

                        <p>Java Yes</p>

                        <p>JavaScript Yes</p>


        </body>
</html>

```




模板继承：

```
func HelloTemplate(writer http.ResponseWriter, request *http.Request) {

	content := `
{{define "content"}}
			{{if .Content}}
				{{range .Content}}
				<p>{{ . | handle}}</p>
				{{end}}
			{{ else}}
			<button>No Result</button>
			{{end}}
{{end}}
`
	html := `
	<html>
		<head>
			<title>{{.Title}}</title>
		</head>
		<body>
			{{template "content" .}}
		</body>
	</html>
`
	temp := template.New("Hello")
	t2 := template.Must(template.Must(temp.Funcs(template.FuncMap{"handle": textHandle}).Parse(content)).Clone())
	t3, err := t2.Parse(html)
	if err != nil {
		panic(fmt.Sprintf("template fail : %s", err.Error()))
	}
	text := struct {
		Title   string
		Content []string
	}{
		Title:   "Hello Golang",
		Content: []string{"Go", "Python", "Java", "JavaScript"},
	}
	err = t3.Execute(writer, text)
	if err != nil {
		panic(err)
	}
}


```

模板的定义使用`{{define "name"}} ... {{ end}}` 语法，模板的继承使用 `{{template "name" .}}` `.` 表示上下文变量。可以对模板进行复用。


上文模板的使用需要注意加载模板的顺序，就上文这个问题，应该先加载`content` 模板，在加载被继承过的模板。不然会导致信息报错。


甚至可以直接调用当前对象的方法：

```
package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

func main() {

	text := `
<html>
	{{ if not .A.IsZero }}
	<h1>{{.B}}</h1>
	{{else}}
	<h3>{{.CreatedAt.Format "2006-01-02 15:04:05"}}</h3>
	<h3>{{.Hello "Golang"}}</h3>
	{{end}}
</html>
`
	content := Name{
		A :A{},
		B: "s",
		CreatedAt: time.Now(),
	}

	t , _ := template.New("s").Parse(text)
	t.Execute(os.Stdout, content)

}

type Name struct {
	A A
	B string
	CreatedAt time.Time
}

func (n Name) Hello(value string) string {
	return value + " oops!"
}

type A struct {
	Value string
}

func (a A) IsZero() bool{
	if a.Value == "" {
		return true
	}
	return false
}


```

上文中调用了结构体对象 `A` 的 `IsZero` 方法 和 结构体对象 `Name` 的 `Hello` 方法，间接调用了 `time.Time` 的 `Format` 方法，甚至还可以接收传入的参数，比如`Hello` 方法传入参数 `Golang`，比如 `Format` 方法传入参数 `2006-01-02 15:04:05`。



结果输出为：

```
<html>
	
	<h3>2019-05-27 15:03:14</h3>
	<h3>Golang oops!</h3>
	
</html>

```


根据上文的梳理，总结下模版有哪些用法：


**作用：**

静态页面中用一些带 `{{}}` 符号的样式表示占位符，通过传入动态数据，完成数据的填充。

**语法：**

模版中支持一些特定的语法：

- 变量：`{{.}}` 表示当前变量，具体含有根据上下文语境
- 判断和分支：`{{if .}} ... {{else}} ... {{end}}`
- 循环： `{{range .}} ... {{end}}`, `{{with .}}...{{end}}`
- 模版定义与继承：`{{define "子模版名称" }} ... {{end}}`, `{{template "子模版名称" .}} ... {{end}}`
- 管道操作：将当前变量作为参数传入一个处理函数中
- 可以直接调用变量对象的方法

**使用：**

- 根据需要的动态数据定义相应的结构体和方法
- 在模版中使用占位符
- 使用 `template` 的`Parse、ParseFile、Execute` 等方法加载模版




## 1.4 内容开发


开发 `web` 服务的核心是设计：请求方法、路由、响应信息、状态码等信息。根据这些内容，针对性的设计简易的 `web` 服务。


### 1.4.1 首页

所有的页面中导航栏和页脚是固定的，常见的网页的导航栏和页脚也都是固定的，为此需要构建模板继承，将中间页面的展示内容抽离出模板。

为更快的搭建页面内容，使用`Bootstrap 4.3.1`前端框架。


整体的页面模板如下：`index.html`

```
<html lang="en">
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta charset="UTF-8">
    <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet" href="https://cdn.bootcss.com/twitter-bootstrap/4.3.1/css/bootstrap.css" >
    <script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.slim.min.js" ></script>
    <script src="https://cdn.bootcss.com/popper.js/1.14.7/esm/popper.min.js" ></script>
    <script src="https://cdn.bootcss.com/twitter-bootstrap/4.3.1/js/bootstrap.min.js"></script>
    <style>
        .btn.btn-primary{
            background-color: #00bb00;
            border-color: #00bb00;
        }
        .container-small{
            max-width: 500px;
        }
        body{
            max-width: 1200px;
            margin: 25px auto;
            background: #f5f5f5;
        }
    </style>
    <title>BiuBiuBiu</title>
</head>
<body>
<nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <a class="navbar-brand" href="/">
        BiuBiuBiu
    </a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
    </button>

    <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav mr-auto">
            <li class="nav-item active">
                <a class="nav-link" href="/">Home <span class="sr-only"></span></a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/login">Login</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/logout">Logout</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/apis">API</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/progress">Progress</a>
            </li>
        </ul>
        <form class="form-inline my-2 my-lg-0">
            <input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search">
            <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
        </form>
    </div>
</nav>
<br>
<br>
<div class="starter-template">
    {{ template "content" . }}
</div>
<br>
<br>
<footer class="text-muted text-center text-small">
    <p class="text-muted">Copyright©️ 2018 ~ 2019 上海情非得已有限公司. All rights reserved.</p>
    <p class="text-muted">Build By: <a href="https://www.zhihu.com/people/wu-xiao-shen-16/activities" target="_blank">@谢伟</a></p>
</footer>
</body>
</html>

```

导航栏在标签 `nav` 内，页脚在标签 `footer` 内，当模板内容没有加载进内容，静态的网页如下所示：

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3esqpqr66j213v0atdgq.jpg)

导航栏的几个标签，对应：主页、登录、登出、API、进度页面链接。


设计目标主页正文显示文章信息：标签、标题、创建时间、内容概括、评论数、收藏数、点击数。


![](http://ww1.sinaimg.cn/large/741fdb86gy1g3eszo4veyj213v0m5q5r.jpg)


根据对模板的学习，主页正文的设计，只需要加载继承一个模板，模板内填充的数据进行遍历即可：


主页正文模板：`home.html`

```
{{define "content"}}
    <div class="jumbotron mt-3" style="background: #FFFFFF  ">
        <h2 class="text-left"> 欢迎，简易的 Web 教程示例 <span class="badge badge-secondary">New</span></h2>
        <p>
            <a class="btn btn-primary btn-lg" href="https://www.zhihu.com/people/wu-xiao-shen-16/activities">访问主页</a>
        </p>
    </div>
    <br>
    <br>
    <div class="row mb-1">
        {{range .}}
        <div class="col-md-12">
            <div class="card flex-md-row mb-4 shadow-sm h-md-250">
                <div class="card-body d-flex flex-column align-items-start">
                    <strong class="d-inline-block mb-2 text-primary">{{.Tag}}</strong>
                    <h3 class="mb-0">
                        <a class="text-dark" href="#">{{.Title}}</a>
                    </h3>
                    <div class="mb-1 text-muted">{{.Time | timeHandle}}</div>
                    <p class="card-text mb-0 text-muted" >{{ .Content }}</p>
                    <span class="d-inline-block" tabindex="0">
                        <button class="btn btn-primary btn-sm " type="button" style="color:#6a757e;background-color: #e8ecef;border-color: #e8ecef">评论: {{.CommentInt}}</button>
                        <button class="btn btn-primary btn-sm" type="button" style="color:#6a757e;background-color: #e8ecef;border-color: #e8ecef">收藏: {{.CollectionInt}}</button>
                        <button class="btn btn-primary btn-sm" type="button" style="color:#6a757e;background-color: #e8ecef;border-color: #e8ecef">点击: {{.ClickInt}}</button>
                    </span>
                    <br>
                    <a type="button" href="/passage" class="btn btn-primary btn-sm" style="color:#6a757e;background-color: #e8ecef;border-color: #e8ecef" data-container="body" data-toggle="popover" title="Reading More" data-content="It is just Mock">Reading More</a>
                </div>
            </div>
        </div>
        {{end}}
    </div>
    <nav aria-label="Page navigation example">
        <ul class="pagination justify-content-center pagination-sm">
            <li class="page-item disabled">
                <a class="page-link" href="#" tabindex="-1"><<</a>
            </li>
            <li class="page-item">
                <a class="page-link" href="#" tabindex="-1">1</a>
            </li>
            <li class="page-item">
                <a class="page-link" href="#" tabindex="-1"> >> </a>
            </li>

        </ul>
    </nav>
{{end}}


```

将需要动态加载的数据，用结构体来进行表示，模板引擎内的变量需和定义的结构图的字段的名称一致，否则模板引擎会报错。

```

type content struct {
	Tag           string    `json:"tag"`
	Title         string    `json:"title"`
	Time          time.Time `json:"time"`
	Content       string    `json:"content"`
	CommentInt    int       `json:"comment_int"`
	CollectionInt int       `json:"collection_int"`
	ClickInt      int       `json:"click_int"`
}

type contents []content


```


对应的逻辑处理函数，只需要加载模板引擎，在加载动态数据即可：

```
func home(writer http.ResponseWriter, req *http.Request) {
	var c contents
	c = []content{
		{
			Tag:           "Go",
			Title:         "How to learn Golang",
			Time:          time.Now(),
			Content:       "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
			CommentInt:    1,
			CollectionInt: 12,
			ClickInt:      100,
		},
		{
			Tag:           "Python",
			Title:         "How to learn Python",
			Time:          time.Now().Add(-24 * time.Hour),
			Content:       "Python is a programming language that lets you work quickly and integrate systems more effectively.",
			CommentInt:    2,
			CollectionInt: 34,
			ClickInt:      1000,
		},
		{
			Tag:           "Java",
			Title:         "How to learn Java",
			Time:          time.Now().Add(-24 * 2 * time.Hour),
			Content:       "Java is a general-purpose programming language that is class-based, object-oriented, and designed to have as few implementation dependencies as possible.",
			CommentInt:    3,
			CollectionInt: 124,
			ClickInt:      900,
		},
		{
			Tag:           "JavaScript",
			Title:         "How to learn JavaScript",
			Time:          time.Now().Add(-24 * 2 * 2 * time.Hour),
			Content:       "JavaScript often abbreviated as JS, is a high-level, interpreted programming language that conforms to the ECMAScript specification. JavaScript has curly-bracket syntax, dynamic typing, prototype-based object-orientation, and first-class functions.",
			CommentInt:    212,
			CollectionInt: 1224,
			ClickInt:      9030,
		},
	}
	currentPath, _ := os.Getwd()
	temp := template.New("index.html")
	t := temp.Funcs(template.FuncMap{"timeHandle": timeHandle})
	t, err := t.ParseFiles(path.Join(currentPath, "chapter10/simple/template/index.html"), 
	path.Join(currentPath, "chapter10/simple/template/home.html"))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(writer, c)
	if err != nil {
		panic(err)
	}
}

func timeHandle(date time.Time) string {
	return date.Format("2006/01/02")
}

func main(){
    
    	http.HandleFunc("/", home)
    	log.Fatal(http.ListenAndServe(":9999", nil))
}

```

上文设计的主页的访问路由是：`localhost:9999`,  调用的逻辑处理函数是：`home`, 启动的服务的端口是: `9999`。

需要注意的点是：

- 对错误的处理信息，本地调试时，最好使用 `panic` 进行错误的捕获，或者使用日志处理，这样方便查看具体的报错信息，初学者没有报错信息，难及时发现问题所在
- 模板的加载注意顺序，先加载哪个，后加载哪个


软件领域中间件是一种独立的系统软件或者服务程序，在 `web` 服务中也常使用中间件，比如日志的处理、认证信息等，这些也称之为中间件。中间件的使用一般是在网络请求逻辑处理之前或者之后，进行相应的操作。

为更好的对 `web` 服务进行处理，这边在所有的网络请求上使用日志处理中间件，程序运行的时候能够看到一些日志信息，方便及时的排除。

中间件服务如何编写？

中间件是发生在真实的网络请求逻辑处理之前或者之后。通常的写法是这样：

```
func middlewareHandler(next http.Handler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        // 执行handler之前的逻辑
        next.ServeHTTP(w, r)
        // 执行完毕handler后的逻辑
    })
}

```

也可以这样编写：

```
func middlewareHandler(next http.HandlerFunc) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request){
        // 执行handler之前的逻辑
        next.ServeHTTP(w, r)
        // 执行完毕handler后的逻辑
    }
}

```
两者的差别在于，接收的参数和返回值不同，一个是 `http.Handler` ，一个是 `http.HandlerFunc`。两者的区别是，一个是构造`http.Handler`, 一个是对函数的处理。两者的调用稍有差别。

源码中 `http.Handler` 和 `http.HandlerFunc` 的定义如下：

```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}


```

可以看出 `http.HanlerFunc` 实现了 `Handler` 接口。

对处理函数添加上日志处理中间件：

```
func logger(next http.HandlerFunc) http.HandlerFunc {
	now := time.Now()
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("[Web-Server]: %s | %s", request.RequestURI, now.Format("2006/01/02 -15:04:05"))
		next.ServeHTTP(writer, request)
	}

}

func middlewareLogger(next http.Handler) http.Handler {
	now := time.Now()
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("[Web-Server]: %s | %s", request.RequestURI, now.Format("2006/01/02 -15:04:05"))
		next.ServeHTTP(writer, request)
	})
}

```

如果需要处理函数进行中间件处理，可以调用上面两者中间件函数：

```
func main(){
    http.HandleFunc("/example", logger(home))
	http.Handle("/exampleTwo", midderlewareLogger(http.HandlerFunc(home)))
}


```


再梳理下首页内容是如何完成的：

- 导航栏和页脚内容固定，固定模板，其他内容继承其他模板
- 正文内容是文章列表，动态加载数据，构建需要填充数据的结构体列表，再遍历即可
- 为方便查看日志信息，使用日志中间件，方便查看访问的路由和时间




### 1.4.2 文章详情


文章详情，包含两大块：文章内容和侧边栏的状态，其中文章内容包含：文章标题、创建时间、作者、文章正文；侧边栏状态包括：用户数、分享数、评论数、收藏数。

基于此定义两部分结构体：文章部分 和 状态部分

```
type passageContent struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Author    string    `json:"author"`
	Detail    string    `json:"detail"`
}

type side struct {
	Tag   string   `json:"tag"`
	Items []string `json:"items"`
}

```

文章详情模板为：`passage.html`

```
{{define "content"}}
 <div class="container">
<div class="row">
        <div class="col-md-9">
            <h1 class="pb-3 mb-4 font-italic border-bottom">{{ .Title }}</h1>
            <p>{{.CreatedAt | time }}<a href="#">{{.Author}}</a></p>
            <p style="line-height: 2">{{.Detail}}</p>
        </div>
     <aside class="col-md-3">
         <div class="p-3">
             <h4 class="font-italic">{{.Tag}}</h4>
             <ol class="list-unstyled mb-0">
                 {{range .Items }}
                     <li >
                         <p class="text-muted">{{ . }}</p>
                     </li>
                 {{end}}
             </ol>
         </div>
     </aside>

</div>
</div>
{{end}}
```

将整个页面内容分隔为`9:3` 其中文章内容占 9 份，侧边栏状态占 3 份，分别对应 `html` 标签 `<div class="col-md-9"></div>` 和 `<div class="col-md-3"></div>`

文章标题的字体、正文的行高等，由前端 `CSS` 样式自定义，后端赋值数据交互。

为此设计文章详情的逻辑处理函数为：`func passage(writer http.ResponseWriter, request *http.Request)`

```

func passage(writer http.ResponseWriter, request *http.Request) {

	var content = struct {
		passageContent
		side
	}{
		passageContent: passageContent{
			Title:     "How to learn golang",
			CreatedAt: time.Now(),
			Author:    "Go Team",
			Detail: `The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.`,
		},
		side: side{
			Tag: "状态",
			Items: []string{
				"用户数: 62",
				"分享数: 27",
				"评论数: 19",
				"收藏数: 12",
			},
		},
	}
	currentPath, _ := os.Getwd()
	temp := template.New("index.html")
	t := temp.Funcs(template.FuncMap{"time": timeFormat})
	t, err := t.ParseFiles(
	path.Join(currentPath, "chapter10/simple/template/index.html"),
	path.Join(currentPath, "chapter10/simple/template/passage.html"))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(writer, content)
	if err != nil {
		panic(err)
	}
}

// 模板中时间的处理函数
func timeFormat(date time.Time) string {
	return fmt.Sprintf(date.Format(time.Stamp) + " By ")
}

```

使用到了模板继承，即将导航栏和页脚栏的内容和文章详情的内容结合在一起。模板加载 `html` 文件的顺序要注意，先加载 `index.html` 再加载 `passage.html`。


结合主页，启动 `web` 服务，则如下的样式：

```
func main() {
	http.HandleFunc("/", logger(home))
	http.Handle("/2", middlewareLogger(http.HandlerFunc(home)))
	http.HandleFunc("/passage", logger(passage))
	log.Fatal(http.ListenAndServe(":9999", nil))

}
```

首页中文章 `Reading More` 按钮即可实现跳转到详情页面。 


再梳理下，文章详情是如何完成的？

- 模板继承，页面中导航栏和页脚内容自动加载，只考虑正文部分
- 正文划分两大块：文章详情和侧边栏的状态
- 定义模板，定义模板中加载数据的结构体
- 编写具体的逻辑处理函数


### 1.4.3 登录/登出界面

登录界面，在前端的展示中其实很简单，1. 输入用户名称或者邮箱 2. 输入用户密码 3. 点击登录按钮。触发动作，真实的场景在点击登录按钮时，向后端发起一个网络请求，后端对用户输入的用户名和密码进行校验，如果该用户存在且密码正确，则跳转网页至相应页面访问，如果不对，提示报错信息。

这边为简单处理，只处理如何读取到用户名和密码，进行相应的参数进行校验。

参照之前的输入，模板内容中动态数据有：错误提示信息。

定义相应的结构体：

```
type loginInfo struct {
	UserName string  `json:"user_name"`
	Password string  `json:"password"`
	Error    []error `json:"error"`
}
```

用户名和密码从表单发起的网络请求中进行读取。后端进行校验，如果出错，错误信息提示出来，在页面中展示。相应的模板如下：`login.html`

```
{{define "content"}}
<div class="container container-small">
<h1>登录</h1>
<form action="/login" method="post" name="login" class="form-signin">
    <div class="alert alert-success">登录界面访问成功～</div>
    <div class="form-group">
            <label>用户名</label>
            <input class="form-control" type="text" placeholder="Email Or Username" name="username">
    </div>
    <div class="form-group">
            <label>密码</label>
            <input class="form-control" type="password" placeholder="Password" name="password">
    </div>
    <div class="form-group">
        <div class="checkbox mb-3">
            <label>
                <input type="checkbox" value="remember-me">
                Remember me
            </label>
        </div>
    </div>
    <div class="form-group">
        <button type="submit" class="btn btn-primary btn-block">Sign in</button>
    </div>
</form>
    {{if .Error}}
    <div class="card">

        <div class="alert-danger">
            <h4 class="alert-heading">Warning</h4>

            {{range .Error}}
                    {{.}}
            {{end}}

        </div>

    </div>
    {{end}}
</div>

{{end}}
```

需要对错误信息进行判断，使用模板的判断语句，如果有值，加载出标签`<div class="card"></div>` 内容。

相应的逻辑处理函数如下：

```
func login(writer http.ResponseWriter, req *http.Request) {

	currentPath, _ := os.Getwd()
	
	temp, _ := template.ParseFiles(
	path.Join(currentPath,"chapter10/simple/template/index.html"),
	path.Join(currentPath, "chapter10/simple/template/login.html"))
	
	var lgInfo loginInfo
	
	if req.Method == http.MethodGet {
		temp.Execute(writer, lgInfo)
		return
	}

	if err := req.ParseForm(); err != nil {
		return
	}
	UserName := req.PostFormValue("username")
	Password := req.PostFormValue("password")

	lgInfo.UserName = UserName
	lgInfo.Password = Password

	errFunc := func(values string) error {
		v := values
		// 校验密码长度
		if len(v) < 8 || len(v) == 0 {
			return fmt.Errorf("the length should be larger 8")
		}
		// 校验密码包含的字符
		if unicode.IsNumber(rune(v[0])) {
			return fmt.Errorf("should not start number")
		}
		return nil
	}

	if errFunc(lgInfo.UserName) == nil && errFunc(lgInfo.Password) == nil {
		http.Redirect(writer, req, "/", http.StatusSeeOther)
		return
		
	} else {
		if errFunc(lgInfo.UserName) != nil {
			lgInfo.Error = append(lgInfo.Error, fmt.Errorf("username is not suitable"))
		}
		if errFunc(lgInfo.Password) != nil {
			lgInfo.Error = append(lgInfo.Error, fmt.Errorf("password is not suitable"))
		}
		log.Println(lgInfo)
		temp.Execute(writer, lgInfo)
	}
}

```

具体的登录动作是 `POST` 请求，第一步需要根据请求 `request` 判断请求的方法是 `GET` 还是 `POST`, 如果是 `GET` 动作，直接加载登录界面，此时，用户名和密码判断得出的错误信息为空，不加载标签 `<div class="card"></div>` 内的内容，如果是 `POST` 请求，则根据网络请求，获取得到用户名和密码的值，对用户名和密码的格式进行判断，比如长度小于 8 或者内容为空，则产生错误信息，否则校验通过，重定向至主页。


这里有几个新知识点：

- 如何获取到网络请求的方法和参数
- 如何重定向至其他逻辑处理函数

网络请的具体逻辑处理函数的格式如下： `func (writer http.ResponseWriter, request *http.Request)`

- writer 负责写入响应信息：提供的方法有：

```

type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}    
```

常用的写入响应信息的方法有：

```
- writer.Writer([]byte(string(content))
- fmt.Fprintf(writer, content)
- io.WriterString(writer, content)
- template.Execute(writer, content)
- json.NewEncoder(writer).Encode(content)


```

主要使用到了 `Go` 的 `Interface` 接口特性，ResponseWriter interface 实现 Write([]byte) (int, error) 的方法，所以也就实现了 io.Writer 方法，所以可以作为 io.Writer 的类型。

```
// io.go , package io
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

- request 负责处理网络请求，网络请求遵循 `HTTP` 协议，有请求方法、请求路径、请求头部信息、请求参数等内容

```

type Request struct {

	Method string

	URL *url.URL

	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0


	Header Header


	Body io.ReadCloser


	GetBody func() (io.ReadCloser, error)


	ContentLength int64


	TransferEncoding []string

	Close bool


	Host string

	Form url.Values


	PostForm url.Values


	MultipartForm *multipart.Form

	Trailer Header


	RemoteAddr string


	RequestURI string


	TLS *tls.ConnectionState

	Cancel <-chan struct{}

	Response *Response


	ctx context.Context
}
```


结构体 `Request` 包含诸多和网络请求相关的属性和方法。

根据 `HTTP` 协议的规则，网络请求的重点在于：路由、请求参数、头部信息、请求方法。

结构体 `Request` 的重点也是在处理这些内容，那么如何操作这些对象？

头部信息：

```
type Request struct {

    Header Header
    // 其他
    
}

type Header map[string][]string

```

请求头部信息是一个 `hash` 字典，提供`Add、Del、Get、Set` 等方法，在构建 `web` 服务中，一般用来设置头部信息，即使用 `request.Header.Add(key string, value string)` 方法。

请求方法：

```

type Request struct{
    Method string
    // 其他
}

```

原生的内置库，设计 `web` 服务，不能够一眼看出路由对应的请求方法是：`POST`、`GET` 或者其他。可以通过 `request.Method` 取到值，再判断请求方法。内置的方法以常量的形式定义好了：

```
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

```

请求参数：

对于网络请求的参数，有这么几种情况，如果请求是`GET`那么获取路径中请求参数，`GET` 请求无需传入消息体，比如请求是：`localhost:9999/person?query=age&name=x`, `query=age&name=x` 就是路径中的请求参数；如果请求是`POST、PUT、PATCH` 等请求，那么如何获取请求中的消息体的参数？

路径中的请求参数使用：`FormValue` 进行单个字段的解析，也可以一次性获取所以路径中请求参数：`Form`。

```
func HelloWorld(writer http.ResponseWriter, request *http.Request){
	fmt.Println(request.FormValue("query"))
	fmt.Println(request.FormValue("name"))
	fmt.Println(request.Form)
	writer.Write([]byte("Hello"))
}
func main() {
	http.HandleFunc("/hello", HelloWorld)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
```

访问：`localhost:9999/hello?query=age?name=x` 打印出：`age`、`x`、`map[query:[age] name:[x]]`

为什么可以这么操作？

```
type Request struct {

	// Form contains the parsed form data, including both the URL
	// field's query parameters and the POST or PUT form data.
	// This field is only available after ParseForm is called.
	// The HTTP client ignores Form and uses Body instead.
	Form url.Values
	
	// 其他
    
}

type Values map[string][]string


```

源码显示，在`POST`或者`PUT`请求方法中，调用了 `ParseForm` 方法才能获取消息体中的参数，否则获取路径中的请求参数。`FormValue` 是解析`Form` 字典中的某个 `key` 的 `value` 值。

消息体中参数：

```
func HelloWorldPost(writer http.ResponseWriter, request *http.Request){
	if err := request.ParseForm();err!=nil{
		panic("request parse form fail")
	}
	param1 := request.PostForm["username"][0] // request.PostFormValue("username")
	param2 := request.PostForm["password"][0]
	params := request.PostForm
	paramsQueryField := request.Form
	fmt.Println(param1)
	fmt.Println(param2)
	fmt.Println(params)
	fmt.Println(paramsQueryField)
	writer.Write([]byte("POST"))
}

func main() {
    
    http.HandleFunc("/hello/post", HelloWorldPost)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
```

访问：`localhost:9999/hello/post?query=age&name=x`, 传入消息体参数：`username:Go, password:Hello Golang`。

输出：

```
Go
Hello Golang
map[username:[Go] password:[Hello Golang] undefined:[]]
map[username:[Go] password:[Hello Golang] undefined:[] query:[age] name:[x]]

```

结论是：`POST`或者`PUT` 操作，获取消息体操作需要先调用 `ParseForm` 方法，获取单个消息体参数： `PostForm[key]`，返回是个列表，也可以直接调用 `PostFormValue` 获取当个值。`PostForm` 返回所有消息体参数，`Form` 既包含消息体参数，也包含路径中请求参数。


结论：

对网络请求的处理包含两个层面：`http.ResponseWriter`、`http.Request`, 一个负责响应信息的处理，比如写入头部信息，写入响应信息；一个负责网络请求的方法、路径、请求参数、消息体参数等的解析。

再回到登录页面的逻辑处理函数中来：

```
func login(writer http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {
		temp.Execute(writer, lgInfo)
		return
	}

	if err := req.ParseForm(); err != nil {
		return
	}
	UserName := req.PostFormValue("username")
	Password := req.PostFormValue("password")
    //	其余内容省略
}	

```

调用请求的方法，判断是否`POST`请求，如果不是，显示登录界面，否则解析用户名和密码，获取到两者的值，再进行校验，比如长度，以及是否包含非法字符等。

校验通过，直接调用 `http.Redirect` 方法，重定向至符合路由的逻辑处理函数中去，比如这边是重定向至主页。


再总结下：登录页面的设计

- 抽象出一个模版，模版内的 `html` 值是表单标签
- 通过 `http.Request` 的方法，和对请求参数、消息体参数的处理，完成参数的校验
- 不成功，输出报错信息。
- 否则，调用 `http.Redirect` 重定向至主页


### 1.4.4 API 文档界面



主页、文章详情页、登录页的一般流程是：

- 明确设计的界面是怎样的，由前端设计
- 抽象出动态加载的内容，定义相应的结构体
- 加载模版，将抽象出的动态数据加载进模版内
- 启动服务，访问后，浏览器渲染出内容


![](http://ww1.sinaimg.cn/large/741fdb86gy1g3eo0rgrztj213x0m3gns.jpg)

API 文档界面包含两部分内容：1. 表格内容 2. 接口详情

对两部分内容再进行分解，表格内容，抽象出动态数据：`ID(序号)、Name(名称)、Author(作者)、Time(时长)、Album(专辑) `, 定义相应的结构体列表，再循环遍历即可。

接口详情，抽象出动态数据：`Title(接口标题)、Content(接口内容)、Method(接口请求方法)、Path(接口路径)、Comment(接口备注)`。

定义动态数据的结构体：

```

// 表格内容
type singleSong struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Time   string `json:"time"`
	Album  string `json:"album"`
}

type songs []singleSong

// api 内容
type api struct {
	Title   string `json`
	Content string `json:"content"`
	Method  string `json:"method"`
	Path    string `json:"path"`
	Comment string `json:"comment"`
}

type apis []api

```

定义模版内容：`song.html`

```
{{define "content"}}
    <div class="container">
        <div class="container">
            <h4 class="text-muted">歌曲排行</h4>
            <div class="row">
                <div class="container">
                    <table class="table table-striped">
                        <tbody >
                        {{range .Songs}}
                            <tr>
                                <th >{{.ID}}</th>
                                <td class="text-muted">{{.Name}}</td>
                                <td class="text-muted">{{.Author}}</td>
                                <td class="text-muted">{{.Time}}</td>
                                <td class="text-muted">{{.Album}}</td>
                            </tr>
                        {{end}}

                        </tbody>
                    </table>

                </div>
            </div>
        </div>
    </div>
    <br>
    <div class="container">

        <div class="container">
            <h4 class="text-muted">接口详情</h4>
            {{range .APis}}
            <div class="card">
                <div class="card-header" style="background: #6c757d; color: #f5f5f5">
                    {{.Title}}
                </div>
                <div class="card-body">
                    <h6 class="card-title text-muted">
                        {{.Content}}
                    </h6>
                    <button type="button" class="btn btn-default card-text text-muted" style="border-color: #e1e1e1">
                        {{.Method}}
                    </button>
                    <button type="button" class="btn btn-default card-text text-muted" style="border-color: #e1e1e1">
                        {{.Path}}
                    </button>
                </div>
                <div class="card-footer text-muted">
                    {{.Comment}}
                </div>
            </div>
            <br>
            {{end}}
            <div class="card">
                <div class="card-header"  style="background: #6c757d; color: #f5f5f5">
                    意见反馈
                </div>
                <div class="card-body">
                    邮箱: xie_wei_shu@shu.edu.cn
                </div>

            </div>
        </div>
    </div>

    <nav aria-label="Page navigation example">
        <ul class="pagination justify-content-center">
            <li class="page-item disabled">
                <a class="page-link" href="#" tabindex="-1">Previous</a>
            </li>
            <li class="page-item"><a class="page-link" href="#">1</a></li>
            <li class="page-item"><a class="page-link" href="#">2</a></li>
            <li class="page-item"><a class="page-link" href="#">3</a></li>
            <li class="page-item">
                <a class="page-link" href="#">Next</a>
            </li>
        </ul>
    </nav>
{{end}}

```

编写逻辑处理函数：

```

func song(writer http.ResponseWriter, req *http.Request) {
	var ss songs
	ss = []singleSong{
		{
			ID:     1,
			Name:   "全部都是你",
			Author: "DP龙猪",
			Time:   "03:23",
			Album:  "全部都是你",
		},
		{
			ID:     2,
			Name:   "对你的感觉",
			Author: "DP龙猪",
			Time:   "04:23",
			Album:  "对你的感觉",
		},
		{
			ID:     3,
			Name:   "我可不可以",
			Author: "DP龙猪",
			Time:   "05:23",
			Album:  "我可不可以",
		},
		{
			ID:     4,
			Name:   "围绕",
			Author: "DP龙猪",
			Time:   "03:23",
			Album:  "围绕",
		},
	}
	var aps apis
	aps = []api{
		{
			Title:   "获取人员信息",
			Content: "通过ID获取人员信息",
			Method:  http.MethodGet,
			Path:    fmt.Sprintf("/person/get?id=ID"),
			Comment: fmt.Sprintf("ID 选择 1 或者 2"),
		},
		{
			Title:   "获取所有人员信息",
			Content: "获取内置所有人员的信息",
			Method:  http.MethodGet,
			Path:    fmt.Sprintf("/persons"),
			Comment: fmt.Sprintf("无须传入请求参数"),
		},
		{
			Title:   "创建人员信息",
			Content: "传入参数 id 和 telephone 创建新人",
			Method:  http.MethodPost,
			Path:    fmt.Sprintf("/person/post"),
			Comment: fmt.Sprintf("传入参数 id 或者 telephone"),
		},
		{
			Title:   "更新人员信息",
			Content: "传入参数 id 更新人员 telephone 信息",
			Method:  http.MethodPatch,
			Path:    fmt.Sprintf("/person/patch?id=ID"),
			Comment: fmt.Sprintf("传入路径参数 id 和请求参数 telephone"),
		},
	}
	var all = struct {
		Songs songs
		APis  apis
	}{
		Songs: ss,
		APis:  aps,
	}
	currentPath, _ := os.Getwd()
	temp, err := template.ParseFiles(
	path.Join(currentPath, "chapter10/simple/template/index.html"),
	path.Join(currentPath, "chapter10/simple/template/song.html"))
	if err != nil {
		log.Println(err)
		return
	}
	err2 := temp.Execute(writer, all)
	if err2 != nil {
		log.Println(err2)
		return
	}
}

```

访问：`localhost:999/apis` 浏览器渲染接口详情页面。




### 1.4.5 进度条页面

进度条的内容比较简单，包含标题栏：提示当前年份已过的百分比、进度条：根据百分比，自动加载占比。


进度条的内容主要使用了 `BootStrap4.3.1` 的进度条 `<div class="progress"></div>` 标签。


按照之前的处理步骤：1. 抽象出动态数据：`Year(当前年份)、Now(当前占比)` 2. 编写前端模版。

动态数据结构体：

```
type progressStatus struct {
	Now  float64 `json:"now"`
	Year int     `json:"year"`
}
```

前端模版：

```
{{define "content"}}
<div class="row">
    <div class="container">
        <h1 class="text-warning" >{{.Year}} 年已经过去 {{ .Now }}%! <small class="text-muted">计划都实现了吗？</small> </h1>
        <br>
        <div class="progress">
            <div class="progress-bar progress-bar-striped" role="progressbar" style="width: {{.Now}}%;" aria-valuenow={{.Now}}% aria-valuemin="0" aria-valuemax="100">{{.Now}}%</div>
        </div>
    </div>

</div>

{{end}}
```

这里的重点是：如何通过当前的时间计算出已经过去的天数，再根据天数除以当年的总天数，得出占比。

逻辑处理函数：

```
func progress(writer http.ResponseWriter, req *http.Request) {

	var proStatus progressStatus

	monthDays := []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	
	now := time.Now()
	
	y, m, d := now.Date()
	ok := 0
	sum := monthDays[time.Month(m)-1] + d
	
	if (y%400 == 0) || ((y%4 == 0) && (y%100 != 0)) {
		ok = 1
	}
	if (ok == 1) && (time.Month(m) > 2) {
		sum += 1
	}
	
	proStatus.Now, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(sum)/float64(365)*100), 64)
	
	proStatus.Year = y
	currentPath, _ := os.Getwd()
	temp, _ := template.ParseFiles(
	path.Join(currentPath, "chapter10/simple/template/index.html"),
	path.Join(currentPath, "chapter10/simple/template/progress.html"))
	temp.Execute(writer, proStatus)
}

```




### 1.4.6 接口测试

上文较为详细的描述了整个 `web` 服务页面的编写，步骤可以抽象为这么几个步骤：

- 抽象动态数据，定义结构体
- 定义前端模版
- 编写具体的逻辑处理函数，动态加载数据



至此：所有的路由和相应的逻辑处理函数编写完成。

```
func main() {
	http.HandleFunc("/", logger(home))
	http.Handle("/2", middlewareLogger(http.HandlerFunc(home)))
	http.HandleFunc("/3", logger(Hello))
	http.HandleFunc("/persons", logger(getHandler))
	http.HandleFunc("/person/post", logger(postHandler))
	http.HandleFunc("/person/patch", logger(patchHandler))
	http.HandleFunc("/person/get", logger(getProfile))
	http.HandleFunc("/login", logger(login))
	http.HandleFunc("/logout", logger(logout))
	http.HandleFunc("/apis", logger(song))
	http.HandleFunc("/progress", logger(progress))
	http.HandleFunc("/passage", logger(passage))
	http.HandleFunc("/template", logger(HelloTemplate))
	log.Fatal(http.ListenAndServe(":9999", nil))
}
```

那么问题是，如何本地测试接口？尤其是需要传入的消息体参数的请求？

下面介绍几个常用的方法：

- curl : 一个利用URL语法在命令行下工作的文件传输工具
- HTTPie : 是一个 HTTP 的命令行客户端,目标是让 CLI 和 web 服务之间的交互尽可能的人性化
- postman : 是一款功能强大的网页调试与发送网页 HTTP 请求工具
- 其他：根据自己的开发环境，比如 集成开发环境，可以安装相应的接口测试插件

**curl**

安装: https://curl.haxx.se/download.html 根据自己的操作系统选择对应的安装包，即可。


为演示方便，选择一个专门用来测试网络请求的网站：httpbin： http://www.httpbin.org/   ，网站提供的请求可以在线访问，也可以在本地启动容器的形式访问本地的请求。

命令行的使用方式是在终端中输入：`curl + 路由 + 参数`

```
// 命令
curl http://httpbin.org/ip

// 响应
{
  "origin": "116.227.77.197, 116.227.77.197"
}

// 命令
curl http://httpbin.org/get?query=name

// 响应
{
  "args": {
      qy
  },
  "headers": {
    "Accept": "*/*",
    "Host": "httpbin.org",
    "User-Agent": "curl/7.54.0"
  },
  "origin": "116.227.77.197, 116.227.77.197",
  "url": "https://httpbin.org/get"
}

```

默认是使用 `GET` 的请求方法。

```
// 命令
curl -o get.json http://httpbin.org/get?query=name

// 响应

同目录下，将请求的响应内容保存在 get.json 文件内

```

- `-o` 参数，将请求的响应保存在命名的文件内

```
// 命令
curl -H "content-type: application/json" http://httpbin.org/get

// 响应

{
  "args": {},
  "headers": {
    "Accept": "*/*",
    "Content-Type": "application/json",
    "Host": "httpbin.org",
    "User-Agent": "curl/7.54.0"
  },
  "origin": "116.227.77.197, 116.227.77.197",
  "url": "https://httpbin.org/get"
}
```

- `-H` 指定请求的头部信息参数

```
// 命令
curl -X POST -d "data=name&language=go" http://httpbin.org/post

// 响应

{
  "args": {},
  "data": "",
  "files": {},
  "form": {
    "data": "name",
    "language": "go"
  },
  "headers": {
    "Accept": "*/*",
    "Content-Length": "21",
    "Content-Type": "application/x-www-form-urlencoded",
    "Host": "httpbin.org",
    "User-Agent": "curl/7.54.0"
  },
  "json": null,
  "origin": "116.227.77.197, 116.227.77.197",
  "url": "https://httpbin.org/post"
}

// 命令

curl -X POST -H "content-type: application/json" -d "data=name&language=go" http://httpbin.org/post

// 响应

{
  "args": {},
  "data": "data=name&language=go",
  "files": {},
  "form": {},
  "headers": {
    "Accept": "*/*",
    "Content-Length": "21",
    "Content-Type": "application/json",
    "Host": "httpbin.org",
    "User-Agent": "curl/7.54.0"
  },
  "json": null,
  "origin": "116.227.77.197, 116.227.77.197",
  "url": "https://httpbin.org/post"
}

```

- `-X` 显式的指定请求方法，`-H` 指定请求的头部参数，默认以表单方式传给服务器：`"Content-Type": "application/x-www-form-urlencoded"` , 指定以 `json` 字符串的形式传给服务器：`"Content-Type": "application/json"`



其他短参数：

- `-v` 显示具体的请求信息，包括请求参数、响应参数，可以对照着学习 `HTTP` 协议
- `-I` 显示头部信息
- `-F` 表单参数


更多示例，查看官方网站内容：https://curl.haxx.se/docs/manual.html


**HTTPie**

curl 命令行工具面对复杂的网络请求，编写请求参数比较繁琐，有其他替代方案，更高效。HTTPie 就是一款可以替代 curl 命令行工具的 HTTP 命令行客户端，让命令行和服务端交互更为便捷，比如高亮、格式化、支持 HTTPS, 代理和授权验证等。

文档地址：https://httpie.org/

`HTTPie` 是 Python 编写的，各平台都可以使用。读者选择适合自己操作系统的版本，下载安装即可。

下面和 curl 命令行一致，请求相同的网络请求：

```
// 命令
http http://httpbin.org/ip

// 响应
HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Encoding: gzip
Content-Length: 57
Content-Type: application/json
Date: Tue, 28 May 2019 16:07:59 GMT
Referrer-Policy: no-referrer-when-downgrade
Server: nginx
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
X-XSS-Protection: 1; mode=block

{
    "origin": "116.227.77.197, 116.227.77.197"
}

```

```
// 命令
http http://httpbin.org/get?query=name

// 响应
HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Encoding: gzip
Content-Length: 194
Content-Type: application/json
Date: Tue, 28 May 2019 16:08:34 GMT
Referrer-Policy: no-referrer-when-downgrade
Server: nginx
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
X-XSS-Protection: 1; mode=block

{
    "args": {
        "query": "name"
    },
    "headers": {
        "Accept": "*/*",
        "Accept-Encoding": "gzip, deflate",
        "Host": "httpbin.org",
        "User-Agent": "HTTPie/0.9.9"
    },
    "origin": "116.227.77.197, 116.227.77.197",
    "url": "https://httpbin.org/get?query=name"
}
```

```
// 命令: 以 json 字符串的形成传给服务器
http POST http://httpbin.org/post data=name language=go

// 响应
HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Encoding: gzip
Content-Length: 276
Content-Type: application/json
Date: Tue, 28 May 2019 16:10:39 GMT
Referrer-Policy: no-referrer-when-downgrade
Server: nginx
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
X-XSS-Protection: 1; mode=block

{
    "args": {},
    "data": "{\"data\": \"name\", \"language\": \"go\"}",
    "files": {},
    "form": {},
    "headers": {
        "Accept": "application/json, */*",
        "Accept-Encoding": "gzip, deflate",
        "Content-Length": "34",
        "Content-Type": "application/json",
        "Host": "httpbin.org",
        "User-Agent": "HTTPie/0.9.9"
    },
    "json": {
        "data": "name",
        "language": "go"
    },
    "origin": "116.227.77.197, 116.227.77.197",
    "url": "https://httpbin.org/post"
}
```

```
// 命令: 以表单的形式传递给服务器
http -f  POST http://httpbin.org/post data=name language=go

// 响应
HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Encoding: gzip
Content-Length: 279
Content-Type: application/json
Date: Tue, 28 May 2019 16:11:34 GMT
Referrer-Policy: no-referrer-when-downgrade
Server: nginx
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
X-XSS-Protection: 1; mode=block

{
    "args": {},
    "data": "",
    "files": {},
    "form": {
        "data": "name",
        "language": "go"
    },
    "headers": {
        "Accept": "*/*",
        "Accept-Encoding": "gzip, deflate",
        "Content-Length": "21",
        "Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
        "Host": "httpbin.org",
        "User-Agent": "HTTPie/0.9.9"
    },
    "json": null,
    "origin": "116.227.77.197, 116.227.77.197",
    "url": "https://httpbin.org/post"
}
```

更多用法：https://httpie.org/

本质上两款命令行工具没有本质的差别，都是用来进行网络请求，`HTTPie` 相比 `curl` 使用起来更为高效，且高亮的语法等，对用户更为友好。读者根据自己实际情况选择。



**postman**

`HTTPie` 和 `curl` 都是命令行工具，需要在终端中输入命令，那有没有图形化界面的工具呢？有的，`Google` 出品的专门用来 `API` 测试的图形化工具。深受前后端开发人员的喜爱。如果一定要推荐的测试 `API` 工具的话，`Postman` 必居其一。

官网：https://www.getpostman.com/


读者选择适合自己操作系统的版本，下载安装，为方便管理 `API`, 下载安装之后，需要注册账户再登录。

下面示例和之前的示例一致：

选择：`GET` 方法，填入：`http://httpbin.org/get` 即可
![](http://ww1.sinaimg.cn/large/741fdb86gy1g3hhvdvdhuj213u0msgow.jpg)

选择：`GET` 方法，填入：`http://httpbin.org/get?query=name` 即可
![](http://ww1.sinaimg.cn/large/741fdb86gy1g3hhvdx3sdj213u0mon0o.jpg)

选择：`POST` 方法，填入：`http://httpbin.org/post` 传入 `json` 字符串即可
![](http://ww1.sinaimg.cn/large/741fdb86gy1g3hhvdx38cj213t0mr78b.jpg)


选择：`POST` 方法，填入：`http://httpbin.org/post` 表单内填入参数即可
![](http://ww1.sinaimg.cn/large/741fdb86gy1g3hhvdzu54j213t0msq74.jpg)



**其他插件**


`vscode` (https://code.visualstudio.com/) 是微软出品的一款非常好用的免费开源的现代化轻量级代码编辑器，越来越多的程序员选择这款工具编写代码，其生态非常丰富，插件非常繁多，极大的便利了开发者。

`vscode` 中有一款插件，让使用者以文本的形式只输入网络请求和参数，就能完成 `API` 的测试。

下载 vscode 安装插件：REST Client(https://github.com/Huachao/vscode-restclient/)。


![](http://ww1.sinaimg.cn/large/741fdb86gy1g3hidadm2lj213v0ku0z2.jpg)

同样，以上文的示例，说明如何使用：

- 任意目录创建任意名称的文件，且文件后缀名为 `.http`，比如：`api.http`
- 每个测试的 API, 以 `###` 开头，换行后，自动显示：`Send Request` 且可点击
- 以 `GET`、`POST` 请求方法开头，后接请求路由
- 如果需要指定头部信息，在请求路由的下一行编写头部信息
- 如果需要指定请求消息体，头部信息后，换行，再编写请求消息体

```
// api.http
###
GET http://httpbin.org/get HTTP/1.1

###
POST http://httpbin.org/post HTTP/1.1
Content-Type: application/json

data=name
&language=Go

###
POST http://httpbin.org/post HTTP/1.1
Content-Type: application/x-www-form-urlencoded

data=name
&language=Go

```

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3hidacqhoj213v0lsjvs.jpg)

点击`Send Request`, 右边显示响应信息，左边显示请求路由和参数等。



总结：

接口测试的目的，是为了检验，编写的服务端的网络请求是否正确，合理的使用这些接口测试工具，能够一定程度上保证服务端编写的网络请求的正确性，是非常好的调试工具。

本环节分别从：命令行和图像化界面两个角度，分别介绍了：`curl`、`HTTPie` 和 `Postman` 、`REST Client`。

在使用的过程中，进一步明白 `HTTP` 协议的标准。

- 请求方法
- 头部信息
- 状态码
- 响应信息
- 协议版本
- ...


参考代码：https://github.com/wuxiaoxiaoshen/GopherBook/tree/master/chapter10/simple


## 2. `net/http` 构建爱鲜蜂 `web` 服务


第一节使用内置的模版引擎，开发了一些简单的页面，上文的情况存在什么问题呢？

- 前端和后端耦合：即前端页面中混合了后端业务代码
- 代码未按照某种方式组织，不易于拓展
- 页面中设计的数据都是 `Mock` 数据，未使用到持久化存储：数据库



真实的互联网流行的开发方式是：前后端分离，前端由前端人员负责，后端又后端人员负责，前端开发选择合适的开发框架，比如`React` `Vue` 等，后端选择合适的编程语言，比如`Go`、`Java`等，两者之间的交互通过应用接口。这种前后端开发的方式：职责分明，耦合度低，容易及时发现潜在的问题等。


作为后端开发，之后我们的重点在设计具有 `Restful` 风格的应用接口(`API`)。

为尽量贴近企业开发，整体流程划分为四个部分：

- 需求流程梳理
- 模型设计：数据库结构设计
- 代码开发
- 持续集成(`CI`)和持续部署(`CD`)



### 2.1 需求流程梳理

产品经理在前期调研、规划、竞品分析、和用户调研等之后，得出产品的定位，比如是生鲜产品、还是社交产品、或是新闻客户端等。

产品定位之后，产品经理做出产品需求文档（PRD），包括产品的定位、产品需求的描述、目标市场、目标用户等。对各阶段的功能都有清晰的描述，包括产品原型，在和设计人员沟通之后，得出产品设计稿。产品设计稿通常划分为：web 端和客户端（Android 和 IOS），加上需求澄清，结合设计稿，相应的开发人员明确自己的目标是什么，能针对性的开发功能。

为讲述方便，这边我们列举市面上已经存在的产品：爱鲜蜂。


 爱鲜蜂是一款定位为生鲜配送的产品，采用众包模式，致力于解决最后一公里的生鲜配送。产品的属性决定了产品的内容是一些生鲜：蔬菜、水果、饮料酒水等。

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3tka8mesij20u01ffgpk.jpg)

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3tka8lym8j20u01f8413.jpg)

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3tka8l0imj20u01f8tao.jpg)

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3tka8o24mj20u01f3abn.jpg)

### 2.2  模型设计

模型设计是后端开发人员的一个重要的能力，同时好的模型的设计能够给开发带来很多的便利，需求是会不断变更，模型设计的好，对需求的不断变更也能灵活的适应。模型设计对开发来说，最重要的是关系型数据库表的设计能力，市场上包含诸多类型的数据库，但关系型数据库依然占据绝大多数份额。

简单的看过爱鲜蜂产品的界面。下面尝试仅从产品界面的角度来进行模型设计的讲解。模型设计主要分为三步：
1. 列出产品设计的实体 
2. 根据设计稿（界面）划分功能 
3. 数据库表设计：字段设计、类型设计、1对1、1对多、多对多实体设计

另外需要注意的是，需求会变动，模型的设计也是会变动的，比如增加字段，使其符合产品功能。


**爱鲜蜂模型设计:**

列出产品实体：首页包含一些运营活动，将一些生鲜产品组合起来进行推广，活动是针对地区的，首页也存在位置管理功能。

分类页包含各种分类的产品：产品实体是：分类的标签、和具体的产品型号（名称、规则、保质期等）

购物车页：产品实体是: 配送地址管理；送货时间管理；产品列表、价格、数量、总价等

个人页：账户信息、积分、劵信息、订单信息、等级说明等。

根据产品设计界面，和产品实体，总结的思维导图如下：

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3tlmpftx6j211b0mi782.jpg)

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3tlo5agw6j20lb0me0ul.jpg)

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3tlok73qrj213q0mcgq1.jpg)

根据实体:设计的模型如下：（将所有涉及的实体，抽象出数据库表结构）

```
// activities_model.go
package model

import (
	"time"

	"github.com/jinzhu/gorm"
)
// 活动结构体，定义字段
type Activity struct {
	gorm.Model
	Title    string    `gorm:"type:varchar" json:"title"`
	FromDate time.Time `gorm:"type:timestamp with time zone" json:"from_date"`
	ToDate   time.Time `gorm:"type:timestamp with time zone" json:"to_date"`
	Products []Product `gorm:"type:many2many: activity2products" json:"products"`
}

// product_model.go
package model

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)
 
// 商品结构体，定义字段
type Product struct {
	gorm.Model
	Name          string          `gorm:"type:varchar" json:"name"`
	Avatar        string          `gorm:"type:varchar" json:"avatar"`
	Price         sql.NullFloat64 `json:"price"`
	Amount        int             `gorm:"type:integer" json:"amount"`
	Specification string          `gorm:"type:varchar" json:"specification"`
	Period        int             `gorm:"type:integer" json:"period"`
	BrandID       uint
	UintID        uint
	TagID         uint
}

// brand_model.go
package model

import "github.com/jinzhu/gorm"

type Brand struct {
	gorm.Model
	EnName string `gorm:"type:varchar" json:"en_name"`
	ChName string `gorm:"type:varchar" json:"ch_name"`
}

// unit_model.go
package model

type Uint struct {
	Name string `gorm:"type:varchar" json:"name"`
}

// tag_model.go
package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar" json:"name"`
}

// shopping_cart_model.go
package model

import "github.com/jinzhu/gorm"

type ShoppingCart struct {
	gorm.Model
	AccountID     uint
	ReceiptDateID uint
	OrderID       uint
	Order         Order
}

// order_model.go
package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	OrderStatus    []OrderStatus
	Status         string `gorm:"type:varchar" json:"status"`
	ShoppingCartID uint
}

type OrderStatus struct {
	gorm.Model
	Product Product
	Amount  int `gorm:"type:integer" json:"amount"`
}

// account_model.go
package model

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	LevelID  uint
	Phone    string    `gorm:"type:varchar" json:"phone"`
	Avatar   string    `gorm:"type:varchar" json:"avatar"`
	Name     string    `gorm:"type:varchar" json:"name"`
	Gender   int       `gorm:"type:integer" json:"gender"` // 0 男 1 女
	Birthday time.Time `gorm:"type:timestamp with time zone" json:"birthday"`
	Points   sql.NullFloat64
}

// receipt_date_model.go
package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ReceiptDate struct {
	gorm.Model
	ReceiveDateID uint
	ReceiveDate   ReceiveDate
	FormTime      time.Time `gorm:"type:timestamp with time zone" json:"form_time"`
	ToTime        time.Time `gorm:"type:timestamp with time zone" json:"to_time"`
}

type ReceiveDate struct {
	gorm.Model
	Item string `gorm:"type:varchar" json:"item"`
}
```

​    

```
// admin_model.go
package model

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	AccountID       uint
	AccountBalance  sql.NullFloat64
	ExchangesNumber int `gorm:"type:integer" json:"exchanges_number"`
	CouponsNumber   int `gorm:"type:integer" json:"coupons_number"`
	Exchanges       []Exchange
	Coupons         []Coupon
}

//exchange_model.go
package model

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type Exchange struct {
	gorm.Model
	Name     string    `gorm:"type:varchar" json:"name"`
	ZeroTime time.Time `gorm:"type:timestamp with time zone" json:"zero_time"`
	EndTime  time.Time `gorm:"type:timestamp with time zone" json:"end_time"`
	Price    sql.NullFloat64
}

// coupons_model.go
package model

import (
	"github.com/jinzhu/gorm"
)

type Coupon struct {
	gorm.Model
	Exchange
	Token string `gorm:"type:varchar" json:"token"`
}

// level_model.go
package model

import "github.com/jinzhu/gorm"

type Level struct {
	gorm.Model
	Name      string `gorm:"type:varchar" json:"name"`
	ZeroValue int    `gorm:"type:integer" json:"zero_value"`
	EndValue  int    `gorm:"type:integer" json:"end_value"`
	Privilege string `gorm:"type:varchar" json:"privilege"`
	Validity  string `gorm:"type:varchar" json:"validity"`
}
```

可以看到，越是功能复杂的项目，涉及的实体越多，涉及的表结构也就越复杂、字段越多。


总结下，上文我们是如何抽象出的数据库模型？

1. 根据设计稿（产品界面）抽象出各种实体，比如活动，产品
2. 根据设计稿（产品界面）抽象出实体的各种字段，比如活动的时间，产品的价格等，这些字段就构成了实体的属性
3. 根据上两步的处理，定义模型



### 2.3 代码开发

上文的模型定义使用到了 ORM 技术，即将结构体映射成数据库表。使用到的库是 GORM。

模型设计的数据库的表设计完备之后，后续的工具就是进行相应的操作：查询记录、删除记录、更改记录、新增记录等，即常说的 CURD。

**原生数据库操作：**

内置库database/sql  提供了标准的接口，开发者可以根据定义的接口来开发相应的数据库驱动。用法大同小异，主要包括这些动作：1. 创建连接 （连接数据库）2. 进行操作：增删改查，事务等

原生的内置库，开发者需要在代码中嵌入 SQL 语句。

具体示例：

```
// 表结构：

CREATE TABLE `wechat_persons` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `nick_name` varchar(10) DEFAULT NULL,
  `account_string` varchar(15) DEFAULT NULL,
  `account_qr` varchar(255) DEFAULT NULL,
  `gender` int(11) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  `signal_person` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_wechat_persons_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8mb4_0900_ai_ci;
```

根据表结构，定义相应的结构体，原生的不支持转换关系，即创建表、表的定义、数据库的定义都需要开发者编写相应的语句来完成。后续的 ORM 只需要定义相应的结构体即可，大大精简了开发的流程，提升了速率。

```
package model

import "time"

type Base struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
// 定义人的属性
type Person struct {
	Base
	Avatar        string
	NickName      string
	AccountString string
	AccountQR     string
	Gender        int
	Location      string
	Signal        string
	Addresses     []Address
	Receipts      []Receipt
}

// 表的名称
func (Person) TableName() string {
	return "wechat_persons"
}

// 序列化结构体
type PersonSerializer struct {
	ID            uint      `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Avatar        string    `json:"avatar"`
	NickName      string    `json:"nick_name"`
	AccountString string    `json:"account_string"`
	Gender        string    `json:"gender"`
	Location      string    `json:"location"`
	Signal        string    `json:"signal"`
}

// 序列化函数
func (p Person) JSONSerializer() PersonSerializer {
	genderString := func(gender int) string {
		if gender == 0 {
			return "男"
		}
		return "女"
	}
	return PersonSerializer{
		ID:            p.ID,
		CreatedAt:     p.CreatedAt.Truncate(time.Hour),
		UpdatedAt:     p.UpdatedAt.Truncate(time.Second),
		Avatar:        p.Avatar,
		NickName:      p.NickName,
		AccountString: p.AccountString,
		Gender:        genderString(p.Gender),
		Location:      p.Location,
		Signal:        p.Signal,
	}
}
```

创建连接对象，只有连接了数据库之后，才可以进行数据库的相应操作：

注：相应的第三库不存在本地时，需要先使用 go get 先获取到本地。

```
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/Orm/origin_example/model"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	var config mysql.Config
	config = mysql.Config{
		User:      "root",
		Passwd:    "admin123",
		DBName:    "person",
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		ParseTime: true,
	}
	dsn := config.FormatDSN()
	originURL := "root:admin123@/person?charset=utf8&parseTime=true&loc=Local"
	fmt.Println(dsn, originURL)

	db, err = sql.Open("mysql", originURL)
	if err != nil {
		fmt.Println(err)
		return
	}
}
```

本地启动数据库的方式，可以从官方上下载 MySQL 及其相应的驱动，也可以使用容器的形式，推荐使用 Docker 启动相应的数据库服务。可以直接使用 Docker 命令直接启动，也可以使用 docker-compose 编排启动。下文的意思是：启动一个 MySQL 容器，创建用户名为 root, 密码为 admin123 的账户，数据库的名称为：person，可以看出上文的连接配置中的信息和下文设置的一致。

```
// 在 docker-compose.yml 文件目录下执行命令，容器在后台启动

docker-compose up -d 

// docker-compose
version: "3"

services:
  mysql:
    image: mysql:latest
    container_name: localMySQL
    volumes:
      - $PWD/mysql-data:/var/lib/mysql
    expose:
      - 3306
    ports:
      - 3306:3306
    environment:
      - MYSQL_ROOT_PASSWORD=admin123
      - MYSQL_DATABASE=person
      - MYSQL_USER=root
```

创建连接对象之后，可以直接调数据库对象相应的方法：

![](http://ww1.sinaimg.cn/large/741fdb86gy1g3tnbzqdrtj20o50azjs5.jpg)

原生的 database/sql  在代码中需要嵌入 SQL 语句，比如删除表命令。

```
// SQL 语句

DELETE FROM tableName;

func deleteTable() {
	stmt, _ := db.Prepare("DELETE  from  wechat_persons")
	stmt.Exec()

}
```

随机准备 10 条记录，插入数据库中：

```
// SQL 语句

INSERT INTO tableName (field1, field2...) VALUES ( value1, value2...);
```

​    

```
func records() []model.Person {
	rand.Seed(time.Now().UnixNano())
	var result []model.Person
	for i := 0; i < 10; i++ {
		one := model.Person{
			Base: model.Base{
				ID:        uint(rand.Int31()),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Avatar:        "https://images.pexels.com/photos/2326961/pexels-photo-2326961.jpeg?auto=format%2Ccompress&cs=tinysrgb&dpr=1&w=500",
			NickName:      fmt.Sprintf("%d", uint(rand.Int31())),
			AccountString: strconv.Itoa(int(rand.Int31n(100000))),
			Gender:        1,
			Location:      fmt.Sprintf("北京: %d", uint(rand.Int31n(100000))),
			Signal:        fmt.Sprintf("走自己的路，让别人去说吧: %d", uint(rand.Int31n(100))),
		}
		result = append(result, one)
	}
	return result
}

func insert() {
	for _, i := range records() {
		_, err := db.Exec("INSERT INTO wechat_persons (id, created_at, updated_at, avatar, nick_name, account_string, gender,location,signal_person ) values ( ?,?,?,?,?,?,?,?,? )",
			i.ID, i.CreatedAt, i.UpdatedAt, i.Avatar, i.NickName, i.AccountString, i.Gender, i.Location, i.Signal)
		if err != nil {
			log.Println(err)
			return
		}

	}
}
```

查询数据库记录，再将数据载入对应的结构体内，调用结构体的序列化方法 JSONSerializer ：

```
// SQL 语句

SELECT field1, field2 ... FROM tableName;
```

​    

```
func query() []model.PersonSerializer {
	var (
		result []model.PersonSerializer
	)
	rows, err := db.Query("SELECT id, created_at, updated_at, avatar, nick_name, account_string,gender,location,signal_person FROM  wechat_persons")
	if err != nil {
		log.Println(err)
		return result
	}
	for rows.Next() {
		var one model.Person
		err := rows.Scan(&one.ID, &one.CreatedAt, &one.UpdatedAt, &one.Avatar, &one.NickName, &one.AccountString, &one.Gender, &one.Location, &one.Signal)
		if err != nil {
			break
		}
		result = append(result, one.JSONSerializer())

	}
	return result
}
```

将数据的交互以接口的形式，供前端或者客户端相应的开发人员调用：

```
func apiGet(writer http.ResponseWriter, request *http.Request) {
	var result []model.PersonSerializer
	result = query()
	var values = make(map[string]interface{})
	values["data"] = result
	//writer.Write(Json(values))
	json.NewEncoder(writer).Encode(values)
}

func apiPatch(writer http.ResponseWriter, request *http.Request) {
	var result = make(map[string]interface{})
	request.Header.Add("content-type", "application/json")
	if request.Method != http.MethodPatch {
		result["code"] = http.StatusBadRequest
		result["detail"] = fmt.Sprintf("method is not allow")
		json.NewEncoder(writer).Encode(Json(result))
		return
	}
	request.ParseForm()

	id := request.FormValue("id")
	nickName := request.PostFormValue("nick_name")
	location := request.PostFormValue("location")
	log.Println(id, nickName, location)

	row := db.QueryRow("SELECT id, nick_name, location FROM wechat_persons  where id = ?", id)
	var person model.Person
	err := row.Scan(&person.ID, &person.NickName, &person.Location)
	if err != nil {
		fmt.Println(err)
		return
	}

	if id == strconv.Itoa(int(person.ID)) && nickName != "" && location != "" {

		stmt, err := db.Prepare("UPDATE wechat_persons set nick_name = ? , location = ?, updated_at = ? where id = ?")
		if err != nil {
			log.Println(err)
			json.NewEncoder(writer).Encode(Json(result))
			return
		}
		_, err = stmt.Exec(nickName, location, time.Now(), id)
		row := db.QueryRow("SELECT id, created_at, updated_at, avatar, nick_name, account_string,gender, location, signal_person FROM wechat_persons  where id = ?", id)
		var one model.Person
		err = row.Scan(&one.ID, &one.CreatedAt, &one.UpdatedAt, &one.Avatar, &one.NickName, &one.AccountString, &one.Gender, &one.Location, &one.Signal)
		if err != nil {
			log.Println(err)
			json.NewEncoder(writer).Encode(Json(result))
			return
		}

		result["data"] = one.JSONSerializer()
		result["code"] = http.StatusOK
		json.NewEncoder(writer).Encode(result)
		return
	}
}

func Json(value map[string]interface{}) []byte {
	content, _ := json.Marshal(value)
	return content
}
```

启动 web 服务：

```
func main() {
	http.HandleFunc("/", apiGet)
	http.HandleFunc("/person", apiPatch)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
```

调用接口形式：使用上文介绍的接口调用工具： curl 、httpie、Postman 中的一种

```
http http://localhost:8080

HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Sat, 08 Jun 2019 04:54:12 GMT
Transfer-Encoding: chunked

{
    "data": [
        {
            "account_string": "16348",
            "avatar": "https://images.pexels.com/photos/2326961/pexels-photo-2326961.jpeg?auto=format%2Ccompress&cs=tinysrgb&dpr=1&w=500",
            "created_at": "2019-06-06T18:00:00+08:00",
            "gender": "女",
            "id": 250770615,
            "location": "上海上海",
            "nick_name": "xieweiwei",
            "signal": "走自己的路，让别人去说吧: 48",
            "updated_at": "2019-06-06T19:21:16+08:00"
        },
				// 省略 8 条记录
        {
            "account_string": "81028",
            "avatar": "https://images.pexels.com/photos/2326961/pexels-photo-2326961.jpeg?auto=format%2Ccompress&cs=tinysrgb&dpr=1&w=500",
            "created_at": "2019-06-06T18:00:00+08:00",
            "gender": "女",
            "id": 1823960809,
            "location": "北京: 88257",
            "nick_name": "1612472921",
            "signal": "走自己的路，让别人去说吧: 79",
            "updated_at": "2019-06-06T18:47:03+08:00"
        }
    ]
}
```

更新 ID  为 250770615 的记录的 location 和 nick_name 字段。

```
http --form patch http://localhost:8080/person\?id\=250770615 nick_name=佩奇 location=北京

HTTP/1.1 200 OK
Content-Length: 392
Content-Type: text/plain; charset=utf-8
Date: Sat, 08 Jun 2019 05:15:50 GMT

{
    "code": 200,
    "data": {
        "account_string": "16348",
        "avatar": "https://images.pexels.com/photos/2326961/pexels-photo-2326961.jpeg?auto=format%2Ccompress&cs=tinysrgb&dpr=1&w=500",
        "created_at": "2019-06-06T18:00:00+08:00",
        "gender": "女",
        "id": 250770615,
        "location": "北京",
        "nick_name": "佩奇",
        "signal": "走自己的路，让别人去说吧: 48",
        "updated_at": "2019-06-08T13:15:51+08:00"
    }
}
```

注：记录的 ID 为随机生成，读者使用时，看到的记录略有不同。

**总结**：上文使用原生 database/sql 和 原生的 net/http 构建了两个接口，一个获取数据库内记录，一个更新数据库指定的记录，响应的结果以  json 的形式展现出来。

构建 Restful API 风格的接口，主要有下面几个要点：

- 路由设计：即访问什么路径获取什么样的资源，比如上文只操作了 person 表，不能操作其他的表
- 状态码、错误码：访问接口会出现错误，这个时候需要显示状态码和错误码，方便及时定位问题
- 响应信息：json 是一种非常流行的数据交换格式，以 json 展现形式，不管是对 web 端还是客户端，都比较友好，企业级的应用接口开发也多以 json 的形式。

使用原生的内容，可以完成任务，但并没有让开发者聚焦在迭代开发任务上，比如充斥着 SQL 语句，比如解析请求参数，需要写很多的代码，进行多次的逻辑判断。

基于此，企业内多使用 web 框架 和 ORM 技术解决这两个痛点，开发者聚焦在业务层面，能够即正确又快速的完成任务。

**ORM 技术：**

ORM 将结构体对象和数据库表之间进行相应的映射，开发者只需要操作结构体对象即可，完成和数据库内记录之间的映射，精简了开发代码。

ORM 技术只要分为这几种操作：
1. 连接数据库对象 
2. 结构体定义，完成对数据库表的映射 
3. 表的操作：创建、删除、是否存在等 
4. 记录的操作：增删改查，索引等操作

开源社区比较流行的两个第三方库是： GORM 和 XORM 。开发者需要使用需提前使用 go get 下载相应的库至本地。

使用同一个示例来比对下两个库的使用异同：

XORM（[https://github.com/go-xorm/xorm](https://github.com/go-xorm/xorm)）:  模型定义

```
package model

import "time"

type Base struct {
	ID        uint       `xorm:"pk 'id'"`
	CreatedAt time.Time  `xorm:"created"`
	UpdatedAt time.Time  `xorm:"updated"`
	DeletedAt *time.Time `xorm:"deleted index"`
}
type Person struct {
	Base          `xorm:"extends"`
	Avatar        string `xorm:"varchar(255)" json:"avatar"`
	NickName      string `xorm:"varchar(10)" json:"nick_name"`
	AccountString string `xorm:" varchar(15)" json:"account_string"`
	AccountQR     string `xorm:" varchar(255)" json:"account_qr"`
	Gender        int    `xorm:" integer" json:"gender"`
	Location      string `xorm:" varchar(255)" json:"location"`
	Signal        string `xorm:" varchar(64)" json:"signal"`
	Addresses     []Address
	Receipts      []Receipt
}

func (Person) TableName() string {
	return "wechat_persons"
}
```

创建连接，同步数据库表：

```
package main

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/Orm/xorm_example/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)


var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:admin123@/xorm_example?charset=utf8")
	if err != nil {
		log.Println(err)
		panic("mysql connect fail")
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "wechat_")
	engine.SetTableMapper(tbMapper)
	engine.Logger().SetLevel(core.LOG_WARNING)
	dropTable()
	syncTable()

}
func dropTable() {
	engine.DropTables(&model.Receipt{})
}
func syncTable() {
	engine.Sync2(new(model.Person))
}

func main() {
	tables, err := engine.DBMetas()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range tables {
		fmt.Println(i)
	}
}
```

查询数据库：只操作结构体，数据库内就存在相应的表。

```
>> mysql -u root -p

>> show database;
+--------------------+
| Database           |
+--------------------+
| gorm_example       |
| information_schema |
| mysql              |
| performance_schema |
| person             |
| person2            |
| sys                |
| xorm_example       |
+--------------------+
8 rows in set (0.00 sec)

>> use xorm_example;
>> show tables;
+------------------------+
| Tables_in_xorm_example |
+------------------------+
| wechat_persons         |
+------------------------+
1 row in set (0.00 sec)
>> describe wechat_persons;
+----------------+--------------+------+-----+---------+-------+
| Field          | Type         | Null | Key | Default | Extra |
+----------------+--------------+------+-----+---------+-------+
| id             | int(11)      | NO   | PRI | NULL    |       |
| created_at     | datetime     | YES  |     | NULL    |       |
| updated_at     | datetime     | YES  |     | NULL    |       |
| deleted_at     | datetime     | YES  | MUL | NULL    |       |
| avatar         | varchar(255) | YES  |     | NULL    |       |
| nick_name      | varchar(10)  | YES  |     | NULL    |       |
| account_string | varchar(15)  | YES  |     | NULL    |       |
| account_q_r    | varchar(255) | YES  |     | NULL    |       |
| gender         | int(11)      | YES  |     | NULL    |       |
| location       | varchar(255) | YES  |     | NULL    |       |
| signal         | varchar(64)  | YES  |     | NULL    |       |
| addresses      | text         | YES  |     | NULL    |       |
| receipts       | text         | YES  |     | NULL    |       |
+----------------+--------------+------+-----+---------+-------+
13 rows in set (0.01 sec)
```

GORM （[https://github.com/jinzhu/gorm](https://github.com/jinzhu/gorm)）: 模型定义

```
package model

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Avatar        string `gorm:"type:varchar(255)" json:"avatar"`
	NickName      string `gorm:"type:varchar(10)" json:"nick_name"`
	AccountString string `gorm:"type:varchar(15)" json:"account_string"`
	AccountQR     string `gorm:"type:varchar(255)" json:"account_qr"`
	Gender        int    `gorm:"type:integer" json:"gender"`
	Location      string `gorm:"type:varchar(255)" json:"location"`
	Signal        string `gorm:"type:varchar(64)" json:"signal"`
	Addresses     []Address
	Receipts      []Receipt
}

func (Person) TableName() string {
	return "wechat_persons"
}
```

创建数据库连接，同步数据库表：

```
package main

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/Orm/gorm_example/model"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var engine *gorm.DB

func init() {
	var err error
	engine, err = gorm.Open("mysql", "root:admin123@/gorm_example?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Print(err)
		panic("gorm connect to mysql failed")
	}
	engine.LogMode(true)
	syncTables()
	defer engine.Close()

}

func syncTables() {
	engine.DropTableIfExists(
		&model.Person{},
	)
	engine.AutoMigrate(
		&model.Person{},
	)
}

func main() {

}
```

查询数据库：

```
>>mysql -u root -p

>>show databases;
+--------------------+
| Database           |
+--------------------+
| gorm_example       |
| information_schema |
| mysql              |
| performance_schema |
| person             |
| person2            |
| sys                |
| xorm_example       |
+--------------------+
8 rows in set (0.00 sec)

>> use gorm_example;
>> show tables;

+------------------------+
| Tables_in_gorm_example |
+------------------------+
| wechat_persons         |
+------------------------+
1 row in set (0.01 sec)

>> show create table wechat_persons\G
*************************** 1. row ***************************
       Table: wechat_persons
Create Table: CREATE TABLE `wechat_persons` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `nick_name` varchar(10) DEFAULT NULL,
  `account_string` varchar(15) DEFAULT NULL,
  `account_qr` varchar(255) DEFAULT NULL,
  `gender` int(11) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  `signal` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_wechat_persons_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
1 row in set (0.04 sec)

>> show columns from wechat_persons;
+----------------+------------------+------+-----+---------+----------------+
| Field          | Type             | Null | Key | Default | Extra          |
+----------------+------------------+------+-----+---------+----------------+
| id             | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| created_at     | timestamp        | YES  |     | NULL    |                |
| updated_at     | timestamp        | YES  |     | NULL    |                |
| deleted_at     | timestamp        | YES  | MUL | NULL    |                |
| avatar         | varchar(255)     | YES  |     | NULL    |                |
| nick_name      | varchar(10)      | YES  |     | NULL    |                |
| account_string | varchar(15)      | YES  |     | NULL    |                |
| account_qr     | varchar(255)     | YES  |     | NULL    |                |
| gender         | int(11)          | YES  |     | NULL    |                |
| location       | varchar(255)     | YES  |     | NULL    |                |
| signal         | varchar(64)      | YES  |     | NULL    |                |
+----------------+------------------+------+-----+---------+----------------+
11 rows in set (0.01 sec)
```

可以看出：ORM 聚焦的是让开发者专注在业务开发上，而不是编写繁琐的 SQL 语句，当然这两个第三方库都支持原生 SQL 语句的命令，这也不是让开发者不用学习 SQL 语句，至少在业务代码内，不应该包含过多的 SQL 语句。

两者将结构体映射成数据库表都从结构体的 Tag 入手：Tag 内可以定义字段的类型、列的名称、索引、多对多关系等。两者的 Tag 语法稍有不同。从语义角度分析，GORM 更为清晰。

那如何选择呢? 1. 是否支持用户使用的数据库类型 2. 从社区活跃度角度考虑，避免开发过程中遇到过多的问题。

**Web 框架：**

内置 net/http 可以构建 web 服务，但也存在这么几个问题：
1. 对参数的处理比较复杂 
2. 对路由的设计不够简便。比如上文的接口，需要写诸多的代码来出来请求参数

企业级项目，如果现有的框架不能满足需求，可以自己开发 web 框架，主流的第三方 web 框架部分在处理性能问题，部分框架在处理路由层面的设计。

常见的 web 框架有：

- gin ([https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)) : 性能较好，处理请求参数也比较便捷
- echo ([https://github.com/labstack/echo](https://github.com/labstack/echo)): 性能较好，和 gin 的用法差不多
- iris ([https://github.com/kataras/iris](https://github.com/kataras/iris)): 目前性能最好的框架，完备的 MVC
- beego（[https://github.com/astaxie/beego](https://github.com/astaxie/beego)）：高性能 web 框架
- go-restful ([https://github.com/emicklei/go-restful](https://github.com/emicklei/go-restful)) 快速构建 Restful 风格的框架

当然还有很多，主流的推荐这几款。在性能、请求参数、路由层面都具有优雅的设计。

下文使用 Iris 框架，对爱鲜蜂这款产品，进行具有 Restful API 接口的后台开发。

### 2.4 项目组织

为提高系统的可扩展性，项目必须具备良好的项目组织，整个的项目组织也是在需求开发的过程中不断进行微调，以满足不断变化的需求。web 框架一般采用标准的 MVC 架构，即: Model（模型层）、View（视图层）、Controller（控制层）。

```
├── Makefile
├── cmd
│   └── root_cmd.go
├── configs
│   └── config.yml
├── deployments
│   └── Dockerfile
├── main.go
├── model
│   ├── example_model.go
├── pkg
└── src
    └── account
        ├── assistance.go
        ├── controller.go
        ├── param.go
        └── router.go
```

各文件的功能说明：

- Makefile:  项目构建，提供简易的命令（类 Unix 操作系统支持）
- cmd：命令行工具，包括数据库表的创建、迁移、数据导入等
- configs：项目的配置文件，包括数据库的配置文件等
- deployments：容器相关的文件
- main.go：项目主入口
- model： 模型文件
- pkg：项目的使用库
- src： 项目的核心逻辑
  - account: 产品实体的抽象
    - assistance.go 辅助函数
    - controller.go 控制器的核心处理
    - param.go 请求参数
    - router.go 路由

### 2.4.1 设计稿和需求文档

这一步的目的明确你需要做什么，你可能需要反复和产品经理磨合，确保你理解的需求和产品经理的一致，减少后期的修改，甚至推倒重做。

当然阅读此书的人绝大多数可能并没有接触企业中的项目开发，更别说查看设计师的设计稿、产品经理的需求文档。那是不是没有办法了? 其实我们可以多关注一些已经在上线的 APP，把已经上线的 APP  当成是设计稿，你的目的是从设计稿的角度，分析出对象实体，进而进行模型设计，代码开发，完成最终的开发目的。

本环节就是使用已经上线的 “爱鲜蜂” 生鲜平台的客户端，分析出对象实体，进而开发出具有 Restful 风格的 API。

### 2.4.2 模型设计

模型设计是整个项目中最重要的一环之一，模型设计简单来说就是：数据库表设计，包括字段设计、字段类型设计、数据库表名设计等。

尽管市面上存在着各种各样的数据库，包括：关系型数据库、基于键值对的数据库、基于文档的数据库... , 但关系型数据库仍然是首选，一方面，关系型数据库诸如 MySQL、PostgreSQL 等，都是开源免费的，另一方面，关系型数据库对数据的组织非常友好，能够满足绝大多数应用场景，在业务量不是很大的场景下，关系型数据库是完全足够满足需求。

而关系型数据库的使用的前提是数据库表的设计，表设计的好，利于数据存储、完成开发目的、面对多变的需求 ，能够很好的进行扩展。

那么如何进行模型设计？

模型的设计的在于表的设计，表设计包括两个方面，表名的设计，表中字段以及字段类型的设计。模型是对实体的抽象，意味着设计表结构，首先你得明确真实的实体是什么？如何知道这些实体，从设计图、需求文档中了解这些。

表名的设计统一规范：数据库 + 实体的形式，比如 ：beeQuick_accout

下面参照着 “爱鲜蜂” APP 已经成型的产品反推出模型设计。

**首页**：

![产品稿-首页.jpeg](http://ww1.sinaimg.cn/large/741fdb86gy1g9yny8jjvrj20u01f5n16.jpg)

首页的逻辑是：选中某家店，加载出对应店的活动以及相应的商品。本质上活动是为了销售出商品。所以首页相关的实体是：

1. 店铺 
2. 活动

店铺涉及到具体的地址，涉及到实体是：
1. 省市区 
2. 具体的街道地址

首页的模型设计如下：

```
// base_model.go
// 主键和时间的字段，每个表都需要，单独提取出一个结构体来
type base struct {
	ID        uint      `xorm:"pk autoincr notnull 'id'" json:"id"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" json:"deleted_at"`
}

// shop_model.go
// 商店模型设计
type Shop struct {
	base       `xorm:"extends"`
	Location   string   `xorm:"varchar(255)" json:"location"`
	ProvinceId int64    `xorm:"index"`
	Province   Province `xorm:"-" json:"—"`
	Name       string   `xorm:"varchar(64)"`
}

// 商店名称
func (c Shop) TableName() string {
	return "beeQuick_shop"
}

type ShopSerializer struct {
	Id         int64              `json:"id"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	ProvinceId int64              `json:"province_id"`
	Province   ProvinceSerializer `json:"province"`
	Name       string             `json:"name"`
	Location   string             `json:"location"`
}

func (c Shop) Serializer() ShopSerializer {
	return ShopSerializer{
		Id:         int64(c.ID),
		CreatedAt:  c.CreatedAt.Truncate(time.Second),
		UpdatedAt:  c.UpdatedAt.Truncate(time.Second),
		Province:   c.Province.Serializer(),
		ProvinceId: c.ProvinceId,
		Name:       c.Name,
		Location:   c.Location,
	}
}
```

每一个模型单独写一个序列化的结构体，这些是用于前端或者客户端调用接口呈现出的字段。

```
// province_model.go

type Province struct {
	base     `xorm:"extends"`
	Name     string `xorm:"varchar(10)" json:"name"`
	AdCode   string `xorm:"varchar(10)" json:"ad_code"`
	CityCode string `xorm:"varchar(6)" json:"city_code"`
	Center   string `xorm:"varchar(32)" json:"center"`
	Level    string `xorm:"varchar(10)" json:"level"`
}

func (p Province) TableName() string {
	return "beeQuick_province"
}

type ProvinceSerializer struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	AdCode    string    `json:"ad_code"`
	CityCode  string    `json:"city_code"`
	Center    string    `json:"center"`
	Level     string    `json:"level"`
}

func (p Province) Serializer() ProvinceSerializer {
	return ProvinceSerializer{
		Id:        int(p.ID),
		CreatedAt: p.CreatedAt.Truncate(time.Second),
		UpdatedAt: p.UpdatedAt.Truncate(time.Second),
		Name:      p.Name,
		AdCode:    p.AdCode,
		Center:    p.Center,
		Level:     p.Level,
		CityCode:  p.CityCode,
	}
}
```

省市区的资源一旦导入数据库中，不太会变更，所以，针对省市区的模型，后续的接口，一般不会有更新或者删除操作。

活动模型：

```
// 活动模型

type Activity struct {
	base    `xorm:"extends"`  // 组合基础字段
	Name    string    `xorm:"varchar(32)" json:"name"`
	Title   string    `xorm:"varchar(32)" json:"title"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Avatar  string    `xorm:"varchar(255)" json:"avatar"`
	ShopIds []int     `xorm:"blob" json:"shop_ids"`
	Status  int       `xorm:"varchar(10)"`
}

// 活动模型的表名

func (a Activity) TableName() string {
	return "beeQuick_activity"
}

// 活动模型序列化结构体
type ActivitySerializer struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Avatar    string    `json:"avatar"`
	ShopIds   []int     `json:"shop_ids"`
	Status    string    `json:"status"`
}

func (a Activity) Serializer() ActivitySerializer {
	return ActivitySerializer{
		Id:        a.ID,
		CreatedAt: a.CreatedAt.Truncate(time.Second),
		UpdatedAt: a.UpdatedAt.Truncate(time.Second),
		Name:      a.Name,
		Title:     a.Title,
		Start:     a.Start,
		End:       a.End,
		Avatar:    a.Avatar,
		ShopIds:   a.ShopIds,
		Status:    ActivityStatus[a.Status],
	}
}

// 提供一些状态，用于表示活动的状态
const (
	DOING = iota
	PROGRESSING
	CANCEL
	FINISH
	ADVANCE
)

var ActivityStatus = make(map[int]string)
var ActivityStatusEn = make(map[int]string)

func init() {
	ActivityStatus[DOING] = "未开始"
	ActivityStatus[PROGRESSING] = "进行中"
	ActivityStatus[CANCEL] = "取消"
	ActivityStatus[FINISH] = "结束"
	ActivityStatus[ADVANCE] = "提前"

	ActivityStatusEn[DOING] = "DOING"
	ActivityStatusEn[PROGRESSING] = "PROGRESSING"
	ActivityStatusEn[CANCEL] = "CANCEL"
	ActivityStatusEn[FINISH] = "FINISH"
	ActivityStatusEn[ADVANCE] = "ADVANCE"

}

// 活动 和 商品 ： 多对多关系
type Activity2Product struct {
	ProductId  int64 `xorm:"index"`
	ActivityId int64 `xorm:"index"`
}

func (s Activity2Product) TableName() string {
	return "beeQuick_activity2Product"
}

// 商铺 和 活动： 多对多的关系
type Shop2Activity struct {
	ShopId     int64 `xorm:"index"`
	ActivityId int64 `xorm:"index"`
}

func (s Shop2Activity) TableName() string {
	return "beeQuick_shop2Activity"
}
```

关于首页的模型，大概核心就是：活动和店铺，其中店铺，设计省市区的一些资源，所以，模型中又包括省市区实体。

可以看到，我们的做法是：
1. 根据设计稿和需求，列出实体包含哪些内容 
2. 定义模型，模型提供一个序列化方法，用于前端或者客户端的展示

**分类**：

![BeeQuick2.jpeg](http://ww1.sinaimg.cn/large/741fdb86gy1g9yo52fce8j20u01f8413.jpg)

从设计稿中可以看出，分类项主要包括：
1. 分类标签（热销榜、整箱购、优选购等） 
2. 商品（名称、价格、品牌、规格、数量等）

涉及的模型包括：标签、商品、品牌、数量单位

```
// 商品模型
type Product struct {
	base          `xorm:"extends"`
	ShopId        int64   `xorm:"index"`
	Name          string  `xorm:"varchar(128) 'name'" json:"name"`
	Avatar        string  `xorm:"varchar(255) 'avatar'" json:"avatar"`
	Price         float64 `xorm:"double 'price'" json:"price"`
	Discount      float64 `xorm:"double default(1) 'discount'" json:"discount"` // 默认为 1
	Specification string  `xorm:"varchar(128) 'specification'" json:"specification"`
	BrandId       int64   `xorm:"index"`
	TagsId        int64   `xorm:"index"`
	Period        string  `xorm:"varchar(64)" json:"period"`
	UnitsId       int64   `xorm:"index"`
	Units         Units   `xorm:"-"`
	Brands        Brands  `xorm:"-"`
	Shop          Shop    `xorm:"-"`
	Tags          Tags    `xorm:"-"`
}

// 商品表名称
func (p Product) TableName() string {
	return "beeQuick_products"
}
```

​    

```
// 商品序列化
type ProductSerializer struct {
	Id            uint      `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ShopId        int64     `json:"shop_id"`
	Name          string    `json:"name"`
	Avatar        string    `json:"avatar"`
	Price         float64   `json:"price"`
	DiscountPrice float64   `json:"discount_price"`
	Period        string    `json:"period"`
	BrandId       int64     `json:"brand_id"`
	TagsId        int64     `json:"tags_id"`
	UnitsId       int64     `json:"units_id"`
	ShopName      string    `json:"shop_name"`
	UnitsName     string    `json:"units_name"`
	BrandsName    string    `json:"brands_name"`
}

func (p Product) Serializer() ProductSerializer {
	return ProductSerializer{
		Id:            p.ID,
		CreatedAt:     p.CreatedAt.Truncate(time.Second),
		UpdatedAt:     p.UpdatedAt.Truncate(time.Second),
		ShopId:        p.ShopId,
		Name:          fmt.Sprintf("%s%s/%s", p.Name, p.Specification, p.Units.Name),
		Avatar:        p.Avatar,
		Price:         p.Price,
		DiscountPrice: p.Price * p.Discount,
		Period:        p.Period,
		BrandId:       p.BrandId,
		TagsId:        p.TagsId,
		UnitsId:       p.UnitsId,
		ShopName:      p.Shop.Name,
		UnitsName:     p.Units.Name,
		BrandsName:    p.Brands.ChName,
	}
}

// 数量单位
type Units struct {
	base      `xorm:"extends"`
	Name      string `xorm:"unique" json:"name"`
	EnName    string `xorm:"unique" json:"en_name"`
	ShortCode string `xorm:"unique" json:"short_code"`
}

// 单位表名称
func (u Units) TableName() string {
	return "beeQuick_units"
}

type UnitsSerializer struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	EnName    string    `json:"en_name"`
	ShortCode string    `json:"short_code"`
}

func (u Units) Serializer() UnitsSerializer {
	return UnitsSerializer{
		Id:        int64(u.ID),
		CreatedAt: u.CreatedAt.Truncate(time.Second),
		UpdatedAt: u.UpdatedAt.Truncate(time.Second),
		Name:      u.Name,
		EnName:    u.EnName,
		ShortCode: u.ShortCode,
	}
}

// 品牌表
type Brands struct {
	base   `xorm:"extends"`
	ChName string `xorm:"unique" json:"ch_name"`
	EnName string `xorm:"unique" json:"en_name"`
}

// 品牌表名称
func (b Brands) TableName() string {
	return "beeQuick_brands"
}

type BrandsSerializer struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ChName    string    `json:"ch_name"`
	EnName    string    `json:"en_name"`
}

func (b Brands) Serializer() BrandsSerializer {
	return BrandsSerializer{
		Id:        int64(b.ID),
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		ChName:    b.ChName,
		EnName:    b.EnName,
	}
}

// 分类标签模型
type Tags struct {
	base `xorm:"extends"`
	Name string `xorm:"unique" json:"name"`
}

// 分类表名称
func (t Tags) TableName() string {
	return "beeQuick_tags"
}

type TagSerializer struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func (t Tags) Serializer() TagSerializer {
	return TagSerializer{
		Id:        int64(t.ID),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		Name:      t.Name,
	}
}

// 商铺和标签：多对多关系
type Shop2Tags struct {
	TagsId int64 `xorm:"index"`
	ShopId int64 `xorm:"index"`
}

func (s2t Shop2Tags) TableName() string {
	return "beeQuick_shop2Tags"
}

// 商品和标签：多对多关系
type Product2Tags struct {
	TagsId    int64 `xorm:"index"`
	ProductId int64 `xorm:"index"`
}

func (p2t Product2Tags) TableName() string {
	return "beeQuick_product2Tags"
}
```

基本步骤和首页的设计一致：根据设计稿和需求，列出实体对象，进行模型设计，为何的1对1关系、多对多关系等，通过相关的记录的 id  来进行维护。

**购物车**：

![BeeQuick3.jpeg](http://ww1.sinaimg.cn/large/741fdb86gy1g9yo6agfryj20u01f8tao.jpg)

从设计稿可以看出，为完成这个需求，涉及的实体：
1. 配送地址和时间 
2. 订单（商品、个数、总价、状态），另外，只有在登录状态才可以看到整个购物车页面。所以对资源的操作，涉及到登录状态和未登录状态，后续需要使用中间件的形式进行区分。

```
const (
	// 准备状态、未付款状态、已付款状态
	READINESS = iota
	BALANCE
	PAID
)

var (
	STATUS_MAP    = make(map[int]string)
	STATUS_MAP_EN = make(map[int]string)
)

func init() {
	STATUS_MAP[READINESS] = "准备状态"
	STATUS_MAP[BALANCE] = "未付款状态"
	STATUS_MAP[PAID] = "已付款状态"
	STATUS_MAP_EN[READINESS] = "readiness"
	STATUS_MAP_EN[BALANCE] = "balance"
	STATUS_MAP_EN[PAID] = "paid"

}

// 订单表
type Order struct {
	base       `xorm:"extends"`
	ProductIds []int `xorm:"blob"`
	Status     int
	AccountId  int64
	Account    Account `xorm:"-"`
	Total      float64
}

// 订单表名称
func (o Order) TableName() string {
	return "beeQuick_order"
}

type OrderSerializer struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Status     string    `json:"status"`
	Phone      string    `json:"phone"`
	AccountId  uint      `json:"account_id"`
	Total      float64   `json:"total"`
	ProductIds []int     `json:"product_ids"`
}

func (o Order) Serializer() OrderSerializer {
	return OrderSerializer{
		Id:         o.ID,
		CreatedAt:  o.CreatedAt.Truncate(time.Second),
		UpdatedAt:  o.UpdatedAt.Truncate(time.Second),
		Status:     STATUS_MAP[o.Status],
		AccountId:  o.Account.ID,
		Phone:      o.Account.Phone,
		Total:      o.Total,
		ProductIds: o.ProductIds,
	}
}
```

一个订单包含多个商品，所以1对多的关系，使用一个数组的形式，包含多个商品。

**个人中心**

![BeeQuick4.jpeg](http://ww1.sinaimg.cn/large/741fdb86gy1g9yo6xc2iuj20u01f3abn.jpg)

个人中心页面的逻辑是登录之后才能看到这些内容，如果不是登录状态，提示先登录，涉及到的实体比较多，核心的是：个人账户相关的内容：账号、优惠券、兑换券、会员体系等。

```
const (
	MEMBER      = "会员"
	ADMIN       = "管理员"
	SUPEARADMIN = "超级管理员"
)

// 账户模型
type Account struct {
	base     `xorm:"extends"`
	Phone    string    `xorm:"varchar(11) notnull unique 'phone'" json:"phone"`
	Password string    `xorm:"varchar(128)" json:"password"`
	Token    string    `xorm:"varchar(128) 'token'" json:"token"`
	Avatar   string    `xorm:"varchar(128) 'avatar'" json:"avatar"`
	Gender   string    `xorm:"varchar(1) 'gender'" json:"gender"`
	Birthday time.Time `json:"birthday"`

	Points      int       `json:"points"`
	VipMemberID uint      `xorm:"index"`
	VipMember   VipMember `xorm:"-"`
	VipTime     time.Time `json:"vip_time"`
}

func (Account) TableName() string {
	return "beeQuick_account"
}

type AccountSerializer struct {
	ID        uint                `json:"id"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Phone     string              `json:"phone"`
	Password  string              `json:"-"`
	Token     string              `json:"token"`
	Avatar    string              `json:"avatar"`
	Gender    string              `json:"gender"`
	Age       int                 `json:"age"`
	Points    int                 `json:"points"`
	VipMember VipMemberSerializer `json:"vip_member"`
	VipTime   time.Time           `json:"vip_time"`
}

func (a Account) Serializer() AccountSerializer {

	gender := func() string {
		if a.Gender == "0" {
			return "男"
		}
		if a.Gender == "1" {
			return "女"
		}
		return a.Gender
	}

	age := func() int {
		if a.Birthday.IsZero() {
			return 0
		}
		nowYear, _, _ := time.Now().Date()
		year, _, _ := a.Birthday.Date()
		if a.Birthday.After(time.Now()) {
			return 0
		}
		return nowYear - year
	}

	return AccountSerializer{
		ID:        a.ID,
		CreatedAt: a.CreatedAt.Truncate(time.Minute),
		UpdatedAt: a.UpdatedAt.Truncate(time.Minute),
		Phone:     a.Phone,
		Password:  a.Password,
		Token:     a.Token,
		Avatar:    a.Avatar,
		Points:    a.Points,
		Age:       age(),
		Gender:    gender(),
		VipTime:   a.VipTime.Truncate(time.Minute),
		VipMember: a.VipMember.Serializer(),
	}
}
```

​    

```
type AccountGroupVip struct {
	Account   `xorm:"extends"`
	VipMember `xorm:"extends"`
}

func (AccountGroupVip) TableName() string {
	return "beeQuick_account"
}
func (a AccountGroupVip) SerializerForGroup() AccountSerializer {
	result := a.Account.Serializer()
	result.VipMember = a.VipMember.Serializer()
	return result
}
```

```
const (
	// 兑换券，优惠券
	EXCHANGE = iota
	COUPON
)

var CouponType = make(map[int]string)

// 兑换券、优惠券
type ExchangeCoupon struct {
	base  `xorm:"extends"`
	Name  string    `xorm:"varchar(32) unique" json:"name"`
	Price float64   `json:"price"`
	Total float64   `json:"total"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Token string    `json:"token"`
	Type  int       `json:"type"` // 0,1 : 兑换券 抵消价格，优惠券 类似几折
}

func (exchange ExchangeCoupon) TableName() string {
	return "beeQuick_exchange_coupons"
}

type Account2ExchangeCoupon struct {
	AccountId        int64 `xorm:"index"`
	ExchangeCouponId int64 `xorm:"index"`
	Status           int   `json:"status"` // 0,1,2:未使用，已使用，已过期
}

func (a2e Account2ExchangeCoupon) TableName() string {
	return "beeQuick_account2exchange_coupon"
}

const (
	// 未使用、已使用、已过期
	NEW = iota
	USED
	EXPIRE
)

var StatusMap = make(map[int]string)

func init() {
	StatusMap[NEW] = "未使用"
	StatusMap[USED] = "已使用"
	StatusMap[EXPIRE] = "已过期"

	CouponType[EXCHANGE] = "兑换券"
	CouponType[COUPON] = "优惠券"
}

type ExchangeCouponSerializer struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Total     float64   `json:"total"`
	Start     string    `json:"start"` //  格式：2006/01/02
	End       string    `json:"end"`   // 格式：2006/01/02
	Status    string    `json:"status"`
	Type      string    `json:"type"`
}

func (exchange ExchangeCoupon) Serializer(status string) ExchangeCouponSerializer {

	return ExchangeCouponSerializer{
		ID:        exchange.ID,
		CreatedAt: exchange.CreatedAt.Truncate(time.Second),
		UpdatedAt: exchange.UpdatedAt.Truncate(time.Second),
		Name:      exchange.Name,
		Price:     exchange.Price,
		Total:     exchange.Total,
		Start:     exchange.Start.Format("2006-01-02 15:04:05"),
		End:       exchange.End.Format("2006-01-02 15:04:05"),
		Status:    status,
		Type:      CouponType[exchange.Type],
	}
}

// 会员规则
type RuleForExchangeOrCoupon struct {
	base     `xorm:"extends"`
	Question string `xorm:"unique"`
	Answer   string
	Type     int
}

func (RuleForExchangeOrCoupon) TableName() string {
	return "beeQuick_rule_coupon"
}

type RuleForExchangeOrCouponSerializer struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	Type      string    `json:"type"`
}

func (r RuleForExchangeOrCoupon) Serializer() RuleForExchangeOrCouponSerializer {
	return RuleForExchangeOrCouponSerializer{
		Id:        r.ID,
		CreatedAt: r.CreatedAt.Truncate(time.Second),
		UpdatedAt: r.UpdatedAt.Truncate(time.Second),
		Question:  r.Question,
		Answer:    r.Answer,
		Type:      CouponType[r.Type],
	}
}

const (
	V0 = iota
	V1
	V2
	V3
	V4
)

var validity = make(map[int]struct {
	Time    int
	Comment int
})

func init() {
	validity[V0] = struct {
		Time    int
		Comment int
	}{Time: 0, Comment: 0}

	validity[V1] = struct {
		Time    int
		Comment int
	}{Time: 1, Comment: 30}

	validity[V2] = struct {
		Time    int
		Comment int
	}{Time: 1, Comment: 169}

	validity[V3] = struct {
		Time    int
		Comment int
	}{Time: 1, Comment: 300}

	validity[V4] = struct {
		Time    int
		Comment int
	}{Time: 1, Comment: 500}
}

// 会员模型
type VipMember struct {
	base      `xorm:"extends"`
	LevelName string `xorm:"varchar(2) notnull unique 'level_name'" json:"level_name"`
	Start     int    `json:"start"`
	End       int    `json:"end"`
	Points    float64
	Comment   string `xorm:"varchar(128) notnull" json:"comment"`
	Period    int    `json:"period"`
	ToValue   int    `json:"to_value"`
}

func (VipMember) TableName() string {
	return "beeQuick_vip_member"
}

type VipMemberSerializer struct {
	ID        uint    `json:"id"`
	LevelName string  `json:"level_name"`
	Start     int     `json:"start"`
	End       int     `json:"end"`
	Comment   string  `json:"comment"`
	Period    int     `json:"period"`
	ToValue   int     `json:"to_value"`
	Points    float64 `json:"points"`
}

func (vip VipMember) Serializer() VipMemberSerializer {
	return VipMemberSerializer{
		ID:        vip.ID,
		LevelName: vip.LevelName,
		Start:     vip.Start,
		End:       vip.End,
		Comment:   vip.Comment,
		Period:    vip.Period,
		ToValue:   vip.ToValue,
		Points:    vip.Points,
	}
}

// 默认的会员级别和相关的积分、特权
func DefaultVipMemberRecord() []*VipMember {
	return []*VipMember{
		{
			LevelName: strings.ToUpper("v0"),
			Start:     0,
			End:       29,
			Points:    0.5,
			Comment:   fmt.Sprintf("获取%.2f倍积分", 0.5),
			Period:    0,
			ToValue:   0,
		},
		{
			LevelName: strings.ToUpper("v1"),
			Start:     30,
			End:       198,
			Points:    1.0,
			Comment:   fmt.Sprintf("获取%.2f倍积分", 1.0),
			Period:    1,
			ToValue:   30,
		},
		{
			LevelName: strings.ToUpper("v2"),
			Start:     199,
			End:       498,
			Points:    1.5,
			Comment:   fmt.Sprintf("获取%.2f倍积分", 1.5),
			Period:    1,
			ToValue:   169,
		},
		{
			LevelName: strings.ToUpper("v3"),
			Start:     499,
			End:       998,
			Points:    2.0,
			Comment:   fmt.Sprintf("获取%.2f倍积分", 2.0),
			Period:    1,
			ToValue:   300,
		},
		{
			LevelName: strings.ToUpper("v4"),
			Start:     999,
			End:       0,
			Points:    3.0,
			Comment:   fmt.Sprintf("获取%.2f倍积分", 3.0),
			Period:    1,
			ToValue:   500,
		},
	}
}
```

当然每个人设计的模型，有可能细节层面不一致，字段也可能不同，但目标都是一个，为了完成需求，最好的理解这些模型的设计的方法是，下载个 “爱鲜蜂” 的APP，进行页面的查看，看看每个页面上存在哪些对象实体，再进行模型设计的修正，看能不能符合要求。

四个主要的页面的模型设计，我们都遵循了一致的方法：

- 根据设计稿抽象出实体对象，列出实体需要的哪些字段
- 模型的表，统一使用 xorm 对字段和类型的设计
- 模型的表名称，统一使用：数据库 + 实体 进行命名
- 模型统一使用一个序列化的方法，把需要暴露给调用者的字段序列化
- 多对多关系，需要维护第三张表，把 id  维护起来即可

项目中 model 层，即模型的设计完成，后续需要把这些模型的定义转换成数据库表等。

### 2.4.3 数据库连接

数据库选择 MySQL，使用容器启动即可。推荐使用 Docker-compose 启动。docker-compose.yml 文件内容如下：

```
version: "3"

services:
  mysql:
    image: mysql:latest
    container_name: localMySQL
    volumes:
      - $PWD/mysql-data:/var/lib/mysql
    expose:
      - 3306
    ports:
      - 3306:3306
    environment:
      - MYSQL_ROOT_PASSWORD=admin123
      - MYSQL_DATABASE=beeQuick_dev
      - MYSQL_USER=root
```

上文的意思是：

- 拉取最新的 mysql 镜像
- 容器启动的别名为：localMySQL
- 挂载的目录为：$PWD/mysql-data:/var/lib/mysql  表示主机上的当前目录的 mysql-data 和容器内的 /var/lib/mysql 一致
- 暴露端口：3306
- 环境变量设置：数据库的密码、 数据库名称、数据库用户

启动命令：(启动之后CONTAINER ID值会有所不同）

```
docker-compose up -d

// 查看

docker ps
CONTAINER ID        IMAGE                 COMMAND                  CREATED             STATUS              PORTS                                            NAMES
a7115c72aa48        mysql:latest          "docker-entrypoint.s…"   2 weeks ago         Up 6 days           0.0.0.0:3306->3306/tcp, 33060/tcp                localMySQL
```

本地即启动的数据库，你也可以直接像在本地安装了 MySQL 一样进行数据库的操作：

```
docker exec -it a7115c72aa48 bash
root@a7115c72aa48:/# mysql -u root -p
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 47
Server version: 8.0.15 MySQL Community Server - GPL

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| beequick_dev       |
| gorm_example       |
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| xorm_example       |
+--------------------+
9 rows in set (0.28 sec)
```

那么在项目中如何进行相关的连接呢？

> 通过配置文件

根据不同的场景，连接不同的数据库，一般区分：生产、开发、测试环境

```
production:
   mysql:
     db: beequick_production
     user: root
     password: admin123
   postgres:
     db: beequick_production
     user: root
     password: admin123
dev:
   mysql:
     db: beequick_dev
     user: root
     password: admin123
   postgres:
     db: beequick_production
     user: root
     password: admin123
test:
   mysql:
     db: beequick_test
     user: root
     password: admin123
   postgres:
     db: beequick_production
     user: root
     password: admin123
```

configs/config.go

```
package configs

var ENV string

// ENV 用于区分不同的环境
```

pkg/database.v1 进行数据库的连接动作

```
package database_v1

import (
	"fmt"

	"github.com/go-xorm/core"

	"github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

var BeeQuickDatabase *xorm.Engine

var (
	drivenName string
	db         string
	port       string
	password   string
	user       string
	dsn        string
)

func init() {
	drivenName = "mysql"
	if drivenName == "mysql2" {
		config := mysql.NewConfig()
		config = &mysql.Config{
			User:   user,
			Passwd: password,
			DBName: db,
			Addr:   port,
		}
		dsn = config.FormatDSN()
	} else {
		//dsn = fmt.Sprintf("%s:%s@/%s?chartset=utf8&parseTime=true&loc=Local", user, password,db)
		dsn = fmt.Sprintf("root:admin123@/beequick_dev?charset=utf8&parseTime=true&loc=Local")
	}

}

func DataBaseInit() {
	var err error
	BeeQuickDatabase, err = xorm.NewEngine(drivenName, dsn)
	if err != nil {
		panic(err)
		return
	}
	BeeQuickDatabase.ShowSQL(true)
	BeeQuickDatabase.Logger()
	BeeQuickDatabase.Charset("utf8")
	BeeQuickDatabase.SetMapper(core.GonicMapper{})
	BeeQuickDatabase.SetTableMapper(core.SameMapper{})

}
```

具体的连接设置等，需要查看 xorm 的文档：[https://github.com/go-xorm/xorm](https://github.com/go-xorm/xorm)

### 2.4.4 iris 的简单使用

本项目选择的是 iris web 框架，在性能层面、参数的处理层面都非常到位，对于业务开发来说，极大的精简了流程。

构建 Restful API 风格的 web 项目，核心包括三点：路由的设计、参数的处理（校验等）、响应的处理。

如何使用 iris 启动项目呢？

```
package main

import "github.com/kataras/iris"

func main() {

		// 默认的服务引擎
    app := iris.Default()

    // 路由：/ping
		// 控制逻辑：返回 json 格式的内容
    app.Get("/ping", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "message": "pong",
        })
    })
    // listen and serve on http://0.0.0.0:8080.
    app.Run(iris.Addr(":8080"))
}
```

最简易的使用 iris 启动的 web 服务大概就是这样，当然，框架提供了更为丰富的功能

- 方法：Get, Post, Put, Patch, Delete and Options

- 路径参数：直接设置类型，或者自定义类型 

```
/get/{id:int}

```

那如何获取到路径中的参数 ? 通过 iris.Context

```
ctx.Params().GetInt("id)
```

其他的类型都提供了相应的方法。

- 路径中的请求参数: 形如 ?name=value&first=1

```
/get?name=go


ctx.URLParamDefault("name", "python")
ctx.URLParam("name")
```

既可以使用默认值的形式，也可以直接获取传入的值。

- 请求参数: 这种一般使用的 HHTP 方法是：Post 、Patch 或者 Put

```
func main() {
  app := iris.Default()


app.Post("/form_post", func(ctx iris.Context) {
  message := ctx.FormValue("message")
  nick := ctx.FormValueDefault("nick", "anonymous")

  ctx.JSON(iris.Map{
      "status":  "posted",
      "message": message,
      "nick":    nick,
  })
})

app.Run(iris.Addr(":8080"))


}

POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=manu&message=this_is_great
```

像上文，需要把参数带入请求中，获取服务器的响应。可以获取单个值，也可以获取默认值。

当然也可以直接使用 ctx.ReadJSON 方法，一次性载入所以请求参数

- 响应信息，一般我们选择 json 作为数据交换格式，所以响应 统一使用 ctx.JSON 即可


总结：上文主要是介绍 iris 的简单使用， web 框架的使用要抓住核心：

- 如何启动服务
- 如何根据不同的请求方法处理相应的请求参数，比如是 URL 中的参数还是请求体中的参数
- 如何返回相应格式的响应



### 2.4.5 项目开发

在 iris 的初步了解之上，构建我们的整个项目，根据之前的整个项目结构的设计，为方便扩展和维护，对各个项目都具有约束，都带有特定的目的，参照之前的项目结构，对项目进行开发。

cmd/root_cmd.go 以命令行的形式启动服务：

启动服务需要涉及哪些内容？

- 数据库连接
- web 服务启动


```
var rootCMD = &cobra.Command{
  Use:   "root command",
  Short: "root command",
  Long:  "run web server",
  Run:   runRootCMD,
}

func runRootCMD(cmd *cobra.Command, args []string) {


database_v1.DataBaseInit() // 数据库连接动作
iris.RegisterOnInterrupt(func() {
  database_v1.BeeQuickDatabase.Close()
})
app := router_v1.ApplyRouter() //  路由集合
err := app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
if err != nil {
  log.Fatal(err.Error())
}


}

func Execute() {
  if err := rootCMD.Execute(); err != nil {
  	log.Println(err.Error())
  	os.Exit(1)
  }
}
```
main.go 

```
package main

import (
	"log"

import (
	"log"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/cmd"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/configs"
)

var ENV string

func main() {
	if ENV == "" {
		configs.ENV = "dev"
	} else {
		configs.ENV = ENV
	}
	log.Println("Running Web Server")
	cmd.Execute()

}
```

在开发过程中，会频繁的编译、更新、迭代新版本，当然开发人员可以不断的 go build 编译，在 go run 启动新版本，这样的操作，比较繁琐，企业级的项目一般使用 Makefile 进行构建。

Makefile 是构建工具，类似于 shell 脚本，命令格式如下：

```
<target> : <prerequisites> 
[tab]  <commands>
```

- target 是提供的命令
- prerequisties 表示前置条件，即执行 命令之前会触发的命令
- commands 是真实的命令

go 项目，一般的会频繁操作，编译动作、启动程序动作，所以把这些命令浓缩到 Makefile 文件中，之后，只需要执行 Makefile 提供的命令即可。

```
BINARY=BeeQuick

VERSION=1.0.0

BUILD=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.Env=production -s -w"

DEV_LDFLAGS=-ldflags "-X main.Env=dev"

TEST_LDFLAGS=-ldflags "-X main.Env=test"

default:
	go build -o ${BINARY} -v ${DEV_LDFLAGS} -tags=jsoniter

production:
	go build -o ${BINARY} -v ${LDFLAGS} -tags=jsoniter

dev:
	go build -o ${BINARY} -v ${DEV_LDFLAGS} -tags=jsoniter

test:
	go build -o ${BINARY} -v ${TEST_LDFLAGS} -tags=jsoniter

run:
	go run -v ${DEV_LDFLAGS} -tags=jsoniter main.go

.PHONY: default production dev test run
```

存在该文件之后，执行的命令是：make 、make production、make dev 等。即将一些大段的执行命令，精简为几个单词，但达到了一样的效果。默认执行第一个命令，即 make default， 如果你什么参数也不带。

使用 Makefile 构建工具，之后的启动服务，只要在 Makefile 所在目录下，执行 make 命令即可。

```
// 启动程序
make run 
```

使用命令行工具，对表进行迁移：即将定义的模型转换成 MySQL 中具体的表。

```
var syncCMD = &cobra.Command{
	Use:   "sync2",
	Short: "xorm.Syn2(model)",
	Run:   sync2,
}

func sync2(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Panic("You should add one argument at least")
		return
	}
	database_v1.DataBaseInit()
	if args[0] == "db" {

		for _, i := range tables() {
			if err := database_v1.BeeQuickDatabase.Sync2(i); err != nil {
				fmt.Println(err)
			}
		}
	}
	if args[0] == "vip" {
		vipMember()
	}

	if args[0] == "drop" {
		database_v1.BeeQuickDatabase.DropTables(new(model_v1.Order))
	}
}

func tables() []interface{} {
	return []interface{}{
		new(model_v1.Account),
		new(model_v1.VipMember),
		new(model_v1.ExchangeCoupon),
		new(model_v1.Account2ExchangeCoupon),
		new(model_v1.RuleForExchangeOrCoupon),
		new(model_v1.Shop),
		new(model_v1.Province),
		new(model_v1.Activity),
		new(model_v1.Activity2Product),
		new(model_v1.Shop2Activity),
		new(model_v1.Product),
		new(model_v1.Product2Tags),
		new(model_v1.Tags),
		new(model_v1.Shop2Tags),
		new(model_v1.Brands),
		new(model_v1.Units),
		new(model_v1.Order),
	}
}

func vipMember() bool {
	if _, err := database_v1.BeeQuickDatabase.Insert(model_v1.DefaultVipMemberRecord()); err != nil {
		return false
	}
	return true
}
```

- 提供了sync2 命令，接受三个子命令：db、vip、drop

- db 表示对数据库表的创建，核心是使用 xorm.Engine.Sync2 方法

- vip 是导入会员体系，即默认的会员 DefaultVipMemberRecord

- drop 是演示如何删除表结构

```
  // 修改 root_cmd.go 文件，使 sync2 成为其子命令
  // 核心是使用了 github.com/spf13/cobra 构建命令行工具，非常简便

  func Execute() {
  	rootCMD.AddCommand(syncCMD)
  	if err := rootCMD.Execute(); err != nil {
  		log.Println(err.Error())
  		os.Exit(1)
  	}
  }

  // rootCMD.AddCommand 使 syncCMD 成为其子命令

```

表结构创建命令：

```
make default // 编译
./BeeQuick sync2 db // 创建表

// 创建命令日志已开启：BeeQuickDatabase.ShowSQL(true)
// 终端中可以看到类似的日志

[xorm] [info]  2019/06/23 21:32:26.230609 [SQL] SELECT `COLUMN_NAME`, `IS_NULLABLE`, `COLUMN_DEFAULT`, `COLUMN_TYPE`, `COLUMN_KEY`, `EXTRA`,`COLUMN_COMMENT` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? [beequick_dev beeQuick_tags]
[xorm] [info]  2019/06/23 21:32:26.235318 [SQL] SELECT `INDEX_NAME`, `NON_UNIQUE`, `COLUMN_NAME` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? [beequick_dev beeQuick_tags]
[xorm] [info]  2019/06/23 21:32:26.239562 [SQL] SELECT `COLUMN_NAME`, `IS_NULLABLE`, `COLUMN_DEFAULT`, `COLUMN_TYPE`, `COLUMN_KEY`, `EXTRA`,`COLUMN_COMMENT` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? [beequick_dev beeQuick_units]
[xorm] [info]  2019/06/23 21:32:26.244750 [SQL] SELECT `INDEX_NAME`, `NON_UNIQUE`, `COLUMN_NAME` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? [beequick_dev beeQuick_units]
[xorm] [info]  2019/06/23 21:32:26.250876 [SQL] SELECT `COLUMN_NAME`, `IS_NULLABLE`, `COLUMN_DEFAULT`, `COLUMN_TYPE`, `COLUMN_KEY`, `EXTRA`,`COLUMN_COMMENT` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? [beequick_dev beeQuick_vip_member]
[xorm] [info]  2019/06/23 21:32:26.258744 [SQL] SELECT `INDEX_NAME`, `NON_UNIQUE`, `COLUMN_NAME` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? [beequick_dev beeQuick_vip_member]
```

至此，表结构的创建命令等已经完成，剩下的便是进行核心的业务开发（内容多，但大体步骤一致，仅以资源：账户接口为例）

路由集合：

```
// pkg/router.v1/router.go 

var (
	VERSION = "v0.1.0"
)

func ApplyRouter() *iris.Application {
	app := iris.Default()

	notFound(app)

	app.Handle("GET", "/", func(context iris.Context) {
		_, _ = context.JSON(iris.Map{
			"data": time.Now().Format("2006-01-02 15:04:05"),
			"code": http.StatusOK,
		})
	})

	app.Handle("GET", "/heart", func(c iris.Context) {
		c.JSON(iris.Map{
			"data": time.Now().Format("2006-01-02 15:04:05"),
			"code": http.StatusOK,
		})
	})

	v1 := app.Party("/v1")
	v1.Get("/version", func(context iris.Context) {
		context.JSON(
			iris.Map{
				"code":    http.StatusOK,
				"version": VERSION,
			},
		)
		return
	})

	app.UseGlobal(middleware.LoggerForProject)
	{

		account.Default.RegisterWithOut(app, "/v1")
		rule.Default.RegisterWithout(app, "/v1")
		province.Default.RegisterWithOut(app, "/v1")
		shop.Default.RegisterWithout(app, "/v1")
		activity.Default.Register(app, "/v1", false)
		unit.Default.Register(app, "/v1")
		brand.Default.Register(app, "/v1")
		tags.Default.Register(app, "/v1")
		product.Default.Register(app, "/v1")
	}

	app.Use(middleware.TokenForProject)

	{
		account.Default.RegisterWith(app, "/v1")
		vip_member.Default.Register(app, "/v1")
		exchange_coupons.Default.Register(app, "/v1")
		activity.Default.Register(app, "/v1", true)
		order.Default.Register(app, "/v1")
	}

	app.Logger().SetLevel("debug")
	return app
}

func notFound(app *iris.Application) {
	app.OnErrorCode(http.StatusNotFound, func(context iris.Context) {
		context.JSON(iris.Map{
			"code":   http.StatusNotFound,
			"detail": context.Request().URL.Path,
			"error":  "error found",
		})
	})
	return
}
```

这里的逻辑是：路由的集合，包括未匹配到路由（notFound）的报错信息，路由分组（app.Party)，使用中间件（app.Use)。

一方面，登录之后才能才看某些资源，因此，需要对接口的部分内容进行限制，即，使用中间件的形式对接口进行区分。

**中间件**

```
//pkg/middleware/middleware.go
package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/model/v1"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/database.v1"

	"github.com/kataras/iris"
)

// 日志中间件
func LoggerForProject(c iris.Context) {
	c.Application().Logger().Debugf("Path: %s | IP: %s | Time: %s",
		c.Path(), c.RemoteAddr(), time.Now().Format("2006-01-02 15:04:05"))
	c.Next()
}
```

​    

```
// 认证中间件
func TokenForProject(c iris.Context) {
	token := c.GetHeader("Authorization")
	tokenList := strings.Split(token, " ")
	if len(tokenList) != 2 || tokenList[0] != "Bearer" {
		c.JSON(iris.Map{
			"code": http.StatusNotFound,
			"err":  "Header Add Authorization: Bearer xxx",
		})
		return
	}
	realToken := tokenList[1]
	var account model_v1.Account
	if _, err := database_v1.BeeQuickDatabase.Where("token = ?", realToken).Get(&account); err != nil {
		c.JSON(iris.Map{
			"code": http.StatusNotFound,
			"err":  err.Error(),
		})
		return
	}
	c.Values().Set("current_admin", account)
	c.Values().Set("current_admin_id", account.ID)
	c.Next()

}
```

- 日志中间件，即访问路由的过程中会打印某些日志，定义了日志的格式
- 认证中间件，即用户需要在头部信息内带上 Authorization: Bearer XXX 格式的头部信息，具体服务端会去数据库内进行内容的收拾，确保头部信息带的认证信息在数据库中存在，即用户已登录

之前已经约定过：核心业务的开发项目的组织如下：

```
// src/account.go
├── assistance.go // 辅助函数
├── controller.go // 核心控制逻辑
├── param.go // 请求参数处理
├── response.go // 响应信息
└── router.go // 路由设计
```

**router.go 设计**

```
package account

import "github.com/kataras/iris"

type ControllerForAccount struct {
}

var Default = ControllerForAccount{}

func (controller ControllerForAccount) RegisterWithOut(app *iris.Application, path string) {
	middleware := func(context iris.Context) {
		context.Next()
	}

	account := app.Party(path, middleware)
	{
		account.Post("/register", registerHandle)
		account.Post("/sign", signHandle)

	}

}

func (controller ControllerForAccount) RegisterWith(app *iris.Application, path string) {
	middleware := func(context iris.Context) {
		context.Next()
	}

	account := app.Party(path, middleware)
	{
		account.Post("/logout", logoutHandle)
		account.Get("/account/{id:uint}", getAccountHandle)
	}
}
```

账户信息相关的接口包括：

- 登录：即使用用户名和密码，匹配数据库中是否存在记录 ：/sign
- 登出：即退出信息:  /logout
- 注册：即使用手机号，注册新账户: /register
- 获取账户信息: /account/id

**controller.go** 控制核心逻辑

- 注册：/v1/register

- 注册时请求参数：phone、password

```
  func registerProcessor(param RegisterParam) (model_v1.AccountGroupVip, error) {
  	var (
  		account model_v1.AccountGroupVip
  		errV1   error_v1.ErrorV1
  	)
  	if err := param.Valid().Struct(param); err != nil {
  		return account, error_v1.ErrorV1{
  			Code:    http.StatusBadRequest,
  			Message: "param not valid",
  			Detail:  "请求参数校验不通过，请检查参数",
  		}
  	}
    var vipMember model_v1.VipMember
    
    if _, dbErr := database_v1.BeeQuickDatabase.Where("level_name = ?", strings.ToUpper("v0")).Get(&vipMember); dbErr != nil {
      return account, error_v1.ErrorV1{
      	Code:    http.StatusBadRequest,
      	Message: dbErr.Error(),
      	Detail:  "会员等级未存在",
      }
    }
    account.VipMember = vipMember
    
    hashPassword, _ := generateFromPassword(param.Password, 8)
    hashToken := generateToken(20)
    
    account.Account = model_v1.Account{
      Phone:       param.Phone,
      Password:    string(hashPassword),
      Token:       hashToken,
      Points:      0,
      VipMemberID: vipMember.ID,
      VipTime:     time.Now(),
    }
    
    if _, err := database_v1.BeeQuickDatabase.InsertOne(&account.Account); err != nil {
      errV1 = error_v1.ErrorV1{
      	Code:    http.StatusBadRequest,
      	Message: err.Error(),
      	Detail:  "用户注册发生错误",
      }
      return account, errV1
    }
    return account, nil
}

func registerHandle(ctx iris.Context) {
  var param RegisterParam
  err := ctx.ReadJSON(&param)
  if err != nil {
  	ctx.JSON(makeResponse(http.StatusBadRequest, err, true))
  	return
  }
  account, err := registerProcessor(param)
  if err != nil {
  	ctx.JSON(makeResponse(http.StatusBadRequest, err, true))
  	return
  }
  ctx.JSON(iris.Map(makeResponse(http.StatusOK, account.SerializerForGroup(), false)))

}
```


整体的处理步骤是：

- 读取请求参数，检验请求参数：比如手机号位数不正确，密码不符合格式
- 提供默认的会员等级：v0
- 如果都符合，即在数据库中生成一条记录，记录当前注册账户的信息

**param.go**

```
type RegisterParam struct {
	Phone    string `form:"phone" json:"phone" validate:"required,len=11"`
	Password string `form:"password" json:"password"`
}

// 参数校验
func (param RegisterParam) suitable() (bool, error) {
	if param.Password == "" || len(param.Phone) != 11 {
		return false, fmt.Errorf("password should not be nil or the length of phone is not 11")
	}
	if unicode.IsNumber(rune(param.Password[0])) {
		return false, fmt.Errorf("password should start with number")
	}
	return true, nil
}

// 参数校验使用 Tag 检查
func registerValidation(sl validator.StructLevel) {
	param := sl.Current().Interface().(RegisterParam)
	if param.Phone == "" && param.Password == "" {
		sl.ReportError(param.Password, "Password", "password", "password", param.Password)
		sl.ReportError(param.Phone, "Phone", "phone", "phone", param.Phone)
	}
}

func (param RegisterParam) Valid() *validator.Validate {
	validate := validator.New()
	validate.RegisterStructValidation(registerValidation, RegisterParam{})
	return validate
}
```

参数在请求中需要注意的是：校验

- 比如类型
- 比如长度
- 是否符合特定的规则，比如邮箱，需要包含@ 等

校验也有许多方法：

- 自定义结构体，设置结构体的方法进行校验，比如： suitable 方法
- 使用 Tag 进行校验，主要使用了 "[gopkg.in/go-playground/validator.v9](http://gopkg.in/go-playground/validator.v9)" 库

**response.go**

Restful API 风格的数据交换格式一般选择：json

```
func makeResponse(code int, value interface{}, isError bool) map[string]interface{} {
	result := make(map[string]interface{})
	result["code"] = code
	if isError {
		result["error"] = value
	} else {
		result["data"] = value
	}
	return result
}
```

响应区分：处理正确的、处理出错的。

- 正确的使用：code, data 的格式
- 错误的使用：code, error 的格式

**assistance.go**

辅助函数一般是给核心处理逻辑提供帮助的，即核心处理逻辑只专注在业务层面。

```
package account

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func generateFromPassword(password string, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), cost)
}

func compareHashAndPassword(hashed []byte, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashed, password); err != nil {
		return false
	}
	return true
}

func generateToken(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
```

辅助函数包括:

- 根据密码加密生成密钥（不明文存储密码，不安全）： generateFromPassword
- 根据密码比对 Token：compareHashAndPassword
- 生成 Token，即密码和 Token 唯一对应： generateToken

其他处理逻辑：

**登录**

登录的逻辑是：用户已经注册，即数据库包含用户的某些信息，进行匹配即可

请求参数依然是：手机号 和 密码

为丰富响应信息，这边将账户表和会员表进行了 join

```
func signProcessor(param RegisterParam) (model_v1.AccountGroupVip, error) {

	var (
		account model_v1.AccountGroupVip
		err     error
	)

	if err := param.Valid().Struct(param); err != nil {
		err = error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "登录参数校验失败",
			Message: err.Error(),
		}
		return account, err
	}
	if _, err := database_v1.BeeQuickDatabase.Join("INNER", "beeQuick_vip_member", "beeQuick_vip_member.id = beeQuick_account.vip_member_id").Get(&account); err != nil {
		err = error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "账号未注册",
			Message: err.Error(),
		}
		return account, err
	}
	if !compareHashAndPassword([]byte(account.Password), []byte(param.Password)) {
		err = error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "密码错误",
			Message: "password not correct",
		}
		return account, err
	}

	return account, nil
}

func signHandle(ctx iris.Context) {
	var param RegisterParam
	err := ctx.ReadJSON(&param)
	if err != nil {
		ctx.JSON(iris.Map(makeResponse(http.StatusBadRequest, err, true)))
		return
	}
	account, err := signProcessor(param)
	if err != nil {
		ctx.JSON(iris.Map(makeResponse(http.StatusBadRequest, err, true)))
		return
	}

	ctx.JSON(iris.Map(makeResponse(http.StatusOK, account.SerializerForGroup(), false)))

}
```

**接口测试**

之前的环节提供了几种供读者使用的接口测试：curl、httpie、postman、vscode插件

本环节使用：httpie

注册：

```
http http://127.0.0.1:8080/v1/register phone='18717711830' password='admin123'

// 响应信息

HTTP/1.1 200 OK
Content-Length: 393
Content-Type: application/json; charset=UTF-8
Date: Sun, 23 Jun 2019 14:14:15 GMT
Proxy-Connection: keep-alive

{
    "code": 200,
    "data": {
        "age": 0,
        "avatar": "",
        "created_at": "2019-06-23T22:14:00+08:00",
        "gender": "",
        "id": 8,
        "phone": "18717711830",
        "points": 0,
        "token": "a817317eae338fbc9d09d2e76021afca7d1c3d7e",
        "updated_at": "2019-06-23T22:14:00+08:00",
        "vip_member": {
            "comment": "获取0.50倍积分",
            "end": 29,
            "id": 1,
            "level_name": "V0",
            "period": 0,
            "points": 0.5,
            "start": 0,
            "to_value": 0
        },
        "vip_time": "2019-06-23T22:14:00+08:00"
    }
}
```

注册时默认生成一些关联的配置，比如积分为0， 会员等级为 V0 等。

获取：

```
// 头部信息需带 Token

http http://127.0.0.1:8080/v1/account/8 'Authorization:Bearer a817317eae338fbc9d09d2e76021afca7d1c3d7e'

// 响应

HTTP/1.1 200 OK
Content-Length: 393
Content-Type: application/json; charset=UTF-8
Date: Sun, 23 Jun 2019 14:18:18 GMT
Proxy-Connection: keep-alive

{
    "code": 200,
    "data": {
        "age": 0,
        "avatar": "",
        "created_at": "2019-06-23T22:14:00+08:00",
        "gender": "",
        "id": 8,
        "phone": "18717711830",
        "points": 0,
        "token": "a817317eae338fbc9d09d2e76021afca7d1c3d7e",
        "updated_at": "2019-06-23T22:14:00+08:00",
        "vip_member": {
            "comment": "获取0.50倍积分",
            "end": 29,
            "id": 1,
            "level_name": "V0",
            "period": 0,
            "points": 0.5,
            "start": 0,
            "to_value": 0
        },
        "vip_time": "2019-06-23T22:14:00+08:00"
    }
}
```

请求得出的响应，是正确的按照定义的格式：code 和 data 字段的格式，这是因为我们模型设计环节，自定义了序列化的响应。

即：将 account 模型的序列化 和 会员体系 vipmember 模型的序列化组合起来

```
type AccountGroupVip struct {
	Account   `xorm:"extends"`
	VipMember `xorm:"extends"`
}

func (AccountGroupVip) TableName() string {
	return "beeQuick_account"
}
func (a AccountGroupVip) SerializerForGroup() AccountSerializer {
	result := a.Account.Serializer()
	result.VipMember = a.VipMember.Serializer()
	return result
}
```

那么如果失败呢，如何响应？

```
http http://127.0.0.1:8080/v1/account/8 'Authorization:Bear no data'

// 响应

HTTP/1.1 200 OK
Content-Length: 57
Content-Type: application/json; charset=UTF-8
Date: Sun, 23 Jun 2019 14:21:58 GMT
Proxy-Connection: keep-alive

{
    "code": 404,
    "err": "Header Add Authorization: Bearer xxx"
}

// Token 不正确，产生报错信息
```

鉴于篇幅，其他的接口处理思路基本一致：

- 对路由进行设计
- 对请求参数进行校验
- 逻辑处理，比如成功了在数据库中记录数据；失败了回滚，进行报错提示

更多详情：参考源代码

```
https://github.com/wuxiaoxiaoshen/GopherBook/tree/master/chapter10/BeeQuick.v1
```

## 2.6 总结

web 开发是各种编程语言中一个非常重要的领域，涉及的知识也非常的繁多，go 原生提供的了网络请求包，是一切网络框架的基石，框架核心的处理在于对路由的处理、请求参数的处理、响应的处理。

学习 web 开发需要了解很多的知识：

- HTTP 协议
- 网络请求的流程
- 如何设计 web 服务

本章的核心在于告诉读者如何构建 Restful API 形式的 web 服务，企业中的开发整体的步骤也大致是如此，但企业中涉及的流程、需求等、体量等都非常的庞大，需要不断的对整体的架构进行调整。

web 服务的开发，包括四个重要的步骤：

- 请求方法：请求方法决定了对服务器上资源的何种操作，比如是创建还是获取
- 请求参数：请求参数校验由具体的需求决定，比如数值型，约定大小关系，比如字符串型，约定长度，是否为空等
- 请求路由：请求路由决定了哪些路径能够获取到服务器上的哪些资源
- 响应信息：响应信息决定了用户层面能够看到的服务器上资源的形式

