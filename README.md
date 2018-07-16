# redis-proxy


# Getting started

    git clone git@github.com:jeveleth/redis-proxy.git
    make test

# Configurable

Running with flags
```go run *.go -help```









Segment Assignment

TODO:
 --> Time: 1hr
The table below defines requirements that the proxy has to meet and against which the implementation would be measured. It allows the proxy to be used as a simple read-through cache. When deployed in this fashion, it is assumed that all writes are directed to the backing Redis instance, bypassing the proxy.

HTTP web service
Clients interface to the Redis proxy through HTTP, with the Redis “GET” command mapped to the HTTP “GET” method. Note that the proxy still uses the Redis protocol to communicate with the backend Redis server.

Single backing instance
Each instance of the proxy service is associated with a single Redis service instance called the “backing Redis”. The address of the backing Redis is configured at proxy startup.



Cached GET
A GET request, directed at the proxy, returns the value of the specified key from the proxy’s local cache if the local cache contains a value for that key. If the local cache does not contain a value for the specified key, it fetches the value from the backing Redis instance, using the Redis GET command, and stores it in the local cache, associated with the specified key.

Sequential concurrent processing
Multiple clients are able to concurrently connect to the proxy (up to some configurable maximum limit) without adversely impacting the functional behaviour of the proxy. When multiple clients make concurrent requests to the proxy, it is acceptable for them to be processed sequentially (i.e. a request from the second only starts processing after the first request has completed and a response has been returned to the first client).

Configuration
The following parameters are configurable at the proxy startup:
    * Address of the backing Redis
    * Cache expiry time
    * Capacity (number of keys)
    * TCP/IP port number the proxy listens on

System tests
    Automated systems tests confirm that the end-to-end system functions as specified.

Platform
The software build and tests pass on a modern Linux distribution or Mac OS installation, with the only assumptions being as follows:
1. The system has the following software installed:
    * make
    * docker
    * docker-compose
    * Bash
2. The system can access DockerHub over the internet.

Single-click build and test
After extracting the source code archive, or cloning it from a Git repo, entering the top-level project directory and executing will build the code and run all the relevant tests. Apart from the downloading and manipulation of docker images and containers, no changes are made to the host system outside the top-level directory of the project. The build and test should be fully repeatable and not requires any of software installed on the host system, with the exception of anything specified explicitly in the requirement.

Documentation
The software includes a README file with:
* High-level architecture overview.
* What the code does.
* Algorithmic complexity of the cache operations.
* Instructions for how to run the proxy and tests.
* How long you spent on each part of the project.
* A list of the requirements that you did not implement and the reasons for omitting them.


Bonus Requirements

The requirements below add some additional complexity to the design and can be implemented as a bonus. However, we strongly encourage candidates who are applying for a role which has a strong backend systems focus to implement these as well.

Parallel concurrent processing
Multiple clients are able to concurrently connect to the proxy (up to some configurable maximum limit) without adversely
impacting the functional behaviour of the proxy. When multiple clients make concurrent requests to the proxy, it would execute a number of these requests (up to some configurable limit) in parallel (i.e. in a way so that one request does not have to wait for another one to complete before it starts processing).

Redis client protocol
Clients interface to the Redis proxy through a subset of the Redis protocol (as opposed to using the HTTP protocol). The proxy should implement the parts of the Redis protocol that is required to meet this specification.

 DONE:
 Makefile
 Git repo