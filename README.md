# particle-go
Golang client library to access the Particle Cloud functionalities.

## Install the library

Run the `go get github.com/matisszilard/particle-go` command.

## Build the library

You can build the library with the provided `makefile` by running `make all` command. The makefile will create a clean build, get the dependencies, and run `go build`.

The library use `glide` as a package manager. For more information please check the `glide` official github page: https://github.com/Masterminds/glide.

### Makefile targets

```
# Create a clean build
all: clean build

# Build the library with the dependencies
build: deps
	go build

# Install the library dependencies
deps:
	glide install

# Remove the vendor folder
clean:
	rm -rf vendor

# Run the linter command
lint:
	gometalinter --vendor --deadline=300s ./...

# Get the required tools
get-tools:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install --update

```

# Licensing
Please check the `LICENSE` file under the source. (MIT)
