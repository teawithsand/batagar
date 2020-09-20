DIRS = . ./events ./fuzz ./internal/pb

PROTO_FILES = main.proto

ci:
	go build $(DIRS)
	go test $(DIRS)

test:
	go test $(DIRS)

build:
	go build $(DIRS)

vet:
	go vet $(DIRS)

fmt:
	go fmt $(DIRS)

generate:
	# Note: protoc compiler must be provided in $PATH already
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install golang.org/x/tools/cmd/stringer
	
	protoc -I=./proto \
			--go_opt=paths=source_relative \
			--go-grpc_opt=paths=source_relative \
			--go_out=./internal/pb/ --go-grpc_out=./internal/pb/ \
			$(PROTO_FILES)

	go generate $(DIRS)