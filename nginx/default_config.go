package nginx

import (
	"text/template"
	"bytes"
	"github.com/oelmekki/leo/config"
)

type ConfigParams struct {
	AppName string
	AppDir  string
}

func DefaultConfig( appName string ) ( content string, err error ) {
	var configB bytes.Buffer
	configParams := ConfigParams{ AppName: appName, AppDir: config.AppDir( appName ) }

	tmpl, err := template.New( "defaultNginxConfig" ).Parse( defaultConfig )
	if err != nil { return }

	err = tmpl.Execute( &configB, configParams )
	if err != nil { return }

	content = configB.String()

	return
}

var defaultConfig = `
server {
  listen      [::]:80;
  listen      80;
  server_name {{.AppName}};
  access_log  {{.AppDir}}/nginx-access.log;
  error_log   {{.AppDir}}/nginx-error.log;

	include {{.AppDir}}/letsencrypt.nginx.conf;

  location    / {
    gzip on;
    gzip_min_length  1100;
    gzip_buffers  4 32k;
    gzip_types    text/css text/javascript text/xml text/plain text/x-component application/javascript application/x-javascript application/json application/xml  application/rss+xml font/truetype application/x-font-ttf font/opentype application/vnd.ms-fontobject image/svg+xml;
    gzip_vary on;
    gzip_comp_level  6;

    proxy_pass  http://{{.AppName}}-upstream;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
    proxy_set_header Host $http_host;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header X-Forwarded-For $remote_addr;
    proxy_set_header X-Forwarded-Port $server_port;
    proxy_set_header X-Request-Start $msec;
  }
}

# Uncomment this block once you ran letsencrypt.
#
# Beware! You need to replace YOUR_FIRST_DOMAIN_NAME with an actual value.
#
# server {
# 	listen      [::]:443 ssl http2;
# 	listen      443 ssl http2;
# 	server_name central.el-mekki.com; 
# 	access_log  /home/leo-deploy/apps/{{.AppName}}/nginx-access.log;
# 	error_log   /home/leo-deploy/apps/{{.AppName}}/nginx-error.log;
# 
# 	# TODO be sure to replace YOUR_FIRST_DOMAIN_NAME, here
# 	ssl_certificate     /home/leo-deploy/apps/{{.AppName}}/letsencrypt/certs/live/YOUR_FIRST_DOMAIN_NAME/fullchain.pem;
# 	ssl_certificate_key /home/leo-deploy/apps/{{.AppName}}/letsencrypt/certs/live/YOUR_FIRST_DOMAIN_NAME/privkey.pem;
# 	ssl_protocols       TLSv1 TLSv1.1 TLSv1.2;
# 
# 	keepalive_timeout   70;
# 
# 	location    / {
# 		gzip on;
# 		gzip_min_length  1100;
# 		gzip_buffers  4 32k;
# 		gzip_types    text/css text/javascript text/xml text/plain text/x-component application/javascript application/x-javascript application/json application/xml  application/rss+xml font/truetype application/x-font-ttf font/opentype application/vnd.ms-fontobject image/svg+xml;
# 		gzip_vary on;
# 		gzip_comp_level  6;
# 
# 		proxy_pass  http://{{.AppName}}-upstream;
# 		proxy_http_version 1.1;
# 		proxy_set_header Upgrade $http_upgrade;
# 		proxy_set_header Connection "upgrade";
# 		proxy_set_header Host $http_host;
# 		proxy_set_header X-Forwarded-Proto $scheme;
# 		proxy_set_header X-Forwarded-For $remote_addr;
# 		proxy_set_header X-Forwarded-Port $server_port;
# 		proxy_set_header X-Request-Start $msec;
# 	}
# }

include {{.AppDir}}/upstream.conf;
`
