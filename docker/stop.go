package docker

import (
	"os"
	"fmt"
	"github.com/oelmekki/leo/config"
)

func Stop( appName string ) ( err error ) {
	if err = os.Chdir( config.AppDir( appName ) ) ; err != nil { return }
	result, err := DockerCompose( "stop" )
	fmt.Println( result )
	return
}
