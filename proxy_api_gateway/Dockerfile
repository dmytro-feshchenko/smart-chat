# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.9

# Install dependencies
RUN go get github.com/graphql-go/graphql

# copy the local package files to the container workspace
ADD . /go/src/github.com/technoboom/smart-chat/proxy_api_gateway

# Setting up working directory
WORKDIR /go/src/github.com/technoboom/smart-chat/proxy_api_gateway

# Build the users command inside the container.
RUN go install github.com/technoboom/smart-chat/proxy_api_gateway

# Run the users microservice when the container starts.
ENTRYPOINT /go/bin/proxy_api_gateway

# Service listens on port 8080.
EXPOSE 8080
