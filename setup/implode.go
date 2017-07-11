package setup

import (
	"os"
	"bufio"
	"fmt"
)

func Implode() ( err error ) {
	reader := bufio.NewReader( os.Stdin )
	fmt.Println(`WARNING!!! If you remove leo, you will lose all your apps volumes and configuration, from /home/leo-deploy.`)
	fmt.Println(`If that's really what you want to do, type exactly "yes":`)
	text, _ := reader.ReadString('\n')

	if text != "yes\n" {
		fmt.Println( "Good call!" )
		fmt.Println( text )
		os.Exit(1)
		return // YOU NEVER KNOW
	}

	if err = removeUser() ; err != nil { return }
	if err = removeNginx() ; err != nil { return }

	fmt.Println( "Done, see you around!" )

	return
}
