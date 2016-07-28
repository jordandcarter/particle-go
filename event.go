package particle

import (
	"encoding/json"

	"github.com/donovanhide/eventsource"
)

// Event represent an event from the Particle Cloud
type Event struct {
	ID   string
	Name string
	Data *Data
}

// Data represent the data struct from the Particle Cloud
type Data struct {
	Data        string `json:"data"`
	TTL         string `json:"ttl"`
	PublishTime string `json:"published_at"`
	PhotonID    string `json:"coreid"`
}

// ToEvent convert an eventsource.Event into an Event struct
func ToEvent(e eventsource.Event) (*Event, error) {
	event := Event{}

	event.ID = e.Id()
	event.Name = e.Event()
	data, err := ToData(e.Data())
	if err != nil {
		return nil, err
	}
	event.Data = data
	return &event, nil
}

// ToData convert a json string to Data struct
func ToData(j string) (*Data, error) {
	data := Data{}
	err := json.Unmarshal([]byte(j), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
