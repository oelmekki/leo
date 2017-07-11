package appinit

import (
	"encoding/json"
	"os"
	"fmt"
	"github.com/oelmekki/leo/leo-cli/config"
	"github.com/oelmekki/leo/leo-cli/remote"
	"github.com/oelmekki/leo/nginx"
)


/*
 * Install leo on server
 */
func Run( appName, serverName string ) ( err error ) {
	if err = checkDockerFile() ; err != nil { return }
	if err = createInitFile( appName, serverName ) ; err != nil { return }
	if err = remote.CheckConnection( serverName ) ; err != nil { return }
	if err = createNginxTemplate( appName ) ; err != nil { return }
	if err = remote.Run( "leo create " + appName ) ; err != nil { return }

	fmt.Println( "We're all set!" )

	return
}

/*
 * This config file will be uploaded on deploy
 */
func checkDockerFile() ( err error ) {
	_, err = os.Stat( "./docker-compose.prod.yml" )
	if err != nil { return fmt.Errorf( "Please add a docker-compose.prod.yml file in this directory" ) }
	return
}

/*
 * Store information about remote server
 */
func createInitFile( appName, serverName string ) ( err error ) {
	filename := "./leo.conf"

	if _, err := os.Stat( filename ) ; err != nil {
		file, err := os.Create( filename )
		if err != nil { return err }

		content := config.Remote{ AppName: appName, ServerName: serverName }
		body, _ := json.Marshal( content )
		_, err = file.Write( body )
		if err != nil { return err }
	} else {
		return fmt.Errorf( "./leo.conf already exists. Seems like this directory is already initialized." )
	}

	return
}

/*
 * Nginx configuration for remote server
 */
func createNginxTemplate( appName string ) ( err error ) {
	filename := "./nginx.conf"

	if _, err := os.Stat( filename ) ; err != nil {
		file, err := os.Create( filename )
		if err != nil { return err }

		config, err := nginx.DefaultConfig( appName )
		if err != nil { return err }

		_, err = file.WriteString( config )
		if err != nil { return err }
	} else {
		return fmt.Errorf( "./nginx.conf already exists. Seems like this directory is already initialized." )
	}

	return
}
