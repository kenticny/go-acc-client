package acc

import (
	pbConfig "github.com/kenticny/go-acc-client/pb/configuration"
)

type AccOptions struct {
	Appid       string
	Appsecret   string
	Host        string
	Environment pbConfig.Environment
}
