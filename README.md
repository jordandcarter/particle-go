# particle-go
Golang client library to access the Particle Cloud functionalities.

## Install the library

Run the `go get github.com/matisszilard/particle-go` command.

## How to use
In the first step the library must be initialized with a valid Particle Token.
```
// Initialize the particle-go library
// Set a proper Particle token to access the Particle Cloud
// Make sure that you don't share any private token in this repository
p := particle.Load(os.Getenv("PARTICLE_TOKEN"))
```

The library provide the following functionalities:
```
// Request all of the claimed particle photons
func (p *Particle) GetPhotons() ([]*Photon, error)

// Request a given Photon by ID
func (p *Particle) GetPhoton(id string) (*Photon, error)

// Subscribe for a specific event. The events are pushed to the "Event" channel set in the function parameters
func (p *Particle) GetEvent(c chan *Event, name string)

// Get a variable information from the Particle Cloud
func (p *Particle) GetVariable(ph *Photon, variable string) (*Variable, error)

// Call a function of a Photon
func (p *Particle) CallFunction(ph *Photon, function string, command string) (*Function, error)
```

NOTE: The library should work with other hardware devices too, but the functionalities are not tested with a different Particle hardware devices, only with Particle Photon.

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

# License
Please check the `LICENSE` file under the source. (MIT)
