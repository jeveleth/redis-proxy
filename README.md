# redis-proxy


Documentation
## High-level architecture overview.
<!-- TODO: -->
## What the code does.
<!-- TODO: -->
## Algorithmic complexity of the cache operations.
<!-- TODO: -->
## Instructions for how to run the proxy and tests.
<!-- TODO: -->
## How long you spent on each part of the project.
    * HTTP web service (1 hr)
    * Single backing instance (2 hrs)
    * Cached GET (2 hrs)
    * System tests (3 hrs)
    * Platform (2 hrs)
    * Single-click build and test (1hr)
    * Documentation ()
    * LRU eviction and fixed key size ()

## A list of the requirements that you did not implement and the reasons for omitting them.
    * Global expiry
    * Sequential concurrent processing
    * Parallel concurrent processing
    * Redis client protocol

## Getting started
To build and test, according to the ```Single-click build and test``` requirement. Open a terminal session and run:

    git clone git@github.com:jeveleth/redis-proxy.git
    make test # This builds the proxy and runs the tests

# Running your own proxy
 Run ```make docker-proxy```, which will put you in an interactive docker container with access to the proxy server. Then run ```./proxy -help```, to see what you can configure. To run the proxy, run ```./proxy``` with any optional flags.

 As an example you can test the service by doing the following (all within the root directory of the project):
  1. (Session 1) Open one terminal session and run: ```make docker-proxy```. Once inside the bash prompt, run ```./proxy -proxy-port 9000```.
  2. (Session 2) Open a separate terminal session and run:
        ```make docker-proxy```
        ```curl localhost:9000/getval/key22```
    You should *not* see any value.

  3. (Session 3) Open an separate terminal session and run ```make redis-cli```, which will drop you into an interactive session with the redis server. For this example, type ```set key22 value22```. Go back to session 2 and run the curl command again, like so: ```bash-4.4# curl localhost:9000/getval/key22```. You should see a response like: ```From Redis: key22 => value22bash-4.4#```. Run the curl command again, and you should see a response like ```From cache: key22 => value1bash-4.4```.

# Running the tests
Run ```make docker-proxy```, which will put you in an interactive docker container with access to the proxy server. Run ```go test -v```.




 <!-- TODO -->
Sequential concurrent processing
Multiple clients are able to concurrently connect to the proxy (up to some configurable maximum limit) without adversely impacting the functional behaviour of the proxy. When multiple clients make concurrent requests to the proxy, it is acceptable for them to be processed sequentially (i.e. a request from the second only starts processing after the first request has completed and a response has been returned to the first client).


Bonus Requirements

The requirements below add some additional complexity to the design and can be implemented as a bonus. However, we strongly encourage candidates who are applying for a role which has a strong backend systems focus to implement these as well.

Parallel concurrent processing
Multiple clients are able to concurrently connect to the proxy (up to some configurable maximum limit) without adversely
impacting the functional behaviour of the proxy. When multiple clients make concurrent requests to the proxy, it would execute a number of these requests (up to some configurable limit) in parallel (i.e. in a way so that one request does not have to wait for another one to complete before it starts processing).

Redis client protocol
Clients interface to the Redis proxy through a subset of the Redis protocol (as opposed to using the HTTP protocol). The proxy should implement the parts of the Redis protocol that is required to meet this specification.
