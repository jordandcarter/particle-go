package model

// Photon model for storing a Photon
type Photon struct {
	// Particle Cloud defined variables
	ID        string `json:"id"`
	Name      string `json:"name"`
	Claimed   bool   `json:"claimed"`
	Connected bool   `son:"connected"`
	Status    string `json:"status"`
	IPAddress string `json:"ipaddress"`

	// Photon software defined variables and functions
	Variables []*Variable `json:"variables"`
	Functions []*Function `json:"functions"`
}

// GetVariable return the Photon Variable
func (p *Photon) GetVariable(name string) *Variable {
	h, i := p.HasVariable(name)
	if !h {
		return nil
	}
	return p.Variables[i]
}

// SetVariable set a variable in the photon, if the variable was already added
// then it is only an update operation, otherwise it is going to add the new variable
func (p *Photon) SetVariable(v *Variable) {
	h, i := p.HasVariable(v.Name)
	if !h {
		p.Variables = append(p.Variables, v)
		return
	}
	p.Variables[i] = v
}

// HasVariable check if the photon has the given variable.
// Return the result of the search and the index of the variable.
func (p *Photon) HasVariable(name string) (bool, int) {
	for i, v := range p.Variables {
		if v.Name == name {
			return true, i
		}
	}
	return false, 0
}

// HasFunction check if the photon has the given function.
// Return the result of the search and the index of the function.
func (p *Photon) HasFunction(name string) (bool, int) {
	for i, f := range p.Functions {
		if f.Name == name {
			return true, i
		}
	}
	return false, 0
}
