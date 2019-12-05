# README.md


## 解析工具

- jquery: goquery 规则
- 正则表达式: regexp
- xpath: 规则


## 文本抓取（获取数据的重要手段）


- 知乎热搜
- v2ex 热搜
- 微博热搜
- 电影
- 新剧
- 独角兽
- APP(mitmproxy、wireshark)
    - 小程序：抽奖助手
    - 抖音
- 拉勾、脉脉、

## 框架

## 解析方式

- Regexp
- jQuery(css)
- xpath

高频 CSS/XPATH 语法示例

## 详情

**Web**

- html 的组成： dom 树

```html
<html>
<h1 id="id-h1" class="class-h1" href="http://www.baidu.com">
 <p>Hello Golang</p>
 <p>Hello Python</p>
</h1>
<h1 id="id-h2" class="class-h2" href="http://www.baidu.com"> Hello Golang</h1>
</html>
```

- 层级：包含关系
- 相邻关系
- 父子关系
- 属性、字段、值

**App**

```json
{
  "data": {
    "language": "go",
    "list": [1,2,3],
    "name": {
      "first_name": "xie",
      "last_name": "wei"
    }
  }
}

```

- string
- k/v （map）
- 整型
- 数组





### chromedp api

- chromedp.Navigate
- chromedp.CaptureScreenshot
- chromedp.Emulate
- chromedp.Click

- chromedp.Value
- chromedp.Text
- chromedp.OutHtml
- chromedp.Node


### 核心

> 提供高级 API 来通过 DevTools(是内嵌在 Chrome 浏览器里的一组用于网页制作和调试的工具) 协议来控制 Chromium 或 Chrome

Headless Chrome 是 Chrome 浏览器的无界面形态，可以在不打开浏览器的前提下，使用所有 Chrome 支持的特性运行你的程序。相比于现代浏览器，Headless Chrome 更加方便测试 web 应用，获得网站的截图，做爬虫抓取信息等。相比于出道较早的 PhantomJS，SlimerJS 等，Headless Chrome 则更加贴近浏览器环境。

Headless 方案：

- Selenium
- Chromedp: golang 版
- Puppeteer: Nodejs 版 （https://zhaoqize.github.io/puppeteer-api-zh_CN/）

Selenium和 Puppeteer的比较：

- Selenium加上不同的WebDriver，可以支援不同的浏览器(例如:chrome、firefox、IE)， 但是Puppeteer只用于Chromium或Chrome。

## chrome 开发者工具

- https://developers.google.com/web/tools/chrome-devtools/?hl=zh-cn
- chrome 版本号 http://omahaproxy.appspot.com/


## 说明

- assistance: 获取网页源代码

```text
chromedep: 使用无界面的浏览器获取网页源代码
content: 使用原生的 net/http 方式获取网页源代码
selenium: 使用 slenium 获取网页源代码
```

- data: 爬虫示例

