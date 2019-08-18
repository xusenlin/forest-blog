# ForestBlog

> ForestBlog 是基于 go 语言开发的，不依赖第三方 go 库，适合用来学习和展示 markdown 文档的精美博客。


示例： [xusenlin.com](http://xusenlin.com) （个人博客，正在使用）

源码： [github.com/xusenlin/ForestBlog](https://github.com/xusenlin/ForestBlog)

---  
## 使用

- 请将你的博客文档克隆到ForestBlog的resources下，
ForestBlog会6个小时自动根据app.json的documentPath key切换到你的博客目录下执行git pull 命令来更新你的文章,
所以正确配置documentPath指向你的博客文档很关键。

- 还有，你的博客文档目录里面最少需要assets、content目录和About.md和Works.md文件(未来完成todo这个将不再是强制性的)。
content目录下的一级目录代表一个分类，如果有多个子级目录也不会产生分类,子级的文档也会属于第一级的分类。
如下：
```
    |-- assets       //博客静态文件，存放一下图片资源，方便显示到文档里
    |-- content
    |   |-- GOLANG   //分类目录
    |       |-- GOLANG基础   //  子分类目录
    |       |--- ForestBlog使用文档.md   
    |   |-- 其他分类
    |       |--- xxx.md
    |-- About.md    //关于
    |-- Works.md    //作品 根
    
```
- 文章
在文章的开头可以配置一些文章的属性，我觉得这个应该是必须的，最少也应该有一个日期，否则默认使用文件创建的时间，
但是当你迁移文档时间就会改变导致文章排序到前面，添加上时间则不会有这个问题。
```
    ```json
    {
        "date":"2019.01.02 14:33"
    }
    ```
```
我会根据关键字```json来解析，不用担心这个会显示到文章内容里面，我在解析的时候就将它去掉了，同时会生成缓存，直到文章被更新。
最后，markdown文章编辑器推荐Typora。

## TODO
- [ ] 1.移动端更好的适配
- [ ] 2.根目录可以添加其他文件生成导航
- [ ] 3.支持主题和添加多套主题
- [ ] 4.支持搜索
- [ ] 5.tags的展示

## 优点 ##

1. 响应迅速  ---没有什么依赖，得益于GOLANG的运行速度，部署在阿里云的博客平均响应在50毫秒内。
2. 迁移方便  ---GOLANG交叉编译可以方便的发布二进制文件到不同的操作系统，执行二进制文件并克隆博客文件即可运行你的博客。
3. 小巧精美  ---非常简单的代码方便学习和改造，即使有一天你厌倦ForestBlog，你的文章也能很好的迁移和阅读。
