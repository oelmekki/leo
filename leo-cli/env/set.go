package env

import (
	"regexp"
	"github.com/oelmekki/leo/leo-cli/config"
	"github.com/oelmekki/leo/leo-cli/remote"
)

func Set( variable, value string ) ( err error ) {
	cfg, err := config.Read()
	if err != nil { return }

	if err = remote.Run( "leo env:set " + cfg.AppName + " " + variable + " " + quotedEnvValue( value ) ) ; err != nil { return }

	return
}

/*
 * We send a command string to server, so we need to make sure to quote values
 * containing spaces
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
