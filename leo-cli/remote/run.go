package remote

import (
	"os/exec"
	"bufio"
	"fmt"
	"github.com/oelmekki/leo/leo-cli/config"
)

/*
 * Execute commands on server, low level function for manipulating leo on server
 */
func Run( cmdStr string ) ( err error ) {
	cfg, err := config.Read()
	if err != nil { return }

	connString := "leo-deploy@" + cfg.ServerName
	fmt.Printf( "Running on %s : %s\n", connString, cmdStr )
	cmd := exec.Command( "ssh", "-o", "PasswordAuthentication=false", connString, cmdStr )
	cmdReader, err := cmd.StdoutPipe()
	if err != nil { return fmt.Errorf( "Can't execute command : %v", err ) }

	scanner := bufio.NewScanner( cmdReader )
	go func() {
		for scanner.Scan() {
			fmt.Println( scanner.Text() )
		}
	}()

	err = cmd.Start()
	if err != nil { return fmt.Errorf( "Can't execute command : %v", err ) }

	err = cmd.Wait()
	if err != nil { return fmt.Errorf( "Can't execute command : %v", err ) }

	return
}
