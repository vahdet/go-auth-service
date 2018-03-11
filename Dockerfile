# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang
MAINTAINER Vahdet Keskin <vahdetkeskin@gmail.com>

# Copy the local package files to the container's workspace.
# Note: COPY vs ADD: COPY is same as 'ADD', but without the tar and remote URL handling.
ADD . /go/src/github.com/vahdet/go-auth-service

# Build the outyet command inside the container.
# RUN set -x && go get github.com/golang/dep/cmd/dep && dep ensure -v
RUN go install github.com/vahdet/go-auth-service

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/go-auth-service

# Document that the service listens on port 5300 .
EXPOSE 5300
