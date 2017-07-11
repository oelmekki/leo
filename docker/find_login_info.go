package docker

import (
	"github.com/oelmekki/leo/env"
	"fmt"
)

func FindLoginInfo( appName string ) ( host, username, password string, err error ) {
	host, err = env.Get( appName, "DOCKER_LOGIN_HOST" )
	if err != nil { err = fmt.Errorf( "No host info for login : %v", err ) ; return }

	username, err = env.Get( appName, "DOCKER_LOGIN_USERNAME" )
	if err != nil { err = fmt.Errorf( "No username info for login : %v", err ) ; return }

	password, err = env.Get( appName, "DOCKER_LOGIN_PASSWORD" )
	if err != nil { err = fmt.Errorf( "No password info for login : %v", err ) ; return }

	return
}
