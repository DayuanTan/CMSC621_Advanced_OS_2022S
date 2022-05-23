# Note
There is 12 commands and their corresponding logs of client side.

Commands are staring with "directory$".

# Command 1: proj2_dayuan/tokenclient$ go run tokenclient.go -create -id 222 -host localhost -port 50051

proj2_dayuan/tokenclient$ go run tokenclient.go -create -id 222 -host localhost -port 50051
2022/05/20 23:28:10 Your input is:
Operation: Create:  true ; Read:  false ; Write:  false ; Drop:  false
Parameters: ID:  222 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  0 ; High:  0 ; Name:

2022/05/20 23:28:10 Client received:
ID: 222, Message: Created ID 222 successed.
# Command 2: proj2_dayuan/tokenclient$ go run tokenclient.go -create -id 222 -host localhost -port 50051

proj2_dayuan/tokenclient$ go run tokenclient.go -create -id 222 -host localhost -port 50051
2022/05/20 23:28:15 Your input is:
Operation: Create:  true ; Read:  false ; Write:  false ; Drop:  false
Parameters: ID:  222 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  0 ; High:  0 ; Name:

2022/05/20 23:28:15 Client: failed to call server CreateOneToken(): rpc error: code = Unknown desc = id 222 already exists
exit status 1
# Command 3: proj2_dayuan/tokenclient$ go run tokenclient.go -create -id 12345 -host localhost -port 50051

proj2_dayuan/tokenclient$ go run tokenclient.go -create -id 12345 -host localhost -port 50051
2022/05/20 23:28:20 Your input is:
Operation: Create:  true ; Read:  false ; Write:  false ; Drop:  false
Parameters: ID:  12345 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  0 ; High:  0 ; Name:

2022/05/20 23:28:20 Client received:
ID: 12345, Message: Created ID 12345 successed.
# Command 4: proj2_dayuan/tokenclient$ go run tokenclient.go -create -id 2222 -host localhost -port 50051

proj2_dayuan/tokenclient$ go run tokenclient.go -create -id 2222 -host localhost -port 50051
2022/05/20 23:28:27 Your input is:
Operation: Create:  true ; Read:  false ; Write:  false ; Drop:  false
Parameters: ID:  2222 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  0 ; High:  0 ; Name:

2022/05/20 23:28:27 Client received:
ID: 2222, Message: Created ID 2222 successed.
# Command 5: proj2_dayuan/tokenclient$ go run tokenclient.go -write -id 2222 -name def -low 0 -mid 10 -high 100 -host localhost -port 50051
proj2_dayuan/tokenclient$ go run tokenclient.go -write -id 2222 -name def -low 0 -mid 10 -high 100 -host localhost -port 50051
2022/05/20 23:28:35 Your input is:
Operation: Create:  false ; Read:  false ; Write:  true ; Drop:  false
Parameters: ID:  2222 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  10 ; High:  100 ; Name:  def

2022/05/20 23:28:35 Client received: Message: Wrote ID 2222 successed.,
ID: 2222,
Name: def,
DomainLow: 0,
DomainMid: 10,
DomainHigh: 100,
StatePartialValue: 7,
StateFinalValue: 0
# Command 6: proj2_dayuan/tokenclient$ go run tokenclient.go -write -id 12345 -name abc -low 0 -mid 10 -high 100 -host localhost -port 50051
proj2_dayuan/tokenclient$ go run tokenclient.go -write -id 12345 -name abc -low 0 -mid 10 -high 100 -host localhost -port 50051
2022/05/20 23:28:42 Your input is:
Operation: Create:  false ; Read:  false ; Write:  true ; Drop:  false
Parameters: ID:  12345 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  10 ; High:  100 ; Name:  abc

2022/05/20 23:28:42 Client received: Message: Wrote ID 12345 successed.,
ID: 12345,
Name: abc,
DomainLow: 0,
DomainMid: 10,
DomainHigh: 100,
StatePartialValue: 4,
StateFinalValue: 0
# Command 7: proj2_dayuan/tokenclient$ go run tokenclient.go -write -id 2 -name def -low 0 -mid 10 -high 100 -host localhost -port 50051
proj2_dayuan/tokenclient$ go run tokenclient.go -write -id 2 -name def -low 0 -mid 10 -high 100 -host localhost -port 50051
2022/05/20 23:29:00 Your input is:
Operation: Create:  false ; Read:  false ; Write:  true ; Drop:  false
Parameters: ID:  2 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  10 ; High:  100 ; Name:  def

2022/05/20 23:29:00 Client: failed to call server WriteOneToken(): rpc error: code = Unknown desc = id 2 was not found
exit status 1
# Command 8: proj2_dayuan/tokenclient$ go run tokenclient.go -read -id 2222 -host localhost -port 50051

proj2_dayuan/tokenclient$ go run tokenclient.go -read -id 2222 -host localhost -port 50051
2022/05/20 23:29:13 Your input is:
Operation: Create:  false ; Read:  true ; Write:  false ; Drop:  false
Parameters: ID:  2222 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  0 ; High:  0 ; Name:

2022/05/20 23:29:13 Client received: Message: Read ID 2222 successed.,
ID: 2222,
Name: def,
DomainLow: 0,
DomainMid: 10,
DomainHigh: 100,
StatePartialValue: 7,
StateFinalValue: 7
# Command 9: proj2_dayuan/tokenclient$ go run tokenclient.go -read -id 12345 -host localhost -port 50051

proj2_dayuan/tokenclient$ go run tokenclient.go -read -id 12345 -host localhost -port 50051
2022/05/20 23:29:51 Your input is:
Operation: Create:  false ; Read:  true ; Write:  false ; Drop:  false
Parameters: ID:  12345 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  0 ; High:  0 ; Name:

2022/05/20 23:29:51 Client received: Message: Read ID 12345 successed.,
ID: 12345,
Name: abc,
DomainLow: 0,
DomainMid: 10,
DomainHigh: 100,
StatePartialValue: 4,
StateFinalValue: 4
# Command 10: proj2_dayuan/tokenclient$ go run tokenclient.go -read -id 1234 -host localhost -port 50051

proj2_dayuan/tokenclient$ go run tokenclient.go -read -id 1234 -host localhost -port 50051
2022/05/20 23:30:05 Your input is:
Operation: Create:  false ; Read:  true ; Write:  false ; Drop:  false
Parameters: ID:  1234 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  0 ; High:  0 ; Name:

2022/05/20 23:30:05 Client: failed to call server ReadOneToken(): rpc error: code = Unknown desc = id 1234 was not found
exit status 1
# Command 11: proj2_dayuan/tokenclient$ go run tokenclient.go -drop -id 1234 -host localhost -port 50051

proj2_dayuan/tokenclient$ go run tokenclient.go -drop -id 1234 -host localhost -port 50051
2022/05/20 23:30:39 Your input is:
Operation: Create:  false ; Read:  false ; Write:  false ; Drop:  true
Parameters: ID:  1234 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  0 ; High:  0 ; Name:

2022/05/20 23:30:39 Client: failed to call server DropOneToken(): rpc error: code = Unknown desc = id 1234 was not found
exit status 1
# Command 12: proj2_dayuan/tokenclient$ go run tokenclient.go -drop -id 12345 -host localhost -port 50051

proj2_dayuan/tokenclient$ go run tokenclient.go -drop -id 12345 -host localhost -port 50051
2022/05/20 23:30:54 Your input is:
Operation: Create:  false ; Read:  false ; Write:  false ; Drop:  true
Parameters: ID:  12345 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  0 ; High:  0 ; Name:

2022/05/20 23:30:54 Client received: Dropped 12345 successed

# Command 13: proj2_dayuan/tokenclient$ go run tokenclient.go -read -id 12345 -host localhost -port 50051

proj2_dayuan/tokenclient$ go run tokenclient.go -read -id 12345 -host localhost -port 50051
2022/05/22 21:15:39 Your input is:
Operation: Create:  false ; Read:  true ; Write:  false ; Drop:  false
Parameters: ID:  12345 ; Host:  localhost ; Port:  50051 ; Low:  0 ; Mid:  0 ; High:  0 ; Name:

2022/05/22 21:15:39 Client: failed to call server ReadOneToken(): rpc error: code = Unknown desc = id 12345 was not found
exit status 1