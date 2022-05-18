deps:
	go mod download
	go mod verify
	go mod tidy

lint:
	gofmt -w=true -s=true -l=true ./example/
	gofmt -w=true -s=true -l=true ./investapi/
	golint ./...
	go vet ./...

check:
	go test -v ./investapi/

generate:
	./generate.sh
