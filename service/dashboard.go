package service

import "github.com/xusenlin/go_blog/config"

func SetThemeColor(index int)  {
	config.Cfg.ThemeColor = config.Cfg.ThemeOption[index]
}
