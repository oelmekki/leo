package nginx

import (
	"fmt"
	"regexp"
	"io/ioutil"
	"github.com/oelmekki/leo/config"
)

/*
 * Retrieve domain name based on nginx configuration for given app.
 *
 * For now, it's very crude and retrieve only the first domain name. With time,
 * it should retrieve all domain names so that they can automatically get passed
 * to certbot.
 */
func FindDomainNamesFor( appName string ) ( domainNames []string, err error ) {
	conf, err := nginxConfForApp( appName )
	if err != nil { return }

	directiveMatcher := regexp.MustCompile( `(?m:^\s*server_name\s+(.*?);$)` )
	matches := directiveMatcher.FindStringSubmatch( conf )

	if len( matches ) > 1 {
		splitter := regexp.MustCompile( `\S+` )
		domainNames = splitter.FindAllString( matches[1], -1 )
	} else {
		err = fmt.Errorf( "Can't find any domain name in this nginx.conf. Is `server_name` directive present?" )
	}

	return
}

/*
 * Find content for nginx configuration for given app
 */
func nginxConfForApp( appName string ) ( conf string, err error ) {
	content, err := ioutil.ReadFile( config.AppDir( appName ) + "/nginx.conf" )
	if err != nil { return }
	conf = string( content )
	return
}
