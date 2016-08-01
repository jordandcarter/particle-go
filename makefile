
all: clean build

build: deps
	go build

deps:
	glide install

clean:
	rm -rf vendor

lint:
	gometalinter --vendor --deadline=300s ./...

get-tools:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install --update
