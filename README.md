# ForestBlog

> ForestBlog 是基于 go 语言开发的，无第三方依赖的，用来展示 markdown 文档的博客。


示例： [xusenlin.com](http://xusenlin.com) （个人博客，正在使用）

源码： [github.com/xusenlin/ForestBlog](https://github.com/xusenlin/ForestBlog)

--- 

- 请将你的博客文档克隆到ForestBlog的resources下，ForestBlog会3个小时自动根据app.json的documentPath key切换到你的博客目录下执行git pull 命令来更新你的文章。所以正确配置documentPath指向你的博客文档很关键。


- 还有，你的博客文档目录里面最少需要content目录、About.md和Works.md文件。content目录下每一个目录代表一个分类。如下：

## 目录结构 ##

    |-- content
    |   |-- GOLANG   //分类目录
    |       |--- ForestBlog使用文档.md     
    |-- About.md    //关于
    |-- Works.md    //作品           
