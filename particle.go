package particle

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
	"github.com/donovanhide/eventsource"
)

// Particle struct to access the Particle Cloud
type Particle struct {
	token      string
	devicesURL string
	access     string
}

type response struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	LastApp       string   `json:"last_app"`
	LastIPAddress string   `json:"last_ip_address"`
	LastHeard     string   `json:"last_heard"`
	Status        string   `json:"status"`
	ProductID     int      `json:"product_id"`
	PlatformID    int      `json:"platform_id"`
	Connected     bool     `json:"connected"`
	Cellular      bool     `json:"cellular"`
	Functions     []string `json:"functions"`
}

func (resp *response) toPhoton() *Photon {
	var photon Photon
	photon.ID = resp.ID
	photon.Name = resp.Name
	photon.Claimed = false
	photon.Connected = resp.Connected
	photon.Status = resp.Status
	photon.IPAddress = resp.LastIPAddress

	photon.Functions = make([]*Function, len(resp.Functions))
	for i, s := range resp.Functions {
		f := &Function{Name: s}
		photon.Functions[i] = f
	}

	return &photon
}

const (
	particleAPI = "https://api.particle.io"
	version     = "/v1"
	devices     = "/devices"
	events      = "/events"
	accessToken = "?access_token="
)

var (
	errParticleCloud = errors.New("Error while requesting the photons from the Particle Cloud")
	errBadInputNil   = errors.New("The given input is nil")
)

// Load initialize the particle token
func Load(particleToken string) *Particle {
	log.Infof("Particle: Load: Creating Particle with token: %q", particleToken)
	return &Particle{token: particleToken, devicesURL: particleAPI + version + devices, access: accessToken + particleToken}
}

// GetPhotons from the Particle Cloud
func (p *Particle) GetPhotons() ([]*Photon, error) {
	var r []response
	err := getInformationFromParticle(p.devicesURL+p.access, &r)
	if err != nil {
		log.Error("Particle: GetPhotons: read body failed: ", err)
		return nil, err
	}
	var photons []*Photon
	photons = make([]*Photon, len(r))
	for i := 0; i < len(r); i++ {
		log.Debug("Particle: GetPhotons: Device name: ", r[i].Name, " device id: %q", r[i].ID)
		photons[i] = r[i].toPhoton()
	}
	return photons, nil
}

// GetPhoton get a photon information from the Particle Cloud
func (p *Particle) GetPhoton(id string) (*Photon, error) {
	var resp response
	err := getInformationFromParticle(p.devicesURL+"/"+id+p.access, &resp)
	if err != nil {
		log.Error("Particle: GetPhoton: read body failed: ", err)
		return nil, err
	}
	return resp.toPhoton(), nil
}

// GetEvent get a given event from the particle cloud
func (p *Particle) GetEvent(c chan *Event, name string) {
	stream, err := eventsource.Subscribe(particleAPI+version+events+"/"+name+p.access, "")
	if err != nil {
		return
	}
	for {
		ev := <-stream.Events
		event, err := ToEvent(ev)
		if err == nil {
			log.WithFields(log.Fields{"event": event}).Debug("Particle: GetEvents: received particle event")
			c <- event
		}
	}
}

// GetVariable get a given photon variable information from the Particle Cloud
func (p *Particle) GetVariable(ph *Photon, variable string) (*Variable, error) {
	var v particleVariable
	if p == nil {
		log.Error("Particle: getInformationFromParticle: the given photon is nil")
		return nil, errBadInputNil
	}
	err := getInformationFromParticle(p.devicesURL+"/"+ph.ID+"/"+variable+p.access, &v)
	if err != nil {
		log.Error("Particle: getInformationFromParticle: error while requesting the information from the Particle Cloud: ", err)
		return nil, err
	}
	return v.toVariable(), nil
}

// CallFunction call a given Particle function on the Photon
func (p *Particle) CallFunction(ph *Photon, function string, command string) (*Function, error) {
	u := p.devicesURL + "/" + ph.ID + "/" + function
	log.WithFields(log.Fields{"url": u}).Debug("CallFunction called with url")
	resp, err := http.PostForm(u, url.Values{"arg": {command}, "access_token": {p.token}})
	if err != nil {
		log.Error("Particle: CallFunction: error while requesting the information from the Particle Cloud: ", err)
		return nil, errParticleCloud
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Particle: CallFunction: cannot read the data from the body: ", err)
		return nil, err
	}
	var r particleFunctionResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Error("Particle: CallFunction: json marshalling failed: ", err)
		return nil, err
	}
	return r.toFunction(), nil
}

func getInformationFromParticle(url string, data interface{}) error {
	log.WithFields(log.Fields{"url": url}).Debug("getInformationFromParticle called with url")
	resp, err := http.Get(url)
	if err != nil {
		log.Error("Particle: getInformationFromParticle: error while requesting the information from the Particle Cloud: ", err)
		return errParticleCloud
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Particle: getInformationFromParticle: cannot read the data from the body: ", err)
		return err
	}
	log.Info(string(body))
	err = json.Unmarshal(body, data)
	if err != nil {
		log.Error("Particle: getInformationFromParticle: json marshalling failed: ", err)
		return err
	}
	return nil
}
