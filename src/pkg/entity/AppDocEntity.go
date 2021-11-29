package entity

type AppDoc struct {
	CompanyName  string `bson:"company_name" json:"company_name"`
	AppName      string `bson:"app_name" json:"app_name"`
	AppVersion   string `bson:"app_version" json:"app_version"`
	Domain       string `bson:"domain" json:"domain"`
	EmailAddress string `bson:"email_address" json:"email_address"`
	IpAddress    string `bson:"ip_address" json:"ip_address"`
	Url          string `bson:"url" json:"url"`
	Country      string `bson:"country" json:"country"`
}
