httpcheck
==================

**Requirements:** We have health checks for every service in our system. For third-party applications like Cassandra we write the health checks in Go. We use Go mostly because it doesn't add any run-time dependencies.

In Go, implement a simple Cassandra health check that makes its results available via http.  It should return an http status of 200 if Cassandra is available, and 503 otherwise.

Lean on libraries as heavily as you feel appropriate - the idea is to produce something small, reliable, and easy to read.

You can assume Cassandra is available on 127.0.0.1.

The service should run forever, stay in the foreground, and listen on :8080.  It should log the result of each http request to stdout.


**Usage:** This can be run stand alone. It starts up an http server that runs a basic 'nodetool status' on Cassandra that is running locally. Each time a request hits the web server, a Cassandra service check is ran. Depending on the output of 'nodetool status', the web server will either respond with a status code of 200 if Cassandra is available, or a status code of 503 if Cassandra is not running. 

The output of the service check is also logged to stdout.