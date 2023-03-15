package authentication

import (
	_ "context"

	"golang.org/x/oauth2"
)

type Authentication_data struct {
	Config *oauth2.Config
	Token  *oauth2.Token
	
}
