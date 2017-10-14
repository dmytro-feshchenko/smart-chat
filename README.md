# Smart Chat - Realtime chat built with Go, Node.js, Docker and GraphQL

Overview
========

The backend is powered by 4 microservices, all of witch happen to be written in
Go as RESTful APIs.

 * Channels Service: Provides information about existing channels.
 * Messages Service: Provides communicating information.
 * Users Service: Provides messages for users by communicating with other services.
 * Web Service: Represents web interface for customer.

We also have Proxy API Gateway written in Node.js + GraphQL to create one endpoint
to communicate with microservices.

Commands
========
Starting services: `docker-compose up --build` or
using script `./bin/start.sh`

Stopping services: `docker-compose stop`

Deploying lates changes: `docker-compose build`

Proxy API Gateway
=================
Written in Node.js + GraphQL (using Apollo Server).


Services
========
1. Users Service: uses Golang with GraphQL to provide endpoint to work with users data. All the data stored into PostgreSQL database.
2. Web Service: represents React + Redux + Apollo Client to provide web
interface for customers.
