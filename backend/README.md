## run on local (in $GOPATH)
```
$ dep ensure
$ go run main.go
```
## run on docker
```
$ docker build . -t backend
$ docker run backend:latest
