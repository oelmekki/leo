package setup

import (
	"os"
	"io/ioutil"
	"fmt"
)

func Install() ( err error ) {
	fmt.Println( "Installing /usr/local/bin/leo..." )
	content, err := ioutil.ReadFile( os.Args[0] )
	if err != nil { return fmt.Errorf( "Can't read binary that we're currently running (wtf) : %v", err ) }

	err = ioutil.WriteFile( "/usr/local/bin/leo", content, 0755 )
	if err != nil { return fmt.Errorf( "Can't write binary to /usr/local/bin : %v", err ) }

	return
}
