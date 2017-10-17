# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.9

# Install dependencies
RUN go get github.com/jinzhu/gorm

# copy the local package files to the container workspace
ADD . /go/src/github.com/technoboom/smart-chat/channels

# Setting up working directory
WORKDIR /go/src/github.com/technoboom/smart-chat/channels

# Build the users command inside the container.
RUN go install github.com/technoboom/smart-chat/channels

# Run the users microservice when the container starts.
ENTRYPOINT /go/bin/channels

# Service listens on port 8080.
EXPOSE 8080
