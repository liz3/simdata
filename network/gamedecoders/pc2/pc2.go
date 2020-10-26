package pc2

import (
	"simdata/models"
	"simdata/utils"
)

func GenerateDecoder(handler *utils.GameHandler)  {
	var p = &models.CarTelemetry{}
	var counter = 0
	handler.PacketHandler = func(data [1500]uint8, rlen int) {

		var header = DecodeHeader(data)
		if header.PacketType == 0 {
			p.PacketNumber = header.PacketNumber
			DecodeTelemetry(p, data)
			counter++
			if counter == 2 {
				handler.FinishHandler(p)
				p = &models.CarTelemetry{}
				counter = 0
			}
		}
		if header.PacketType == 3 {
			DecodeTiming(p, data)
			counter++
			if counter == 2 {
				handler.FinishHandler(p)
				p = &models.CarTelemetry{}
				counter = 0
			}
		}
		if header.PacketType == 4 {
			handler.UtilsHandler("GAME_STATE", data[15])
		}
		if header.PacketType == 1 {
			var info = DecodeRaceInfo(data)
			handler.UtilsHandler("TRACK_NAME", info.TrackName)

		}
		if header.PacketType == 8 && header.PartialPacketIndex == 1 {
			var info = DecodeVehicleNameInfo(data)
			handler.UtilsHandler("CAR_NAME", info.Name)

		}
	}
}
