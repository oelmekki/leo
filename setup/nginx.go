package setup

import (
	"os"
	"fmt"
	"github.com/oelmekki/leo/nginx"
)

func SetupNginx() ( err error ) {
	fmt.Println( "Setting up nginx..." )

	if err = addGlobalNginxConfig() ; err != nil { return }
	if err = nginx.Reload() ; err != nil { return }

	return
}

/*
 * The global config for nginx is responsible for loading
 * app specific configuration files.
 */
func addGlobalNginxConfig() ( err error ) {
	filename := "/etc/nginx/conf.d/leo.conf"
	fmt.Printf( "Installing %s...\n", filename )

	if _, err := os.Stat( filename ) ; err != nil {
		file, err := os.Create( filename )
		if err != nil { return err }

		_, err = file.WriteString( nginxConfig )
		if err != nil { return err }
	} else {
		fmt.Println( "Global Nginx configuration already exists. Skipping creation." )
	}

	return
}

/*
 * Who needs a webserver anyway?
 * (only removes leo's conf in nginx)
 */
func removeNginx() ( err error ) {
	fmt.Println( "Removing /etc/nginx/conf.d/leo.conf..." )
	if err = os.Remove( "/etc/nginx/conf.d/leo.conf" ) ; err != nil { return }

	return
}

var nginxConfig = `
include /home/leo-deploy/apps/*/nginx.conf;

server_tokens off;

ssl_session_cache shared:SSL:20m;
ssl_session_timeout 10m;

ssl_ciphers EECDH+AES128:RSA+AES128:EECDH+AES256:RSA+AES256:EECDH+3DES:RSA+3DES:!MD5;
`
