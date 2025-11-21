# Specific version of go
FROM golang:1.25.4

# where we want to start the application in the docker container
WORKDIR /app

# loads these for caching so it only runs when they change
COPY go.mod ./

RUN go mod download

COPY . .

RUN make .

CMD ["./app/bin/app"]
