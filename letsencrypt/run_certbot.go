package letsencrypt

import (
	"fmt"
	"github.com/oelmekki/leo/docker"
	"github.com/oelmekki/leo/config"
	"github.com/oelmekki/leo/setup"
)

func runCertbot( appName, email string, domainNames []string ) ( err error ) {
	wwwPath := config.AppDir( appName ) + "/letsencrypt/www"
	certsPath := config.AppDir( appName ) + "/letsencrypt/certs"

	params := []string{
		"run",
		"--rm",
		"-v", wwwPath + ":/var/www/letsencrypt",
		"-v", certsPath + ":/etc/letsencrypt",
		"certbot/certbot",
		"certonly",
		"--agree-tos",
		"--non-interactive",
		"--expand",
		"-m", email,
		"--webroot",
		"-w", "/var/www/letsencrypt",
	}

	for _, domainName := range domainNames {
		params = append( params, "-d" )
		params = append( params, domainName )
	}

	result, err := docker.Docker( params... )
	if err != nil { return }
	fmt.Println( result )

	if err = setup.ReclaimOwnership() ; err != nil { return }

	return
}
