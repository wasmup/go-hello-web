all: 
	go mod tidy -v
	go clean
	time go build -ldflags=-s

test:
# go test -timeout 30s -cover	
	go test -race .