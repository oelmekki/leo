package docker

/*
 * Strategy to rotate a service running through docker-compose:
 *
 * - pull new images
 * - scale the service up: docker-compose scale web=2
 * - wait for the new container to be ready
 * - update nginx upstream to point to new container
 * - reload nginx
 * - wait for old container to have processed its requests
 * - stop the previous container with docker: docker stop <app>_web_<container_number>
 * - remove the previous container with docker : docker rm <app>_web_<container_number>
 * - scale the service down : docker-compose scale web=1
 * 
 * Regarding the part about waiting for containers, it's just blind timeouts,
 * for now: `newContainerGraceTime` secs waiting for new container to boot,
 * `oldContainerGraceTime` secs waiting for old container to be done with its
 * old requests.
 *
 * So it's imperative your container takes less than `newContainerGraceTime`
 * secs to boot up to avoid downtime.
 */

import (
	"os"
	"os/exec"
	"time"
	"strings"
	"fmt"
	"github.com/oelmekki/leo/nginx"
	"github.com/oelmekki/leo/config"
	"github.com/oelmekki/leo/env"
)

/*
 * Those variables tell how long we wait before considering a new container
 * should have booted, and how long before considering an old container
 * should have finished processing active requests (in seconds).
 */
var newContainerGraceTime = 5
var oldContainerGraceTime = 5

/*
 * Perform a zero-downtime deploy of app
 *
 * The actual code update is done by pulling latest
 * image through docker-compose.
 */
func Rotate( appName string ) ( err error ) {
	if err = os.Chdir( config.AppDir( appName ) ) ; err != nil { return }

	if err = PullImages( appName ) ; err != nil { return }

	// TODO check nginx config
	// TODO check docker-compose config

	running, err := servicesRunning()
	if err != nil { return }

	if running {
		if err = rotateWeb( appName ) ; err != nil { return }

		serviceNames := servicesToRotate( appName )

		for _, serviceName := range serviceNames {
			if err = rotateService( serviceName ) ; err != nil { return err }
		}
	} else {
		if err = startServices( appName ) ; err != nil { return }
	}

	fmt.Println("rotated!")
	return
}

/*
 * Return the list of services the user asked to rotate.
 * We're ok if they provided none (we always rotate the "web" service).
 */
func servicesToRotate( appName string ) ( services []string ) {
	services = make( []string, 0 )
	servString, err := env.Get( appName, "LEO_ROTATE_SERVICES" )
	if err != nil { return services }

	if len( servString ) > 0 {
		for _, service := range strings.Split( servString, "," ) {
			services = append( services, service )
		}
	}

	return
}

func servicesRunning() ( running bool, err error ) {
	out, err := DockerCompose( "top" )
	if err != nil { return }

	running = out != ""
	return
}

func rotateWeb( appName string ) ( err error ) {
	fmt.Println( "Rotating web container..." )

	if _, err = DockerCompose( "scale", "web=2" ) ; err != nil { return }
	fmt.Println( "Waiting for new container..." )
	time.Sleep( time.Duration( newContainerGraceTime ) * time.Second )
	oldContainer, newContainer, err := containerIdsFor( "web" )
	if err != nil { return }

	if err = updateUpstream( appName, newContainer ) ; err != nil { return }
	if err = reloadNginx() ; err != nil { return }

	fmt.Println( "Waiting for old container to be done with requests..." )
	time.Sleep( time.Duration( oldContainerGraceTime ) * time.Second )

	out, err := Docker( "stop", oldContainer )
	if err != nil { fmt.Println( out ) ; return }

	out, err = Docker( "rm", oldContainer )
	if err != nil { fmt.Println( out ) ; return }

	out, err = DockerCompose( "scale", "web=1" )
	if err != nil { fmt.Println( out ) ; return }

	return
}

func rotateService( name string ) ( err error ) {
	fmt.Println( "Rotating " + name + " container..." )

	if _, err = DockerCompose( "scale", name + "=2" ) ; err != nil { return }
	fmt.Println( "Waiting for new container..." )
	time.Sleep( time.Duration( newContainerGraceTime ) * time.Second )
	oldContainer, _, err := containerIdsFor( name )
	if err != nil { return }

	fmt.Println( "Waiting for old container to be done with requests..." )
	time.Sleep( time.Duration( oldContainerGraceTime ) * time.Second )

	out, err := Docker( "stop", oldContainer )
	if err != nil { fmt.Println( out ) ; return }

	out, err = Docker( "rm", oldContainer )
	if err != nil { fmt.Println( out ) ; return }

	out, err = DockerCompose( "scale", name + "=1" )
	if err != nil { fmt.Println( out ) ; return }

	return
}

func startServices( appName string ) ( err error ) {
	if _, err = DockerCompose( "up", "-d" ) ; err != nil { return }

	return
}

/*
 * The output order for `docker-compose ps -q` is not guaranteed: sometime the oldest
 * container is listed first, sometime it's listed second.
 */
func containerIdsFor( service string ) ( oldContainer, newContainer string, err error ) {
	out, err := DockerCompose( "ps", "-q", service )
	if err != nil { return }

	lines := strings.Split( string( out ), "\n" )
	const template = "2006-01-02T15:04:05"

	out, err = Docker( "inspect", "-f", "{{.Created}}", lines[0] )
	if err != nil { return }
	firstDate, _ := time.Parse( template, strings.Split( string( out ), "." )[0] )

	out, err = Docker( "inspect", "-f", "{{.Created}}", lines[1] )
	if err != nil { return }
	secondDate, _ := time.Parse( template, strings.Split( string( out ), "." )[0] )

	if secondDate.Unix() > firstDate.Unix() {
		oldContainer = lines[0]
		newContainer = lines[1]
	} else {
		oldContainer = lines[1]
		newContainer = lines[0]
	}

	return
}

func updateUpstream( appName, container string ) ( err error ) {
	out, err := Docker( "inspect", "-f", "{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}", container )
	ip := strings.Split( string( out ), "\n" )[0]

	filename := config.AppDir( appName ) + "/upstream.conf"

	if _, err := os.Stat( filename ) ; err != nil { os.Remove( filename ) }

	file, err := os.Create( filename )
	if err != nil { return err }

	nginxConfig, err := nginx.Upstream( appName, ip )
	if err != nil { return err }

	_, err = file.WriteString( nginxConfig )
	if err != nil { return err }

	file.Close()

	return
}

func reloadNginx() ( err error ) {
	fmt.Println( "reloading nginx..." )
	out, err := exec.Command( "sudo", "/usr/sbin/nginx", "-s", "reload" ).CombinedOutput()
	fmt.Println( string( out ) )
	if err != nil { return fmt.Errorf( "Error while reloading nginx: %v", err ) }

	return
}
