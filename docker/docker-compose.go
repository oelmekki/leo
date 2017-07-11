package docker

import (
	"os/exec"
	"fmt"
)

func DockerCompose( args ...string ) ( result string, err error ) {
	out, err := exec.Command( "docker-compose", args... ).CombinedOutput()
	result = string( out )
	if err != nil { fmt.Println( result ) ; return }

	return
}
