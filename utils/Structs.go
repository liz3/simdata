package utils

import "simdata/models"

type GameHandler struct {
	PacketHandler func(data [1500]byte, rlen int)
	FinishHandler func(entry *models.CarTelemetry)
	UtilsHandler func(name string, entry interface{})
}