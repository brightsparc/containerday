# Docker compose golang redis

Example docker container with linked redis container

## Prerequisites

Following is required to build and deploy.

* [Docker 17.05](https://docs.docker.com/engine/installation) or higher.
* [Go 1.8](https://blog.golang.org/go1.8) or higher if you want to build locally.

## Getting started

Make sure your go environment is setup and docker engine is running.

### Build

Rebuild the docker images with `docker-compose`:

```
$ docker-compose build
```

### Run

Run the docker container for redis db and web app with `docker-compose`:

```
$ docker-compose up
```

### Testing

Test the API with curl 

```
$ curl localhost:5000

```

## Authors

* Julian Bright - [brightsparc](https://github.com/brightsparc/)
