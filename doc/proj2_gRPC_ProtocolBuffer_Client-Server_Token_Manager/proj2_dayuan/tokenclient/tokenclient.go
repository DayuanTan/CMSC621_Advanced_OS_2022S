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

const (
	defaultName = "Dayuan"
)

var (
	// id               int    = -1
	// name string = "default no name"
	// low              uint64 = 0
	// mid              uint64 = 0
	// high             uint64 = 0
	// host             string = "localhost"
	// port             int    = 50051
	// neededArgsAmount = map[string]int{
	// 	"-create": 6,
	// 	"-read":   6,
	// 	"-write":  14,
	// 	"-drop":   5,
	// }
	createOperPtr = flag.Bool("create", false, "Claim create operation. Is bool.")
	readOperPtr   = flag.Bool("read", false, "Claim read operation. Is bool.")
	writeOperPtr  = flag.Bool("write", false, "Claim write operation. Is bool.")
	dropOperPtr   = flag.Bool("drop", false, "Claim drop operation. Is bool.")
	idPtr         = flag.Int("id", -1, "The id (int) for your token")
	namePtr       = flag.String("name", defaultName, "The name (string) of your token. Default is 'Dayuan'")
	lowPtr        = flag.Uint64("low", 0, "The low value (uint64) of your token")
	midPtr        = flag.Uint64("mid", 0, "The mid value (uint64) of your token")
	highPtr       = flag.Uint64("high", 0, "The high value (uint64) of your token")
	hostPtr       = flag.String("host", "localhost", "The host (string) to connect to. Default is localhost")
	portPtr       = flag.String("port", "50051", "The port (string) to connect to. Default is '50051'")
)

// func parseArgsAndHandleErr(argsWithoutProg []string, caseNumber int) (int, string, int, uint64, uint64, uint64, string) {
// 	switch caseNumber {
// 	case 6:
// 		if len(argsWithoutProg) == 7 && argsWithoutProg[1] == "-id" && argsWithoutProg[3] == "-host" && argsWithoutProg[5] == "-port" {
// 			id, err := strconv.Atoi(argsWithoutProg[2])
// 			if err != nil {
// 				log.Fatalf("Client: cannot convert string to int: %v", err)
// 			}
// 			host = argsWithoutProg[4]
// 			port, err := strconv.Atoi(argsWithoutProg[6])
// 			if err != nil {
// 				log.Fatalf("Client: cannot convert string to int: %v", err)
// 			}
// 			return id, host, port, 0, 0, 0, ""
// 		} else {
// 			log.Fatalf("Client: wrong args amount. Example: -create -id 1234 -host localhost -port 50051")
// 		}
// 	case 14:
// 		if len(argsWithoutProg) == 15 && argsWithoutProg[1] == "-id" && argsWithoutProg[3] == "-name" && argsWithoutProg[5] == "-low" && argsWithoutProg[7] == "-mid" && argsWithoutProg[9] == "-high" && argsWithoutProg[11] == "-host" && argsWithoutProg[13] == "-port" {
// 			id, err := strconv.Atoi(argsWithoutProg[2])
// 			if err != nil {
// 				log.Fatalf("Client: cannot convert string to int: %v", err)
// 			}
// 			name = argsWithoutProg[4]
// 			low, err := strconv.ParseUint(argsWithoutProg[6], 10, 64)
// 			if err != nil {
// 				log.Fatalf("Client: cannot convert string to int: %v", err)
// 			}
// 			mid, err := strconv.ParseUint(argsWithoutProg[8], 10, 64)
// 			if err != nil {
// 				log.Fatalf("Client: cannot convert string to int: %v", err)
// 			}
// 			high, err := strconv.ParseUint(argsWithoutProg[10], 10, 64)
// 			if err != nil {
// 				log.Fatalf("Client: cannot convert string to int: %v", err)
// 			}
// 			host = argsWithoutProg[12]
// 			port, err := strconv.Atoi(argsWithoutProg[14])
// 			if err != nil {
// 				log.Fatalf("Client: cannot convert string to int: %v", err)
// 			}
// 			return id, host, port, low, mid, high, name
// 		} else {
// 			log.Fatalf("Client: wrong args amount. Example: -write -id 1234 -name abc -low 0 -mid 10 -high 100 -host localhost -port 50051")
// 		}
// 	case 5:
// 		if len(argsWithoutProg) == 6 && argsWithoutProg[2] == "-host" && argsWithoutProg[4] == "-port" {
// 			id, err := strconv.Atoi(argsWithoutProg[1])
// 			if err != nil {
// 				log.Fatalf("Client: cannot convert string to int: %v", err)
// 			}
// 			host = argsWithoutProg[3]
// 			port, err := strconv.Atoi(argsWithoutProg[5])
// 			if err != nil {
// 				log.Fatalf("Client: cannot convert string to int: %v", err)
// 			}
// 			return id, host, port, 0, 0, 0, ""
// 		} else {
// 			log.Fatalf("Client: wrong args amount. Example: -drop 1234 -host localhost -port 50051")
// 		}
// 	}
// 	return -1, "", -1, 0, 0, 0, ""
// }

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
	// Deal with command line arguments
	// argsWithoutProg := os.Args[1:]
	// fmt.Println(argsWithoutProg)
	// operation := argsWithoutProg[0]
	// id, host, port, low, mid, high, name = parseArgsAndHandleErr(argsWithoutProg, neededArgsAmount[operation])
	// fmt.Println(id, host, port, low, mid, high, name)

	flag.Parse()
	fmt.Println(*createOperPtr, *readOperPtr, *writeOperPtr, *dropOperPtr)
	fmt.Println(*idPtr, *hostPtr, *portPtr, *lowPtr, *midPtr, *highPtr, *namePtr)
	checkFlagsValidation()
	addr := getAddr()

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client: did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTokenServiceClient(conn)

	// Call server methods
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateOneToken(ctx, &pb.Token{Name: *namePtr})
	if err != nil {
		log.Fatalf("Client: could not greet: %v", err)
	}
	log.Printf("Client received: %s", r.GetName())
}
