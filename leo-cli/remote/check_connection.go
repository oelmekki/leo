package remote

import (
	"os/exec"
	"fmt"
)

/*
 * Allow to make sure we can ssh to server before trying anything
 */
func CheckConnection( serverName string ) ( err error ) {
	connString := "leo-deploy@" + serverName
	err = exec.Command( "ssh", "-o", "PasswordAuthentication=false", connString, "ls"  ).Run()
	if err != nil {
		return fmt.Errorf( "Can't connect to " + connString + ". Did you install your public key there?" )
	}

	return
}
