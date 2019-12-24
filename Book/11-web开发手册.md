
上一章主要讲解了如何构建 web 服务。主要包括：

- template 模版的使用
- http 协议
- 原生 net/http 构建 web 服务
- 接口测试工具
- web 框架 iris 构建 web 服务


这一章，针对 web 开发会遇到的问题进行方法论的总结，让读者对 web 开发形成自己的一套知识体系，之后不管是使用原生的 net/http 还是 web 框架进行 web 开发都能知道知识点是什么？遇到问题如何解决？web 开发的难点在什么地方等...



### 1. 再谈HTTP

web 开发是在 HTTP 协议的基础之上进行构建的，用于完成客户端和服务端之间的通信，发起请求的一方称之为客户端，提供资源响应的一方称之为服务端。客户端常见的形式是：web 网页（也就是我们通常说的前端）、Andriod 客户端、ios 客户端；服务端一般都是服务器实体，将程序运行在服务器上，提供资源的访问。

客户端向服务端发起请求，遵循 HTTP 协议，建立起通信之后，服务端返回响应信息，传输给客户端，客户端以网页等的形式展示出来。

协议即一些约束，HTTP 协议约束了客户端和服务端之间通信需要遵守的规定。要知道 HTTP 协议包含哪些知识，最简单的办法是使用浏览器的调试功能（比如 Chrome 的 F12 进行审查元素），查看请求的具体信息。

```

GET /v1/api/vote?vote_id=1 HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Authorization: Beaer d757e670d62d16921edb
Connection: keep-alive
Host: localhost:7201
User-Agent: HTTPie/0.9.9



HTTP/1.1 200 OK
Content-Length: 573
Content-Type: application/json
Date: Mon, 15 Jul 2019 03:41:49 GMT

{
    "code": 200,
    "data": {
        "admin_id": 1,
        "choice": [
            {
                "choice_title": "学习",
                "created_at": "2019-07-14T11:42:33+08:00",
                "id": 1,
                "number": 1,
                "ratio": "100.0%",
                "updated_at": "2019-07-14T15:49:22+08:00",
                "vote_title": "辞职"
            },
            {
                "choice_title": "Java",
                "created_at": "2019-07-14T11:42:33+08:00",
                "id": 2,
                "number": 0,
                "ratio": "0.0%",
                "updated_at": "2019-07-14T11:42:33+08:00",
                "vote_title": "辞职"
            }
        ],
        "created_at": "2019-07-14T11:42:33+08:00",
        "dead_line": "2020-07-14T11:42:33+08:00",
        "description": "",
        "id": 1,
        "is_anonymous": false,
        "is_single": true,
        "title": "辞职",
        "updated_at": "2019-07-14T11:42:33+08:00"
    }
}

```

发起请求规定：访问方法、访问资源的路径、HTTP协议的版本、头部信息

返回信息规定：协议版本、状态码、头部信息、响应消息的主体

涉及 HTTP 协议中有哪些知识点？


1. HTTP 的请求方法分类和异同
2. HTTP 的状态码的分类和含义
3. 路由：即访问资源的具体路径
4. 响应信息


### 2. 设计 RESTful API


REST (Reporesentational State Transfer) 中文含义表现层状态转化，符合 REST 设计规则的架构称之为 RESTful 架构。该设计原则实际上是对资源数据的转化和状态的变化。

资源表示的是网络上的一个实体，比如 Github 资源有仓库、用户、代码等，所以 RESTful API 是对这些资源的操作，比如：删除、查看、修改、增加等动作。不同的动作需要使用不同的方法：比如：DELETE、GET、PUT、POST等动作的方法。资源以不同的形式展现出来，比如：纯文本（text/plain)、HTML、XML、JSON、String 等格式。


在设计前后端分离的开发模式的过程中，经常使用 RESTful 设计架构。后端开发人员负责开发出资源的操作的接口，前端人员提供一个交互、展示等，直接调用接口，完成对资源的操作。RESTful API 是前后端之间的桥梁。

RESTful API 有一定的设计规范。这些设计规范内容和 HTTP 协议的内容高度重合。

- 设计资源的访问动作
- 设计路由，即资源的唯一访问地址 URL
- 设计状态码
- 设计资源的响应格式


#### 2.1 资源的访问动作

首先得知道，HTTP 协议支持哪些动作？

- HEAD、GET、POST、PUT、DELETE、OPTIONS


具体的选择方案如下：

- 获取资源：GET 
- 创建资源：POST
- 更新资源：PUT(PATCH 通常部分更新)
- 删除资源：DELETE
- 查询服务器支持的方法：OPTIONS
- 仅查询头部信息：HEAD


在前端开发过程中发起网络请求，经常会接触到 CORS (Cross-Origin Resource Sharing) 跨域，指的是浏览器除可以访问当前页面的域之外，还可以访问其他的域，简单的说就是在开发过程中可以访问其他 URL。跨域是如何实现的？一般都是靠服务端对HTTP 头部信息的设置运行跨域，达到跨域的目的。

简单的说：web 端像服务端发起一个 OPTIONS 动作的请求，服务端将支持的请求方法和一些头部信息返回给 web 端，运行跨域之后，再真正的发起网络请求，操作资源。

```
// 经常在请求头部信息中存在这么一个字段，表示的是本域，要在页面内请求其他网址，就需要跨域

origin: https://www.google.com

```

这些设置大概如下：

```
Access-Control-Allow-Headers: Origin,Content-Length,Content-Type,Authorization
Access-Control-Allow-Methods: GET,POST,PUT,HEAD,OPTIONS,DELETE,PATCH
Access-Control-Allow-Origin: *
Access-Control-Max-Age: 43200
```

- Access-Control-Allow-Origin: 表示允许的域名设置

服务端有时候需要在响应的头部信息中自定义设置某些值，也需要进行设置。


```
// 类似于这种
X-RateLimit-Limit: 5000
X-RateLimit-Remaining: 4998
```


#### 2.2 路由的设计


网络上的资源，都是通过 URL (Uniform Resource Locator,统一资源定位符) 唯一定位到指定服务器上的资源。路由对应的是操作服务器上的资源，所以 RESTful 风格的路由设计，一律抽象出资源的实体（选择名词）。

具体来说，比如设计一个投票系统，如何设计路由？

抽象出投票的资源实体，定义为：vote 

发起一个投票：

```
POST /v1/api/vote
```

获取所有投票信息：

```
GET /v1/api/votes
```

获取单个投票信息：`{vote_id}` 表示路径参数中的变量

```
GET /v1/api/vote/{vote_id}
```

更新一个投票信息：

```
PATCH /v1/api/vote/{vote_id}
```

删除一个投票信息：

```
DELETE /v1/api/vote/{vote_id}
```

设计这些路由的原则是：清晰、简洁，让使用者看到路由就明确对服务器上的哪种资源做哪些操作。

另外，需求是不断的变更的，随着时间的推移，某些功能可能被废弃、某些功能又被添加，为了考虑兼容和升级需求，一般也会在路由中增加版本信息。

```
/v1/api/*
```

当然也有选择在响应的头部信息中显示版本信息。

```
X-GitHub-Media-Type: github.v3
```


#### 2.3 参数

RESTful 设计中还涉及一个重要的话题，即如何处理参数。根据请求方法的不同又分为不同类型的参数。


**查询参数**

```
GET /v1/api/votes?search=golang&return=all_list&page=1&per_page=10
```

这种依靠`&` 连接的参数，称之为查询参数。查询参数通过解析 url 获取得到。通常将这些获取查询参数进行函数封装，方便调用。

```
// 根据key 获取值
func Query(request *http.Request, key string) string {
	return request.URL.Query().Get(key)
}

// 根据key 获取值，为空则设置默认值
func QueryAndDefault(request *http.Request, key string, defaultValue string) string {
	value := Query(request, key)
	if value == "" {
		return defaultValue
	}
	return value
}

// 获取所有请求参数
func Vars(request *http.Request) map[string]string {
	all := request.URL.Query()
	var results = make(map[string]string)
	for k, i := range all {
		results[k] = i[0]
	}
	return results
}
```

上文的三个函数也是常见的 web 框架的提供查询参数的方法。


**路径参数**

```
GET /v1/api/vote/{vote_id}
```

像上文这种带 `{vote_id}` 的参数称之为路径参数，表示变量，常见的 Web 框架对这个路径参数进行了大量的封装。经常还能看到指定类型或者包含通配符的路径参数。

```
// 指定类型为字符串
GET /v1/api/vote/{name:string} 

// 指定类型为整型
GET /v1/api/vote/{id:int}
```

这种具体的路径参数如何使用，根据不同的 web 框架使用不同。

比如 gin 框架，路径参数不指定类型，而是调用相应方法进行转化，默认是开发者知道路径参数的类型。

```
router.GET("/user/:name/*action", func(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
})

```

比如 echo 框架，做法和 gin 类似。

```
e.GET("/users/:name", func(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
})
```

再比如 iris 框架，就可以指定数据类型和通配符，且支持类型繁多。

```
app.Get("/limitchar/{name:string range(1,200) else 400}", func(ctx iris.Context) {
    name := ctx.Params().Get("name")
    ctx.Writef(`Hello %s | the name should be between 1 and 200 characters length
    otherwise this handler will not be executed`, name)
})

app.Get("/users/{id:uint64}", func(ctx iris.Context) {
    id := ctx.Params().GetUint64Default("id", 0)
    ctx.Writef("User with ID: %d", id)
})
```

原生的 net/http 并不支持这类路径参数，开发者可以自定义数据类型，对路由进行重新组织，使其支持，当然，另外一种做法是不使用路径参数，而直接使用查询参数。


比如获取 vote_id 为 1 的记录，你完全可以将路径参数转化为查询参数。

```
GET /v1/api/vote/1

// 转化为查询参数

GET /v1/api/vote?vote_id=1
```

**JSON 数据**

这类通常使用在 POST、PATCH、PUT 等需要传递参数给网络请求的路由中，而且可以带层级结构。

```
curl -X POST \
  http://localhost:8888/v1/api/vote \
  -d '{
	"data":{
		"title":"Golang",
		"description":"Golang Web"
	}
}'
```

对于这类Json 请求数据，我们一般将其绑定在某一个结构体上，方便后续操作。包括后面会讲到的参数校验。

一般使用下面的函数进行数据绑定。

```
func BindJson(request *http.Request, param interface{}) error {
	err := json.NewDecoder(request.Body).Decode(param)
	if err != nil {
		return err
	}
	return nil
}

```

上文的示例对应的结构体为：

```
type CreateVoteParam struct {
	Data struct {
		Title       string `json:"title" form:"title" validate:"required"`
		Description string `json:"description" form:"description" validate:"required"`
	} `json:"data" validate:"required"`
}
```

这样将请求中的 json 数据绑定在 `CreateVoteParam` 结构体上。

```
func (c ControllerVote) PostOneVote(writer http.ResponseWriter, request *http.Request) {
	var param CreateVoteParam
	if err := make_request.BindJson(request, &param); err != nil {
		log.Println("err: ", err.Error())
		return
	}
	log.Println("Param: ", param)
}
```

#### 2.4 参数校验

针对获取到的参数，不管是请求参数、JSON 数据，再具体的业务处理中一般都还需要进行一步操作：参数校验，意思是参数需要符合相应的业务场景。

举个具体的例子，登录注册是个很常见的功能，比如使用手机号或者邮箱注册，比如业务约束密码使用大于8位、小于16位的不是纯数字字符串作为密码。

就这个例子，可以把参数抽象成结构体：

```
type Register struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

```

业务场景中限定了如下几个约束：

- 手机号：纯数字 11 位
- 邮箱：包含 `@`
- 密码长度最小8，最大为16，且不是纯数字


像这些描述，就需要对获取的参数进行校验的动作。

对参数校验一般的处理方式是：

- 对结构体提供相应的方法，对参数校验
- 使用结构体的标签 Tag 对参数校验
- 两者混合使用


方法一：结构体方法进行校验

```

// 方法一：结构体方法进行校验
type Register struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (r Register) ValidAccount() error {
	// 具体的 r.Account 校验
	return nil
}

func (r Register) ValidPassword() error {
	// 具体的 r.Password 校验
	return nil
}
```

方法二：结构体标签 Tag 进行校验

```
type Register struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"min=8,max=16"`
}

```
只需简单的定义 Tag 就能完成大部分的参数校验动作。这些校验具体的逻辑，是使用到了反射机制，获取到Tag，再进行校验。

具体的可以查看 https://godoc.org/gopkg.in/go-playground/validator.v9


这些参数的校验，一方面和自身的数据类型有关，一方面和具体的业务挂钩。但都有规律可循。

比如，整型的参数校验，一般是大小关系约束：最大、最小等。

比如，数组的参数校验，一般是数组的长度等。

比如，字符串的参数校验，一般是长度、限定选项等。


参数检验的处理一般是在具体的业务之前，如果不符合要求，无需执行后续动作，否则再处理后续的逻辑处理。




#### 2.5 响应信息

RESTful 风格的响应信息，推荐使用 JSON 格式，易读，易理解，解析起来也非常方便。当然响应，有正确的也有错误的响应，错误的信息也需要返回相应的信息。一方面是便于开发者明确这是 Bug, 还只是给用户呈现的报错信息。


推荐的返回格式如下：

包含状态码和具体的信息：

```
{
    "code": 200,
    "data": {
        "id": 1,
        "created_at": "2019-07-14T11:42:33+08:00",
        "updated_at": "2019-07-14T11:42:33+08:00",
        "title": "辞职",
        "admin_id": 1,
        "description": "",
        "choice": [
            {
                "id": 1,
                "created_at": "2019-07-14T11:42:33+08:00",
                "updated_at": "2019-07-14T15:49:22+08:00",
                "vote_title": "辞职",
                "choice_title": "学习",
                "number": 1,
                "ratio": "100.0%"
            },
            {
                "id": 2,
                "created_at": "2019-07-14T11:42:33+08:00",
                "updated_at": "2019-07-14T11:42:33+08:00",
                "vote_title": "辞职",
                "choice_title": "Java",
                "number": 0,
                "ratio": "0.0%"
            }
        ],
        "dead_line": "2020-07-14T11:42:33+08:00",
        "is_anonymous": false,
        "is_single": true
    }
}
```

报错时，显示具体的报错信息和状态码：

```
{
    "code": 400,
    "error": "record not found"
}
```

当然具体的格式可以有所差异，但根据开发经验，一定要包含这些信息：状态码、正确时的响应信息，或者报错时的报错信息。

状态码对应HTTP协议提供的状态码：

- 1XX 相关信息
- 2XX 成功
- 3XX 重定向
- 4XX 客户端错误
- 5XX 服务端错误


状态码具体的个数非常的多，使用起来难以做到非常精准，但是对开发者而已，大方向不能错，比如说，请求成功，一定是返回 2XX 状态码等。状态码只能用来显示请求的状态，所以响应信息中需要包含另一个字段，成功时 `data`, 错误时 `error`。



具体代码层面如下：

```
func Response(w http.ResponseWriter, code int, data interface{}) {
	var results = make(map[string]interface{})
	results["code"] = code
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	if code == http.StatusOK {
		results["data"] = data
	} else {
		results["error"] = data
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "")
	err := enc.Encode(results)
	if err != nil {
		log.Println("err : ", err.Error())
		return
	}

}
```

web 框架比如：echo、gin、iris、mux 等都支持各种类型的响应信息

- String 字符串
- HTML 数据
- XML 数据
- JSON 数据
- Blob 二进制长文件 
- Stream 流文件


针对 RSETful 架构的 API 设计，一般都选择 JSON 数据，方便调用和解析。

### 3. 数据模型

RESTful API 中路由的设计，采用的方法是将具体的资源抽象出实体，比如投票系统，设计了这些 API。

```
// 获取多个投票信息
GET /v1/api/votes

// 获取单个投票信息
GET /v1/api/vote?vote_id=1

// 创建一个投票信息
POST /v1/api/vote

// 更新一个投票信息
PATCH /v1/api/vote?vote_id=1

// 删除一个投票信息
DELETE /v1/api/vote?vote_id=1
```

抽象出了实体：vote， 那么 vote 资源包含哪些具体的信息呢？具体的信息都应该和具体的业务挂钩，而不是凭空想象的。

![](http://ww1.sinaimg.cn/large/741fdb86gy1g50xwyrkrwj20990f0aap.jpg)


假设产品经理和设计人员最终的设计稿如上，那么对开发人员而已，设计的字段可以如下：

```

type Vote struct {
	Title       string // 标题
	AdminId     uint // 用户
	Description string // 描述
	Choice      []Choice // 选项
	DeadLine    time.Time // 截止日期
	IsAnonymous bool // 是否匿名
	IsSingle    bool // 投票单选
}

type Choice struct {
	VoteId uint // 某个投票
	Title  string // 选项标题
	Number int    // 选项投票人数
}
```
这种对资源的设计称之为模型的设计，模型和资源一一对应，不同的领域抽象出的模型各不相同，模型最后都会持久化存储在数据库中，尽管市面上数据库的种类非常的繁多，但是关系型数据库依然是主流。模型设计依然是使用关系型数据库中非常重要的一环，好的模型，利于后续的需求的持续迭代和拓展。

模型是由字段组成，每个字段对应数据库中的一列，字段包含数据类型，每个模型对应数据库中的数据表。

模型的具体设计步骤如下：

- 明确具体需求
- 针对需求抽象出资源实体
- 将资源实体进行字段划分
- 遵守数据库设计三范式，允许适量数据冗余


为什么模型设计非常重要？web 开发中服务器上的资源不可避免的需要持久化存储，项目正常运行上线后，数据量逐渐增加，不可避免的会遇到 SQL 优化查询。这些设计要点，最好在设计之初都考虑进去。

关于数据库，web 开发会接触到哪些知识呢？

- 表设计：表名、字段设计、字段类型、创建、索引设置
- 搜索资源：结构化查询
- 更新资源：更新记录
- 删除资源：删除记录

关系型数据库支持多种数据类型，选择正确的数据类型是获取数据库查询高性能的第一步，选择数据类型一般依据下面几个要点：

- 更小的数据类型：即选择可以正确存储数据的最小数据类型，主要的考量是：第一：符合业务需求，第二：占用更小的磁盘空间
- 简单：简单的数据类型需要的 CPU 调度周期更少，比如：整型操作代价就比字符串低
- 尽量避免 NULL : 比如定义时，决定好字段是否非空，对于索引键，比如主键索引等，不允许为空
- 时间数据类型：关系型数据库关于时间的类型有：DATETIME、TIMESTAMP、YEAR、DATE 等，通常选择TIMESTAMP, 占用更少的内存空间，且支持时区

数据库操作在项目代码层面中一般不会直接写 SQL 语句进行操作，一般选择的方案ORM, 在 Golang 中比较优秀的 ORM 方案有：GORM 和 XORM，两者方案在 web 项目开发中使用非常广泛。

关于数据库操作：在项目中会涉及哪些操作？

- 创建表
- 操作记录

ORM 使这些操作非常便捷。下面分别讲解如何使用 ORM 操作记录。

根据上文的投票系统定义的字段如何创建表：

**GORM 方案**

```
package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Vote struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(32)"`
	AdminId     uint   `json:"admin_id"`
	Description string `json:"description" gorm:"type:varchar(64)"`
	Choice      []Choice
	DeadLine    time.Time
	IsAnonymous bool
	IsSingle    bool
}

type Choice struct {
	gorm.Model
	VoteId uint
	Title  string `gorm:"type:varchar(32)" json:"title"`
	Number int    `gorm:"type:integer(4)" json:"number"`
}

```

- 字段的定义是根据结构体的Tag 来定义的（结构体的Tag 在 Go中使用频繁，比如 ORM 数据库表字段定义、类型定义；比如 JSON 序列化显示的字段；比如用来检验参数有效性等）
- 一个结构体对应一张数据表
- 不设置结构体 Tag时，使用默认的数据类型



**XORM 方案**

```
package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 基础字段
type Base struct {
	Id        int64      `xorm:"pk notnull"`
	CreatedAt time.Time  `xorm:"created" json:"created_at"`
	UpdatedAt time.Time  `xorm:"updated" json:"updated_at"`
	DeletedAt *time.Time `xorm:"deleted" json:"deleted_at"`
}

// tag 定义字段类型
type VoteByXORM struct {
	Base        `xorm:"extends"` // 继承字段
	Title       string `xorm:"varchar(10) notnull" json:"title"`
	AdminId     uint   `xorm:"index" json:"admin_id"`
	Description string `xorm:"varchar(64) default(null)" json:"description"`
	Choice      []Choice
	DeadLine    time.Time `xorm:"timestamp" json:"dead_line"`
	IsAnonymous bool
	IsSingle    bool
}

// 显式定义数据表名称
func (v VoteByXORM) TableName() string {
	return "vote_by_xorm"
}

type ChoiceByXORM struct {
    Base `xorm:"extends"`
	VoteId uint
	Title  string `xorm:"varchar(10) notnull" json:"title"`
	Number int    `xorm:"int(4)" json:"number"`
}


func (v ChoiceByXORM) TableName() string {
	return "choice_by_xorm"
}


```


上文定义了结构体，结合结构体的 Tag 的定义，相当于定义了数据表的结构，但还需要转换到数据库内。

为了更好的理解关系型数据库，需要理解下关系型数据库的一般架构，在明确架构的基础之上理解，需要做什么，为什么这么做。

- 数据库是：客户端、服务端端（C/S）架构
- 本地启动数据库服务（可以下载相应的软件启动，也可以直接使用容器启动数据库服务）
- 客户端通过连接器和服务端连接
- 服务端分析器对 SQL 语句进行语法分析，明确需要执行什么命令，命令有没有错误
- 服务端优化器对 SQL 语句进行优化分析，使用最优的方式执行 SQL 语句
- 服务引擎执行器执行 SQL 语句，将结构返回给客户端


可以看到，使用数据库的第一步是和数据库服务建立连接。GORM 和 XORM  方案的建立连接方式为：

```

package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
)

// gorm 方案
var EngineMySQLGORM *gorm.DB

func EngineGORMInit() {
	db, err := gorm.Open("mysql", "root:admin123@/votes?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("CONNECT DB FAIL: ", err.Error())
		return
	}
	db.LogMode(true)
	EngineMySQLGORM = db
}

// xorm 方案
var EngineMySQLXORM *xorm.Engine

func EngineXORMInit() {
	db, err := xorm.NewEngine("mysql", "root:admin123@/votes?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("CONNECT DB FAIL :", err.Error())
		return
	}
	db.ShowSQL(true)
	EngineMySQLXORM = db
}

```
建立连接需要提供：用户名（root）, 密码（admin123），端口（3306默认），数据库（votes）。建立连接之后，客户端和服务端之间打通，可以执行SQL 语句。

将定义好的结构体模型转换为数据库内的数据表：

```

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter11/pkg/database"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter11/pkg/middleware"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter11/web/model"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter11/web/vote"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

func Execute() {
	err := RootCMD.Execute()
	if err != nil {
		log.Println("FAIL")
		return
	}
}

var RootCMD = &cobra.Command{
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Println("Start Execute Command")
		database.EngineGORMInit()
		database.EngineXORMInit()

	},
	Run: func(cmd *cobra.Command, args []string) {
		MigrateByGORM()
		MigrateByXORM()

	},
	PostRun: func(cmd *cobra.Command, args []string) {
		database.EngineMySQLGORM.Close()
		database.EngineMySQLXORM.Close()
	},
}

func MigrateByGORM() {
	database.EngineMySQLGORM.AutoMigrate(model.Vote{}, model.Choice{})

}

func MigrateByXORM() {
	database.EngineMySQLXORM.CreateTables(model.ChoiceByXORM{}, model.VoteByXORM{})

}
```

- gorm 使用数据库引擎的 AutoMigrate 方法
- xorm 使用数据库引擎的 CreateTables 方法


创建数据库、数据库表之后，可以直接操作结构体对象完成数据库的增删改查动作。



**总结**

ORM 将数据表和结构体对象直接相互映射，开发者只需操作结构体对象，便可以操作数据库表，完成对数据的搜索、增加、修改、删除动作。

数据库模型的定义需要遵循的要点有哪些？

数据库模型定义和需求紧密结合，模型设计需要遵循一些基本的设计要点：
1. 需要有主键，且主键的选择是和业务无关的自增的整型 
2. 字段划分最小化，即每一列不可再细分 
3. 字段的个数不要太多，建议最大上限 20。

操作ORM 完成数据库的操作的一般步骤是：

- 定义结构体字段和Tag 完成数据库表的定义（gorm 或者 xorm 任选其一方案即可）
- 创建数据库连接
- 迁移数据库，即每次更新表结构，需要重新执行下创建表命令
- 使用数据库对象操作，完成业务需求







### 4. 中间件


web 开发过程中，经常需要在请求和响应之间嵌入一些操作，比如日志、认证等，这些操作函数称之为中间件。中间件是“可插拔式”，无需改变原有业务逻辑，不影响你的编码方式。web 开发过程中各框架都提供了一些默认的中间件：主要的分类有：

- 日志中间件
- 认证中间件
- 恢复重启中间件
- 跨域中间件

中间件所处位置的不同又可划分为不同的级别：

- RootLevel 请求前处理路由的一些中间件，比如处理`/`; 请求后处理参数之类的
- GroupLevel 可以使用分组中间件，对一组资源统一的操作
- RouteLevel 对单个路由的中间件

上述三个级别的中间件，作用范围依次减少。
中间件的开发非常简单，但在 web 开发过程中必不可少，原生 net/http 提供了一些默认的中间件，阅读下源码，看如何定义的？

```
// 前缀中间件: 过滤掉特定的前缀

func StripPrefix(prefix string, h Handler) Handler {
	if prefix == "" {
		return h
	}
	return HandlerFunc(func(w ResponseWriter, r *Request) {
		if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
			r2 := new(Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			h.ServeHTTP(w, r2)
		} else {
			NotFound(w, r)
		}
	})
}



// 超时中间件
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler {
	return &timeoutHandler{
		handler: h,
		body:    msg,
		dt:      dt,
	}
}

```

可以看成中间件的开发，只需要返回值是 `http.Handler` 接口即可：

```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

```

比如开发自己定义的日志中间件：

```
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		format := fmt.Sprintf("[ http_log ]: %s | %s | %s | %s", request.URL, request.Host, request.Method, time.Now().Format(time.RFC3339))
		log.Printf("%s", Red(format))
		next.ServeHTTP(writer, request)
	}
}
func Red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}

```
该日志中间件，每次网络请求打印出路由、服务器地址、请求方法、请求时间，在终端中显示红色。

这里为什么可以返回 `http.HandlerFunc`？ 按理说不是应该返回 `http.Handler` 接口吗? 继续查看下 `http.HandlerFunc` 源码的定义。

```
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

自定义的中间件如何使用？直接作用在路由控制器上。

```
func main() {
	http.HandleFunc("/ping", middleware.Logger(func(writer http.ResponseWriter, request *http.Request) {
		var results = make(map[string]interface{})
		results["code"] = http.StatusOK
		results["data"] = "ping"
		writer.Header().Set("Content-type", "application/json;charset=UTF-8")
		enc := json.NewEncoder(writer)
		enc.SetIndent("", "")
		enc.Encode(results)

	}))
	prefix := "/v1/api"

	var v vote.ControllerVote
	http.HandleFunc(fmt.Sprintf("%s/votes", prefix), middleware.Logger(v.GetAllVote))
	http.HandleFunc(fmt.Sprintf("%s/vote", prefix), middleware.Logger(v.VoteHandler))

	//服务启动
	go func() {
		if err := http.ListenAndServe(":8888", nil); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	_, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	log.Println("shutting down")
	os.Exit(0)
    
}
```

```
http http://localhost:8888/ping

HTTP/1.1 200 OK
Content-Length: 27
Content-Type: application/json;charset=UTF-8
Date: Sat, 20 Jul 2019 09:30:06 GMT

{
    "code": 200,
    "data": "ping"
}

// 启动服务终端上显示日志：
2019/07/20 17:30:06 [ http_log ]: /ping | localhost:8888 | GET | 2019-07-20T17:30:06+08:00

```


明确了其基本原理之后，开发者可以根据自己的业务需求，定义相应的中间件。

比如跨域的中间件：

```
func CORS(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(writer, request)
	}
}

```

比如基本的认证中间件：

```
func BasicAuth(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userName := request.Header.Get("username")
		password := request.Header.Get("password")
		if userName != "Go" && len(password) == 0 {
			var results = make(map[string]interface{})
			results["code"] = http.StatusBadRequest
			results["error"] = fmt.Sprintf("Add username and password in requests header")
			if err := json.NewEncoder(writer).Encode(results); err != nil {
				log.Println(err)
			}
			return
		}
		next.ServeHTTP(writer, request)
	}
}

// 请求中需要带 username 和 password, 且指定 username = Go， password 不为空。
```



事实上各流行的 web 框架的中间件的实现方式也大同小异。

**gin 的中间件**

```
func main(){
	app :=gin.New()
	app.Use(gin.Logger()) // 全局日志中间件
	app.Run(":9999")
}
```
gin.Logger 源码

```
func Logger() HandlerFunc {
	return LoggerWithConfig(LoggerConfig{})
}
// LoggerWithConfig instance a Logger middleware with config.
func LoggerWithConfig(conf LoggerConfig) HandlerFunc {
	formatter := conf.Formatter
	if formatter == nil {
		formatter = defaultLogFormatter
	}

	out := conf.Output
	if out == nil {
		out = DefaultWriter
	}

	notlogged := conf.SkipPaths

	isTerm := true

	if w, ok := out.(*os.File); !ok || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd())) {
		isTerm = false
	}

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			param := LogFormatterParams{
				Request: c.Request,
				isTerm:  isTerm,
				Keys:    c.Keys,
			}

			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(ErrorTypePrivate).String()

			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path

			fmt.Fprint(out, formatter(param))
		}
	}
}
```

**echo 中间件**

```
func main(){
    app := echo.New()
	app.Use(middleware.Logger())
	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	app.Logger.Fatal(app.StartServer(s))
    
}
```

middleware.Logger() 源码

```
func Logger() echo.MiddlewareFunc {
	return LoggerWithConfig(DefaultLoggerConfig)
}

type MiddlewareFunc func(HandlerFunc) HandlerFunc

```

**iris 中间件**

```
func main(){
    app := iris.New()
	app.Use(logger.New(logger.DefaultConfig()))
	app.Run(iris.Addr(":8080"))
}
```

logger.New源码

```
func New(cfg ...Config) context.Handler {
	c := DefaultConfig()
	if len(cfg) > 0 {
		c = cfg[0]
	}
	c.buildSkipper()
	l := &requestLoggerMiddleware{config: c}

	return l.ServeHTTP
}
```



**总结**

中间件主要在请求和响应之间，本质上是一个函数。中间件在 web 开发中非常常见，能处理诸多的任务。开发中间件也非常简单，即可以使用默认的中间件，也可以自定义中间件。一般流行的中间件主要是处理这几方面：
1. 跨域 
2. 日志 
3. 认证信息

- 跨域主要是为了解决前端开发的同源策略
- 日志主要是为了更友好的查看网络请求
- 认证信息主要是为了解决对资源的限制问题



### 5. 响应信息

前后端分离的开发模式，前端通过 API 调用服务端的资源，后端负责开发 RESTful API 风格的接口。数据的交换格式一般选择 JSON。 

之前的章节我们讨论过，响应信息，最好包含两个信息特征：

- 状态码
- 具体的信息

正确的响应格式可以如下：

```
{
    "code": 200,
    "data": "message..."
}
```

发生错误的响应格式可以如下：

```
{
    "code": 400,
    "error": "error message..."
}
```

那么正确时的响应具体是哪种格式呢？返回的字段是如何定义的？

**正确时的响应信息：**

一般正确时的响应具体格式及信息和模型定义环节紧密相关。返回的响应信息一般和模型定义内的字段一致，当然也可以根据前端的需求在恰当的变更。

具体的示例更容易理解：

路由：

```
GET /v1/api/votes
```

模型设计：

```
// 投票的模型
type Vote struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(32)"`
	AdminId     uint   `json:"admin_id"`
	Description string `json:"description" gorm:"type:varchar(64)"`
	Choice      []Choice
	DeadLine    time.Time
	IsAnonymous bool
	IsSingle    bool
}
```

一般给模型定义个方法：`Serializer` 

```
// 指定前端需要的字段和显示的格式
type VoteSerializer struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	AdminId     uint      `json:"admin_id"`
	Description string    `json:"description"`
	DeadLine    string    `json:"dead_line"`
	IsAnonymous string    `json:"is_anonymous"`
	IsSingle    string    `json:"is_single"`
}

// 模型的 Serializer 方法
func (v Vote) Serializer() VoteSerializer {

	var isAnonymous = func(key bool) string {
		if key {
			return "匿名"
		}
		return "公开"

	}
	var isSingle = func(key bool) string {
		if key {
			return "单项选择"
		}
		return "多项选择"
	}

	return VoteSerializer{
		ID:          v.ID,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		Title:       v.Title,
		AdminId:     v.AdminId,
		Description: v.Description,
		DeadLine:    v.DeadLine.Format("2006-01-02 15:04:05"),
		IsAnonymous: isAnonymous(v.IsAnonymous),
		IsSingle:    isSingle(v.IsSingle),
	}
}
```

具体的做法是：

- 定义模型的字段：（比如：Vote）
- 定义个同模型类似的结构体，具体的字段和前端或者需求需要的返回值对应：（比如：VoteSerializer）
- 给模型定义个序列化方法，序列化方法返回的值是同模型类型的结构体：（比如：Vote.Serializer 方法返回 VoteSerializer）

这是一种正确响应时的处理方式。当然可以存在其他处理方式，但是还是那条规则，编写易于理解的代码，注意代码的风格的统一。

```
// 响应处理函数
func Response(w http.ResponseWriter, code int, data interface{}) {
	var results = make(map[string]interface{})
	results["code"] = code
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	if code == http.StatusOK {
		results["data"] = data
	} else {
		results["error"] = data
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "")
	err := enc.Encode(results)
	if err != nil {
		log.Println("err : ", err.Error())
		return
	}

}
```





**错误时的响应信息：**


错误时的信息，只看到状态码，并不能知道报错的具体详情，所以规则这类响应格式如下，其中 error 字段可以更丰富点，提供报错的具体信息。

```
{
    "code": 400,
    "error": "error message"
}
```

这类错误响应，只需要实现 Error 接口即可，可以自定义错误类型：

```
type ErrorForVotes struct {
	Code      int    `json:"code"`
	Detail    string `json:"detail"`
	Message   string `json:"message"`
	MessageZh string `json:"message_zh"`
}


func (e ErrorForVotes) Error() string {
	return fmt.Sprintf("Code: %d, Detail: %s, Message: %s, MessageZh: %s",
		e.Code, e.Detail,e.Message,e.MessageZh)
}

```


对于一些频繁使用的错误信息，可以统一处理：命名以 `Error_` 开头，方便识别

```
var (
	ErrorParam  = ErrorForVotes{
		Code: http.StatusBadRequest, 
		Detail: "param fail", 
		Message: "param invalid", 
		MessageZh: "参数校验失效"}
	ErrorInsert = ErrorForVotes{
		Code: http.StatusBadRequest, 
		Detail: "insert data fail", 
		Message: "insert data fail", 
		MessageZh: "记录插入失败"}
)

```


按照上述步骤，即可完成错误信息时的响应处理。

**总结：**

针对 RESTful 风格的响应，一般选择的数据交换格式选择 JSON。API 响应有成功，也有失败的。成功时的响应信息和模型设计的字段几乎一致，开发者只需要定义同名的结构体，并提供方法即可，当然字段可以不同模型一致，具体和需求挂钩。

针对错误时的响应信息，项目自定义符合项目的错误类型，实现 Error 接口，自定义字段，针对常见场景的错误信息，统一处理，统一调用即可。




### 6. 项目组织

根据 web 开发的特性，对项目进行组织，每个开发者对项目的组织各不相同，但依然建议保持风格统一。

```
├── cmd
│   └── root_cmd.go
├── deployments
│   └── keep
├── docs
│   └── keep
├── logs
│   └── keep
├── main.go
├── pkg
│   ├── database
│   │   └── database.go
│   ├── error
│   │   └── error.go
│   ├── log
│   │   └── log.go
│   └── middleware
│       ├── auth.go
│       ├── cors.go
│       └── logger.go
├── scripts
│   └── run.sh
└── web
    ├── make_request
    │   └── request.go
    ├── make_response
    │   └── response.go
    ├── model
    │   ├── vote_gorm.go
    │   └── vote_xorm.go
    └── vote
        ├── assistance.go
        ├── controller.go
        ├── param.go
        └── router.go
```

为什么总是在强调选择一个较好的项目组织方案？其实这些经验来自于热门的项目，热门的 web 开发项目推荐的代码组织方式，是实践证明能够适应多变的需求变化的项目组织方式。


**Makefile**

提供命令，自动化管理和编译项目，有一套语法规范，类似于 shell 脚本。

**cmd**

主要的功能是提供命令行工具，项目中，有时候的需要提供命令行的方式迁移数据表结构，或者导入一些外部的数据，命令行的方式非常适合这种形式。推荐库：	"github.com/spf13/cobra"


**configs**

项目配置文件，主要是一些环节变量或者数据配置信息。

**deployments**

主要的功能是完成镜像构建的文件，web 开发非常适合适应容器对其进行部署。主要使用到的技术是 Docker 容器。

**docs**

主要是提供项目的API 文档，方便和前端或者客户端查看，推荐的方式是 Swagger。

**logs**

主要是提供项目的日志记录，持久化的日志存储目录，日志对实际的项目开发非常重要，后端开发人员接触到的各种技术，数据库 MySQL， Redis 等，可以发现都有日志，不管是报错日志，归档日志等，事实上某些数据的恢复就是靠磁盘上持久化的日志信息，对于 web 开发，持久化的日志一般是一些关键的信息点，方便开发人员遇到问题及时排除问题。

**main.go**

函数主入口。

**pkg**

项目的库文件，包括：数据库连接、中间件、自定义错误类型、日志级别等。供项目中使用。

**scripts**

主要提供一些脚本，项目中经常会使用一些 shell 脚本，完成某些自动化功能。

**web**

实际的业务开发核心内容，主要完成四个任务：1. 模型的定义（当然有些开发者当然把 model 层和 web 同一层级处理） 2. 参数处理，主要是处理路径参数、请求参数、JSON 数据等 3. 响应，提供JSON 化响应 4. 核心的业务开发

核心业务开发，根据笔者的经验又划分为四个文件：

```
└── vote
    ├── assistance.go
    ├── controller.go
    ├── param.go
    └── router.go
```

每个资源起一个文件夹，主要完成对该资源的操作，比如 投票信息的投票信息资源。

- assistance.go 业务代码辅助函数，主要是一些和业务无关的函数处理，当然，如果整个项目都需要的辅助函数，可以抽取出来，和 pkg 同一层级，供项目使用
- controller.go 完成资源的操作，即路由对应的控制器
- param.go 参数处理，对应资源的操作需要的参数，包括参数的字段的定义和有效性的校验
- router.go 主要提供路由和控制器的对应关系


![](http://ww1.sinaimg.cn/large/741fdb86gy1g57hqwhci8j213w0mqwho.jpg)

参考资源：

- https://github.com/golang-standards/project-layout



### 8. 代码管理和托管

作为开发者，如何管理自己的代码，并记录一些提交记录，代码还可以随意的切换回滚等操作？

Git 就是分布式版本控制系统中一种，绝大多数互联网公司都使用 Git 对代码进行管理。著名的程序员社区 GitHub 就是这类代码托管平台的典型代表，同时也是最大的托管网站。许许多多的开源项目都托管在 GitHub 上，比如 Golang 源代码、Python 源代码等。

对于个人开发者而言，如何使用这个巨大的代码托管平台。

- 学习 git 的使用，这几乎是入行程序员的标配，甚至是开发者的“个人名片”
- 维护个人项目


当然和此类相似的托管平台有很多：

- github : 最大的代码托管平台： https://github.com
- gitlab : 支持无限公有、私有项目 https://about.gitlab.com/ 
- Bitbucket : 免费支持5个开发成员的团队创建无限私有代码托管库 https://bitbucket.org/
- 开源中国：http://git.oschina.net/
- coding.net : https://coding.net/home.html
- 码云：https://gitee.com/


你甚至可以在自己的服务器上进行部署，私有化管理自己的或者公司的项目。

日常开发过程中，如何使用 Git 进行开发？

- 维护主分支 master
- 维护开发分支 dev

维护三个分支： master 分支负责线上，稳定运行；新功能在 dev 分支上维护；开发人员各自维护自己的开发分支，待开发完成，合入测试分支，待测试人员测试通过，将开发人员的开发分支合入 master 分支。

日常开发过程，需要不断的开发新功能，建议分支命名具有一定的规范性。提供一些命名规范。

- 功能开发：`operator_feat_date` 即 开发者—功能开发-开发日期
- 修复功能：`operator_fix_date` 即 开发者-修复功能-修复日期

代码提交记录，提供一些提交记录规范：

- 功能开发：`git commit -m "feat: message"`
- 修复功能：`git commit -m  "fix: message"`
- 重构代码：`git commit -m "refactor: message"`



通过使用 Github 托管平台结合其他服务，比如：Travis CI 完成持续集成， Codecov 测试结果可视化。

建议每个开发者都需要维护好自己的 GitHub 账户，维护一些个人的项目。尽管提供了各种托管平台，但仍然建议使用 Github，社区氛围浓厚，代码量多，是个学习的好平台。

### 9. Make 构建工具 


make 是最常用的构建工具，Makefile（makefile 小写也可以） 是一个文本文件，遵守一套语法规范，可以用来对复杂项目进行构建、编译等流程，定义一系列规则指定执行命令，类似于 shell 脚本，通过 Makefile 文件的定义，最好执行 `make command` 便可完成相应的指令。

**1. 语法规范**

```
<target> : <prerequisites> 
[tab]  <commands>

```

- target :  命令
- prerequistirtes : 前置条件，即执行 target 命令执行的动作
- [tab] : 每个命令之前必须有一个 tab 键，当然可以自定义其他的符号
- commands : 具体的执行命令

makefile 文件就是由这样一套语法规则构成，使用的命令即是：`make <target>`。


当然还支持一些其他的语法规范，整体使用起来和 shell 脚本非常像。

注释： `#`
```
# 注释
```
变量：

```
BINARY=go
```
不回显命令:
```
<target>: <prerequisites>
[tab]  @<commands>
```


示例：

- 项目内新建 makefile 文件
- 编写命令

```
PROJECT=go

default:
        go env

version:
        go version

noecho:
        @go version


PHONY: default version noecho

```

makefile 所在目录下执行命令：

- make 什么参数也不带则默认执行第一个，即`make default`
```
make

go env
GOARCH="amd64"
GOBIN="/Users/xiewei/go/bin"
GOCACHE="/Users/xiewei/Library/Caches/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
...
```

- make version 则实质是执行 `go version`

```
make version

go version  // 回显的具体命令语句
go version go1.12.7 darwin/amd64 // 执行命令的结果

```

- make noecho 则实质执行 `go version`，不显示具体的命令

```
make noecho

go version go1.12.7 darwin/amd64
```

- `.PHONY` 表示伪目标，内置的字段，一般前置命令是所有的命令


make 构建工具，使对项目的构建、编译等更为简便，使用在项目中，简化各种执行命令，一旦makefile 文件写好，只需要执行 make 命令。


**2. Go 项目的命令**


makefile 文件在 Github 上的许多开源项目中也存在它的身影，用于项目编译，构建，配合自动化流水线，可以完成一整套的持续集成、持续部署动作。针对 Go 项目，开发者应该提供哪些指令来操作项目呢？

首先 Go 内置的命令行工具，提供哪些命令？

```
Go is a tool for managing Go source code.

Usage:

	go <command> [arguments]

The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         download and install packages and dependencies
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

	buildmode   build modes
	c           calling between Go and C
	cache       build and test caching
	environment environment variables
	filetype    file types
	go.mod      the go.mod file
	gopath      GOPATH environment variable
	gopath-get  legacy GOPATH go get
	goproxy     module proxy protocol
	importpath  import path syntax
	modules     modules, module versions, and more
	module-get  module-aware go get
	packages    package lists and patterns
	testflag    testing flags
	testfunc    testing functions

Use "go help <topic>" for more information about that topic.

```

比较重要且常用的是：

- 测试
- 编译
- 静态检查
- 运行
- 安装下载库
- 格式化代码

所以一个适用于 Go 项目的makefile 文件也应该支持这些动作，具体的根据项目的实际需求而定，当然也可以自定义命令，完成相应的动作。

```
make default 编译
make install 下载安装
make vet 静态检查
make fmt 格式化代码
make clean 移除编译的文件

```


比如：下面这个示例：

```
BINARY="votes"
VERSION=1.0.0
BUILD=`date +%FT%T%z`

PACKAGES=`go list ./... | grep -v /vendor/`
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default:
	@go build -o ${BINARY} -tags=jsoniter

list:
	@echo ${PACKAGES}
	@echo ${VETPACKAGES}
	@echo ${GOFILES}

fmt:
	@gofmt -s -w ${GOFILES}

fmt-check:
	@diff=$$(gofmt -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

install:
	@govendor sync -v

test:
	@go test -cpu=1,2,4 -v -tags integration ./...

vet:
	@go vet $(VETPACKAGES)

docker:
	@docker build . -t wuxiaoxiaoshen/votes:latest -f chapter11/deployments/Dockerfile

clean:
	@if [[ -f ${BINARY} ]] ; then rm ${BINARY} ; fi

.PHONY: default fmt fmt-check install test vet docker clean


```


makefile 不旦方便于开发人员本地编译、构建项目，同时配合自动化流水线，可以做更多的前置任务。

比如在镜像构建环节，在构建之前进行测试等校验，进一步确保镜像构建的正确性。


### 10. 容器化部署

容器是一种新型的虚拟化方式，比传统的虚拟机方式具有诸多的优势。

- 部署方便
- 轻量级
- 相同的运行环境
- 持续交互、持续部署


在容器化技术未出现以前，web 服务的部署方式较为传统，比如将代码拷贝至服务器上，安装相关的依赖，启动服务。有新功能合入代码，再拷贝至服务器，启动服务。传统的方式较为繁琐。容器技术的诞生使得部署更为简便，只需要构建相应的镜像，启动镜像的同时运行相应代码，启动服务即可。


Docker 容器有三大组件：镜像(image)、容器(container) 和 仓库(repository)。要了解容器技术，建议还是查看相关的文档、知识（ https://www.docker.com/ ），这边仅作简单的介绍。

- 镜像：是一个特殊的文件系统，提供容器运行时的程序、库、配置等，属于静态数据
- 容器：实质是进程，和宿主机上进程有所不同，容器进程有属于自己的独立的命名空间，有自己的 root 文件系统、网络配置等
- 仓库：存储各种镜像文件


对后端开发人员而言，需要掌握如何操作镜像、如何操作容器、如何操作仓库。镜像是文件系统，也是一种资源，镜像的操作包括：删除、构建、获取、列表等；容器是进程，容器的操作包括：启动、停止、查看等；仓库是存储镜像的地方，仓库的操作包括：推送、获取等。

Docker 容器采用C/S(客户端/服务端)架构，要使用Docker，需要安装软件，启动Docker 进程，之后可以对镜像、容器、仓库进行操作。由于一些网络访问速度较慢的原因，访问 Docker 官网安装或者获取镜像速度慢，可以使用国内的一些加速器进行加速，这边推荐 Daocloud (https://www.daocloud.io/) ，可以访问这个网站进行下载安装 Docker 并设置Registry (集中存储、分发镜像) (https://www.daocloud.io/mirror)。这样可以方便的获取镜像。后端开发用到的所有技术，比如：ubuntu 系统、Nginx 服务、MySQL、Go、PostgreSQL等，都存在相应的镜像。你需要使用到什么技术，有 Docker 之后，你再也不用费时费力的下载安装各种软件，就可以使用到相应的技术。
比如你需要使用到：PostgreSQL数据库，可以简单的执行下面的命令。

```
// 拉取远程镜像仓库中的 postgres 置本地
docker pull postgres

// 启动容器
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres

// 查看容器
docker ps 

获取得到postgres 容器的ID为: a5d42af1e361

// 进入容器
docker exec -it a5d42af1e361 bash

// 切换用户
su posgres

// 进入数据库
psql -h localhost

postgres@a5d42af1e361:/$ psql -h localhost
psql (11.4 (Debian 11.4-1.pgdg90+1))
Type "help" for help.

postgres=#


```

上文仅以 postgres 为例，简单的使用，事实上各种服务都可以使用 Docker 操作它们，无需在本地安装和繁杂的配置。之后开发者想学习任何服务，优先使用容器版本，快速上手。

获取到镜像之后，怎么使用它们？各种镜像的配置参数不同，如何使用它们？可以查看 DockerHub 最大的镜像托管平台（https://hub.docker.com）。网站上托管了许许多多的官方和个人的镜像，查看相应的文档即可。


对个人开发者而言，如何构建自己的镜像？无需从零开始，在官方的镜像基础之上，构建自己的镜像即可，这也是个人或者企业构建镜像的核心步骤，如何构建镜像？答案是编写 Dockerfile，官方提供了这样一套语法规范，按照规范编写 Dockerfile 文件即可。

```
# 基础镜像
FROM golang:1.13.4
# 维护者信息
LABEL maintainer="XieWei"

# 工作目录
WORKDIR /go/GopherBook/chapter11
# 暴露端口
EXPOSE 8888
# 设置环境变量
ENV GO111MODULE=on
# volume 目录
VOLUME [ "/go/GopherBook/chapter11/logs" ]

# 下载安装依赖
RUN apt-get update && apt-get install -q -y vim git openssh-client cron bash && apt-get clean;
# 拷贝文件
COPY . .
# 执行命令
RUN make install
RUN make prod

# 容器启动时执行命令
CMD [ "bash", "-c", "/go/GopherBook/chapter11/votes;" ]


```

开发者只需要熟悉这样一套语法规范，命令不多，十几个 （）。

```
# 构建镜像

docker build 命令

```

为什么这样操作？构建镜像是为了什么目的？

开发者构建 web 服务，在本地开发的时候，直接本地启动服务即可，比如 go web 项目，执行 go run 命令，启动服务。那么想要在远程服务器上部署这套代码，怎么启动服务？答案是构建镜像，启动镜像的同时，启动服务，这样开发者只需要提供 Dockerfile 文件，就可以构建镜像，运维或者自动的 CI/CD 系统在远程服务器上直接执行构建镜像命令，再执行 docker run 命令启动容器。即可启动 web 服务。这就是为什么容器为什么这么火的原因，一套代码，多处部署非常方便。


个人开发者，如果自己有私有服务器，可以配合 Daocloud(https://www.daocloud.io/) 搭建流水线，自动同步远端代码，进行测试，构建等环节，自动构建镜像，部署容器。




**总结**

不管是个人项目还是公司项目，如何既维持一致的环境，又快速的部署系统？答案是使用 Docker。 容器技术既轻量，又方便快捷。开发者了解了镜像、容器和仓库三大组件，就了解了容器的整个生命周期，生产环境中多使用容器技术进行服务的部署，方便管理和维护。比如修复了功能，只需提交新代码，再次构建新镜像，再次启动容器即可。


### 11. 自动CI/CD


开发人员编写完代码，需要将代码构建成镜像，服务器上拉取镜像，启动容器即可，当然生产环境的容器，并不是单一，容器相互之间又有联系，构建起复杂的系统架构。问题是，每次更新代码，有没有什么方式能够将代码自动构建新的镜像，自动完成部署？整个的字段构建新的镜像、部署新的镜像的过程，称之为CI(持续集成)/CD(持续部署)。有个很明显的优势是，开发人员只需要设置一次自动CI/CD流程，之后只关注在开发功能，完成业务代码上，只需要提交新代码，自动CI/CD 自动拉取新代码，构建新镜像，完成部署，即完成一次功能开发的过程。

对个人开发者而言，Github 是最大的代码托管平台，配合 Github 的代码，个人开发者的开源项目，可以完成自动CI/CD 流程。其中的典型代表就是：TravisCI，当然还存在更多的免费或者收费的持续集成、持续部署工具，具体参考：https://github.com/marketplace，下文仅以 TravisCI 为例讲解整体步骤。

**1. TravisCI**

TravisCI 是持续集成，持续部署中典型代表，使用 GitHub 账户，对开源项目免费，支持绝大多数编程语言，意味着你的绝大多数代码，都可以通过TravisCI 进行持续集成。先阶段对托管在 Github 上的公开和私有的项目都能对接。公开的项目使用：https://travis-ci.org/ 私有的项目使用：https://travis-ci.com/ 。

**2. 使用步骤**

- 使用 GitHub 账号登入网站，会同步你的Github 上托管的代码，想要对哪个项目进行持续集成，勾选即可。

公开项目：
![](http://ww1.sinaimg.cn/large/741fdb86gy1g5a8pa01jbj213y0h7q5e.jpg)

私有项目：
![](http://ww1.sinaimg.cn/large/741fdb86gy1g5a8psf37sj213z0gt40f.jpg)

- 项目中的根目录下需要存在 `.travis.yml` 配置文件，采用 YAML 格式，指定了项目的语言类型，以及相关的执行动作，需要遵守一套语法规范。


**3. 语法规范**

- YAML 语法： 键值对、列表、字符串、数值型

```
language: go
go:
  - "1.12"
  - "1.12.x"
env:
  - GO111MODULE=on

script:
  - echo "Hello Golang"
  - echo "Hi!"
```


Go 版本的 travis.yml 语法规范：https://docs.travis-ci.com/user/languages/go/

**4. 处理流程**

可以看出，配置文件是项目中的一个不可或缺的项目，比如Go web 项目，我们也会使用配置文件进行数据库等配置信息，一般采用 json 格式或者 yaml 格式，可读性高。

镜像构建环节，构建镜像 Dockerfile 也需要编写配置文件，遵守一套 Docker 约定的规范。

Travis 对关联或者对接的项目，首先检测是否存在`.travis.yml` 文件，如果存在，执行：`install` 和 `script` 两个环节，如果不显式的编写这两个环节，Travis 会执行默认的 `install` 和 `script` 环节，`install` 环节用来进行库的安装，对 Go 项目，可以下载依赖库等，可以存在多个命令，`script` 环节执行命令环节，一般用来进行测试、静态检查、编码规范等环节，也可以存在多个命令。

```
install: go get ./...

script:
    - echo "Hello Golang"
    - go fmt $(go list ./... | grep -v vendor)
    - go vet $(go list ./... | grep -v vendor)

```

这两个环节一定会存在，配置文件不存在，会执行默认的。

```
install: go get ${gobuild_args} ./...

script: go test ${gobuild_args} ./...

```

这两个环节又存在相应的钩子环节，即在环节之前或者环节之后执行的动作。

```
before_install：install 阶段之前执行
before_script：script 阶段之前执行
after_failure：script 阶段失败时执行
after_success：script 阶段成功时执行
before_deploy：deploy 步骤之前执行
after_deploy：deploy 步骤之后执行
after_script：script 阶段之后执行

```


**5. Go 项目的.travis.yml**

> 注册 dockerhub (https://hub.docker.com/)，为进行镜像托管

- 代码托管在 Github 上
- 根据是公开还是私密的项目，选择对应的 travis (公开travis-ci.org，私密travis-ci.com )
- 项目跟目录下创建 `.travis.yml`
- 指定编程语言，根据项目执行 `install`、`script` 环节，或者相应的钩子环节
- travis 网站配置 `DOCKER_USERNAME` 和 `DOCKER_PASSWORD` 环境变量
- 项目中 deployments 下编写 Dockerfile 文件
- 执行 docker build 构建镜像
- 所有步骤结束后，将镜像推送至 DockerHub


比如本项目：https://github.com/wuxiaoxiaoshen/GopherBook

对应的 `.travis.yml` 配置文件

```

language: go
go:
  - "1.12"
  - "1.12.x"
env:
  - GO111MODULE=on
notifications:
  email:
    recipients:
      - wuxiaoshen@shu.edu.cn
    on_success: change # default: change
    on_failure: always # default: always


before_script:
  - echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin


script:
  - echo "Hello"
  - go fmt $(go list ./... | grep -v vendor)
  - go vet $(go list ./... | grep -v vendor)
  - make fmt
  - make vet
  - make

after_success:
  - docker build . -t wuxiaoshen/beequick:latest -f Chapter5/BeeQuick.v1/deployments/Dockerfile
  - docker build . -t wuxiaoshen/votes:latest -f chapter11/deployments/Dockerfile
  - docker push wuxiaoshen/beequick:latest
  - docker push wuxiaoshen/votes:latest

```

环境变量配置：
![](http://ww1.sinaimg.cn/large/741fdb86gy1g5a9gyvu72j213x0lzq65.jpg)

提交新代码，travis 会拉取最新代码，执行 travis 内的步骤，且可以看到执行动作日志。

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5a9ikhomij213x0m277w.jpg)


运行成功，镜像推送至 DockerHub，但是镜像是公开的，免费用户暂时只能私密一个镜像。

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5a9kk7wdpj213y0m5tcu.jpg)


构建环节的成功，一定程度上能保证，此次的合入的新代码，没有异常。之后，开发者只需要专注业务即可，使整个过程自动化，简化了流程。TravisCI 的语法约束，我们会把很多的检验，比如测试、静态检测等放在 script 环节，只有环节通过了，才允许此次进行镜像构建，否则失败。

虽然支持各种功能，但仍然给出个人建议：

- `.travis.yml` 文件保持精简，复杂的处理流程，推荐使用 `make` 或者 `shell` 脚本进行处理，进行简化，比如使用 Makefile 对项目进行编译，提供出简单的命令。


TravisCI 在各种开源项目中频繁的使用，对提交代码的进行检测，是否满足覆盖率、测试等。希望对大家有所启发。

真实的商业项目，一般会选择各种云服务，比如阿里云、华为云等各种云厂商提供的服务，这类云厂商都会提供流水线对合入的新代码进行自动化CI，满足条件，自动构建镜像，在服务器上进行镜像替换。当然也可以自己搭建一套符合公司项目的流水线，完成类似的功能，鉴于人力和学习成本，绝大多数中小企业仍然会选择云服务，专注在业务上，另外部署等项目偏运维，后端开发也需要了解整体的流程。

**总结**

个人开源项目的自动CI/CD流程，可以选择 Travis 服务，镜像托管在 DockerHub 上。整体的流程，就是为了自动化处理，尽量做到人力不干预，完成持续集成和持续部署。希望后端开发人员对整体的流程有一定的认识。





### 12. 路线


本章介绍了 Go Web 的各个方面，希望能从整体上把握 Web 开发的整体流程。涉及的概念、技术、知识很多，需要反复的不断自我总结，形成自己的 Web 开发知识体系，在真实的商业项目中，每个环节会有所差异，但大体的思路是一致。

- RESTful 风格的 API 设计是重点
- 模型的设计关系到业务的实现、查询性能等
- 项目的良好组织是为了应对后续多变的需求
- 容器化部署配合自动化 CI/CD 可以快速发布版本，修复功能
