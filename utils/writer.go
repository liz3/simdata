package utils

import (
	"encoding/json"
	"os"
	"simdata/models"
	"strconv"
	"time"
)

type Instance struct {
	Path      string
	entries   []*models.CarTelemetry
	created   bool
	CarName   string
	Time      string
	Game      string
	TrackName string
	Ready     bool
}

func CreateInstance() *Instance {
	var in = &Instance{
		created:   false,
		Time:      strconv.FormatInt(time.Now().Unix(), 10),
		Ready:     false,
		Game:      "",
		TrackName: "",
		CarName:   "",
	}
	return in
}

func (instance *Instance) WriteHeader() {
	if instance.created {
		return
	}
	f, err := os.OpenFile(instance.Path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		f.Write([]byte("----\ngame " + instance.Game + "\ntime " + instance.Time + "\ntrack " + instance.TrackName + "\ncar " + instance.CarName + "\n----\n\n"))
		f.Close()
		instance.created = true
	}
}

func (instance *Instance) Flush()  {
	f, err := os.OpenFile(instance.Path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		for _, entry := range instance.entries {
			bytes, err := json.Marshal(entry)
			if err == nil {
				f.Write(bytes)
				f.Write([]byte(",\n"))
			}
		}
		instance.entries = []*models.CarTelemetry{}
		f.Close()
	}
}

func (instance *Instance) Push(entry *models.CarTelemetry) {
	instance.entries = append(instance.entries, entry)
	if len(instance.entries) >= 300 {
		instance.WriteHeader()
		instance.Flush()
	}
}
