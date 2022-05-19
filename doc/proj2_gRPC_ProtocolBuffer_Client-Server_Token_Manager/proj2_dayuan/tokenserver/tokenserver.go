package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"net"
	"sort"

	"proj2_dayuan/lib"
	pb "proj2_dayuan/token"

	"google.golang.org/grpc"
)

type Token struct {
	Id                string
	Name              string
	DomainLow         uint64
	DomainMid         uint64
	DomainHigh        uint64
	StatePartialValue uint64
	StateFinalValue   uint64
	Message           string
}

var (
	port      = flag.Int("port", 50051, "The server port. Default is 50051.")
	tokenList = []Token{}
)

type server struct {
	pb.UnimplementedTokenServiceServer
}

func argminxHash(name string, low uint64, mid uint64) uint64 {
	nonceAndHashcode := make([][]uint64, mid-low)
	for i := range nonceAndHashcode {
		nonceAndHashcode[i] = []uint64{0, math.MaxUint64}
	}

	i := 0
	for nonce := low; nonce < mid; nonce++ {
		nonceAndHashcode[i][0] = nonce
		nonceAndHashcode[i][1] = lib.Hash(name, nonce)
		i++
	}

	sort.SliceStable(nonceAndHashcode, func(i, j int) bool {
		return nonceAndHashcode[i][1] < nonceAndHashcode[j][1] // sort by second col (hashcode)
	})
	return nonceAndHashcode[0][0] // return the x which hash min hashcode
}

func PrintTokenList(ctx context.Context) {
	log.Printf("Current tokenList is:\n")
	for _, token := range tokenList {
		log.Printf("ID: %v; \n	Name: %v; \n	DomainLow: %v; \n	DomainMid: %v; \n	DomainHigh: %v; \n	StatePartialValue: %v; \n	StateFinalValue: %v \n\n", token.Id, token.Name, token.DomainLow, token.DomainMid, token.DomainHigh, token.StatePartialValue, token.StateFinalValue)
	}
}

func (s *server) CreateOneToken(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	onetoken := Token{
		Id: in.GetId(),
	}
	log.Printf("Server received: \nID: %v", onetoken.Id)
	tokenList = append(tokenList, onetoken)
	PrintTokenList(ctx)

	return &pb.Token{
		Id: tokenList[len(tokenList)-1].Id,
	}, nil
}

func (s *server) WriteOneToken(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	log.Printf("Server received: \nID: %v,\nName: %v, \nDomainLow: %v, \nDomainMid: %v, \nDomainHigh: %v\n", in.GetId(), in.GetName(), in.GetDomainLow(), in.GetDomainMid(), in.GetDomainHigh())
	for i := range tokenList {
		if tokenList[i].Id == in.GetId() {
			tokenList[i].Name = in.GetName()
			tokenList[i].DomainLow = in.GetDomainLow()
			tokenList[i].DomainMid = in.GetDomainMid()
			tokenList[i].DomainHigh = in.GetDomainHigh()
			tokenList[i].StatePartialValue = argminxHash(in.GetName(), in.GetDomainLow(), in.GetDomainMid())
			// token.StateFinalValue = in.GetStateFinalValue()

			PrintTokenList(ctx)
			return &pb.Token{
				Id:         in.GetId(),
				Name:       in.GetName(),
				DomainLow:  in.GetDomainLow(),
				DomainMid:  in.GetDomainMid(),
				DomainHigh: in.GetDomainHigh(),
			}, nil
		}
	}
	return nil, errors.New("id not found")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Server: failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTokenServiceServer(s, &server{})
	log.Printf("Server: listening at %v\n\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server: failed to serve: %v", err)
	}
}
