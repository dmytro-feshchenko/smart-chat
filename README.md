# Smart Chat - Realtime chat built with Go, Node.js, Docker, React and GraphQL

Overview
========

The backend is powered by 4 microservices, all of witch happen to be written in
Go as RESTful APIs.

 * Channels Service: Provides information about existing channels. [Not yet]
 * Messages Service: Provides communicating information. [Not yet]
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
Responsibilities:
* Authentication
* Managing accounts
* Restore password
2. Web Service: represents React + Redux + Apollo Client to provide web
interface for customers.
3. Messages Service: stores all data about communication between customers. Provides
a RESTful API to give an access to this information.
4. Channels Service: stores all data about public channels. Provides REST API.
Responsibilities:
* Search public channels
* Join to some public/private channel

Web Interface
=============
Built with create-react-app and Apollo Client.
