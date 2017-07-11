package letsencrypt

import (
	"fmt"
	"github.com/oelmekki/leo/config"
	"github.com/oelmekki/leo/env"
)

/*
 * Let's Encrypt require to register an email address, so that we are
 * warn when certificate is about to expire. We'll expect user to
 * store this in their env file.
 */
func checkLetsEncryptMail( appName string ) ( email string, err error ) {
	email, err = env.Get( appName, "LETS_ENCRYPT_EMAIL" )
	if err != nil { return email, fmt.Errorf( "Can't find LETS_ENCRYPT_EMAIL variable. Please make sure to put it in %s/env\n%v\n", config.AppDir( appName ), err ) }
	return
}
