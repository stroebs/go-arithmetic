package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"

	pb "github.com/stroebs/go-arithmetic/proto"
	"google.golang.org/grpc"
)

var (
	host = flag.String("host", "localhost", "Server hostname, defaults to localhost")
	port = flag.Uint("port", 4040, "Server listen port, defaults to 4040")
)

func processArgs(args []string, client pb.ArithmeticClient) {
	// fmt.Printf(args)
	var inputArgs []float64
	for i, v := range args[1:] {
		it, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Printf("Invalid parameter at position %v: %v\n", (i + 1), v)
			fmt.Printf("Please only input numbers\n")
			return
		}
		inputArgs = append(inputArgs, it)
	}
	input := pb.InputMultiple{
		X: inputArgs,
	}
	switch args[0] {
	case "add":
		fmt.Printf("Performing gRPC AddMultiple for numbers (%v)\n", inputArgs)
		responseMessage, error := client.AddMultiple(context.Background(), &input)
		if error == nil {
			fmt.Printf("Addition result for numbers (%v): %v\n", inputArgs, responseMessage.Result)
		} else {
			fmt.Printf("Service error occurred: \n%v\n", error)
		}
	case "subtract":
		fmt.Printf("Performing gRPC SubtractMultiple for numbers (%v)\n", inputArgs)
		responseMessage, error := client.SubtractMultiple(context.Background(), &input)
		if error == nil {
			fmt.Printf("Subtraction result for numbers (%v): %v\n", inputArgs, responseMessage.Result)
		} else {
			fmt.Printf("Service error occurred: \n%v\n", error)
		}
	case "multiply":
		fmt.Printf("Performing gRPC MultiplyMultiple for numbers (%v)\n", inputArgs)
		responseMessage, error := client.MultiplyMultiple(context.Background(), &input)
		if error == nil {
			fmt.Printf("Multiplication result for numbers (%v): %v\n", inputArgs, responseMessage.Result)
		} else {
			fmt.Printf("Service error occurred: \n%v\n", error)
		}
	case "divide":
		fmt.Printf("Performing gRPC DivideMultiple for numbers (%v)\n", inputArgs)
		responseMessage, error := client.DivideMultiple(context.Background(), &input)
		if error == nil {
			fmt.Printf("Division result for numbers (%v): %v\n", inputArgs, responseMessage.Result)
		} else {
			fmt.Printf("Service error occurred: \n%v\n", error)
		}
	default:
		fmt.Println("Please specify operation as one of: add, subtract, multiply, divide")
		return
	}
}

// Main function
func main() {
    // Parse CLI flags
    flag.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
        fmt.Fprintf(flag.CommandLine.Output(), "First argument is the operator: add, subtract, divide, multiply\n")
        fmt.Fprintf(flag.CommandLine.Output(), "Examples:\n"+
            "add 1 2 3 4 5\n"+
            "subtract 5 7 9 1\n"+
            "multiply 2.7 2.8 9.9\n"+
            "divide 9 4.1 7.777 9.999\n")
        flag.PrintDefaults()
    }
    flag.Parse()
    if len(os.Args) < 2 {
        fmt.Println("One of add, subtract, multiply or divide sub-commands are required")
        os.Exit(1)
    }

	// Connect to the gRPC server
	server := fmt.Sprintf("%s:%d", *host, *port)
	conn, e := grpc.Dial(server, grpc.WithInsecure())
	if e != nil {
		panic(e)
	}
	defer conn.Close()
	client := pb.NewArithmeticClient(conn)
    // Process CLI arguments
	processArgs(flag.Args(), client)
}
