package env

import (
	"regexp"
)

/*
 * Remove empty lines from given content
 */
func cleanup( content string ) ( newContent string ) {
	replacer := regexp.MustCompile( `\n\s*\n` )
	newContent = replacer.ReplaceAllString( content, "\n" )

	return
}
