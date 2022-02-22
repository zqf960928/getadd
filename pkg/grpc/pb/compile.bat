:: Install proto3.
:: https://github.com/google/protobuf/releases
:: Update protoc Go bindings via
::  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
::  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
::
:: See also
::  https://github.com/grpc/grpc-go/tree/master/examples

protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative GetAdd