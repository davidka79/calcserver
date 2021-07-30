package main

import (
	"context"
	"fmt"
	"github.com/davidka79/calcpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

var srv calcpb.UnimplementedCalcServiceServer

type server struct {
}

func (*server) CalcServiceServer(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("Recived som RPC: %v", req)
	firstNum := req.Firstnum
	secNum := req.Secnum
	sum := firstNum + secNum
	result := &calcpb.SumResponse{
		Sumresult: sum,
	}
	return result, nil
}

func main() {
	fmt.Println("It is a Calc Server ...")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed %v", err)
	}
	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &srv)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
