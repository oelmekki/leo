package setup

import (
	"os"
	"os/exec"
	"os/user"
	"fmt"
	"github.com/oelmekki/leo/config"
)

func SetupUser() ( err error ) {
	if err = createUser() ; err != nil { return }
	if err = createFiles() ; err != nil { return }

	return
}

/*
 * User will allow to add ssh public keys and connect to manipulate
 * leo on server side.
 */
func createUser() ( err error ) {
	fmt.Println( "Creating user..." )

	_, err = user.Lookup( "leo-deploy" )
	if err == nil {
		fmt.Println( "User leo-deploy already exists" )
		return
	}

	err = exec.Command( "useradd", "-m", "-G", "docker", "leo-deploy" ).Run()
	if err != nil { return fmt.Errorf( "Can't create leo-deploy user: %v", err ) }

	return
}

/*
 * This directory will hold apps configuration, volumes, etc.
 */
func createFiles() ( err error ) {
	fmt.Println( "Creating apps dir..." )

	if _, err = os.Stat( config.LeoDir() ) ; err != nil {
		err = fmt.Errorf( "Can't create files: user leo-deploy exists but not its home directory, " + config.LeoDir() + ". Something is incredibly wrong." )
		return
	}

	/* there's a catch, here. Just try and find it. */
	if _, err := os.Stat( config.LeoDir() + "/apps/" ) ; err != nil {
		if err = os.Mkdir( config.LeoDir() + "/apps/", os.FileMode( 0755 ) ) ; err != nil { return err }
	}

	if err = ReclaimOwnership() ; err != nil { return }

	return
}

/*
 * Our contribution to the apocalypse
 */
func removeUser() ( err error ) {
	fmt.Println( "Removing user leo-deploy..." )
	_, err = user.Lookup( "leo-deploy" )
	if err == nil {
		out, err := exec.Command( "userdel", "-r", "leo-deploy" ).CombinedOutput()
		if err != nil { return fmt.Errorf( "%s\n\nCan't remove leo-deploy user: %v", out, err ) }
	}

	return
}
