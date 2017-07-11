package deploy

import (
	"fmt"
	"regexp"
	"io/ioutil"
	"github.com/oelmekki/leo/leo-cli/config"
	"github.com/oelmekki/leo/leo-cli/remote"
)

func Run() ( err error ) {
	cfg, err := config.Read()
	if err != nil { return }

	if err = checkEnvInComposeFile() ; err != nil { return }

	if err = remote.Upload( "./nginx.conf", "nginx.conf" ) ; err != nil { return }
	if err = remote.Upload( "./docker-compose.prod.yml", "docker-compose.yml" ) ; err != nil { return }
	if err = remote.Run( "leo rotate " + cfg.AppName ) ; err != nil { return }

	fmt.Println( "up and running!" )

	return
}

func checkEnvInComposeFile() ( err error ) {
	content, err := ioutil.ReadFile( "./docker-compose.prod.yml" )
	if err != nil { return fmt.Errorf( "Missing docker-compose.prod.yml file : %v", err ) }

	envMatcher := regexp.MustCompile( `env_file` )
	match := envMatcher.Find( content )
	if len( match ) == 0 {
		return fmt.Errorf( "You need to source the env file using `env_file: ./env` in your container in docker-compose.prod.yml. Do it now, it does no harm and will prevent later headaches." )
	}

	return
}
