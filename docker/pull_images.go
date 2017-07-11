package docker

import (
	"os"
	"os/exec"
	"fmt"
	"github.com/oelmekki/leo/config"
)

func PullImages( appName string ) ( err error ) {
	host, username, password, err := FindLoginInfo( appName )
	if err == nil {
		if err = Login( host, username, password ) ; err != nil { return }
	} else {
		fmt.Printf( "Won't attempt docker login : %v\n", err )
	}

	if err = os.Chdir( config.AppDir( appName ) ) ; err != nil { return }

	cmd := exec.Command( "docker-compose", "pull" )
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil { return fmt.Errorf( "Can't pull images : %v", err ) }

	return
}
