test:
	go test -v ./...

bench:
	go test -bench=./... -benchmem

cover:
	go test -cover