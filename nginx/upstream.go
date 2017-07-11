package nginx

import (
	"text/template"
	"bytes"
)

type UpstreamParams struct {
	AppName string
	Ip      string
}

func Upstream( appName, ip string ) ( upstream string, err error ) {
	var upstreamB bytes.Buffer
	upstreamParams := UpstreamParams{ AppName: appName, Ip: ip }

	tmpl, err := template.New( "upstreamNginxConfig" ).Parse( upstreamConfig )
	if err != nil { return }

	err = tmpl.Execute( &upstreamB, upstreamParams )
	if err != nil { return }

	upstream = upstreamB.String()

	return
}

var upstreamConfig = `
upstream {{.AppName}}-upstream {
	server {{.Ip}}:5000;
}
`
