
generate:
	go run cmd/autogenerate/autogenerate.go

clean:
	-rm cf

cf: generate
	go build

all: clean cf

.PHONY: all
