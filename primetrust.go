package primetrust

import (
	"encoding/base64"
	"fmt"
)

const (
	Version             = "1.0.0"
	SandboxAPIPrefix    = "https://sandbox.primetrust.com/v2"
	ProductionAPIPrefix = "https://api.primetrust.com/v2"
)

var _apiPrefix string
var _authHeader string

func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func Init(sandbox bool, login string, password string) {
	if sandbox {
		_apiPrefix = SandboxAPIPrefix
	} else {
		_apiPrefix = ProductionAPIPrefix
	}
	_authHeader = fmt.Sprintf("Basic %s", basicAuth(login, password))
}
