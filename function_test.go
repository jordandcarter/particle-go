package particle

import "testing"

func TestToFunction(t *testing.T) {
	var r particleFunctionResponse
	r.ID = "id"
	r.Connected = true
	r.Name = "functionName"
	r.LastApp = "lastApplicationName"
	r.ReturnValue = 0
	f := r.toFunction()
	if f.Name != r.Name {
		t.Error("Function name is not equal.")
	}
	if f.PhotonID != r.ID {
		t.Error("Photon ID is not equal.")
	}
	if f.ReturnValue != r.ReturnValue {
		t.Error("Return value is not equal.")
	}
}
