package env

import (
	"os/exec"
	"fmt"
	"github.com/oelmekki/leo/config"
)

func Show( appName string ) ( err error ) {
	out, err := exec.Command( "cat", config.AppDir( appName ) + "/env" ).CombinedOutput()
	fmt.Println( string( out ) )
	return
}
