package authentication

import (
	_ "context"
)

type Authentication interface {
	StartAuthenticateServer()
	Load_Info(clid string, clisec string)
	Return_Info() *Authentication_data
}
