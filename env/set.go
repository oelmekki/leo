package env

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"github.com/oelmekki/leo/config"
)

/*
 * Allow to add or update env variable that will be used in containers,
 * much like `heroku config:set` (the syntax differs, though).
 */
func Set( appName, variable, value string ) ( err error ) {
	_, err = Get( appName, variable )
	if err == nil {
		err = updateVariable( appName, variable, value )
	} else {
		err = addVariable( appName, variable, value )
	}

	return
}

/*
 * append variable and its value in env file
 */
func addVariable( appName, variable, value string ) ( err error ) {
	fmt.Println( "Adding new variable" )
	content, err := ioutil.ReadFile( config.AppDir( appName ) + "/env" )
	if err != nil { return }

	contentStr := fmt.Sprintf( "%s\n%s=%s\n", string( content ), variable, quotedEnvValue( value ) )

	err = ioutil.WriteFile( config.AppDir( appName ) + "/env", []byte( cleanup( contentStr ) ), 0644 )
	if err != nil { return }

	return
}

/*
 * replace value for already existing variable
 */
func updateVariable( appName, variable, value string ) ( err error ) {
	fmt.Println( "Updating existing variable" )
	content, err := ioutil.ReadFile( config.AppDir( appName ) + "/env" )
	if err != nil { return }

	replacer := regexp.MustCompile( `(?m:^` + regexp.QuoteMeta( variable ) + `=(.*?)$)`  )
	contentStr := replacer.ReplaceAllString( string( content ), variable + "=" + quotedEnvValue( value ) )

	err = ioutil.WriteFile( config.AppDir( appName ) + "/env", []byte( cleanup( contentStr ) ), 0644 )

	return
}


/*
 * We need to quote values with spaces for env file format
 */
func quotedEnvValue( value string ) ( quotedValue string ) {
	matcher := regexp.MustCompile( `\s+` )
	match := matcher.FindString( value )

	if len( match ) > 0 {
		matcher = regexp.MustCompile( `"` )
		match = matcher.FindString( value )

		if len( match ) > 0 {
			quotedValue = `'` + value + `'`
		} else {
			quotedValue = `"` + value + `"`
		}
	} else {
		quotedValue = value
	}

	return
}
