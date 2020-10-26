package main

import (
	"simdata/models"
	"simdata/network"
	"simdata/utils"
	"strconv"
	"time"
)

func main() {

	go func() {
		var name = "logs/" + strconv.FormatInt(time.Now().Unix(), 10) + ".sd"
		var instance = utils.CreateInstance()
		instance.Game = "F12020"
		instance.Path = name
		var sessionId uint64 = 0
		network.StartProcessor(utils.F12020, 20777, func(telemetry *models.CarTelemetry) {
			instance.Push(telemetry)
		}, func(name string, entry interface{}) {
			if name == "CAR_NAME" && instance.CarName == "" {
				instance.CarName = entry.(string)
			}
			if name == "TRACK_NAME" && instance.TrackName == "" {
				instance.TrackName = entry.(string)
			}
			if name == "GAME_STATE"  {
				var passed = entry.(uint64)
				if passed != sessionId && sessionId != 0 {
					name = "logs/" + strconv.FormatInt(time.Now().Unix(), 10) + ".sd"
					instance = utils.CreateInstance()
					instance.Game = "F12020"
					instance.Path = name
				}
				sessionId = passed

			}
			if instance.CarName != "" && instance.TrackName != "" {
				instance.WriteHeader()
			}
		})
	}()
	go func() {
		var name = "logs/" + strconv.FormatInt(time.Now().Unix()+2, 10) + ".sd"
		var instance = utils.CreateInstance()
		instance.Game = "PC2"
		instance.Path = name
		var lastState uint8 = 0
		network.StartProcessor(utils.ProjectCars2, 5606, func(telemetry *models.CarTelemetry) {
			instance.Push(telemetry)
		}, func(name string, entry interface{}) {
			if name == "CAR_NAME" && instance.CarName == "" {
				instance.CarName = entry.(string)
			}
			if name == "TRACK_NAME" && instance.TrackName == "" {
				instance.TrackName = entry.(string)
			}
			if name == "GAME_STATE" {
				var passed = entry.(uint8)
				if lastState == 0 && passed == 30 {
					name = "logs/" + strconv.FormatInt(time.Now().Unix()+2, 10) + ".sd"
					instance = utils.CreateInstance()
					instance.Game = "PC2"
					instance.Path = name
				}
				lastState = passed
			}
			if instance.CarName != "" && instance.TrackName != "" {
				instance.WriteHeader()
			}
		})
	}()

	for {

		time.Sleep(2 * time.Second)

	}
}
