package setup

import (
	"os/exec"
	"fmt"
)

func ReclaimOwnership() ( err error ) {
	err = exec.Command( "sudo", ChownScriptPath ).Run()
	if err != nil { return fmt.Errorf( "Can't change permissions of leo directory: %v", err ) }
	return
}
