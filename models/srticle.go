package models

import "time"

type Article struct {
	// 文章的标题
	Title string

	// 作者姓名
	Author string

	// 创建时间
	CreatedAt time.Time `toml:"created_at"`

	// 最后更新时间
	UpdatedAt time.Time `toml:"updated_at"`

	// 标签
	Tags []string

	// 所属分类的名称
	Category string

	// 头部图片 URL 地址
	HeadImg string `toml:"head_img"`

	// 作者的个人主页
	HomePage string `toml:"home_page"`

	// 简短的描述
	Description string

	// 文章主题内容， markdown
	Body string

	// 文章在服务器上的文件路由
	Path string
}
