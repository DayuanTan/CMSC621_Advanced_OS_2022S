package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "proj2_dayuan/token"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	createOperPtr = flag.Bool("create", false, "Claim create operation. Is bool.")
	readOperPtr   = flag.Bool("read", false, "Claim read operation. Is bool.")
	writeOperPtr  = flag.Bool("write", false, "Claim write operation. Is bool.")
	dropOperPtr   = flag.Bool("drop", false, "Claim drop operation. Is bool.")
	idPtr         = flag.String("id", "-1", "The id (string) for your token")
	namePtr       = flag.String("name", "", "The name (string) of your token. Default is 'Dayuan'")
	lowPtr        = flag.Uint64("low", 0, "The low value (uint64) of your token")
	midPtr        = flag.Uint64("mid", 0, "The mid value (uint64) of your token")
	highPtr       = flag.Uint64("high", 0, "The high value (uint64) of your token")
	hostPtr       = flag.String("host", "localhost", "The host (string) to connect to. Default is localhost")
	portPtr       = flag.String("port", "50051", "The port (string) to connect to. Default is '50051'")
)

func checkFlagsValidation() {
	if *createOperPtr {
		if *readOperPtr || *writeOperPtr || *dropOperPtr {
			log.Fatalf("Client: Flag error. Only one operation is allowed at same time.")
		}
	} else if *readOperPtr {
		if *createOperPtr || *writeOperPtr || *dropOperPtr {
			log.Fatalf("Client: Flag error. Only one operation is allowed at same time.")
		}
	} else if *writeOperPtr {
		if *createOperPtr || *readOperPtr || *dropOperPtr {
			log.Fatalf("Client: Flag error. Only one operation is allowed at same time.")
		}
	} else if *dropOperPtr {
		if *createOperPtr || *readOperPtr || *writeOperPtr {
			log.Fatalf("Client: Flag error. Only one operation is allowed at same time.")
		}
	}

}

func getAddr() string {
	return *hostPtr + ":" + *portPtr
}

func main() {
	flag.Parse()
	log.Printf("Your input is:")
	fmt.Println("Operation: Create: ", *createOperPtr, "; Read: ", *readOperPtr, "; Write: ", *writeOperPtr, "; Drop: ", *dropOperPtr)
	fmt.Println("Parameters: ID: ", *idPtr, "; Host: ", *hostPtr, "; Port: ", *portPtr, "; Low: ", *lowPtr, "; Mid: ", *midPtr, "; High: ", *highPtr, "; Name: ", *namePtr)
	fmt.Println()
	checkFlagsValidation()
	addr := getAddr()

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client: did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTokenServiceClient(conn)

	// Call server methods to deal with token(s)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if *createOperPtr {
		r, err := c.CreateOneToken(ctx, &pb.Token{
			Id: *idPtr,
		})
		if err != nil {
			log.Fatalf("Client: failed to call server CreateOneToken(): %v", err)
		}
		log.Printf("Client received: \nID: %v, Message: %v\n", r.GetId(), r.GetMessage())
	}
	if *writeOperPtr {
		r, err := c.WriteOneToken(ctx, &pb.Token{
			Id:         *idPtr,
			Name:       *namePtr,
			DomainLow:  *lowPtr,
			DomainMid:  *midPtr,
			DomainHigh: *highPtr,
		})
		if err != nil {
			log.Fatalf("Client: failed to call server WriteOneToken(): %v", err)
		}
		log.Printf("Client received: Message: %v, \nID: %v,\nName: %s, \nDomainLow: %v, \nDomainMid: %v, \nDomainHigh: %v, \nStatePartialValue: %v, \nStateFinalValue: %v", r.GetMessage(), r.GetId(), r.GetName(), r.GetDomainLow(), r.GetDomainMid(), r.GetDomainHigh(), r.GetStatePartialValue(), r.GetStateFinalValue())
	}
	if *readOperPtr {
		r, err := c.ReadOneToken(ctx, &pb.Token{
			Id: *idPtr,
		})
		if err != nil {
			log.Fatalf("Client: failed to call server ReadOneToken(): %v", err)
		}
		log.Printf("Client received: Message: %v, \nID: %v,\nName: %s, \nDomainLow: %v, \nDomainMid: %v, \nDomainHigh: %v, \nStatePartialValue: %v, \nStateFinalValue: %v", r.GetMessage(), r.GetId(), r.GetName(), r.GetDomainLow(), r.GetDomainMid(), r.GetDomainHigh(), r.GetStatePartialValue(), r.GetStateFinalValue())
	}
	if *dropOperPtr {
		r, err := c.DropOneToken(ctx, &pb.Token{
			Id: *idPtr,
		})
		if err != nil {
			log.Fatalf("Client: failed to call server DropOneToken(): %v", err)
		}
		log.Printf("Client received: %v \n", r.GetMessage())
	}
}
