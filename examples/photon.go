package main

import (
	"fmt"
	"os"

	"github.com/matisszilard/particle-go"
)

func main() {
	// Initialize the particle-go library
	// Set a proper Particle token to access the Particle Cloud
	// Make sure that you don't share any private tokin in this repository
	p := particle.Load(os.Getenv("PARTICLE_TOKEN"))

	// Get all Photons claimed by the current user (token)
	ps, err := p.GetPhotons()

	// Always cneck the return values of the functions
	if err != nil {
		fmt.Println("Error while requesting photons : ", err)
		return
	}

	// Iterate over the requested photons
	for _, ph := range ps {
		fmt.Println("Photon - [", ph.ID, "]")

		// Get again the photons one-by-one
		pr, err := p.GetPhoton(ph.ID)
		if err != nil {
			fmt.Println("Error while getting Photon information - ", ph)
			return
		}
		fmt.Println("Requested Photon - [", pr, "]")
	}

}
