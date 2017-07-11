package validator

import (
	"os"
	"os/user"
	"fmt"
	"github.com/oelmekki/leo/config"
)

/*
 * Check if environment is ready for leo installation
 */
func TestConfigSetup() ( err error ) {
	if err = TestNginx() ; err != nil { return }
	if err = TestDocker() ; err != nil { return }
	if err = TestRoot() ; err != nil { return }

	return
}

/*
 * Check if environment is ready for common usage
 */
func TestConfigRun() ( err error ) {
	if err = TestNginx() ; err != nil { return }
	if err = TestDocker() ; err != nil { return }
	if err = TestLeo() ; err != nil { return }

	return
}

/*
 * Some tasks need to be ran as privileged user
 */
func TestRoot() ( err error ) {
	currentUser, err := user.Current()
	if err != nil { return }

	if currentUser.Username != "root" {
		err = fmt.Errorf( "This command needs to be ran as root" )
	}

	return
}

/*
 * We use nginx as a global webserver on host
 *
 * Not containerizing it makes it easier for user to manually
 * edit it and manage non-leo apps.
 */
func TestNginx() ( err error ) {
	if _, err = os.Stat( "/etc/nginx/" ) ; err != nil {
		err = fmt.Errorf( "Can't find nginx configuration in /etc/nginx/. Is nginx installed?" )
		return
	}

	return
}

/*
 * We use docker to deploy apps
 *
 * We have no swarm support for now. Will that socket ever not be present?
 */
func TestDocker() ( err error ) {
	if _, err = os.Stat( "/var/run/docker.sock" ) ; err != nil {
		err = fmt.Errorf( "Can't find docker socket in /var/run/docker.sock. Is docker installed?" )
		return
	}

	return
}

/*
 * Leo manages a few config files for apps, mainly nginx and docker-compose config
 */
func TestLeo() ( err error ) {
	if _, err = os.Stat( config.LeoDir() ) ; err != nil {
		fmt.Println( "Seems like leo has not been setup yet. Please run:" )
		fmt.Println( "\n  sudo leo setup" )
		err = fmt.Errorf( "\nError: Can't find leo files in " + config.LeoDir() + "." )
		return
	}

	return
}
