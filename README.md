# swapi-web
Web interface for SWAPI written in Go

## Running from source code
### Requisites
- [Go](https://golang.org/doc/install) 1.8.3 with your workspace defined in $GOPATH

```sh
 go get github.com/victormlourenco/swapi-web
 cd $GOPATH/src/github.com/victormlourenco/swapi-web && go run swapi-web.go
```

### Testing

```sh
  go test $(go list ./... | grep -v /vendor/)
```

## Running from Docker Hub
### Requisites
- [Docker](https://get.docker.com/)
```sh
  docker run -d -p 6060:6060 --name swapi-web victormlourenco/swapi-web
```

## Creating Docker image from source code
### Requisites
- [Docker](https://get.docker.com/)

```sh
 git clone https://github.com/victormlourenco/swapi-web.git && cd swapi-web
 docker build -t swapi-web .
 docker run -d -p 6060:6060 swapi-web
```

## Demo:
http://138.197.196.80/
