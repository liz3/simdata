package f12020

import (
	"simdata/models"
	"simdata/utils"
)

func GenerateDecoder(handler *utils.GameHandler)  {
	var p = &models.CarTelemetry{}
	var last float32 = 0
	handler.PacketHandler = func(data [1500]uint8, rlen int) {
		var header = DecodeHeader(data)
		handler.UtilsHandler("GAME_STATE", header.SessionUid)
		if header.PacketId == 0 {
			if last != 0 && header.SessionTime != last {
				handler.FinishHandler(p)
				p = &models.CarTelemetry{}
				p.PacketNumber = header.FrameIdentifier
			}
			DecodeMotion(p, data)
			last = header.SessionTime
		}
		if header.PacketId == 6 {
			if last != 0 && header.SessionTime != last {
				handler.FinishHandler(p)
				p = &models.CarTelemetry{}
				p.PacketNumber = header.FrameIdentifier

			}
			DecodeTelemetry(p, data)
			last = header.SessionTime
		}
		if header.PacketId == 2 {
			if last != 0 && header.SessionTime != last {
				handler.FinishHandler(p)
				p = &models.CarTelemetry{}
				p.PacketNumber = header.FrameIdentifier

			}
			DecodeTiming(p, data)
			last = header.SessionTime
		}
		if header.PacketId == 7 {
			if last != 0 && header.SessionTime != last {
				handler.FinishHandler(p)
				p = &models.CarTelemetry{}
				p.PacketNumber = header.FrameIdentifier

			}
			DecodeStatus(p, data)
			last = header.SessionTime
		}
		if header.PacketId == 1 {
			var info = DecodeRaceInfo(data)
			handler.UtilsHandler("TRACK_NAME", info.TrackName)
			handler.UtilsHandler("CAR_NAME", info.Class)

		}
	}
}
