package setup

import "fmt"

func Run() ( err error ) {
	if err = SetupChownScript() ; err != nil { return }
	if err = SetupUser() ; err != nil { return }
	if err = SetupNginx() ; err != nil { return }
	if err = Install() ; err != nil { return }

	fmt.Println( "Done!" )

	return
}
