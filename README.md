# gin_xiaomi_shop

## 总结
### iframe
iframe标签实现导航栏状态保存。
- 场景：左侧导航栏，右侧显示区，当跳转导航时，右侧显示区显示不同当内容，左侧导航栏状态需要保持不变。
- 使用：
```html
<iframe src="/admin" width="100%" height="500" name="rightMain"></iframe>
<a href="/default" target="rightMain"></a>
<!--
    1. src：默认会和本地地址拼接，内嵌当前页面url地址
    2. name：可以和a标签当target结合使用
    3. 当a标签的target和iframe的name一样时，点击a标签，那么iframe内嵌的就是a标签跳转的页面
-->
```

### base64Captcha库的使用
```shell
go get -u github.com/mojocn/base64Captcha
```
[base64Captcha](https://github.com/mojocn/base64Captcha)是一个验证码生成和校验的库。具体使用规则见`common/captcha.go`。

### meta标签
meta标签用来提供有关html文档的元数据，元数据不会显示给用户看，但浏览器可以识别，同时对SEO起着重要作用。
- author 用来表示网页作者名字，例如某个组织或结构。
```html
<meta name="author" content="aaa@ail.asdf.com" />
```
- description 对页面内容的描述
```html
<meta name="description" content="全球领先的中文搜索引擎、致力于让网民更便捷地获取信息，找到所求。百度超过千亿的中文网页数据库，可以瞬间找到相关的搜索结果。" />
```
- keywords 页面内容相关的关键词，以逗号隔开，有利于SEO优化
```html
<meta name="description" content="aaa,bbb,ccc" />
```
- viewport
```html
 <meta name="viewport" content="width=device-width, initial-scale=1.0">
```
- referrer 控制由当前文档发出的请求的 HTTP Referer 请求头，详细的配置属性见 https://developer.mozilla.org/zh-CN/docs/Web/HTML/Element/meta/name
```html
<meta name="referrer" content="no-referrer">
```
- X-UA-Compatible 用来做IE浏览器适配 必须具有值 "IE=edge"
```html
<meta http-equiv="X-UA-Compatible" content="ie=edge,chrome=1">
```
- charset 该属性声明了文档的字符编码。如果存在该属性，则其值必须是字符串 "utf-8" 的不区分 ASCII 大小写的匹配，因为 UTF-8 是 HTML5 文档的唯一有效编码。声明字符编码的 <meta> 元素必须完全位于文档的前 1024 个字节内。
```html
<meta charset="UTF-8">
```
- content-type 该属性声明了文档类型和字符集。其值必须是字符串"text/html; charset=utf-8"。
```html
<meta http-equiv="content-type" content="text/html;charset=utf-8">
```
- refresh 
- 页面重新加载的秒数：仅当 content 属性包含非负整数时。
- 页面重定向到指定链接的秒数：仅当 content 属性包含非负整数后跟字符串“;url=”和有效的 URL 时。
```html
<!-- 5s后页面重新加载 -->
<meta http-equiv="refresh" content="5">  
<!-- 5s后页面自动跳转到百度首页 -->
<meta http-equiv="refresh" content="5;url=https://www.baidu.com">
```

