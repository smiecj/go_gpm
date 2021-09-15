g_trace:
	go build -o main trace/main.go
	./main

g_debug:
	go build -o main debug/main.go
	GODEBUG=schedtrace=1000 ./main