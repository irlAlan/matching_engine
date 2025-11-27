# Specific version of go
FROM golang:1.25.4

# where we want to start the application in the docker container
WORKDIR /app

# loads these for caching so it only runs when they change
COPY go.mod go.sum ./

RUN go mod download

RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      protobuf-compiler-grpc protoc-gen-go-grpc protoc-gen-go



COPY . .

RUN make .

CMD ["./app/bin/app"]
