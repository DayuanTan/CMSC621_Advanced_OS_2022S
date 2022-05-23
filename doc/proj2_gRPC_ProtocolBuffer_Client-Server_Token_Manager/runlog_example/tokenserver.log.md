
# Note
There is only  one command then all others are log of server side.

Commands are staring with "directory$".

# Command 1: proj2_dayuan/tokenserver$ go run tokenserver.go -port 50051

proj2_dayuan/tokenserver$ go run tokenserver.go -port 50051
2022/05/20 23:28:03 Server: listening at [::]:50051

2022/05/20 23:28:10 Server received: to create ID: 222
2022/05/20 23:28:10 Creating ID 222 successed.
2022/05/20 23:28:10 Current tokenList is:
	ID: 222;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

2022/05/20 23:28:15 Server received: to create ID: 222
2022/05/20 23:28:15 Creating ID 222 failed.

2022/05/20 23:28:20 Server received: to create ID: 12345
2022/05/20 23:28:20 Creating ID 12345 successed.
2022/05/20 23:28:20 Current tokenList is:
	ID: 222;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

	ID: 12345;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

2022/05/20 23:28:27 Server received: to create ID: 2222
2022/05/20 23:28:27 Creating ID 2222 successed.
2022/05/20 23:28:27 Current tokenList is:
	ID: 222;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

	ID: 12345;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

	ID: 2222;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

2022/05/20 23:28:35 Server received: to write ID: 2222,
Name: def,
DomainLow: 0,
DomainMid: 10,
DomainHigh: 100
2022/05/20 23:28:35 Writing ID 2222 successed.
2022/05/20 23:28:35 Current tokenList is:
	ID: 222;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

	ID: 12345;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

	ID: 2222;
	Name: def;
	DomainLow: 0;
	DomainMid: 10;
	DomainHigh: 100;
	StatePartialValue: 7;
	StateFinalValue: 0

2022/05/20 23:28:42 Server received: to write ID: 12345,
Name: abc,
DomainLow: 0,
DomainMid: 10,
DomainHigh: 100
2022/05/20 23:28:42 Writing ID 12345 successed.
2022/05/20 23:28:42 Current tokenList is:
	ID: 222;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

	ID: 12345;
	Name: abc;
	DomainLow: 0;
	DomainMid: 10;
	DomainHigh: 100;
	StatePartialValue: 4;
	StateFinalValue: 0

	ID: 2222;
	Name: def;
	DomainLow: 0;
	DomainMid: 10;
	DomainHigh: 100;
	StatePartialValue: 7;
	StateFinalValue: 0

2022/05/20 23:29:00 Server received: to write ID: 2,
Name: def,
DomainLow: 0,
DomainMid: 10,
DomainHigh: 100
2022/05/20 23:29:00 Writing ID 2 failed.
2022/05/20 23:29:00 Current tokenList is:
	ID: 222;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

	ID: 12345;
	Name: abc;
	DomainLow: 0;
	DomainMid: 10;
	DomainHigh: 100;
	StatePartialValue: 4;
	StateFinalValue: 0

	ID: 2222;
	Name: def;
	DomainLow: 0;
	DomainMid: 10;
	DomainHigh: 100;
	StatePartialValue: 7;
	StateFinalValue: 0

2022/05/20 23:29:13 Server received: to read ID: 2222
2022/05/20 23:29:13 Reading ID 2222 successed and StateFinalValue updated.
2022/05/20 23:29:13 Current tokenList is:
	ID: 222;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

	ID: 12345;
	Name: abc;
	DomainLow: 0;
	DomainMid: 10;
	DomainHigh: 100;
	StatePartialValue: 4;
	StateFinalValue: 0

	ID: 2222;
	Name: def;
	DomainLow: 0;
	DomainMid: 10;
	DomainHigh: 100;
	StatePartialValue: 7;
	StateFinalValue: 7

2022/05/20 23:29:51 Server received: to read ID: 12345
2022/05/20 23:29:51 Reading ID 12345 successed and StateFinalValue updated.
2022/05/20 23:29:51 Current tokenList is:
	ID: 222;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

	ID: 12345;
	Name: abc;
	DomainLow: 0;
	DomainMid: 10;
	DomainHigh: 100;
	StatePartialValue: 4;
	StateFinalValue: 4

	ID: 2222;
	Name: def;
	DomainLow: 0;
	DomainMid: 10;
	DomainHigh: 100;
	StatePartialValue: 7;
	StateFinalValue: 7

2022/05/20 23:30:05 Server received: to read ID: 1234
2022/05/20 23:30:05 Reading ID 1234 failed.

2022/05/20 23:30:39 Server received: to drop ID: 1234
2022/05/20 23:30:39 Dropping ID 1234 failed.
2022/05/20 23:30:39 Current tokenList is:
	ID: 222;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

	ID: 12345;
	Name: abc;
	DomainLow: 0;
	DomainMid: 10;
	DomainHigh: 100;
	StatePartialValue: 4;
	StateFinalValue: 4

	ID: 2222;
	Name: def;
	DomainLow: 0;
	DomainMid: 10;
	DomainHigh: 100;
	StatePartialValue: 7;
	StateFinalValue: 7

2022/05/20 23:30:54 Server received: to drop ID: 12345
2022/05/20 23:30:54 Dropping ID 12345 successed.
2022/05/20 23:30:54 Current tokenList is:
	ID: 222;
	Name: ;
	DomainLow: 0;
	DomainMid: 0;
	DomainHigh: 0;
	StatePartialValue: 0;
	StateFinalValue: 0

	ID: 2222;
	Name: def;
	DomainLow: 0;
	DomainMid: 10;
	DomainHigh: 100;
	StatePartialValue: 7;
	StateFinalValue: 7

2022/05/22 21:15:39 Server received: to read ID: 12345
2022/05/22 21:15:39 Reading ID 12345 failed.


^Csignal: interrupt
proj2_dayuan/tokenserver$