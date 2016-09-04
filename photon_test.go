package particle

import "testing"

func TestHasVariableSuccess(t *testing.T) {
	var p Photon
	h, _ := p.HasVariable("variable")
	if h {
		t.Error("HasVariable - failed")
	}
}

func TestGetVariableFail(t *testing.T) {
	var v Variable
	v.Name = "pi"
	v.Value = 3.14
	var p Photon
	p.SetVariable(&v)

	vv := p.GetVariable("var")
	if vv != nil {
		t.Error("GetVariable - failed")
	}
}

func TestSetAndGetVariableSuccess(t *testing.T) {
	var v Variable
	v.Name = "pi"
	v.Value = 3.14
	var p Photon
	p.SetVariable(&v)
	h, i := p.HasVariable("pi")
	if !h {
		t.Error("HasVariable")
	}
	if i != 0 {
		t.Error("HasVariable - get index failed")
	}
	vv := p.GetVariable("pi")
	if vv.Name != "pi" {
		t.Error("GetVariable - get name failed")
	}
	if vv.Value != 3.14 {
		t.Error("GetVariable - get value failed")
	}
}

func TestHasFunction(t *testing.T) {
	var p Photon
	h, _ := p.HasFunction("SayHello")
	if h {
		t.Error("HasFunction - failed")
	}
}
