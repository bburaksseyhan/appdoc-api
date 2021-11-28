package model

type AppDoc struct {
	CompanyName  string `json:"company_name"`
	AppName      string `json:"app_name"`
	AppVersion   string `json:"app_version"`
	Domain       string `json:"domain"`
	EmailAddress string `json:"email_address"`
	IpAddress    string `json:"ip_address"`
	Url          string `json:"url"`
	Country      string `json:"country"`
}
