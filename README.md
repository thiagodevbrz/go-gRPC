* It is necessary to install protoc to compile protobuf


## To install protobuf features execute the follow commands:

sudo apt install protobuf-compiler

# Install go
https://go.dev/doc/install


# To initialize the project
go mod init github.com/thiagodevbrz/go-gRPC


# Add golang proto and gRPC features
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
sudo apt install golang-goprotobuf-dev


### Para compilar os arquivos protos dentro da pasta proto e gerar os arquivos na pasta pb

protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.


## Use Evans to test gRPC call by CI 
Project: https://github.com/ktr0731/evans

