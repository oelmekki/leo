package remote

import (
	"os"
	"os/exec"
	"fmt"
	"github.com/oelmekki/leo/leo-cli/config"
)

/*
 * Execute commands on server, low level function for manipulating leo on server
 */
func RunInteractive( service, command string ) ( err error ) {
	cfg, err := config.Read()
	if err != nil { return }

	cmdStr := fmt.Sprintf( "leo run %s %s %s", cfg.AppName, service, command )

	connString := "leo-deploy@" + cfg.ServerName
	fmt.Printf( "Running on %s : %s\n", connString, cmdStr )
	cmd := exec.Command( "ssh", "-t", "-o", "PasswordAuthentication=false", connString, cmdStr )
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil { return fmt.Errorf( "Can't execute command : %v", err ) }

	return
}
