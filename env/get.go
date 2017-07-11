package env

import (
	"io/ioutil"
	"strings"
	"regexp"
	"fmt"
	"github.com/oelmekki/leo/config"
)

func Get( appName, key string ) ( value string, err error ) {
	content, err := ioutil.ReadFile( config.AppDir( appName ) + "/env" )
	if err != nil { return value, fmt.Errorf( "Can't read env file : %v", err ) }

	keyMatcher := regexp.MustCompile( `^` + key + `=["']?(.+?)["']?$` )

	for _, line := range strings.Split( string( content ), "\n" ) {
		match := keyMatcher.FindStringSubmatch( line )
		if len( match ) == 2 {
			return match[1], err
		}
	}

	err = fmt.Errorf( "Can't find variable %s in env file.", key )

	return
}
