package entities

type Config struct {
	AppName string `json: "appname"`
	Host    string `json: "host"`
	Port    int    `json: "port"`
	Static string `json:"static"`
}
