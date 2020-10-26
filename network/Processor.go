package network

import (
	"simdata/models"
	"simdata/network/gamedecoders/f12020"
	"simdata/network/gamedecoders/pc2"
	"simdata/utils"
)



func StartProcessor(game utils.GameName, port int, packetHandler func(telemetry *models.CarTelemetry),  utilsHandler  func(name string, entry interface{}))  {
	var handler = utils.GameHandler{FinishHandler: func(entry *models.CarTelemetry) {
		packetHandler(entry)

	}, UtilsHandler: func(name string, entry interface{}) {
		utilsHandler(name, entry)
	}}
	if game == utils.F12020 {
		f12020.GenerateDecoder(&handler)
	}
	if game == utils.ProjectCars2 {
		pc2.GenerateDecoder(&handler)
	}
	Listen(port, func(data [1500]byte, rlen int) {
		handler.PacketHandler(data, rlen)
	})
}