module github.com/stmcodes/go-proto-gql

go 1.16

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/danielvladco/go-proto-gql v0.8.3
	github.com/golang/protobuf v1.5.2
	github.com/jhump/protoreflect v1.8.2
	github.com/mitchellh/mapstructure v1.3.0
	github.com/nautilus/gateway v0.1.4
	github.com/nautilus/graphql v0.0.9
	github.com/vektah/gqlparser/v2 v2.1.0
	google.golang.org/grpc v1.33.2
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.1
	google.golang.org/protobuf v1.26.0
)

replace github.com/99designs/gqlgen v0.13.0 => github.com/danielvladco/gqlgen v0.13.0-inputs

replace github.com/danielvladco/go-proto-gql => ./
