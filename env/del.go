package env

import (
	"io/ioutil"
	"regexp"
	"github.com/oelmekki/leo/config"
)

/*
 * Allow to remove env variable from env file
 */
func Del( appName, variable string ) ( err error ) {
	content, err := ioutil.ReadFile( config.AppDir( appName ) + "/env" )
	if err != nil { return }

	replacer := regexp.MustCompile( `(?m:^` + regexp.QuoteMeta( variable ) + `=(.*?)$)`  )
	contentStr := replacer.ReplaceAllString( string( content ),  "" )

	err = ioutil.WriteFile( config.AppDir( appName ) + "/env", []byte( cleanup( contentStr ) ), 0644 )

	return
}
