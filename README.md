# go-arithmetic
Arithmetic in Go using gRPC  

![](https://github.com/stroebs/go-arithmetic/workflows/Go%20Staging/badge.svg) | ![](https://github.com/stroebs/go-arithmetic/workflows/Go%20Production/badge.svg)  

This is a simple microservice which contains a few methods to perform basic arithmetic.  
The package was designed with TDD, CI/CD and 12factor application principles in mind.  
The microservice utilises gRPC for communication and has been tested locally and with Kubernetes.  

## Endpoints

Two endpoints for the solution are live on GKE:
* **Staging**: <removed>
* **Production**: <removed>

Deployment is entirely automated and controlled by Git. Pull Requests are tested, built and deployed to staging and production is deployed when a push to master occurs. This is all handled by Github Actions.

## Usage

The client interface is really simple to use. It relies on the server running locally, inside Docker with port 4040 running on localhost or externally in a Kubernetes cluster. See Endpoints above). You can configure the host and port the client uses through command-line arguments or environment variables.  

It's best to build the client before using these commands for fast testing.  
  
**Arguments:**
* `add`: Use `add` followed by any number of integers or floats to add the numbers together
* `subtract`: Use `subtract` followed by any number of integers or floats to subtract the numbers from each other, in order from left to right
* `multiply`: Use `multiply` followed by any number of integers or floats to multiply the numbers by each other, in order from left to right
* `divide`: Use `divide` followed by any number of integers or floats to divide the numbers by each other, in order from left to right

## Configuration:

Configuration from `.env`, environment variables and CLI flags are supported.  

### Server:
`SERVER_LISTEN_PORT` | `--port`: Important for running in a Kubernetes environment for flexibility of port selection.

### Client:
`ARITH_SERVER_URI` | `--uri`: This is in the `host:port` format for the gRPC endpoint you wish to use for the client.


#### Examples:

**Firstly:**
`go build client.go`

**Input:** `./client add 1 2 3 4 5`  
**Output:**
```
Performing gRPC AddMultiple for numbers ([1 2 3 4 5])
Addition result for numbers ([1 2 3 4 5]): 15
```

**Input:** `./client subtract 2.3 4.5 7.7 9`  
**Output:**
```
Performing gRPC SubtractMultiple for numbers ([2.3 4.5 7.7 9])
Subtraction result for numbers ([2.3 4.5 7.7 9]): -18.9
```

**Input:** `./client multiply 1.1 7.1 6 2 3 1 6 2`  
**Output:**
```
Performing gRPC MultiplyMultiple for numbers ([1.1 7.1 6 2 3 1 6 2])
Multiplication result for numbers ([1.1 7.1 6 2 3 1 6 2]): 3373.9199999999996
```

**Input:** `./client divide 1.1 7.1`  
**Output:**
```
Performing gRPC DivideMultiple for numbers ([1.1 7.1])
Division result for numbers ([1.1 7.1]): 0.15492957746478875
```

**Input:** `./client -uri 35.187.4.200:80 add 1 4 5 6 7`  
**Output:**
```
Performing gRPC AddMultiple for numbers ([1 4 5 6 7])
Addition result for numbers ([1 4 5 6 7]): 23
```

### Caveats:

* Due to the use of floating point numbers, results may not be exact. No rounding is in place on the server or client side
* Division by zero will result in infinity and is not handled in this service
* At present, gRPC API error handling is non-existent

## Development

### Running tests

The tests have been embedded within the Dockerfile for ease of running, so you only have to execute a single command to ensure tests are run.  
If you want to run them outside of Docker, simply use: `go test ./...` or `cd server; go test`.  

### Building the Docker image

You can build the Docker image with the following command:  

`docker build -t stroebs/go-arithmetic .`  
The Docker image is auto-built and contained on Docker Hub here: https://hub.docker.com/repository/docker/stroebs/go-arithmetic  
  
The image build is a multi-stage build which performs unit tests and builds the server into a `scratch` image containing only the binary required. This provides for an extremely fast and highly-scalable 12MB image containing all required dependencies (thanks Golang!).

## 12factor App Alignment

The application was written to work with the principles of a [12factor](https://12factor.net/) application in mind.  

**The 12 factors:**  
**I. Codebase**  
_One codebase tracked in revision control, many deploys_  
The entire `go-arithmetic` codebase and infrastructure is contained in this single repository.

**II. Dependencies**  
_Explicitly declare and isolate dependencies_  
Dependencies are managed through Godep with files `go.mod` and `go.sum`. As Golang treats everything as a package, this is quite easy to manage.

**III. Config**  
_Store config in the environment_  
The package `godotenv`, which is an extension of the Ruby implementation `dotenv` allows specification of configuration in both `.env` and as environment variables.

**IV. Backing services**  
_Treat backing services as attached resources_  
No backing services present or required in this case. This would change if something like an eventstore was added.

**V. Build, release, run**  
_Strictly separate build and run stages_  
The Github Actions separate the pipeline into test, build and deploy actions. Releasing is handled through a merge to the master branch.

**VI. Processes**  
_Execute the app as one or more stateless processes_  
The app runs as a single process inside a Docker container (best practice), contains no state and is horizontally-scalable.

**VII. Port binding**  
_Export services via port binding_  
The gRPC service is built-in and bound to an external port through the environment.

**VIII. Concurrency**  
_Scale out via the process model_  
The app runs as a single process inside a Docker container (best practice), contains no state and is horizontally-scalable.

**IX. Disposability**  
_Maximize robustness with fast startup and graceful shutdown_  
The Docker image is minimal and contains a single process.
Graceful shutdown has not been tested at this point but it would be easy to implement cleanup based on SIGTERM if there was an attached backing service that required graceful termination.

**X. Dev/prod parity**  
_Keep development, staging, and production as similar as possible_  
As the Github Actions pipeline is nearly identical for both staging and production, parity is held through multiple deploys/service updates.

**XI. Logs**  
_Treat logs as event streams_  
The logging package in Golang does this quite well and sends logs to stdout/stderr and contain a timestamp for easy event identification.

**XII. Admin processes**  
_Run admin/management tasks as one-off processes_  
No one-off admin/management tasks are required in such a small application but could exist if an eventstore were present to perform migrations or other related operations.

### Adding an eventstore

In this simple example, adding an event store would be trivial and can be achieved on the gRPC method level to store each call that is received and replied to as an event in an immutable storage system. The gRPC methods on the server make method calls to their like-named counterparts so adding event store code in the upper-level method calls would be simple. This could be enhanced by adding a middleware package that could handle this instead of making individual event store calls per method.
