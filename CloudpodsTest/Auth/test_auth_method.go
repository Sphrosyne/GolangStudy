package auth

import "fmt"

var (
	AUTH_METHODS = []string{AUTH_METHOD_PASSWORD, AUTH_METHOD_TOKEN, AUTH_METHOD_AKSK, AUTH_METHOD_CAS}
)

const (
	AUTH_METHOD_PASSWORD = "password"
	AUTH_METHOD_TOKEN    = "token"
	AUTH_METHOD_AKSK     = "aksk"
	AUTH_METHOD_CAS      = "cas"
	AUTH_METHOD_SAML     = "saml"
	AUTH_METHOD_OIDC     = "oidc"
	AUTH_METHOD_OAuth2   = "oauth2"
	AUTH_METHOD_VERIFY   = "verify"
)

func authMethod(method string) {
	for i := range AUTH_METHODS {
		if AUTH_METHODS[i] == method {
			fmt.Println("-----", byte(i)+1)
		}
	}
}
