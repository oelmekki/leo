package remote

import (
	"os/exec"
	"fmt"
	"github.com/oelmekki/leo/leo-cli/config"
)

/*
 * Send file to server in app dir, used for making docker-compose and nginx
 * configuration available.
 */
func Upload( src, dest string ) ( err error ) {
	cfg, err := config.Read()
	if err != nil { return }

	connString := "leo-deploy@" + cfg.ServerName
	dest = connString + ":apps/" + cfg.AppName + "/" + dest
	fmt.Printf( "Copying %s into %s\n", src, dest )
	out, err := exec.Command( "scp", src, dest ).CombinedOutput()
	fmt.Println( string( out ) )
	if err != nil {
		return fmt.Errorf( "Can't upload to " + connString + ". Did you install your public key there?" )
	}

	return
}
