all: 
	time go build -ldflags=-s
# real    0m0.421s
# user    0m0.701s
# sys     0m0.176s

test:
# go test -timeout 30s -cover	
	go test -race .