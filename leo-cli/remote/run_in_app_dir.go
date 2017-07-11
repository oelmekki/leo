package remote

import (
	"fmt"
)

/*
 * Helper to execute a command in the directory where app config is,
 * eg for running docker-compose commands.
 */
func RunInAppDir( appName, cmd string ) ( err error ) {
	return Run( fmt.Sprintf( `/bin/bash -c "cd ~/apps/%s && %s"`, appName, cmd ) )
}
