# cli interface for leo

This is the cli interface for [leo](https://github.com/oelmekki/leo), and probably what you'll be using on a daily basis.

Once leo is installed on server, leo-cli makes it easy to deploy and manage
your docker-compose based applications on your server.

Here is the normal flow :

1. you build an app using docker-compose
2. most probably, one of the images used is a custom image built on gitlab or dockerhub
3. you have a `docker-compose.prod.yml` file in your repos
4. after initializing leo, you have a nginx.conf as well
5. you use `leo deploy` to push those files and launch an app using them on production
6. on further deploy, containers are rotated doing zero downtime deploy
7. you can add letsencrypt certificate with a simple command
8. you can manage env variables through leo as well


## Install

Assuming your `$GOPATH` is properly set up and that `$GOPATH/bin` is in your
`$PATH`:

```
go get github.com/jwaldrip/odin
go get github.com/oelmekki/leo/leo-cli
```

## Expectations

* you MUSTN'T use host directory as volumes, as leo-deploy home is regularly recursively chown. Use volumes.
* the app name you use on init should works as dir name and as nginx upstream name. Be standard.
* you have a `web` service in your docker-compose file
* your web app listens on port 5000
* you don't define port mapping for `web` service in dockerfile (needed for rotating services, nginx will point to container ip directly)


## Example session

Let's say you have a `docker-compose.prod.yml` in a directory containing this:

```
version: '2'
services:
  web:
    image: my_web_image
    env_file: ./env  # this is required if you want to use `leo env`
  background:
    image: my_background_image
```

You first need to let leo know when it's supposed to deploy app, and what is its name:

```
leo-cli my_app my_server.com
```

This will create a `nginx.conf` file with sensible default, edit it to change the domain name in `server_name`.

When ready, deploy your application:

```
leo-cli deploy
``` 

Now, make a few changes to your web image. Running deploy again would rotate
the container to use the new image, with zero downtime deploy.

But you probably want to rotate the background image as well, here. By default,
leo only rotate the web container, so you need to give the list of services you
want to rotate:

```
leo-cli env:set LEO_ROTATE_SERVICES "web,background"
```

You can now deploy changes to both services:

```
leo-cli deploy
```

Want ssl? Easy enough:

```
leo-cli letsencrypt
```

Your default nginx configuration contains a block to use ssl. Uncomment it,
adjust certificate path and deploy changes.


## Usage

> Note : it's perfectly fine to alias `leo-cli` to `leo` if you don't intend
> to run leo server on your computer


### `leo-cli init <app name> <server address>`

Initialize the current repository as a leo project.

App name must be standard enough so it can be used as directory name and nginx
upstream name.

Server address is the domain of the server where leo is installed. You need ssh
access to the `leo-deploy` user on it.

Removing leo from project is easy enough, since it only adds a `nginx.conf` and
a `leo.conf` file.


### `leo-cli deploy`

Upload `nginx.conf`, `docker-compose.prod.yml`, and rotate (or start) services.


### `leo-cli run <service> <command>`

Run `command` in a one off container for `service` on server.


### `leo-cli env`

Displays current custom environment variables.


### `leo-cli env:set <name> <value>`

Set a new environment variable, which will be loaded in containers you set
source it from (using compose instruction `env_file`).

> I stress that: by default, environment file is not sourced anywhere and is
> just a convenience. You're responsible for sourcing it.

Note that there are a few differences with the familiar `heroku config:set` command:

* you can only set one variable in one command
* `FOO=bar` syntax is not supported, pass variable name as first parameter and value as second
* this will not automatically restart the application, just deploy it again


### `leo-cli env:del <name>`

Remove an environment variable.

Note that app is not restarted automatically after changing env file.


### `leo-cli letsencrypt`

Retrieve a ssl certificate valid for each domain name you've set in the first
`server_name` instruction of your nginx configuration file.

You'll be asked for an email address on first time. This is required by
letsencrypt and is used to send notifications about certificate expiration.

Once the certificate is generated, it will be automatically renewed through a
cron task.

You're responsible to set nginx to use the certificates, although most of
what you need is already contained as a commented block in the default generated
nginx configuration (you just need to change the certificate directory name,
which is the same as the first domain name you set on `server_name`).
