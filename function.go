package particle

// Function contains information about a Particle function
type Function struct {
	Name        string
	ReturnValue int
	PhotonID    string
}

type particleFunctionResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	LastApp     string `json:"last_app"`
	Connected   bool   `json:"connected"`
	ReturnValue int    `json:"return_value"`
}

func (r *particleFunctionResponse) toFunction() *Function {
	var f Function
	f.Name = r.Name
	f.PhotonID = r.ID
	f.ReturnValue = r.ReturnValue
	return &f
}
