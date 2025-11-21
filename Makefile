# run: build
#

build: app/main.go
	go build -C app/ -o bin/app main.go

clean:
	rm -rfv app/bin/*
