# Leo, fleet deployment and orchestration

Leo is a webserver manager that makes it easy to do zero downtime deployment
with docker, docker-compose and nginx.

What it does, mainly, is to start new containers, make nginx config point to
them, and shut down old containers after a short delay (that is, rotating
containers).

Its point is to manage production containers, staying the closest possible to
core tools. if leo breaks, the whole production stack is just docker-compose
and nginx, so you can manage it without leo, in urgency.

> Note : this is the documentation for server side leo.
> You may want to look at documentation of [leo-cli](https://github.com/oelmekki/leo/tree/master/leo-cli) instead.


## What leo is not

Leo is not an app you should use as is in your company. I made it for my own
needs and I do not guarantee it will fit your needs, nor do I guarantee I will
continue to maintain it : I may decide an other tool is better any time.

Why did I publish it, then, you'll ask? Finding a way to do zero downtime
deployment with docker-compose was hard enough for me, so I decided to publish
leo's sources so that people wanting to do the same thing have an example about
how to do it. It may also be a good bootstrap base for people wanting to build
their own tool.


## Prerequisites

You need to install on your host system:

* docker
* docker-compose
* nginx


## Magic-free

Leo mainly relies on well known tools (docker, docker-compose and nginx) to do
its tasks. Here is what leo is doing specifically:

1. it runs docker-compose on server
2. it installs nginx.conf for app in `/home/leo-deploy/apps/<app_name>/` and edit its upstream list to list web containers
3. on new deploy, it scales web process up, rotate upstreams in nginx conf, reload nginx and remove old containers
4. optionally, it automates letsencrypt certificate retrieval and installation in `/home/leo-deploy/apps/<app_name>/`

This means that at any point, if something goes wrong, you can manage nginx and docker-compose manually.


## Install

Assuming your `$GOPATH` is properly set up and that `$GOPATH/bin` is in your
`$PATH`:

```
go get github.com/jwaldrip/odin
go get github.com/oelmekki/leo

sudo leo setup
```


This will create a `/home/leo-deploy` directory and add a
`/etc/nginx/conf.d/leo.conf`. Those are the only two things you have to delete
if you want to purge leo's installation.

It is assumed that your nginx configuration is loading all files in
`/etc/nginx/conf.d/` (this is the default on ubuntu). If not, add this in your
nginx configuration, in the `http {}` block:

```
include /etc/nginx/conf.d/*.conf;
```

Additionally, you need to allow `leo-deploy` user to reload nginx and change
certificates permissions. Add in sudoers file:

```
leo-deploy ALL=(ALL) NOPASSWD: /usr/sbin/nginx -s reload
leo-deploy ALL=(ALL) NOPASSWD: /usr/local/bin/chown_leo_deploy
```

Don't forget to add your ssh key in `/home/leo-deploy/.ssh/authorized_keys` to
be able to ssh to leo-deploy user after that (needed to use leo-cli).


## Expectations for managed apps

* you MUSTN'T use host directory as volumes, as leo-deploy home is regularly recursively chown. Use volumes.
* the app name you use on init should works as dir name and as nginx upstream name. Be standard.
* you have a `web` service in your docker-compose file
* your web app listens on port 5000
* you don't define port mapping for `web` service in dockerfile (needed for rotating services, nginx will point to container ip directly)

## Usage

Note that all commands can be run from anywhere, no need to be in a specific directory.


### `leo setup`

Install leo on this server.

This creates a `leo-deploy` user, install `chown_leo_deploy` management script and setup nginx.

Don't forget to add sudo rules:

```
leo-deploy ALL=(ALL) NOPASSWD: /usr/sbin/nginx -s reload
leo-deploy ALL=(ALL) NOPASSWD: /usr/local/bin/chown_leo_deploy
```

`chown_leo_deploy` is called on various occasion to reclaim ownership for leo
home directory. This is needed for example because of letsencrypt management,
which mount host directory as volume to retrieve certificates. We need to make
sure leo owns everything if we want to be able to remove app.

That's also the reason why your own volumes must be data containers instead of
host directories (if you use a host directory for uploaded files and we chown
it to user id 1002, for say, shit will probably happen).


### `leo create <app name>`

Create a new app. It will be created in `/home/leo-deploy/apps/<app name>/`.


### `leo start <app name>`

Start the application if it's not running yet.


### `leo stop <app name>`

Stop the application if it's running.


### `leo run <app name> <service> <command>`

Run `command` in a one off container of `service` from `app name`.

You can use this to run web console, for example.


### `leo rotate <app name>`

The very reason of leo existence. Perform a zero downtime deployment with
docker-compose.

What it does exactly is:

1. start a new container for each required service (only `web` by default,
can be configured through `LEO_ROTATE_SERVICES` variable in env file)
2. find the id of old containers and new containers
3. edit nginx upstream to redirect new requests to new container (`web` only)
4. wait for 10 sec
5. shutdown old containers

This is the reason why you can't assign a hardcoded port to your containers: we
need to have two containers living at the same time during rotate, and only one
can bind to a port (plus, this would probably conflict with other apps anyway).


### `leo env <app name>`

Display environment variables for `app name`.

> Note regarding env management : there's no magic here either, you can just
> edit `/home/leo-deploy/apps/*/env` if you prefer to. You need to restart
> your app to source changes.


### `leo env:set <app name> <name> <value>`

Set environment variable for `app name`.


### `leo env:del <app name> <name>`

Remove environment variable for `app name`.


### `leo letsencrypt <app name>`

Generate or renew ssl certificate for `app name`.

A certificate will be generate for all domains listed in the first
`server_name` instruction in your `nginx.conf` file.

You don't have to stop your app to do this, leo will create a new container and
temporarily edit nginx conf to route letsencrypt challenge request to it.

You do have to edit your nginx configuration to enable ssl. Certificates are in
`/home/leo-deploy/apps/<app name>/letsencrypt/certs/live/<your first domain name>/`.

`your first domain name` is the first domain name appearing in the first `server_name`
instruction.


### `leo implode`

Remove leo from server.

> Beware : this will wipe out all data. Use at your own risk!
