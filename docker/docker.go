package docker

import (
	"os/exec"
	"fmt"
)

func Docker( args ...string ) ( result string, err error ) {
	out, err := exec.Command( "docker", args... ).CombinedOutput()
	result = string( out )
	if err != nil { fmt.Println( result ) ; return }

	return
}
