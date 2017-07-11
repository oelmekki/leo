package docker

import (
	"os"
	"os/exec"
)

func Login( host, username, password string ) ( err error ) {
	cmd := exec.Command( "docker", "login", "-u", username, "-p", password, host )
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	return
}
