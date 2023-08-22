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