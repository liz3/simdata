package main

import (
	"fmt"
	"simdata/models"
	"simdata/network"
	"simdata/utils"
	"strconv"
	"time"
)

func main() {

	go func() {
		fmt.Println("F12020 starting collection")
		var name = "logs/" + strconv.FormatInt(time.Now().Unix(), 10) + ".sd"
		var instance = utils.CreateInstance()
		instance.Game = "F12020"
		instance.Path = name
		var sessionId uint64 = 0
		var lastLap uint8 = 0
		var lastTime float32 = 0
		network.StartProcessor(utils.F12020, 20777, func(telemetry *models.CarTelemetry) {
			//if( telemetry.CurrentLap < lastLap && telemetry.CurrentLap != 0) || (telemetry.CurrentTime < lastTime && telemetry.CurrentTime != 0 && telemetry.CurrentLap != lastLap + 1) {
			//	instance.Flush()
			//	lastLap = 0
			//	lastTime = 0
			//	name = "logs/" + strconv.FormatInt(time.Now().Unix(), 10) + ".sd"
			//	instance = utils.CreateInstance()
			//	instance.Game = "F12020"
			//	instance.Path = name
			//	fmt.Println("F12020 - Reset to new log")
			//}

			instance.Push(telemetry)
			//lastLap = telemetry.CurrentLap
			//lastTime = telemetry.CurrentTime
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
					instance.Flush()
					lastLap = 0
					lastTime = 0
					name = "logs/" + strconv.FormatInt(time.Now().Unix(), 10) + ".sd"
					instance = utils.CreateInstance()
					instance.Game = "F12020"
					instance.Path = name
					fmt.Println("F12020 - Reset to new log")
				}
				sessionId = passed

			}
			if instance.CarName != "" && instance.TrackName != "" {
				instance.WriteHeader()
			}
		})
	}()
	go func() {
		fmt.Println("PC2 starting collection")
		var name = "logs/" + strconv.FormatInt(time.Now().Unix()+2, 10) + ".sd"
		var instance = utils.CreateInstance()
		instance.Game = "PC2"
		instance.Path = name
		var lastState uint8 = 0
		var lastLap uint8 = 0
		var lastTime float32 = 0

		network.StartProcessor(utils.ProjectCars2, 5606, func(telemetry *models.CarTelemetry) {
			//if( telemetry.CurrentLap < lastLap && telemetry.CurrentLap != 0) || (telemetry.CurrentTime < lastTime && telemetry.CurrentTime == -1 && telemetry.CurrentLap != lastLap + 1) {
			//	fmt.Println(telemetry.CurrentTime, telemetry.CurrentLap, lastTime, lastLap)
			//	instance.Flush()
			//	lastLap = 0
			//	lastTime = 0
			//	name = "logs/" + strconv.FormatInt(time.Now().Unix(), 10) + ".sd"
			//	instance = utils.CreateInstance()
			//	instance.Game = "PC2"
			//	instance.Path = name
			//	fmt.Println("PC2 - Reset to new log")
			//}

			instance.Push(telemetry)
			//lastLap = telemetry.CurrentLap
			//lastTime = telemetry.CurrentTime
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
					lastLap = 0
					lastTime = 0
					instance.Game = "PC2"
					instance.Path = name
					fmt.Println("PC2 - Reset to new log")
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
