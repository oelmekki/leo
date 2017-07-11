package letsencrypt

import (
	"text/template"
	"bytes"
	"os"
	"github.com/oelmekki/leo/nginx"
	"github.com/oelmekki/leo/config"
)

/*
 * Create the temporary nginx rules which allow us to answer
 * Let's Encrypt challenge.
 */
func writeTemporaryNginxConf( appName string ) ( err error ) {
	filename := config.AppDir( appName ) + "/letsencrypt.nginx.conf"

	file, err := os.Create( filename )
	if err != nil { return }

	config, err := temporaryConf( appName )
	if err != nil { return }

	if _, err = file.WriteString( config ) ; err != nil { return }
	if err = nginx.SudoReload() ; err != nil { return }

	return
}

type ConfigParams struct {
	AppDir string
}

func temporaryConf( appName string ) ( content string, err error ) {
	var configB bytes.Buffer
	configParams := ConfigParams{ AppDir: config.AppDir( appName ) }

	tmpl, err := template.New( "temporaryNginxConfig" ).Parse( temporaryConfig )
	if err != nil { return }

	err = tmpl.Execute( &configB, configParams )
	if err != nil { return }

	content = configB.String()

	return
}

var temporaryConfig = `
location /.well-known/acme-challenge/ {
	root {{.AppDir}}/letsencrypt/www/;
}
`
