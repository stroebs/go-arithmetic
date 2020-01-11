# go-arithmetic
Arithmetic in Go using gRPC  

This is a simple microservice which contains a few methods to perform basic arithmetic  
The package was designed with TDD, CI/CD and 12factor application principles in mind  
The microservice utilises gRPC for communication and has been tested locally and with Kubernetes  

## Usage

The client interface is really simple to use. It relies on the server running either locally or inside Docker with port 4040 running on localhost. You can configure the host and port the client uses through command-line arguments or environment variables.  
  
**Arguments:**
* `add`: Use `add` followed by any number of integers or floats to add the numbers together
* `subtract`: Use `subtract` followed by any number of integers or floats to subtract the numbers from each other, in order from left to right
* `multiply`: Use `multiply` followed by any number of integers or floats to multiply the numbers by each other, in order from left to right
* `divide`: Use `divide` followed by any number of integers or floats to divide the numbers by each other, in order from left to right

#### Examples:

**Input:** `go run client.go add 1 2 3 4 5`  
**Output:**
```
Performing gRPC AddMultiple for numbers ([1 2 3 4 5])
Addition result for numbers ([1 2 3 4 5]): 15
```

**Input:** `go run client.go subtract 2.3 4.5 7.7 9`  
**Output:**
```
Performing gRPC SubtractMultiple for numbers ([2.3 4.5 7.7 9])
Subtraction result for numbers ([2.3 4.5 7.7 9]): -18.9
```

**Input:** `go run client.go multiply 1.1 7.1 6 2 3 1 6 2`  
**Output:**
```
Performing gRPC MultiplyMultiple for numbers ([1.1 7.1 6 2 3 1 6 2])
Multiplication result for numbers ([1.1 7.1 6 2 3 1 6 2]): 3373.9199999999996
```

**Input:** `go run client.go divide 1.1 7.1`  
**Output:**
```
Performing gRPC DivideMultiple for numbers ([1.1 7.1])
Division result for numbers ([1.1 7.1]): 0.15492957746478875
```

**Input:** `go run client.go -host my-awesome-host.com -port 9191 add 1 4 5 6 7`
**Output:**
```
Performing gRPC AddMultiple for numbers ([1 4 5 6 7])
Addition result for numbers ([1 4 5 6 7]): 23
```

### Caveats:

* Due to the use of floating point numbers, results may not be exact. No rounding is in place on the server or client size
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
