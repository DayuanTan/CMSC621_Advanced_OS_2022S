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

func argminxHash(name string, start uint64, end uint64) uint64 {
	nonceAndHashcode := make([][]uint64, end-start)
	for i := range nonceAndHashcode {
		nonceAndHashcode[i] = []uint64{0, math.MaxUint64}
	}

	i := 0
	for nonce := start; nonce < end; nonce++ {
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
	log.Printf("Current tokenList is:")
	for _, token := range tokenList {
		fmt.Printf("	ID: %v; \n	Name: %v; \n	DomainLow: %v; \n	DomainMid: %v; \n	DomainHigh: %v; \n	StatePartialValue: %v; \n	StateFinalValue: %v \n\n", token.Id, token.Name, token.DomainLow, token.DomainMid, token.DomainHigh, token.StatePartialValue, token.StateFinalValue)
	}
}

func (s *server) CreateOneToken(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	log.Printf("Server received: to create ID: %v\n", in.GetId())
	oldleng := len(tokenList)
	for i := 0; i < oldleng; i++ {
		if tokenList[i].Id == in.GetId() {
			log.Printf("Creating ID " + in.GetId() + " failed.")
			fmt.Printf("\n")
			return nil, errors.New("id " + in.GetId() + " already exists")
		}
	}
	onetoken := Token{
		Id: in.GetId(),
	}
	tokenList = append(tokenList, onetoken)
	log.Printf("Creating ID " + in.GetId() + " successed.\n")
	PrintTokenList(ctx)

	return &pb.Token{
		Id:      tokenList[len(tokenList)-1].Id,
		Message: "Created ID " + in.GetId() + " successed.",
	}, nil
}

func (s *server) WriteOneToken(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	log.Printf("Server received: to write ID: %v,\nName: %v, \nDomainLow: %v, \nDomainMid: %v, \nDomainHigh: %v\n", in.GetId(), in.GetName(), in.GetDomainLow(), in.GetDomainMid(), in.GetDomainHigh())
	oldleng := len(tokenList)
	for i := 0; i < oldleng; i++ {
		if tokenList[i].Id == in.GetId() {
			tokenList[i].Name = in.GetName()
			tokenList[i].DomainLow = in.GetDomainLow()
			tokenList[i].DomainMid = in.GetDomainMid()
			tokenList[i].DomainHigh = in.GetDomainHigh()
			tokenList[i].StatePartialValue = argminxHash(in.GetName(), in.GetDomainLow(), in.GetDomainMid())
			tokenList[i].StateFinalValue = 0

			log.Printf("Writing ID " + in.GetId() + " successed.\n")
			PrintTokenList(ctx)
			return &pb.Token{
				Id:                tokenList[i].Id,
				Name:              tokenList[i].Name,
				DomainLow:         tokenList[i].DomainLow,
				DomainMid:         tokenList[i].DomainMid,
				DomainHigh:        tokenList[i].DomainHigh,
				StatePartialValue: tokenList[i].StatePartialValue,
				StateFinalValue:   tokenList[i].StateFinalValue,
				Message:           "Wrote ID " + tokenList[i].Id + " successed.",
			}, nil
		}
	}
	log.Printf("Writing ID " + in.GetId() + " failed.\n")
	PrintTokenList(ctx)
	return nil, errors.New("id " + in.GetId() + " was not found")
}

func (s *server) ReadOneToken(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	log.Printf("Server received: to read ID: %v\n", in.GetId())
	oldleng := len(tokenList)
	for i := 0; i < oldleng; i++ {
		if tokenList[i].Id == in.GetId() {

			tempfinal := argminxHash(tokenList[i].Name, tokenList[i].DomainMid, tokenList[i].DomainHigh)
			tokenList[i].StateFinalValue = tokenList[i].StatePartialValue
			if tempfinal <= tokenList[i].StatePartialValue {
				tokenList[i].StateFinalValue = tempfinal
			}

			log.Printf("Reading ID " + in.GetId() + " successed and StateFinalValue updated.\n")
			PrintTokenList(ctx)
			return &pb.Token{
				Id:                tokenList[i].Id,
				Name:              tokenList[i].Name,
				DomainLow:         tokenList[i].DomainLow,
				DomainMid:         tokenList[i].DomainMid,
				DomainHigh:        tokenList[i].DomainHigh,
				StatePartialValue: tokenList[i].StatePartialValue,
				StateFinalValue:   tokenList[i].StateFinalValue,
				Message:           "Read ID " + in.GetId() + " successed.",
			}, nil
		}
	}
	log.Printf("Reading ID " + in.GetId() + " failed.\n")
	fmt.Println()
	return nil, errors.New("id " + in.GetId() + " was not found")
}

func (s *server) DropOneToken(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	log.Printf("Server received: to drop ID: %v\n", in.GetId())
	oldleng := len(tokenList)
	for i := 0; i < oldleng; i++ {
		if tokenList[i].Id == in.GetId() {
			tokenList[i] = tokenList[oldleng-1]
			tokenList = tokenList[:oldleng-1]

			log.Printf("Dropping ID " + in.GetId() + " successed.\n")
			PrintTokenList(ctx)
			return &pb.Token{
				Message: "Dropped " + in.GetId() + " successed",
			}, nil
		}
	}
	log.Printf("Dropping ID " + in.GetId() + " failed.\n")
	PrintTokenList(ctx)
	return nil, errors.New("id " + in.GetId() + " was not found")
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
