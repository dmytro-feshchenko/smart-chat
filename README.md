# Smart Chat - Realtime chat built with Go, Docker, MongoDB and GraphQL

Overview
========

Cinema is an example project which demonstrates the use of microservices for a fictional movie theater.

The backend is powered by 4 microservices, all of witch happen to be written in
Go.

 * Channels Service: Provides information about existing channels.
 * Messages Service: Provides communicating information.
 * Users Service: Provides messages for users by communicating with other services.
 * Web Service: Represents web interface for customer.

Commands
========
Starting services: `docker-compose up --build` or
using script `./bin/start.sh`

Stopping services: `docker-compose stop`

Deploying lates changes: `docker-compose build`
