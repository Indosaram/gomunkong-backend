go get -u google.golang.org/grpc
go get google.golang.org/protobuf/reflect/protoreflect@v1.26.0
go get google.golang.org/protobuf/runtime/protoimpl@v1.26.0

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
# Add GOBIN to PATH
export PATH="$PATH:$(go env GOPATH)/bin"