
RESTful 风格的 API 形式的开发，在互联网企业中应用非常广泛，如果你留心观察，绝大多数企业都会使用这种形式构建 web 应用，大型互联网公司的开放平台，也多是使用这种形式。

RESTful 风格的API 本质是对资源的操作，对任一资源可以抽象出四种类型的基本操作，比如构建一个抽奖系统，抽象的资源实体是：抽奖，那么开发人员可能构建出这样的 API：

获取所有抽奖：

```
GET /v1/api/lotteries
```

获取单个抽奖：

```
GET /v1/api/lottery/{lottery_id}
```

创建单个抽奖信息：

```
POST /v1/api/lottery
```

删除单个抽奖信息：

```
DELETE /v1/api/lottery{lottery_id}
```

这种形式的开发主要存在两个大问题：

- 版本管理问题
- API(接口) 增多


随着业务的不断变更，接口为保持向下兼容，需要有版本管理，有些接口可能不适应现在的业务需求，但移除又有可能影响业务，比较稳妥的办法是保留。对新业务、新需求，开发新的接口，每个资源抽象出的四种类型的接口，如果业务很复杂，抽象的资源实体很多，接口也会异常的多。

为解决这个问题，Facebook 于 2015 年开源了一种用于 API 的查询语言：GraphQL。

> GraphQL 既是一种用于 API 的查询语言也是一个满足你数据查询的运行时。 GraphQL 对你的 API 中的数据提供了一套易于理解的完整描述，使得客户端能够准确地获得它需要的数据，而且没有任何冗余，也让 API 更容易地随着时间推移而演进，还能用于构建强大的开发者工具。


简单的说 GraphQL 是一种 API 查询语言，约束了一套关于查询 API 的语法规范，只使用一个路由，能够完成所有的对资源的操作。国外很多公司已经转向 GraphQL，比如：Facebook、Coursera、Github 等，国内广泛使用的并不算多，但 GraphQL 依然具有巨大的潜力，值得投入时间学习。


那么 GraphQL 具有哪些优势呢？

- 无需版本管理
- 类型(schema.graphql)文件即 API 文档
- 极大的便利了前端调用，按需获取
- 减少路由，多个请求合并为一个请求，可减少网络开销
- 可以使用任何编程语言来实现



既然是一种用于 API 的查询语言，那么首先第一点，需要熟悉这套语法规范。


## 0. GraphQL 的使用

RESTful 风格的 API 数据的返回，都是后端处理好，比如哪些字段，而并不在乎调用者是否需要，前端或者客户端调用根据获取到的数据进行业务逻辑的开发，有时候需要调用多个接口，重新组织数据，才能完成任务。使用 GraphQL 不会出现这样的情况，数据的多少，完全跟调用者的查询请求相关，你可以只获取到你需要的字段，也可以在一个路由中获取多个请求的数据，这是 GraphQL 使用形式上最大的不同。


![](http://ww1.sinaimg.cn/large/741fdb86gy1g5qx7msm5bj213y082whk.jpg)

在 GraphQL 中通过预先定义一个 Schema 和声明一些 Type 来完成这些任务，每个 GraphQL 服务都会定义一套类型，用以描述你从服务器上查询到的数据，每当查询时，服务器会根据 schema 验证并执行查询，简单的说 绝大多数 Type 类型，对应的资源实体的序列化结果，即最终接口调用时的响应信息。其中 Query(查询类型)、Mutation(变更类型)、Subscription(订阅类型) 是服务器的操作入口。


### 0.1 类型系统

还是以抽奖系统为例，一个 schema 文件大概是这种样式：

```
type Query {
    # 通过 id  获取抽票信息
    getLottery(id: ID): LotterySerializer
    # 通过 adminId  获取抽奖信息
    getLotteries(adminId: ID): [LotterySerializer]
}

type Mutation {
    # 创建 抽奖信息
    createLottery(name: String, date: String): LotterySerializer
    # 删除 抽奖信息
    deleteLottery(id:ID): LotterySerializer
}

type LotterySerializer {
    # 抽奖信息响应的字段
    id: ID
    adminId: ID
    name: String
    date: String
}

```

- Query 定义查询类型，即对应的资源的查询操作
- Mutatiton 定义变更类型，即对应的资源的创建、修改、删除操作
- 其他类型，对应的资源的响应，该类型由一系列字段组成

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5qymr1g1dj213w0doq4c.jpg)


可以看出整个 schema 文件是由类型系统组成，Query、Mutation、Subscription 是查询入口，其他对应的是资源的响应，响应由字段组成，字段有不同的类型，内置的类型包括：ID(唯一表示，实质是字符串)、Int、Float、String、Boolean 五种类型，这些内置的类型称之为标量类型。

当然，还支持：自定义类型(scalar)、枚举类型(enum)、输入类型(input)、接口类型(interface)、联合类型(union)，各种类型的区分使用关键字区分。

- 自定义类型：实际上自定义类型的场景不多，最终的响应，一般都会处理成 JSON，标量类型足够了

```
scalar Date
```

- 枚举类型：限制在一个可选的集合内，比如抽奖类型，划分为：普通抽奖、现场抽奖、定额抽奖

```
enum LotteryType {
    NORMAL # 普通抽奖
    LUCKY # 现场抽奖
    QUOTA # 定额抽奖
}
```

- 输入类型：输入类型主要是为了解决 Mutation 类型中变量比较复杂的场景，即传入的变量较多，可以抽象出一个输入类型

```
input LotteryData {
    name: String
    date: String
    number: Int
    class: LotteryClass
}

```

- 接口类型：接口是一个抽象类型，包含一些字段，其他类型也具有这些字段，即实现了该接口，接口类型丰富了 GraphQL 具体的查询语句

```
interface LotteryIn {
    name: String
    date: String
}
```

- 联合类型：联合类型和接口十分相似，但是它并不指定类型之间的任何共同字段

```
union SearchResult = Human | Droid | Starship
```

尽管 schema 由类型系统组成，开发者不断在丰富这个文件，定义各种类型系统，但最核心、使用最频繁的依然是这些：

- 内置的标量类型
- 枚举类型
- 输入类型


这个 schema 文件有什么意义？

- 作为后端开发人员的开发指导
- 作为前端或者客户端开发人员的 API 文档

### 0.2 查询语法

假设你已经按照 schema 文件的定义，开发好了接口，那么如何查询呢？接口是供前端或者客户端调用的，理应前端人员应该要更熟悉这类查询语法，当然后端也需要熟悉这类查询语法，毕竟需要接口自测。


为熟悉这套语法，推荐使用 Github的 GraphQL API v4 即第四版的 API，除此之外，还提供了查询 IDE，非常方便，当然你可以通过GitHub GraphQL v4 来学习 schema 的定义，进一步掌握 GraphQL 的查询语法。

- GraphQL v4: https://developer.github.com/v4/
- 查询IDE: https://developer.github.com/v4/explorer/
- 文档：https://developer.github.com/v4/public_schema/

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5r6b4owdfj213w0l00va.jpg)

IDE 包含三个部分：

- 请求
- 变量
- 结果



关于查询语法，需要知道下面的内容，主要围绕 Github GraphQL v4 的IDE 进行示例的学习。

- 字段
- 参数
- 变量
- 操作名称
- 指令
- 片段
- 元字段


首先从 Github GraphQL 文档内抽取出部分类型定义：

```
type Query {
    repository(
    """
    The name of the repository
    """
    name: String!

    """
    The login field of a user or organization
    """
    owner: String!
  ): Repository
  
}

```

上面这个类型文件的含义是：

查询某个仓库的信息，需要输入某个仓库的名称（name，且不能为空）和所有者（owner ，且不能为空），输出结果是个 Repository 类型，该 Repository 内定义了各种字段。


#### 0.2.1 查询字段、参数


Query 类型的定义，左边是：操作名称，包含请求参数的名称和类型；右边是响应类型，包含一系列字段的名称和类型。

查询语句：

```

query {
	repository(name:"go", owner:"golang"){
    id
    name
    createdAt
    url
  }
}
```

- 输出请求参数：name: "go", owner:"golang"
- 输出字段：id、name、createdAt、url

查询结果：

```
{
  "data": {
    "repository": {
      "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
      "name": "go",
      "createdAt": "2014-08-19T04:33:40Z",
      "url": "https://github.com/golang/go"
    }
  }
}
```
可以看出这种查询方法需要的字段个数，由用户指定输出字段，而不像 REASTful 风格的，一股脑的全部给出。

#### 0.2.2 操作名称、变量

查询语句中包含操作名称变量：

```
query GetRepository($NAME: String!, $OWNER: String!){
	repository(name:$NAME, owner:$OWNER){
    id
    name
    createdAt
    url
  }
}

# 变量

{
  "NAME":"go",
  "OWNER": "golang"
}
```


结果：

```
{
  "data": {
    "repository": {
      "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
      "name": "go",
      "createdAt": "2014-08-19T04:33:40Z",
      "url": "https://github.com/golang/go"
    }
  }
}
```

其实是同一个查询，只不过赋予了这个查询一个操作名称：GetRepository， 同时将真实的请求参数，以变量的形式传入，变量以 $ 标示，同样需指明类型。

变量以 JSON 格式传入：

```
{
  "NAME":"go",
  "OWNER": "golang"
}

```

查询动作首先会检查是否符合语法规范，如果不符合语法规范，会给出详细的报错信息，方便开发者排查问题。



#### 0.2.3 指令

指令在 GraphQL 规范中指按照条件动态的显示字段，常用的指令主要包含两个：@include(if: Boolean) 参数为 true 显示字段, @skip(if: Boolean) 参数为 true 跳过字段。

查询语句：

```
query GetRepository($NAME: String!, $OWNER: String!, $INCLUDE: Boolean!, $SKIP: Boolean!){
	repository(name:$NAME, owner:$OWNER){
    id
    name
    createdAt @include(if: $INCLUDE)
    url @skip(if: $SKIP)
  }
}

# 变量

{
  "NAME":"go",
  "OWNER": "golang",
  "INCLUDE": false,
  "SKIP": false
}
```

查询结果：

```

{
  "data": {
    "repository": {
      "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
      "name": "go",
      "url": "https://github.com/golang/go"
    }
  }
}
```

可以看出最终的结果中根据 INCLUDE, SKIP 两个值的布尔类型决定是否显示。需要指出的是这类指令，并不需要开发人员在处理逻辑上进行操作，GraphQL 语法本身就支持。

#### 0.2.4 片段

片段指将可复用的字段抽取出来，需要使用的时候再引入，简化查询, 片段内仍然可以使用变量。

查询语句：

```
query GetRepository($NAME: String!, $OWNER: String!, $INCLUDE: Boolean!, $SKIP:Boolean!){
	repository(name:$NAME, owner:$OWNER){

    createdAt @include(if: $INCLUDE)
    ...commonField

  }
}


fragment commonField on Repository{
  id
  name
  url @skip(if: $SKIP)
}

# 变量

{
  "NAME":"go",
  "OWNER": "golang",
  "INCLUDE": true,
  "SKIP": true
}
```

查询结果：

```
{
  "data": {
    "repository": {
      "createdAt": "2014-08-19T04:33:40Z",
      "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
      "name": "go"
    }
  }
}

```

需要指出的是片段，依然不需要开发人员在处理逻辑上进行任何操作，本身 GraphQL 语法就支持。

#### 0.2.5 元字段

GraphQL 是一种基于类型语法的 API 查询语法规范，查询中如何知道当前操作的具体类型是哪种？关键字：`__typename`，在查询的任何位置皆可使用。

查询：

```

query GetRepository($NAME: String!, $OWNER: String!, $INCLUDE: Boolean!, $SKIP:Boolean!){
	repository(name:$NAME, owner:$OWNER){

    createdAt @include(if: $INCLUDE)
    ...commonField

  }
}


fragment commonField on Repository{
  id
  name
  url @skip(if: $SKIP)
  __typename 
}

# 变量

{
  "NAME":"go",
  "OWNER": "golang",
  "INCLUDE": true,
  "SKIP": true
}
```

结果：

```
{
  "data": {
    "repository": {
      "createdAt": "2014-08-19T04:33:40Z",
      "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
      "name": "go",
      "__typename": "Repository"
    }
  }
}
```

#### 0.2.6 总结

上文以示例的形式学习了 GraphQL 的语法规范，事实上还有许多规范，这边仅做简单的说明。主要围绕在 Query 类型的查询上，包括：字段、变量、操作名称、片段、元字段、指令等。


通过使用 Github 的 GraphQL v4 示例，发现一些定义 schema 的规范：

- Query, Mutation, Subscription 类型固定，且有且只有一个
- 驼峰式命名字段名称，而不使用下划线式
- 类型的定义首字母大些
- 字段的名称，首字母小写
- Query 类型，操作名称尽量使用名称形式
- Mutation 类型，操作名称尽量使用动词形式
- 枚举类型，字段：使用大写




给大家的启示是：通过示例式，快速掌握常用的语法规范。

- GraphQL 中文网：https://graphql.cn/
- GraphQL 官网：https://graphql.org/

## 1. GraphQL 客户端：graphql-go

GraphQL 是一种语法规范，具体的实现并不限定编程语言，本书探讨的是 go, 所以先学习下 go 版本的  graphql-go 如何实现的。

### 1.1 下载安装

```
go get github.com/graphql-go/graphql
go get github.com/graphql-go/handler
```

### 1.2 原理

各种编程语言实现的客户端语法有差异，但核心思想是，实现 GraphQL ，一种用于 API 的查询规范。涉及的知识点有：

- schema 类型集合：其中 Query、Mutation、Subscription 是入口
- 定义各种类型，类型对应的是响应，类型包含字段，字段有类型，GraphQL 内置ID、Int、Float、String、Boolean 五种类型
- 输入类型、枚举类型、接口类型、联合类型
- 变量



编程语言实现，即也是实现这些内容，通过查看 graphql-go 文档，得出：

```
type Schema
type SchemaConfig


type Object
type ObjectConfig

type Fields
type Field
type FieldConfigArgument
type FieldArgument


type Scalar
type ScalarConfig


type Union
type UnionConfig


type Input
type InputObject
type InputObjectConfigFieldMap
type InputObjectFieldConfig


type Enum
type EnumConfig
type EnumValueConfigMap
type EnumValueConfig



```


之所以列出这些 graphql-go 的 API 是想让读者知道，graphql-go 库的核心思想是在定义各种类型对象，最后构造成 schema 文件。


### 1.3 使用

正式编码前，需要定一个 schema 文件，这个文件作为具体的开发的指导，即需要开发哪些内容，定义了哪些类型，哪些响应。

简单使用，这个预期实现一个“心跳” 程序，判断程序是否健康状态，给定任一参数，能将该参数返回。


**schema 文件描述**

```
type Query {
    ping(data: String): ResponsePing 
}


type ResponsePing {
    data: String
    code: Int
}

```

查询示例：

```
query{
  ping(data:"pong"){
    code
    data
  }
}

```

查询结果示例：

```
{
  "data": {
    "ping": {
      "code": "200",
      "data": "pong"
    }
  }
}
```


代码实现：

**定义Query 类型**

```
var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"ping": &graphql.Field{
			Type: Ping,
			Args: graphql.FieldConfigArgument{
				"data": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				ch := make(chan Result, 1)
				var result Result
				go func() {
					defer close(ch)
					result.data = ResponsePing{
						Data: p.Args["data"].(string),
						Code: http.StatusOK,
					}
					ch <- result
				}()
				return func() (interface{}, error) {
					r := <-ch
					return r.data, r.error
				}, nil
			},
		},
	},
})

```

根据 schema 描述，Query 类型包含的字段是 ping，请求参数是 data 类型是 string, 返回的响应是类型 ResponsePing(即代码中的 Ping).

- "ping" 表示字段名称
- Type 定义响应类型
- Args 定义请求参数
- Resolve 是核心的处理逻辑，即如何处理参数，构造响应


**定义 ResponsePing 类型，即代码中的 Ping**

```
type ResponsePing struct {
	Data string `json:"data"`
	Code int    `json:"code"`
}

var Ping = graphql.NewObject(graphql.ObjectConfig{
	Name: "ping",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				ch := make(chan Result, 1)
				go func() {
					defer close(ch)
					if source, ok := p.Source.(ResponsePing); ok {
						ch <- Result{
							data:  source.Data,
							error: nil,
						}
					}
				}()
				return func() (interface{}, error) {
					r := <-ch
					return r.data, r.error
				}, nil

			},
		},
		"code": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				ch := make(chan Result, 1)
				go func() {
					defer close(ch)
					if source, ok := p.Source.(ResponsePing); ok {
						ch <- Result{
							data:  source.Code,
							error: nil,
						}
					}
				}()
				return func() (interface{}, error) {
					r := <-ch
					return r.data, r.error
				}, nil
			},
		},
	},
})

type Result struct {
	data interface{}
	error
}
```

根据 schema 文件描述，Ping(ResponsePing) 类型包含两个字段：data 类型是 String, code 类型是 Int。

- "data", "code" 即字段名称
- Type 即定义字段的类型
- Resolve 即核心的处理逻辑，根据参数，进行处理，返回响应


**构造成 schema**


```
func RegisterSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: Query,
	})
}

```

- schema 的入口是 Query, Mutation 类型


**定义路由，启动服务**

```
func RegisterHandler() *handler.Handler {
	schema, err := RegisterSchema()
	if err != nil {
		log.Println(err)
		return nil
	}
	return handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
}
func main() {
	h := RegisterHandler()
	http.Handle("/graphql", h)
	log.Fatalln(http.ListenAndServe(":9876", nil))
}

```


一旦你理解了 schema 文件是如何定义的，GraphQL 是基于类型的查询API的规范，能够使用各种编程语言的客户端定义 schema , 启动 web 服务，只需要一个路由 `/graphql` 即可完成所有的查询动作。（路由 `/graphql` 路由一般约束如此）


启动服务后，如何测试接口？

这边推荐两个工具，一个是 Postman（专门用于API 测试的工具），一个是 GraphiQL(https://github.com/graphql/graphiql), 两个都很好用，推荐使用 GraphiQL，毕竟是专门用来测试 GraphQL 接口的。


![](http://ww1.sinaimg.cn/large/741fdb86gy1g5rhhhzumpj213v0haabm.jpg)


**总结**

使用 graphql-go 进行开发的核心任务是在处理：定义 schema 文件中定义的类型，因为需要对每个字段进行逻辑处理，整体上看来，如果一个类型字段多，代码量很多，开发者可以合理规范项目组织结构进行优化，使其逻辑更清晰。


一般的 graphql api 的开发步骤如下：

- 定义 schema 文件：作为后端开发人员的开发指导，其实整个 schema 文件的定义，是一个梳理业务逻辑，处理哪些资源，如何组织响应的过程
- 定义 schema 文件中约束的对象类型（graphql.Object），字段类型(graphql.Field)声明，参数(graphl.FieldConfigArgument)处理，逻辑处理(Resolve)。
- 定义 schema 对象(graphql.Schema)
- 定义路由(一般统一定义为 /graphql), 启动服务
- 接口测试：使用 GraphiQL 进行查询请求



## 2. Go GraphQL 构建第十八届国际泳联世锦赛 API


上文的示例比较简单，复杂度不够，很多的用法没使用到，为提升整体的难度，使其符合企业级的项目，笔者对第十八届的国际泳联世锦赛的一些比赛数据进行构建 API。


参照之前的 RESTful 的风格开发指导思想，一般的开发步骤是：

- 需求分析，即明确你需要做些什么，完成什么任务
- 定义模型，方便持续化存储和操作
- 进行项目结构组织
- 项目开发
- 接口测试


尽管 Graphql 的思想和 RESTful 风格的完全不一样，但整体的开发思想依然适用，即依然需要进行需求分析，模型设计，项目组织等步骤，差异在于 RESTful 需要对资源进行抽象，每种资源大概会出现四种类型的接口。而 GraphQL 只需要使用一个接口，定义 schema 作为开发指导，定义schema 约束的类型，完成开发任务。


### 2.1 需求分析


开发的对象是第十八届国际泳联世锦赛，自然而然的用户更关心的主要围绕的是：

- 第 18 届国际泳联世锦赛的标语、含义、吉祥物、场馆、举办时间
- FINA 的历史，总部，设立时间，成员国，赛事项目
- 各届 FINA 比赛的介绍
- 比赛的种类：游泳、跳水、高空跳水等，细分的比赛种类，比如男子 400米自由泳
- 奖牌的分布，比如游泳总共多少奖牌
- 最终各国的奖牌分布，历史上各国的奖牌分布及排行
- 世界记录的保持者及世界记录
- ...


这些内容，如果你不是特意的去关心，其实你只会关注最终的奖牌数，那些多出来的信息，其实就是需求分析的结果，想要构建一个较为丰富且稍微复杂的项目，牵扯的内容不能过少。那么核心数据内容来自于哪呢？

收集信息的渠道有很多，最正确的途径是访问官方网站，进行信息获取，信息多且杂的话，手动操作，那不太符合程序员这行，所以，在不触犯当地法律的情况下，可以使用爬虫对数据进行收集。

爬虫技术对信息的收集主要的步骤是：

- 访问目标网站，借助 chrome 浏览器对网页进行分析，比如数据是嵌入在 html 代码中，还是直接访问接口即可得到，这一步主要熟悉目标网站中资源都在哪些网页中
- 根据你需要得到的数据，进行模型的设计，数据需要持久化，需要数据库操作，模型的设计需要遵循一定的规范
- 进行网络请求，对获取到的网页源代码进行解析，如果数据嵌入在 html 中，借助 css 选择器或者 xpath 操作 dom 树，抽取需要的资源，如果数据是 JSON 格式，直接解析 JSON 数据即可
- 对获取的数据，进行入库操作
- 后续操作


对于网上的资源，一类是 UGC 即用户自生产，比如知乎社区，所有的内容，都是用户产生的。另一类是聚合类，即不生产内容，通过聚合别人已经生产的资源，这类聚合类的主要手段是通过爬虫技术。

爬虫获取数据的主要步骤，如上文所示。按照这样的步骤，对目标网站：（FINA 中文网：https://www.fina-gwangju2019.com/chn/） 进行数据的抓取。

> 鉴于一些读者可能对某些网站访问速度不佳，后续会提供抓取的数据源文件


### 2.2 网页分析，模型定义


> 模型的定义跟实际的需求挂钩，不同的开发者的字段定义稍有差异。

**FINA 介绍**

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5rl00zz34j213v0lywla.jpg)

模型定义：

```
type FiNa struct {
	Base                `xorm:"extends"`
	Description         string `json:"description"`
	Established         string `xorm:"varchar(32) notnull 'established'" json:"established"`
	Headquarters        string `xorm:"varchar(32) notnull 'headquarters'" json:"headquarters"`
	NationalMember      string `xorm:"varchar(24) notnull 'national_member'" json:"national_member"`
	NumberOfDisciplines string `xorm:"varchar(24) notnull 'number_of_disciplines'" json:"number_of_disciplines"`
}

func (F FiNa) TableName() string {
	return "fina"
}
```

主要获取的字段：简介、成立时间、总部、成员国、项目数。


**FINA 历史**

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5rl27cdhoj213x0m179j.jpg)

模型定义：

```
type FiNaHistory struct {
	Base   `xorm:"extends"`
	Year   int    `json:"year"`
	Detail string `json:"detail"`
}

func (F FiNaHistory) TableName() string {
	return "history"
}

```

主要获取的字段是：年份、介绍

**赛事项目**

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5rlf1fxdlj213z0llwje.jpg)


模型定义：

```
type Sports struct {
	Base           `xorm:"extends"`
	Total          int     `xorm:"integer(3) 'total'" json:"total"`
	SportClass     int     `xorm:"'sport_class'"`
	SportName      string  `xorm:"'sport_name'" json:"sport_name"`
	Description    string  `json:"description"`
	CompetitionIds []int64 `xorm:"'competition_ids'" json:"competition_ids"`
	Rule           string  `xorm:"varchar(1024) 'rule'" json:"rule"`
}

func (S Sports) TableName() string { return "sports" }

type Competitions struct {
	Base             `xorm:"extends"`
	CompetitionClass int    `xorm:"'competition_class'" json:"competition_class"`
	Detail           string `xorm:"'detail'" json:"detail"`
}

func (C Competitions) TableName() string { return "competitions" }


```

主要获取的字段是：介绍、金牌总数、具体项目、规则。项目总共有6大类，即 SportClass 可以定义为枚举类型，比赛类型又分为男、女、团队，即CompetitionClass 也可以定义为枚举类型。

**奖牌结果**

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5rlkazbzuj213y0lhq6b.jpg)

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5rlmsy3b9j213x0ld77d.jpg)

模型定义：

```
# 国际奖牌
type CountryMedal struct {
	Base      `xorm:"extends"`
	Year      int   `json:"year"`
	CountryId int64 `xorm:"index 'country_id'" json:"country_id"`
	Gold      int   `xorm:"'gold'" json:"gold"`
	Silver    int   `xorm:"'silver'" json:"silver"`
	Bronze    int   `xorm:"'bronze'" json:"bronze"`
}

func (CC CountryMedal) TableName() string { return "medal" }


# 国家
type Country struct {
	Base  `xorm:"extends"`
	Name  string `xorm:"unique" json:"name"`
	Short string `xorm:"unique" json:"short"`
}

func (Cry Country) TableName() string {
	return "country"
}


```

主要包括的字段：年份、国际、金牌、银牌、铜牌, 可以看到历届奖牌的数据并不嵌入在 html 标签内，而是通过 API 返回的 JSON 数据内。


**世界记录**

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5rlqqt6fij213u0m2jvx.jpg)


模型定义：

```
type RecordMax struct {
	Base             `xorm:"extends"`
	EventName        string    `xorm:"varchar(32)" json:"event_name"`
	Record           string    `xorm:"varchar(32)" json:"record"`
	CountryId        int64     `xorm:"integer(3)" json:"country_id"`
	Date             time.Time `json:"date"`
	Location         string    `json:"location"`
	CompetitionClass int       `json:"competition_class"`
	SportClass       int       `json:"sport_class"`
	Name             string    `json:"name"`
}

func (R RecordMax) TableName() string { return "records" }

```

主要获取的字段：项目名称、国家、时间、地点、比赛项目、名称、记录


**赛事介绍**

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5rlvkn8lmj213v0len2q.jpg)

模型定义：

```
type FiFaChampionships struct {
	Base           `xorm:"extends"`
	NumberOlympic  int       `xorm:"integer(4) 'number_olympic'" json:"number_olympic"`
	ShortSlogan    string    `xorm:"text 'short_slogan'"`
	StartDate      time.Time `xorm:"datetime notnull 'start_date'" json:"start_date"`
	EndDate        time.Time `xorm:"datetime notnull 'end_date'" json:"end_date"`
	DisciplinesIds []int64   `xorm:"blob 'disciplines_ids'" json:"disciplines_ids"`
	VenuesIds      []int64   `xorm:"blob 'venus_ids'" json:"venus_ids"`
}

func (F FiFaChampionships) TableName() string {
	return "championships"
}

type Kinds struct {
	Base  `xorm:"extends"`
	Name  string `xorm:"varchar(32) 'name'" json:"name"`
	Class int    `xorm:"integer(1)" json:"class"`
}

func (K Kinds) TableName() string {
	return "kinds"
}

```

主要获取的第18届世锦赛的举办时间，地点，项目、场馆


**赛事象征**

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5rlwaf1m5j213w0l9tdj.jpg)


模型定义：

```
type Symbol struct {
	Base                    `xorm:"extends"`
	SymbolText              string  `xorm:"varchar(64) 'symbol_text'" json:"symbol_text"` // 标志
	SymbolTextImage         string  `xorm:"varchar(128) 'symbol_text_image'" json:"symbol_text_image"`
	SymbolTextShort         string  `xorm:"varchar(12) 'symbol_text_short'" json:"symbol_text_short"`
	SymbolDescription       string  `xorm:"varchar(64) 'symbol_description'" json:"symbol_description"` // 标语
	SymbolDescriptionImage  string  `xorm:"varchar(128) 'symbol_description_image'" json:"symbol_description_image"`
	SymbolDescriptionShort  string  `xorm:"varchar(12) 'symbol_description_short'" json:"symbol_description_short"`
	SymbolAnimalImage       string  `xorm:"varchar(128) 'symbol_animal_image'" json:"symbol_animal_image"` // 吉祥物
	SymbolAnimalDescription string  `xorm:"varchar(64) 'symbol_animal_description'" json:"symbol_animal_description"`
	SymbolAnimalShort       string  `xorm:"varchar(12) 'symbol_animal_short'" json:"symbol_animal_short"`
	BlueVersions            []int64 `xorm:"'blue_versions'" json:"blue_versions"`
}

func (S Symbol) TableName() string { return "symbol" }

type Blue struct {
	Base        `xorm:"extends"`
	Short       string `xorm:"varchar(1)"json:"short"`
	EnName      string `xorm:"varchar(12) 'en_name'"json:"en_name"`
	ChName      string `xorm:"varchar(32) 'ch_name'"json:"ch_name"`
	Description string `xorm:"varchar(64) 'description'" json:"description"`
	Image       string `xorm:"varchar(128) 'image'" json:"image"`
}

func (B Blue) TableName() string { return "blue" }


```

主要获取的字段是：标语、吉祥物、标语含义等


模型的设计是项目的基础，一方面决定了数据存储的格式，另一方面模型的设计也几乎决定了接口的响应格式，笔者的做法是对每个模型结构体，设计一个 Serializer 方法，该方法的主要目的是对模型进行序列化的操作，即接口的响应字段和格式。

比如 blue 是世锦赛中蓝图的模型，接口的响应，笔者习惯这么操作。

```
# 定义一个序列化结构体
type BlueSerializer struct {
	Id          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Short       string    `json:"short"`
	EnName      string    `json:"en_name"`
	ChName      string    `json:"ch_name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
}

# 模型的 Serializer 方法返回这个序列化结构体

func (B Blue) Serializer() BlueSerializer {
	return BlueSerializer{
		Id:          B.Id,
		CreatedAt:   B.CreatedAt.Truncate(time.Second),
		UpdatedAt:   B.UpdatedAt.Truncate(time.Second),
		Short:       B.Short,
		EnName:      B.EnName,
		ChName:      B.ChName,
		Description: B.Description,
		Image:       B.Image,
	}
}

```

具体的接口的响应信息的处理方式有很多，像笔者的这种处理方式也并不是统一标准，仅从易于理解的角度考量。

对于目标网站的数据的获取，同样因业务和个人的理解而各不相同，笔者认为上文的网页分析中的数据对项目有帮助，对用户有帮助。

### 2.3 数据获取

对爬虫而言，获取数据的一般步骤是：

- 分析网页，定位所需要的资源的位置，是嵌入 html 内，还是JSON 数据内，两者采用的方法不同
- 发起网络请求，解析数据
- 获取数据入数据库，等待后续操作使用


如果目标网站有客户端应用(APP), 借助网络代理，可以优先对客户端进行分析，因为一般的客户端都是调用 API 接口，而且数据格式一般都是 JSON，所以整体分析起来比较容易，当然热门的应用数据获取还是有难度的，大多数APP 都会对数据进行加密。APP 端解决不了，考虑web 端，像这类官方网站的数据，一般都是没有太多的限制，意外着你只要借助 chrome 浏览器的调试功能，绝大多少信息都能分析出来。

网页源代码的解析一般的出来方式是这样的：

- 使用正则表达式：这种对正则表达式的语法要求比较高
- 使用 css 选择器：如果你用过 jQuery，对 css 选择器应该不陌生，能操作 dom 树，定位到资源
- 使用 xpath 路径表达式：能操作 dom 树，定位到资源
- 如果是 JSON 数据，直接 JSNON 解析即可


如果你使用 css 选择器，最佳的调试方式是在浏览器中 Console 中进行调试。

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5s2mweayoj213v0lh4fo.jpg)

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5s2l69gtuj213v0m34bc.jpg)

结合网页分析，定位出资源的位置：`".content.content_wide div div[class=text]"`


如果你使用 xpath 路径表达式，推荐你使用一个 chrome 插件：xpath helper。

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5s2pdky85j213v0m47n5.jpg)

结合网页分析，定位出资源的位置：`//div[@class="content content_wide"]//div[@class="text"]`


两者的语法不尽相同，殊途同归，都完成了目的，至于读者喜好哪个，就使用哪个，当然笔者建议都需要学会，学习成本也不是很高。

- CSS 选择器语法：https://www.w3school.com.cn/cssref/css_selectors.asp
- XPATH 路径表达式语法：https://www.w3school.com.cn/xpath/xpath_syntax.asp


> 当然文本数据获取到了，还需要进行进一步操作，比如你将数据分割，比如你只截取部分内容，比如你需要替换部分内容等...


**相关库下载**

- css 选择器

```
go get github.com/PuerkitoBio/goquery
```

- xpath 路径表达式

```
go get github.com/antchfx/htmlquery
```

- json 解析

```
go get -u github.com/tidwall/gjson
```


**网页基本组成**

学习使用 css 选择器 和 xpath 路径表达式，有必要了解下 html 网页的组成，明确了组成，知道这些层级结构，读者更能明白这些语法的含义，比如属性、子节点、兄弟节点、文本等。

```
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>测试页面</title>
  </head>
  <body>
    <img src="images/firefox-icon.png" alt="测试图片">
  </body>
</html>
```

- <!DOCTYPE html> — 文档类型。
- <html></html> — <html> 元素。这个元素包含了整个页面的内容，有时也被称作根元素。
- <head></head> — <head> 元素。这个元素放置的内容不是展现给用户的，而是包含例如面向搜索引擎的搜索关键字（keywords）、页面描述、CSS 样式表和字符编码声明等。
- <meta charset="utf-8"> — 这个元素指定了当前文档使用 UTF-8 字符编码 ，UTF-8 包括绝大多数人类已知语言的字符。基本上 UTF-8 可以处理任何文本内容，还可以避免以后出现某些问题，我们没有任何理由再选用其他编码。
- <title></title> — <title> 元素。这个元素设置页面的标题，显示在浏览器标签页上，同时作为收藏网页的描述文字。
- <body></body> — <body> 元素。这个元素包含期望让用户在访问页面时看到的内容，可以是文本、图像、视频、游戏、可播放的音轨或其他内容。



![](http://ww1.sinaimg.cn/large/741fdb86gy1g5tgyp0ckmj20ta04cjrz.jpg)


- `<p>` 表示开始标签
- `</p>` 表示结束标签
- `class` 表示属性字段，`editor-note` 表示属性值
- "我的猫咪脾气爆:)" 表示文本内容
- 整个表示：元素


整个的网页解析的过程就是不断的对元素内属性、属性值或者文本解析的过程。下面分本对 html 文件和 json 文件进行解析，为讲解方便，只截取真实数据中的一部分进行说明：

**html 部分文件： content.html**

```
<div class="v_list v_list01">
    <div class="img"><img src="/home/chn/images/sub/vision01.jpg" alt="5.18 기념공원 사진"></div>
    <div class="text xh-highlight">
        <div class=""><span class="point">P</span><span class="title"><em>People</em>共同生活的人类</span></div>
        <ul>
            <li>在民主·人权城市提高人道主义价值</li>
            <li>寄予人类和平共存</li>
        </ul>
    </div>
</div>
```

假设想要获取的内容是：`<li>` 元素内的文本内容：即："在民主·人权城市提高人道主义价值..."




**CSS 选择器**


css 选择器语法则为：`div div ul`

```
func ParseByCss(reader io.Reader) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(doc.Find("div div ul").Text())
}

# 测试

func TestParseCss(test *testing.T) {
	file, err := os.Open("content.html")
	if err != nil {
		log.Println(err)
		return

	}
	fina.ParseByCss(file)

}

# 结果

=== RUN   TestParseCss

            在民主·人权城市提高人道主义价值
            寄予人类和平共存
        
--- PASS: TestParseCss (0.00s)
PASS

```


**xpath 路径表达式**

xpath 路径表达式为：`//div/div/ul`

```
func ParseByXpath(reader io.Reader) {
	doc, err := htmlquery.Parse(reader)
	if err != nil {
		log.Println(err)
		return
	}
	text := htmlquery.FindOne(doc, "//div/div/ul")
	fmt.Println(htmlquery.InnerText(text))
}

# 测试

func TestParseXpath(test *testing.T) {
	file, err := os.Open("content.html")
	if err != nil {
		log.Println(err)
		return

	}
	fina.ParseByXpath(file)
}

# 结果

=== RUN   TestParseXpath

            在民主·人权城市提高人道主义价值
            寄予人类和平共存
        
--- PASS: TestParseXpath (0.00s)
PASS

```

**json 数据: content.json**

```

[
  {
    "n_SportID": 117,
    "c_Sport": "Swimming",
    "c_SportShort": "SWM",
    "n_EventID": 9775,
    "c_Event": "50m Freestyle",
    "c_EventShort": "50m Free",
    "n_EventSort": 3,
    "n_GenderID": 1,
    "c_Gender": "Men",
    "c_GenderShort": "M",
    "c_Participant": "Cesar Cielo",
    "c_ParticipantShort": "Cielo",
    "c_ParticipantFirstName": "Cesar",
    "c_ParticipantLastName": "Cielo",
    "n_ParticipantNatioGeoID": 2235,
    "c_ParticipantNatio": "Brazil",
    "c_ParticipantNatioShort": "BRA",
    "n_TeamID": 906,
    "n_PersonID": 570751,
    "n_NOCID": 16,
    "n_NOCGeoID": 2235,
    "c_NOC": "Brazil",
    "c_NOCShort": "BRA",
    "c_Result": "20.91",
    "d_Date": "\/Date(1261090800000+0100)\/",
    "n_Date": 20091218,
    "c_Date": "2009-12-18T00:00:00",
    "n_LocationGeoID": 25180,
    "c_Location": "São Paulo",
    "n_CountryGeoID": 2235,
    "c_Country": "Brazil",
    "c_CountryShort": "BRA",
    "c_Competition": null,
    "c_CompetitionShort": null
  }
]

```

**json 解析**

```
func ParseByJson(reader io.Reader) {
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println(err)
		return
	}

	type result struct {
		sport       string
		event       string
		gender      string
		participant string
		national    string
		date        string
	}

	doc := gjson.ParseBytes(content).Array()

	var one result
	for _, i := range doc {
		sport := i.Get("c_Sport").String()
		event := i.Get("c_Event").String()
		gender := i.Get("c_Gender").String()
		participant := i.Get("c_Participant").String()
		national := i.Get("c_ParticipantNatio").String()
		date := i.Get("c_Date").String()
		one = result{
			sport:       sport,
			event:       event,
			gender:      gender,
			participant: participant,
			national:    national,
			date:        date,
		}
		fmt.Println(one)
	}

}

# 测试

func TestParseJson(test *testing.T) {
	file, err := os.OpenFile("content.json", os.O_RDWR, 650)
	if err != nil {
		log.Println(err)
		return
	}
	fina.ParseByJson(file)
}


# 结果

=== RUN   TestParseJson
{Swimming 50m Freestyle Men Cesar Cielo Brazil 2009-12-18T00:00:00}
--- PASS: TestParseJson (0.00s)
PASS

```

按照这样的方式，把需要的数据全部抓取入库，鉴于篇幅所限，并没有详细讲解内容的抓取过程，核心思想总结如下：

- 分析网页，内容是 html 还是 json ： 工具 chrome 浏览器
- 模型定义：xorm 、gorm 
- 发起网络请求获取网页源代码： net/http
- 解析网页源代码： css 选择器、xpath 路径表达式、json 解析
- 存储入库
- 后续操作


整个的过程，笔者认为最重要的步骤是：

- 借助各种手段分析资源的内容实体，明确内容在哪
- 解析：熟悉 css 选择器，xpath 语法，json 解析以应对更为复杂的解析


数据获取是作为数据驱动项目的第一步，现实的数据获取，只会比上文解释的复杂，甚至某些互联网公司有专门的爬虫工程师，也有专门的反爬虫工程师。最后需要说明的是，数据获取需要在符合法律法规的情况下进行。



### 2.4 GraphQL 接口开发

模型建立好了，数据也获取到了，所有的数据都存储在数据库中，别忘了，最终的目的是构建 GraphQL 风格的 API。GraphQL 是基于类型的 API 查询语言，最重要的是定义哪些类型（即真实的响应信息），提供哪些接口。

整个的流程模块化成下面几个步骤：

- 定义类型文件，即 schema.graphql， 非必要，但很重要，可以作为开发指导、接口文档
- 每个类型对应资源实体，分别定义字段和类型
- 使用 graphql-go 进行类型定义
- 启动服务
- 接口测试，是否满足需求，否则，修复


类型到底需要定义哪些？一个很简单的准则：有多少资源实体，便定义多少类型（整个的过程和 RESTful 风格的很相似，有多少资源，抽象多少实体，定义多少类接口）。


整个的项目是：第18届国际泳联世锦赛，获取到的数据和也这些相关：金牌数、参赛国家、吉祥物、FINA 历史...


```
mysql> show tables;
+----------------+
| Tables_in_fina |
+----------------+
| blue           |
| championships  |
| competitions   |
| country        |
| fina           |
| history        |
| kinds          |
| medal          |
| records        |
| sports         |
| symbol         |
+----------------+
11 rows in set (0.02 sec)

```

> 对照着定义模型环节看

- blue: 标语 PEACE 意义及解释
- championships: 第18届国际泳联相关信息
- competitions： 具体的比赛项目
- country: 国家信息
- fina: 国际泳联信息，创立时间，成员国等
- history: 国际泳联历届比赛信息
- kinds: 比赛项目、场地
- medal: 各国奖牌数
- records: 游泳比赛世界记录
- sports: 比赛项目金牌数、规则等
- symbol: 吉祥物信息


基于此， schema.graphql 类型文件的内容即是这些资源：

```

# 查询类型
type Query {
    ping(data:String!): ResponseForPing!
    countries(name:String, short: String, all: Boolean):[Country]
    countryMedal(name: String!, year: Int): [Medal]
    countryMeadlRank(year: Int!, sortBy: String!):[Medal]
    history(year: Int): [History]
    histories(orderBy: String): [History]
    fina: FiNa
    blues: [Blue]
    symbol: Symbol
    kinds(class: KindClass!): [Kind]
    competitions(class: CompetitionClass!): [Competition]
    sports(class: SportClass!):[Sport]
    records(name: String, all: Boolean): [Record]
}


# 心跳类型
type ResponseForPing {
    code: Int
    data: String
}

# 国家类型
type Country {
    id:ID!
    createdAt: String
    updatedAt: String
    name: String
    short: String
}

# 奖牌类型
type Medal {
    id: ID!
    createdAt: String
    updatedAt: String
    year: Int
    countryId: ID
    countryName: String
    gold: Int
    silver:Int
    bronze:Int
}

# 历史信息类型
type History {
    id : ID
    createdAt: String
    updatedAt: String
    year: Int
    detail: String
}

# 国家泳联类型
type FiNa {
    id: ID
    createdAt: String
    updatedAt: String
    description: String
    established: String
    headquarters: String
    nationalMember: String
    numberOfDisciplines: String
}

# 标语类型
type Blue {
    id: ID
    createdAt: String
    updatedAt: String
    short: String
    enName: String
    chName: String
    description: String
    image: String
}

# 吉祥物类型
type Symbol {
    id: ID
    createdAt: String
    updatedAt: String
    symbolText: String
    symbolTextImage: String
    symbolTextShort: String
    symbolDescription: String
    symbolDescriptionImage: String
    symbolDescriptionShort: String
    symbolAnimalImage: String
    symbolAnimalDescription: String
    symbolAnimalShort: String
    blueVersions: [Blue]
}

# 种类枚举
enum KindClass {
    DISCIPLINE
    VENUES
}

# 种类类型
type Kind {
    id: ID
    createdAt: String
    updatedAt: String
    name: String
    class: Int
    classString: String
}

# 比赛类型枚举
enum CompetitionClass {
    MAN
    WOMAN
    TEAM
}

# 比赛类型
type Competition {
    id:ID
    createdAt: String
    updatedAt: String
    competitionClass: Int
    competitionClassString: String
    detail: String
}

# 项目类型枚举
enum SportClass {
    SWIMMING
    DIVING
    HIGHDIVING
    ARTISICSWIMMING
    OPENWATER
    WATERPOLO
}

# 项目类型
type Sport {
    id: ID
    createdAt: String
    updatedAt: String
    total: Int
    sportClass: Int
    sportClassString: String
    sportName: String
    description: String
    rule: String
    competitions: [Competition]

}

# 记录类型
type Record {
    id: ID
    createdAt: String
    updatedAt: String
    eventName: String
    record: String
    countryId: ID
    countryName: String
    date: String
    location: String
    competitionClass: Int
    competitionClassString: String
    sportClass: Int
    sportClassString: String
    name: String
}

```


这些字段为什么这么定义，有没有什么规范？

**为什么这么定义？**

模型定义环节，笔者的操作，是给这个模型定义一个序列化结构体（对应：该模型的响应），模型的序列化方法的返回值是这个序列化结构体。

```

# 模型的定义，对应的数据库内的字段
type RecordMax struct {
	Base             `xorm:"extends"`
	EventName        string    `xorm:"varchar(32)" json:"event_name"`
	Record           string    `xorm:"varchar(32)" json:"record"`
	CountryId        int64     `xorm:"integer(3)" json:"country_id"`
	Date             time.Time `json:"date"`
	Location         string    `json:"location"`
	CompetitionClass int       `json:"competition_class"`
	SportClass       int       `json:"sport_class"`
	Name             string    `json:"name"`
}

# 表名称
func (R RecordMax) TableName() string { return "records" }


# 序列化结构体，对应的是 RecordMax 模型的响应信息
type RecordsMaxSerializer struct {
	Id                     int64     `json:"id"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	EventName              string    `json:"event_name"`
	Record                 string    `json:"record"`
	CountryId              int64     `json:"country_id"`
	CountryName            string    `json:"country_name"`
	Date                   time.Time `json:"date"`
	Location               string    `json:"location"`
	CompetitionClass       int       `json:"competition_class"`
	CompetitionClassString string    `json:"competition_class_string"`
	SportClass             int       `json:"sport_class"`
	SportClassString       string    `json:"sport_class_string"`
	Name                   string    `json:"name"`
}

# 模型结构体的方法，返回值是序列化结构体
func (R RecordMax) Serializer() RecordsMaxSerializer {
	var country Country
	database.MySQL.ID(R.CountryId).Get(&country)
	return RecordsMaxSerializer{
		Id:                     R.Id,
		CreatedAt:              R.CreatedAt.Truncate(time.Second),
		UpdatedAt:              R.UpdatedAt.Truncate(time.Second),
		EventName:              R.EventName,
		Record:                 R.Record,
		CountryId:              R.CountryId,
		CountryName:            country.Name,
		Date:                   R.Date,
		Location:               R.Location,
		CompetitionClass:       R.CompetitionClass,
		CompetitionClassString: CompetitionClass[R.CompetitionClass],
		SportClass:             R.SportClass,
		SportClassString:       SportClass[R.SportClass],
		Name:                   R.Name,
	}
}

```


类型文件中的 Record 对应的字段即上 RecordsMaxSerializer 的字段，这就是我们的类型文件定义的参考。

**schema.graphql 文件有什么规范？**

规范主要围绕在命名上，1. 采用驼峰式命名方式 2. 类型名称首字母大些，字段首字母小写 3. 字段和字段的类型使用分隔符冒号。


#### 2.4.1 项目组织

结构清晰项目组织，展现的是开发者的逻辑思维能力，能够较为清晰的组织项目。还记得 RESTful API 设计章节中关于项目的组织吗？不同类型的接口开发是否可以复用之前的项目组织结构？答案肯定的，但还是稍微有点差异，整体差不多，差异主要体现在对应的是不同的系统，RESTful 是对资源的组织，GraphQL 是对类型的组织。


项目地址：https://github.com/wuxiaoxiaoshen/GopherBook/tree/master/chapter12/fina

```
.
├── Makefile
├── cmd
│   ├── data
│   ├── import_cmd.go
│   ├── root_cmd.go
│   └── sync_cmd.go
├── configs
├── data
│   └── fina-2019-08-06.sql
├── deployments
├── go.mod
├── go.sum
├── main.go
├── models
├── pkg
├── schema.graphql
├── scripts
└── web
    ├── blue
    ├── competition
    ├── country
    ├── country_medal
    ├── fina
    ├── history
    ├── kind
    ├── mutation
    │   └── type_mutation.go
    ├── ping
    │   └── type_ping.go
    ├── query
    │   └── type_query.go
    ├── records
    │   ├── curd_records.go
    │   ├── param_records.go
    │   └── type_records.go
    ├── sports
    └── symbol

```


- Makefile 项目构建文件，主要提供出 make 命令
- cmd: 提供项目命令行工具，包括数据的抓取，数据库的表迁移等
- configs: 项目配置文件，主要包括数据库地址等
- data: 项目导出 sql 语句，直接导入，启动项目即可，方便读者学习
- deployments: 容器构建文件
- models: 模型定义文件夹
- pkg: 项目用到的包，主要包括：日志、错误信息、中间件等
- schema.graphql 类型文件，同时作为 API 文档
- scripts: 脚本信息
- web: 项目核心出来逻辑

**web 层的思路**

- 按资源实体划分，比如 blue、competition 等， 其中 query、mutation 对应 GraphQL 查询入口
- 每个资源实体下面，划分多个文件：curd_repords.go 对应实体的增删改查操作（可以换其他名字，关键是要注意统一性），param_records.go 对应的是操作资源实体需要的参数，可以完成参数的校验，type_records.go 类型文件的定义，主要是构造类型文件：字段、字段类型、Reslove 函数等，当然还可能有 enum_records.go 对应的是资源实体需要用到的枚举类型，input_records.go 对应的是资源实体需要用到的输入类型等。


再次说明，命名方式因人而异，但一定要保持统一的风格。


#### 2.4.2 基础模块

**数据库连接: pkg/database/database.go**

主要是创建连接对象，具体使用 xorm, 还是 gorm 或者原生的 database/sql ， 依个人开发者喜好来，推荐使用 orm， 示例以 xorm 为例。其次为方便排查问题，数据库操作的日志最好打开，方便定位问题。



```
var MySQL *xorm.Engine

var (
	dbMySQL    = "fina"
	dbUser     = "root"
	dbPassword = "adminMysql"
)

func MySQLInit() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPassword, dbMySQL))
	if err != nil {
		panic(fmt.Sprintf("CONNECT ENGINE BY XORM FAIL %s", err.Error()))
	}
	engine.SetTableMapper(core.SameMapper{})
	engine.SetMapper(core.GonicMapper{})
	MySQL = engine
	MySQL.ShowSQL(true)
	return MySQL

}
```

**日志模块：pkg/log/logger.go**

日志模块主要是方便查看关键信息点，原生的不支持颜色打印，对原生进行了简单的封装。日志库，开源社区内存在非常非常多的优秀的库，没有特殊的需求，建议还是使用内置的日志模块。

```
package log_for_project

import (
	"fmt"
	"log"
)

func red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}

func Println(message string) {
	log.Println(fmt.Sprintf(red(message)))
}

```

**错误信息模块：pkg/error/error.go**

错误信息一般的处理逻辑是实现 Error 接口，但本项目中 graphql-go 的报错信息非常强大、详尽，几乎能绝对定位错误问题，故只编写可复用的错误信息。

```

package error_for_project

import "errors"

var (
	NotFound   = errors.New("field not found")
	ParamField = errors.New("forget post params ")
)


```

#### 2.4.3 类型定义

本环节以国际泳联的实体资源 Sports 讲述类型定义的方法。还需要明确的是类型定义对标的是响应，即模型结构体 Sports 的 Serializer 方法的返回值。

```
type Query {
    sports(class: SportClass!):[Sport]
}

enum SportClass {
    SWIMMING
    DIVING
    HIGHDIVING
    ARTISICSWIMMING
    OPENWATER
    WATERPOLO
}

type Sport {
    id: ID
    createdAt: String
    updatedAt: String
    total: Int
    sportClass: Int
    sportClassString: String
    sportName: String
    description: String
    rule: String
    competitions: [Competition]
}

```



![](http://ww1.sinaimg.cn/large/741fdb86gy1g5tv962ft6j21300gqtah.jpg)


将 graphql-go 中关于类型对象的抽象出上图的结构，即一个类型对象包含多个字段，每个字段都有类型和解析函数。

**单个字段定义**

```

var id = &graphql.Field{
	Type: graphql.ID,
	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		if sport, ok := p.Source.(models.SportSerializer); ok {
			return sport.Id, nil
		}
		return nil, error_for_project.NotFound
	},
}

```

- Type 对应类型
- Resolve: 解析函数，解析 SportSerializer 中的 Id 字段


对应的 SportSerializer 每个字段都按上文操作：

```
type SportSerializer struct {
	Id               int64                   `json:"id"`
	CreatedAt        time.Time               `json:"created_at"`
	UpdatedAt        time.Time               `json:"updated_at"`
	Total            int                     `json:"total"`
	SportClass       int                     `json:"sport_class"`
	SportClassString string                  `json:"sport_class_string"`
	SportName        string                  `json:"sport_name"`
	Description      string                  `json:"description"`
	Competitions     []CompetitionSerializer `json:"competitions"`
	Rule             string                  `json:"rule"`
}
```


**Sports 类型对象定义**

```

var Sports = graphql.NewObject(graphql.ObjectConfig{
	Name: "Sports",
	Fields: graphql.Fields{
	    "id": id,
	    "createdAt": createdAt,
	    ...
	}
```

**枚举类型定义**

枚举类型是一系列可选值的集合，核心是一系列值。

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5tx8a6xihj20zc0i675s.jpg)


schema.graphql 中枚举类型：

```

enum SportClass {
    SWIMMING
    DIVING
    HIGHDIVING
    ARTISICSWIMMING
    OPENWATER
    WATERPOLO
}


```

在代码中实现如下：

```

var SportEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "sportEnum",
	Values: graphql.EnumValueConfigMap{
		"swimming": &graphql.EnumValueConfig{
			Value: models.SWIMMING,
		},
		"diving": &graphql.EnumValueConfig{
			Value: models.DIVING,
		},
		"highDiving": &graphql.EnumValueConfig{
			Value: models.HIGHDIVING,
		},
		"artisticSwimming": &graphql.EnumValueConfig{
			Value: models.ARTISTICSWIMMING,
		},
		"openWater": &graphql.EnumValueConfig{
			Value: models.OPENWATER,
		},
		"waterPolo": &graphql.EnumValueConfig{
			Value: models.WATERPOLO,
		},
	},
})


```

可以看出枚举类型就是一系列可选值。

**Query 类型**

```
type Query {
    sports(class: SportClass!):[Sport]
}
```

- 字段名称是：sports
- 变量名称是: class, 类型是枚举类型 SportClass
- Query 字段类型是列表：Sport


```

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query"})

func init() {
	Query.AddFieldConfig("sports", &graphql.Field{
		Name: "sports",
		Type: graphql.NewList(sports.Sports),
		Args: graphql.FieldConfigArgument{
			"class": &graphql.ArgumentConfig{
				Type: sports.SportEnum,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			type result struct {
				data []models.SportSerializer
				error
			}
			var param sports.GetSportParam
			param.Class = p.Args["class"].(int)
			controller := sports.Default
			ch := make(chan result, 1)
			go func() {
				defer close(ch)
				data, err := controller.GetSports(param)
				ch <- result{data: data, error: err}
			}()
			return func() (interface{}, error) {
				r := <-ch
				return r.data, r.error
			}, nil

		},
	})
}

```

Query 类型同样是定义字段类型和解析函数，只不过 Query 是请求的入口，整体处理起来稍微复杂点。

- Type: graphql.NewList(sports.Sports) 列表
- Args: 变量名称 class, 类型（枚举类型）：sports.SportEnum
- Resolve: 解析函数，返回一个数组


真实的解析函数是操作模型 Sports 进行资源的获取。

```
type ControllerSports struct {
}

var Default = ControllerSports{}


// 根据传入的类型，查找所有的竞技项目
func (C ControllerSports) GetSports(param GetSportParam) ([]models.SportSerializer, error) {
	var result []models.SportSerializer
	var sports []models.Sports
	if err := param.Valid(); err != nil {
		return result, err
	}
	if dbError := database.MySQL.Where("sport_class = ?", param.Class).Find(&sports); dbError != nil {
		return result, nil
	}
	for _, i := range sports {
		result = append(result, i.Serializer())
	}
	return result, nil
}

```


**总结**

类型的定义，一个类型对象由多个字段组成，每个字段有对应的Type 和 Resolve 解析函数，核心的逻辑是由 Resolve 来处理，完成资源的增删改查。由于需要对每个类型的字段进行定义，整体上 GraphQL 开发下来，代码量会偏多，将资源类型按文件进行项目组织尤为重要。Query 作为核心处理入口，最好也需要对资源进行划分。使其整体看上去如下面的格式：

```
var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query"})
	

// 资源 blue
func init(){
    Query.AddFieldConfig("blue", &graphql.Field{})
}


// 资源 sports
func init() {
    Query.AddFieldConfig("sports", &graphql.Field{})
}

// 资源 competitions
func init(){
   Query.AddFieldConfig("competitions", &graphql.Field{})
}

// 资源 history 
func int {
    Query.AddFieldConfig("history", &graphql.Field{})
}

...
```

### 2.5 启动服务


启动服务主要的操作是构建 Schema 对象，构建路由。

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5ubni35jtj213r0dytaa.jpg)



可以看出 schema 对应主要由 Query, Mutation, Subscription 三者构成，实际上Query 就是个 graphql.Object 对象。

```
// 构建 shcema 对象
func RegisterSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: query.Query,
		//Mutation: mutation.Mutation,
	})
}

// 构建处理器
func RegisterHandler() *handler.Handler {
	schema, err := RegisterSchema()
	if err != nil {
		log.Println(err)
		return nil
	}
	return handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
}

// 启动服务
func StartWeb() {
	h := RegisterHandler()
	http.HandleFunc("/graphql", h)
	log.Fatalln(http.ListenAndServe(":2345", nil))
}
```


所有的请求都绑定在同一个路由上，GraphQL 一个重要的特征是避免路由爆发式增长，当然路由 /graphql 一般约束成这个名称，读者换成其他名称同样可以。


### 2.6 中间件：日志

接口开发其中一个重要的环节是：中间件，比如跨域请求、日志、认证等。同样在 graphql 中也是采用中间件的形式完成这些任务。

日志中间件，一般用来处理网络请求中的一些重要的特征：比如访问的路由、请求的参数、头部信息等，把关键点收集起来，方便及时定位问题。

**日志中间件：pkg/middleware/logger.go**
```


func Logger(ctx context.Context, h *handler.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		message := fmt.Sprintf("%s | %s | %s | %s", request.Method, request.Host, request.RequestURI, time.Now().Format(time.RFC3339))
		log_for_project.Println(message)
		bodyBytes, _ := ioutil.ReadAll(request.Body)
		defer request.Body.Close()

		var opts handler.RequestOptions
		json.Unmarshal(bodyBytes, &opts)
		log_for_project.Println(opts.Query)
		request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		h.ContextHandler(ctx, writer, request)
	}
}

```

上文的中间件接受一个 context 对象和 handler.Handler 对象，将网络请求的方法、主机地址、路由、请求参数在终端上显示，因为已经读过一次 request， 为复用 request, 再次使用 ioutil.NopCloser 将请求参数写入 request， 否则报错。

使用上日志中间件，微调启动服务代码：

```
func StartWeb() {
	h := RegisterHandler()
	ctx := context.TODO()
	http.HandleFunc("/graphql", middleware.Logger(ctx, h))
	log.Fatalln(http.ListenAndServe(":2345", nil))
}
```

其他认证、跨域等中间件，同样是如此操作，之所以使用 context 对象，是方便在不同的协程中并发安全的复用数据。



### 2.7 接口测试


**方式一： Postman**

- 路由：POST localhost:2345/graphql
- 选择 GraphQL 模式
- Headers: content-type: application/json
- 请求
```
query {
  sports(class: swimming) {
    id
    sportName
    sportClass
    sportClassString
    description
  }
}
```

结果：

```
{
	"data": {
		"sports": [
			{
				"description": "游泳项目以自由泳，仰泳，蛙泳，蝶泳等指定游法，50米到1,500米的方式进行，顺序为预赛，半决赛和决赛。\r\n个人项目分为自由泳(50, 100, 200, 400, 800, 1,500米)·仰泳·蛙泳·蝶泳(50, 100, 200米)和个人混合泳(200, 400米)，团体项目分为接力泳(400, 800米)、混合泳接力(400米)、男女接力泳(400米)、男女混合泳接力(400米)。",
				"id": "17",
				"sportClass": "swimming",
				"sportClassString": "游泳",
				"sportName": "游泳"
			}
		]
	}
}
```

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5uckt31gsj20wj0kw0v2.jpg)


**方式二：GraphiQL**

- 路由：POST localhost:2345/graphql
- 请求参数

```
query {
  sports(class: swimming) {
    id
    sportName
    sportClass
    sportClassString
    description
  }
}
```

结果：

```
{
  "data": {
    "sports": [
      {
        "description": "游泳项目以自由泳，仰泳，蛙泳，蝶泳等指定游法，50米到1,500米的方式进行，顺序为预赛，半决赛和决赛。\r\n个人项目分为自由泳(50, 100, 200, 400, 800, 1,500米)·仰泳·蛙泳·蝶泳(50, 100, 200米)和个人混合泳(200, 400米)，团体项目分为接力泳(400, 800米)、混合泳接力(400米)、男女接力泳(400米)、男女混合泳接力(400米)。",
        "id": "17",
        "sportClass": "swimming",
        "sportClassString": "游泳",
        "sportName": "游泳"
      }
    ]
  }
}
```

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5ucmsogzjj213x0jcwha.jpg)


**总结**

两者都非常适合 GraphQL 风格的API 接口测试，个人倾向于 GraphiQL 进行 GraphQL 风格的 API 测试，主要原因在于其具备自动提示字段的功能。



### 2.8 总结


本节围绕的主要是构建 GraphQL 风格的API 的设计，主要的内容包括：

- GraphQL 本身的查询语法
- graphql-go 的使用
- 结合一个示例，构建第18届国际泳联的数据项目：包括数据的获取，类型的设计，字段的设计等


笔者认为，构建 GraphQL 风格的 API 难点在什么地方？

- 把握需求，逻辑清晰的进行模型设计
- 把握需求，逻辑清晰的设计 schema， schema 设计的好，几乎都完成大半的任务
- 合理的进行项目组织，使整体逻辑清晰分明，不管采用何种风格，需要维持一致性


RESTful 风格和 GraphQL 风格的 API 的目的都是一样，对外提供接口的形式，只不过两者采用的实现手段不同，至于如何选择，依赖于团队内的技术选型，市面上采用 GraphQL 风格进行设计的公司不算多，绝大多数还是采用 RESTful, 因为历史项目的原因，绝大多少技术团队技术选型依然会选择 RESTful, 如果没有技术负担，选择 GraphQL 是更好的选择。


## 3. Go GraphQL 开发指南


实例可以使读者整体上了解 GraphQL 风格的开发，但有一些注意的开发准则还需遵守，本节主要围绕开发过程中一些规律性的话题进行总结。


### 3.1 项目组织

在各个环节，都在陈述项目组织的重要性。其一，平时严格要求自己养成良好组织项目的习惯，在真实的企业级开发项目中，不至于逻辑混乱，理不清楚。其二，需求是多变的，良好的项目组织，对多变的需求可以满足其拓展性。

GraphQL 风格的开发，项目组织的一个重要原则是：按资源实体划分（或者按照功能划分）

```
web
├── blue
│   ├── curd_blue.go
│   ├── param_blue.go
│   └── type_blue.go
├── competition
│   ├── curd_competition.go
│   ├── enum_competition.go
│   ├── param_competiton.go
│   └── type_competition.go
├── country
│   ├── curd_country.go
│   ├── param_country.go
│   └── type_country.go
...
```

项目组织有没有什么值得参考的经验？

- 维持统一性
- 保持拓展性


统一性指统一的风格，统一的风格可以体现在多个方面，下面从整体，内部两个方面陈述。

对整体而言，项目按照资源划分，比如遇到新的资源实体，新建文件夹即可，再保持一致的文件命名方式。

```
web
├── blue
...
├── sports
│   ├── curd_sports.go
│   ├── enum_sports.go
│   ├── param_sports.go
│   └── type_sports.go
...
```

对内而言，变量、函数、方法都维持一致的命名方式或者组织方式：

- curd_sports.go: 真实的业务逻辑操作，主要对接的是模型和数据库

```
type ControllerSports struct {}

var Default = ControllerSports{}

func (C ControllerSports) GetSports(param GetSportParam) {}

```

- enum_sports.go: 枚举类型定义，主要的是定义一系列可选值

```
var SportEnum = graphql.NewEnum(graphql.EnumConfig{}
```

- param_sports.go: 参数处理，主要是将参数定义成结构体和参数检验

```
type GetSportParam struct {
	Class int `json:"class" validator:"eq=0|eq=1|eq=2|eq=3|eq=4|eq=5"`
}

func (G GetSportParam) Valid() error {
	return validator.New().Struct(G)
}

```

- type_sports.go: 类型定义，由一系列字段（Fields）组成

```
var Sports = graphql.NewObject(graphql.ObjectConfig{})
```


如果你没有真实项目的组织经验，可以参考一些开源项目的处理，尤其是一些框架的项目组织，希望自己模仿着组织，受到一些启发。


### 3.2 参数检验

参数校验是对请求参数进行部分限制，满足特定的业务需求，一方面 GraphQL 本身的查询会限制参数，不符合要求的会直接报错。

比如请求参数本身限制在某个枚举类型当中，除此之外，一般的做法是解析网络请求，获取对应的字段，构造成对应的结构体，对结构体的相应方法进行参数校验，参数校验的方式和 RESTful 中一致，要么在结构体 Tag 内限定，要么在具体的方法内进行参数限定。

```
type GetSportParam struct {
	Class int `json:"class" validator:"eq=0|eq=1|eq=2|eq=3|eq=4|eq=5"`
}

func (G GetSportParam) Valid() error {
    // do something
	return validator.New().Struct(G)
}
```

```
func init(){
    	Query.AddFieldConfig("sports", &graphql.Field{
		Name: "sports",
		Type: graphql.NewList(sports.Sports),
		Args: graphql.FieldConfigArgument{
			"class": &graphql.ArgumentConfig{
				Type: sports.SportEnum,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			type result struct {
				data []models.SportSerializer
				error
			}
			var param sports.GetSportParam
			param.Class = p.Args["class"].(int)
			controller := sports.Default
			ch := make(chan result, 1)
			go func() {
				defer close(ch)
				data, err := controller.GetSports(param)
				ch <- result{data: data, error: err}
			}()
			return func() (interface{}, error) {
				r := <-ch
				return r.data, r.error
			}, nil

		},
	})
}
```

- 限制请求的参数类型是： Args: sports.SportEnum
- 将网络请求，解析成对应的结构体：var param sports.GetSportParam



### 3.3 并发处理

GraphQL API 的开发，对外呈现只有一个路由，请求中可以根据需求，一次获取多个请求响应(在RESTful 中需要调用多个接口才能实现)

```
type Query {
    sports(class: SportClass!):[Sport]
    records(name: String, competitionClass:CompetitionClass, sportClass: SportClass, all: Boolean): [Record]

}
```

如果一次性进行这两个请求，最好的方式是并发的执行，这样的耗时，取决于最长的那个，而不是两者的叠加。

![](http://ww1.sinaimg.cn/large/741fdb86gy1g5unbms403j213t0joady.jpg)


并发如何处理？通过上文的解释，我们知道，最核心的处理逻辑是在 Resolve 解析函数那块，那么并发的思路是从这块入手。每个Resolve 启动一个协程，并发的执行。

```
Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
	controller := records.Default
	type result struct {
		data []models.RecordsMaxSerializer
		error
	}
	var param records.GetRecordParam
	ch := make(chan result, 1)
	if p.Args["all"] != nil {
		param.All = p.Args["all"].(bool)
	} else {
		param = records.GetRecordParam{
			Name: p.Args["name"].(string),
		}
		if p.Args["sportClass"] != nil {
			param.SportClass = p.Args["sportClass"].(int)
		}
		if p.Args["competitionClass"] != nil {
			param.CompetitionClass = p.Args["competitionClass"].(int)
		}
	}
	// 启动并发协程
	go func() {
		defer close(ch)
		data, err := controller.GetRecords(param)
		ch <- result{data: data, error: err}
	}()
	return func() (interface{}, error) {
		r := <-ch
		return r.data, r.error
	}, nil
},

```

- 解析参数，进行校验等，按照正常的逻辑，核心的业务数据库逻辑操作，使用并发执行



### 3.4 总结


关于 GraphQL API 开发注意的事项，主要围绕在项目组织、参数校验、中间件、并发处理层面。

不管怎么说，业务逻辑的开发还是需要针对具体的问题具体分析，但整体的处理流程大致相似。之所以总是从开发过程中总结出一些核心的流程化的内容，是为了方便快速展开项目，一些规律性的内容，抽象出来，让开发人员聚焦在实现核心的业务价值层面。每个人总结出的内容流程化的东西各不相同，希望读者批判性的看待。


## 4 GraphQL 总结


GraphQL 被誉为替换 RESTful 风格的下一代 API 工具，解决了部分问题，比如路由的爆发式增长，路由版本问题等，这一整套的语法约束极大的方便了前端开发人员，更多的处理逻辑由后端开发人员承担。但是 GraphQL 也不全是优势，最大的问题的是 N + 1 问题带来的性能问题，开源社区也有一些解决方案进行规避。


整体的 GraphQL 的API 开发流程很简单。

- 根据需求，开发出 schema 文件作为开发指导
- 根据需求进行模型设计，模型的字段决定了具体的 API 的响应字段，模型设计是整个项目最为核心的部分，不管是 RESTfule 风格还是 GraphQL 风格的开发中
- 良好的项目组织，按照某个梳理准则进行组织，比如资源的类型
- 定义各种类型，包括字段的定义，解析函数的处理
- 接口测试

参考代码：https://github.com/wuxiaoxiaoshen/GopherBook/tree/master/chapter12/fina