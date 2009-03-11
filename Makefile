all: build

build:
	go build -o flog ./cmd/flog

test:
	go test ./...

grid:
	./flog grid -year 2010

clean:
	rm -f flog

.PHONY: all build test grid clean
