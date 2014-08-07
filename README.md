Default Project [![Build Status](https://drone.io/github.com/elcct/amazingwebsite/status.png)](https://drone.io/github.com/elcct/amazingwebsite/latest)
===============

Provides essentials that most web applications need - MVC pattern and user authorisation that can be easily extended.

It consists of 3 core components:

- Goji - A web microframework for Golang - http://goji.io/
- Gorilla web toolkit sessions - cookie (and filesystem) sessions - http://www.gorillatoolkit.org/pkg/sessions
- mgo - MongoDB driver for the Go language - http://labix.org/mgo

# Dependencies

Default Project requires `Go`, `MongoDB` and few other tools installed.

Instructions below have been tested on `Ubuntu 14.04`.

## Installation

If you don't have `Go` installed, follow installation instructions described here: http://golang.org/doc/install

Then install remaining dependecies:

```
sudo apt-get install make git mercurial subversion bzr
```

MongoDB:

```
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv 7F0CEB10
sudo echo 'deb http://downloads-distro.mongodb.org/repo/debian-sysvinit dist 10gen' | sudo tee /etc/apt/sources.list.d/mongodb.list
sudo apt-get update
sudo apt-get install mongodb-org
```



No go to your GOPATH location and run:

```
go get github.com/elcct/amazingwebsite
```

And then:

```
go install github.com/elcct/amazingwebsite
```

In your GOPATH directory you can create `config.json` file:

```
{
	"secret": "secret",
	"public_path": "./src/github.com/elcct/amazingwebsite/public",
	"template_path": "./src/github.com/elcct/amazingwebsite/views",	
	"database": {
		"hosts": "localhost",
		"database": "defaultproject"
	}
}
```

Finally, you can run:

```
./bin/defaultproject
```

That should output something like:

```
2014/06/19 15:31:15.386961 Starting Goji on [::]:8000
```

And it means you can now direct your browser to `localhost:8000`

# Project structure

`/controllers`

All your controllers that serve defined routes.

`/helpers`

Helper functions.

`/models`

You database models.

`/public`

It has all your static files mapped to `/assets/*` path except `robots.txt` and `favicon.ico` that map to `/`.

`/system`

Core functions and structs.

`/views`

Your views using standard `Go` template system.

`server.go`

This file starts your web application and also contains routes definition.

