package letsencrypt

import (
	"os"
	"github.com/oelmekki/leo/nginx"
	"github.com/oelmekki/leo/config"
)

/*
 * Empty certbot nginx conf so we don't leave possibly sensible files world accessible
 */
func flushTemporaryNginxConf( appName string ) ( err error ) {
	filename := config.AppDir( appName ) + "/letsencrypt.nginx.conf"

	file, err := os.Create( filename )
	if err != nil { return err }

	if _, err = file.WriteString( "" ) ; err != nil { return }

	if err = nginx.SudoReload() ; err != nil { return }

	return
}
