package letsencrypt

import (
	"bufio"
	"os"
	"fmt"
	"github.com/oelmekki/leo/leo-cli/env"
)

func askForEmail() ( err error ) {
	reader := bufio.NewReader( os.Stdin )
	fmt.Println( "Please provide an email address for letsencrypt (expiration notification): " )
	fmt.Printf( "> " )
	email, _ := reader.ReadString( '\n' )

	err = env.Set( "LETS_ENCRYPT_EMAIL", email )
	if err != nil { return }

	return
}
