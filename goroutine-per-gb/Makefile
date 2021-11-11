build:
	docker image build . -t goroutine-per-gb

test: build
	docker container run --rm -it -m 1g --memory-swappiness 0 goroutine-per-gb

clean:
	docker image rm goroutine-per-gb:latest

.PHONY: build test clean
