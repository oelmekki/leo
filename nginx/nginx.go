package nginx

import (
	"os/exec"
	"fmt"
)

/*
 * Load new nginx config, provided it's sane
 */
func Reload() ( err error ) {
	if err = CheckConfig() ; err != nil { return }

	err = exec.Command( "/etc/init.d/nginx", "reload" ).Run()
	if err != nil {
		err = exec.Command( "/etc/init.d/nginx", "start" ).Run()
		if err != nil { return fmt.Errorf( "Can't reload nginx." ) }
	}

	return
}

/*
 * Like Reload(), but using sudo
 */
func SudoReload() ( err error ) {
	fmt.Println( "reloading nginx..." )
	out, err := exec.Command( "sudo", "/usr/sbin/nginx", "-s", "reload" ).CombinedOutput()
	fmt.Println( string( out ) )
	if err != nil { return fmt.Errorf( "Error while reloading nginx: %v", err ) }

	return
}

func CheckConfig() ( err error ) {
	err = exec.Command( "/etc/init.d/nginx", "configtest" ).Run()
	if err != nil { return fmt.Errorf( "Woops: nginx says its config is broken. Please check." ) }

	return
}
