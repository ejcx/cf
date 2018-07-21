
generate:
	go run cmd/autogenerate/autogenerate.go

clean:
	-rm cf
	-rm -rf dist

cf: generate
	go build

all: clean cf

dist-linux:
	GOOS=linux go build -o dist/cf-linux

dist:
	-mkdir dist

dist-osx:
	GOOS=darwin go build -o dist/cf-osx

release: clean dist dist-linux dist-osx

.PHONY: all
