package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "proj2_dayuan/token"

	"google.golang.org/grpc"
)

var (
	port       = flag.Int("port", 50051, "The server port. Default is 50051.")
	tokensList = []pb.Token{}
)

type server struct {
	pb.UnimplementedTokenServiceServer
}

func (s *server) CreateOneToken(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	log.Printf("Server received: %v", in.GetName())
	onetokenPtr := pb.Token{
		Id:                in.GetId(),
		Name:              in.GetName(),
		DomainLow:         in.GetDomainLow(),
		DomainMid:         in.GetDomainMid(),
		DomainHigh:        in.GetDomainHigh(),
		StatePartialValue: in.GetStatePartialValue(),
		StateFinalValue:   in.GetStateFinalValue(),
	}
	tokensList = append(tokensList, onetokenPtr)
	log.Printf("tokensList: %v", tokensList)
	return &pb.Token{Name: "One Token created: " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Server: failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTokenServiceServer(s, &server{})
	log.Printf("Server: listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server: failed to serve: %v", err)
	}
}
