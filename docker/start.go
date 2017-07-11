package docker

import (
	"os"
	"fmt"
	"github.com/oelmekki/leo/config"
)

func Start( appName string ) ( err error ) {
	if err = os.Chdir( config.AppDir( appName ) ) ; err != nil { return }
	result, err := DockerCompose( "start" )
	fmt.Println( result )
	return
}
