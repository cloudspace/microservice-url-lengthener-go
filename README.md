URL Lengthener Microservice (Go)
================================

Very simple microservice written in Go for lengthening URLs.

To use the url lengthener, run:
`./lengthener http://domain.com/tinyurl1 http://domain.com/tinyurl2`

----

Build binary with all dependencies:
`CGO_ENABLED=0 go build -a -ldflags '-s' lengthener.go`

Verify with (should show “not a dynamic executable”):
`ldd lengthener`

Building Dockerfile:
`docker build -t cloudspace/url-lengthener-go .`

Running Dockerfile:
`docker run -ti cloudspace/url-lengthener-g http://domain.com/tinyurl1`

Pushing docker image:
`docker push cloudspace/url-lengthener-go`

Finding docker images:
`docker images -a`

Finding docker containers:
`docker ps -a`

Removing a container:
`docker rm containerid`
