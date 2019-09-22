package config


type userConfig struct {
	SiteName string `json:"siteName"`

	Author string `json:"author"`

	Icp string `json:"icp"`

	TimeLayout string `json:"timeLayout"`

	Port int `json:"port"`

	WebHookSecret string `json:"webHookSecret"`

	PageSize int `json:"pageSize"`

	DescriptionLen int `json:"descriptionLen"`

	DocumentPath string `json:"documentPath"`

	HtmlKeywords string `json:"htmlKeywords"`

	HtmlDescription string `json:"htmlDescription"`

	CategoryListFileNumber int `json:"categoryListFileNumber"`

	ThemeColor string `json:"themeColor"`

	ThemeOption []string `json:"themeOption"`

	DashboardEntrance string `json:"dashboardEntrance"`
}


