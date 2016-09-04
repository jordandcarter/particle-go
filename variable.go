package particle

// Variable contains information about a Particle variable
type Variable struct {
	PhotonID string
	Name     string
	Value    float64
}

type particleVariable struct {
	CMD    string  `json:"cmd"`
	Name   string  `json:"name"`
	Result float64 `json:"result"`
	// TODO handle different response format. The CoreInfo is not the same with the response struct
	//Photon *response `json:"coreInfo"`
}

func (p *particleVariable) toVariable() *Variable {
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
