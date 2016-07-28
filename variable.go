package particle

// Variable contains information about a Particle variable
type Variable struct {
	PhotonID string
	Name     string
	Value    float64
}

// Function contains information about a Particle function
type Function struct {
	Name        string
	ReturnValue int
	PhotonID    string
}

type particleVariable struct {
	CMD    string  `json:"cmd"`
	Name   string  `json:"name"`
	Result float64 `json:"result"`
	// TODO handle different response format. The CoreInfo is not the same with the response struct
	//Photon *response `json:"coreInfo"`
}

type particleFunctionResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	LastApp     string `json:"last_app"`
	Connected   bool   `json:"connected"`
	ReturnValue int    `json:"return_value"`
}

func toVariable(p *particleVariable) *Variable {
	// TODO mszg Handle the Photon information
	//if p.Name == "" || p.Photon == nil {
	//	return nil
	//}
	var v Variable
	v.Name = p.Name
	//v.PhotonID = p.Photon.ID
	v.Value = p.Result
	return &v
}

func toFunction(r particleFunctionResponse) *Function {
	var f Function
	f.Name = r.Name
	f.PhotonID = r.ID
	f.ReturnValue = r.ReturnValue
	return &f
}
