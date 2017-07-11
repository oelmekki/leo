package docker

import (
	"os"
	"os/exec"
	"github.com/oelmekki/leo/config"
)

func Run( appName, service, command string ) ( err error ) {
	if err = os.Chdir( config.AppDir( appName ) ) ; err != nil { return }

	cmd := exec.Command( "docker-compose", "run", service, "/bin/bash", "-c", command )
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil { return err }

	return
}
