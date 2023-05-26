```sh
	go clean
	time go build -ldflags=-s
# real    0m0.340s
# user    0m0.389s
# sys     0m0.164s

	time go build -ldflags=-s
# real    0m0.126s
# user    0m0.187s
# sys     0m0.100s
```