package env

import (
	"regexp"
	"strings"
	"fmt"
	"github.com/oelmekki/leo/leo-cli/config"
	"github.com/oelmekki/leo/leo-cli/remote"
)

func Get( key string ) ( value string, err error ) {
	cfg, err := config.Read()
	if err != nil { return }

	content, err := remote.RunCapture( "leo env " + cfg.AppName )
	if err != nil { return }

	keyMatcher := regexp.MustCompile( `^` + key + `=["']?(.+?)["']?$` )

	for _, line := range strings.Split( string( content ), "\n" ) {
		match := keyMatcher.FindStringSubmatch( line )
		if len( match ) == 2 {
			return match[1], err
		}
	}

	err = fmt.Errorf( "Can't find variable %s", key )

	return
}
