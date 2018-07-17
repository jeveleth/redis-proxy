# redis-proxy

### High-level architecture overview
<!-- TODO: Add chart -->
### What the code does
<!-- TODO: -->
### Algorithmic complexity of the cache operations
<!-- TODO: -->
### Instructions for how to run the proxy and tests

#### Getting started
To build and test, according to the ```Single-click build and test``` requirement. Open a terminal session and run:

    git clone git@github.com:jeveleth/redis-proxy.git
    cd redis-proxy
    make test # This builds the proxy and runs the tests

#### Running your own proxy
(This assumes you have cloned the repo and are in the top-level directory of the project.)

 Run ```make docker-proxy```, which will put you in an interactive docker container with access to the proxy server. Then run ```./proxy -help```, to see what you can configure. To run the proxy, run ```./proxy``` with any optional flags.

 You can run the service by opening three terminal sessions and doing as follows:
* (In **Session 1**) Run:

        make docker-proxy
        ./proxy -proxy-port 9000 # Once inside the bash prompt

* (In **Session 2**) Run:

        make docker-proxy
        curl localhost:9000/getval/key22 # Once inside the bash prompt

You should *not* see any value.

* (In **Session 3**) Run: 
    
        make redis-cli # drops you into an interactive session with the redis server.
        set key22 value22. # sets a value in redis server.
 
Go back to **Session 2**. Run the curl command twice and see a response first from Redis, then from the local cache.

    curl localhost:9000/getval/key22 #Run this.
    From Redis: key22 => value22. #Proxy responds with value from redis 
    curl localhost:9000/getval/key22 # Run this.
    From cache: key22 => value1.  # Proxy responds with value from cache

#### Running the tests
Run ```make docker-proxy```, which will put you in an interactive docker container with access to the proxy server. Run ```go test -v```.

### How long you spent on each part of the project.
* HTTP web service (1 hr)
* Single backing instance (2 hrs)
* Cached GET (2 hrs)
* System tests (3 hrs)
* Platform (2 hrs)
* Single-click build and test (1hr)
* LRU eviction and fixed key size (30 min)
* Global expiry (30 min)
* Documentation ()   <!-- TODO: Time estimate -->

#### A list of the requirements that you did not implement and the reasons for omitting them.
    * Sequential concurrent processing (Confused about setting max connections on server)
    * Parallel concurrent processing (Ran out of time)
    * Redis client protocol (Ran out of time)

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
