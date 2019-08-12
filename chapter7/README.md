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