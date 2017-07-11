package letsencrypt

import (
	"fmt"
	"os/exec"
	"math/rand"
	"regexp"
	"strconv"
	"github.com/oelmekki/leo/nginx"
)

/*
 * To get a certificate from let's encrypt, we need to run certbot allowing
 * it to create a file that will be retrieved by Let's encrypt servers
 * to confirm we have control over the domain we want a certificate for.
 */
func GetCertificate( appName string ) ( err error ) {
	domainNames, err := nginx.FindDomainNamesFor( appName )
	if err != nil { return }

	email, err := checkLetsEncryptMail( appName )
	if err != nil { return }
	fmt.Printf( "email: %s\n", email )

	if err = performRetrieval( appName, email, domainNames ) ; err != nil { return }
	if err = addCronJob( appName ) ; err != nil { return }

	return
}

func performRetrieval( appName, email string, domainNames []string ) ( err error ) {
	if err = writeTemporaryNginxConf( appName ) ; err != nil { return }
	defer flushTemporaryNginxConf( appName )

	if err = runCertbot( appName, email, domainNames ) ; err != nil { return }

	fmt.Println( appName + " is now certified to the world!" )

	return
}

/*
 * Create a crontab to renew letsencrypt. It's ran every 15 days, as per letsencrypt
 * recommendation (it's valid 3 months, but we're supposed to renew every day
 * in case "we accendently deleted your certificates, lol").
 */
func addCronJob( appName string ) ( err error ) {
	out, err := exec.Command( "crontab", "-l" ).CombinedOutput()
	if err != nil { out = []byte( "" ) }
	jobFinder := regexp.MustCompile( `leo\s+letsencrypt\s+` + appName )
	matches := jobFinder.Find( out )
	if len( matches ) == 0 {
		hours := rand.Intn( 24 )
		minutes := rand.Intn( 59 )
		out, err = exec.Command( "bash", "-c", `(crontab -l ; echo "` + strconv.Itoa( minutes ) + ` ` + strconv.Itoa( hours ) + ` */1 * * leo letsencrypt ` + appName + `") 2>&1 | grep -v "no crontab" | sort | uniq | crontab -` ).CombinedOutput()
		if err != nil { return fmt.Errorf( "Can't add cron job :\n%s", out ) }
	}
	return
}
