package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "github.com/stroebs/go-arithmetic/proto"
)

var (
	port = flag.Uint("port", 4040, "Server listen port")
)

type ArithmeticServer struct {
}

// Add (gRPC) accepts protobuf input, processes via Add method and returns result via gRPC
func (s *ArithmeticServer) Add(ctx context.Context, input *pb.Input) (*pb.Result, error) {
	log.Printf("Received Add gRPC request with input: %v", input)
	result := Add(input.X, input.Y)
	log.Printf("Returning Add result %v", result)
	return &pb.Result{
		Result: result,
		Error:  nil,
	}, nil
}

// AddMultiple (gRPC) accepts protobuf input, processes via AddMultiple method and returns result via gRPC
func (s *ArithmeticServer) AddMultiple(ctx context.Context, inputMultiple *pb.InputMultiple) (*pb.Result, error) {
	log.Printf("Received AddMultiple gRPC request with input: %v", inputMultiple)
	result := AddMultiple(inputMultiple.X)
	log.Printf("Returning AddMultiple result %v", result)
	return &pb.Result{
		Result: result,
		Error:  nil,
	}, nil
}

// Subtract (gRPC) accepts protobuf input, processes via Subtract method and returns result via gRPC
func (s *ArithmeticServer) Subtract(ctx context.Context, input *pb.Input) (*pb.Result, error) {
	log.Printf("Received Subtract gRPC request with input: %v", input)
	result := Subtract(input.X, input.Y)
	log.Printf("Returning Subtract result %v", result)
	return &pb.Result{
		Result: result,
		Error:  nil,
	}, nil
}

// SubtractMultiple (gRPC) accepts protobuf input, processes via SubtractMultiple method and returns result via gRPC
func (s *ArithmeticServer) SubtractMultiple(ctx context.Context, inputMultiple *pb.InputMultiple) (*pb.Result, error) {
	log.Printf("Received SubtractMultiple gRPC request with input: %v", inputMultiple)
	result := SubtractMultiple(inputMultiple.X)
	log.Printf("Returning SubtractMultiple result %v", result)
	return &pb.Result{
		Result: result,
		Error:  nil,
	}, nil
}

// Multiply (gRPC) accepts protobuf input, processes via Multiply method and returns result via gRPC
func (s *ArithmeticServer) Multiply(ctx context.Context, input *pb.Input) (*pb.Result, error) {
	log.Printf("Received Multiply gRPC request with input: %v", input)
	result := Multiply(input.X, input.Y)
	log.Printf("Returning Multiply result %v", result)
	return &pb.Result{
		Result: result,
		Error:  nil,
	}, nil
}

// MultiplyMultiple (gRPC) accepts protobuf input, processes via MultiplyMultiple method and returns result via gRPC
func (s *ArithmeticServer) MultiplyMultiple(ctx context.Context, inputMultiple *pb.InputMultiple) (*pb.Result, error) {
	log.Printf("Received MultiplyMultiple gRPC request with input: %v", inputMultiple)
	result := MultiplyMultiple(inputMultiple.X)
	log.Printf("Returning MultiplyMultiple result %v", result)
	return &pb.Result{
		Result: result,
		Error:  nil,
	}, nil
}

// Divide (gRPC) accepts protobuf input, processes via Divide method and returns result via gRPC
func (s *ArithmeticServer) Divide(ctx context.Context, input *pb.Input) (*pb.Result, error) {
	log.Printf("Received Divide gRPC request with input: %v", input)
	result := Divide(input.X, input.Y)
	log.Printf("Returning Divide result %v", result)
	return &pb.Result{
		Result: result,
		Error:  nil,
	}, nil
}

// DivideMultiple (gRPC) accepts protobuf input, processes via DivideMultiple method and returns result via gRPC
func (s *ArithmeticServer) DivideMultiple(ctx context.Context, inputMultiple *pb.InputMultiple) (*pb.Result, error) {
	log.Printf("Received DivideMultiple gRPC request with input: %v", inputMultiple)
	result := DivideMultiple(inputMultiple.X)
	log.Printf("Returning DivideMultiple result %v", result)
	return &pb.Result{
		Result: result,
		Error:  nil,
	}, nil
}

func main() {
	// Parse command-line flags
	flag.Parse()
	// Add datetime and line logging to logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//Initialise server
	log.Printf("Starting server on tcp/%v", int(*port))
	netListener := getNetListener(*port)
	gRPCServer := grpc.NewServer()
	arithmeticServer := &ArithmeticServer{}
	pb.RegisterArithmeticServer(gRPCServer, arithmeticServer)

	// Start the server
	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}

// getNetListener creates a listener on specified port and host
func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}

	return lis
}

// Add takes two float64s and returns the sum of them
func Add(x, y float64) float64 {
	return x + y
}

// AddSlice takes n float64 inputs as a slice and returns the sum of them
func AddMultiple(inputs []float64) float64 {
	sum := inputs[0]
	for _, input := range inputs[1:] {
		sum += input
	}
	return sum
}

// Subtract takes two float64s and returns the subtraction of the second from the first
func Subtract(x, y float64) float64 {
	return x - y
}

// SubtractSlice takes n inputs as a slice and returns the ordered subtraction from left to right
func SubtractMultiple(inputs []float64) float64 {
	sum := inputs[0]
	for _, input := range inputs[1:] {
		sum -= input
	}
	return sum
}

// Multiply takes two float64s and returns the multiplied result of them
func Multiply(x, y float64) float64 {
	return x * y
}

// MultiplySlice takes n float64 inputs as slice and returns the ordered multiplication from left to right
func MultiplyMultiple(inputs []float64) float64 {
	sum := inputs[0]
	for _, input := range inputs[1:] {
		sum *= input
	}
	return sum
}

// Divide takes two float64s and returns the division of the first by the second
func Divide(x, y float64) float64 {
	return x / y
}

// DivideSlice takes n float64 inputs as slice and returns the ordered division from left to right
func DivideMultiple(inputs []float64) float64 {
	sum := inputs[0]
	for _, input := range inputs[1:] {
		sum /= input
	}
	return sum
}
