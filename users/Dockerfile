# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.8

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
RUN go get github.com/graphql-go/graphql

# copy the local package files to the container workspace
ADD . /go/src/github.com/technoboom/smart-chat/users

# Setting up working directory
WORKDIR /go/src/github.com/technoboom/smart-chat/users

# Build the users command inside the container.
RUN go install github.com/technoboom/smart-chat/users

# Run the users microservice when the container starts.
ENTRYPOINT /go/bin/users

# Service listens on port 8080.
EXPOSE 8080
