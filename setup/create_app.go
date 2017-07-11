package setup

import (
	"os"
	"fmt"
	"github.com/oelmekki/leo/config"
)

func CreateApp( name string ) ( err error ) {
	fmt.Println( name )

	if err = createRoot( name ) ; err != nil { return err }
	if err = createEnvFile( name ) ; err != nil { return err }
	if err = createUpstreamFile( name ) ; err != nil { return err }
	if err = createSSLTempConf( name ) ; err != nil { return err }
	if err = createSSLTempDir( name ) ; err != nil { return err }
	if err = ReclaimOwnership() ; err != nil { return }

	fmt.Println( "App created in " + config.AppDir( name ) + "/" )

	return
}

/*
 * This is the place where all configuration files will be placed for this app
 */
func createRoot( name string ) ( err error ) {
	if err = os.Mkdir( config.AppDir( name ), os.FileMode( 0775 ) ) ; err != nil {
		return fmt.Errorf( "Can't create app : %v", err )
	}

	return
}

/*
 * This file can be used to define environment variables for the application container
 */
func createEnvFile( name string ) ( err error ) {
	file, err := os.Create( config.AppDir( name ) + "/env" )
	if err != nil { return err }
	file.Close()

	return
}

/*
 * This file is internal and will to use to set container ip for nginx proxy
 */
func createUpstreamFile( name string ) ( err error ) {
	file, err := os.Create( config.AppDir( name ) + "/upstream.conf" )
	if err != nil { return err }
	file.Close()

	return
}

/*
 * Create file that will be used to temporary redirect letsencrypt challenge url
 * to the certbot container.
 *
 * It should be blank at all time and not edited by user (changes will be discarded
 * anyway).
 */
func createSSLTempConf( name string ) ( err error ) {
	file, err := os.Create( config.AppDir( name ) + "/letsencrypt.nginx.conf" )
	if err != nil { return err }
	file.Close()

	return
}

/*
 * Create directories used by certbot to put its files.
 *
 * `www` is used as public directory where certbot will place files to answer
 * Let's Encrypt challenge.
 *
 * `certs` is the directory where final certificates will be stored. You really
 * don't want to mess with that one.
 *
 */
func createSSLTempDir( name string ) ( err error ) {
	if err = os.Mkdir( config.AppDir( name ) + "/letsencrypt", 0777 ) ; err != nil { return }
	if err = os.Mkdir( config.AppDir( name ) + "/letsencrypt/www", 0777 ) ; err != nil { return }
	if err = os.Mkdir( config.AppDir( name ) + "/letsencrypt/certs", 0777 ) ; err != nil { return }
	return
}
