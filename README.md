# ForestBlog

> ForestBlog 是基于 go 语言开发的，不依赖第三方 go 库，适合用来学习和展示 markdown 文档的精美博客。


示例： [xusenlin.com](http://xusenlin.com) （个人博客，正在使用）

源码： [github.com/xusenlin/ForestBlog](https://github.com/xusenlin/ForestBlog)

--- 

- 请将你的博客文档克隆到ForestBlog的resources下，
ForestBlog会6个小时自动根据app.json的documentPath key切换到你的博客目录下执行git pull 命令来更新你的文章,
所以正确配置documentPath指向你的博客文档很关键。

- 还有，你的博客文档目录里面最少需要assets、content目录和About.md和Works.md文件。
content目录下的一级目录代表一个分类，如果有多个子级目录也不会产生分类,子级的文档也会属于第一级的分类。
如下：
```
    |-- assets       //博客静态文件
    |-- content
    |   |-- GOLANG   //分类目录
    |       |-- GOLANG基础   //  子分类目录
    |       |--- ForestBlog使用文档.md   
    |   |-- 其他分类
    |       |--- xxx.md
    |-- About.md    //关于
    |-- Works.md    //作品 根
    
```

TODO
[] 根目录可以添加其他文件生成导航
[] 支持配置主题
[] 支持搜索

## 优点 ##

1. 速度  ---没有什么依赖，得益于GOLANG的运行速度，部署在阿里云的博客平均响应在50毫秒内。
2. 迁移方便  ---GOLANG交叉编译可以方便的发布二进制文件到不同的操作系统，执行二进制文件并克隆博客文件即可运行你的博客。
