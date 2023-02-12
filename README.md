# Prerequisites

## install Go

`brew install go`

# GRPC

## install protoc

`brew install protobuf`

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## add to PATH

```export PATH="$PATH:$(go env GOPATH)/bin"```

##

# Add dependencies

go get github.com/mel3kings/event-driven-architecture/events
