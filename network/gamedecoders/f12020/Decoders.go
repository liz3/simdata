package f12020

import (
	"encoding/binary"
	"math"
	"simdata/models"
	"strconv"
)

type PacketHeader struct {
	PacketFormat uint16
	MajorVersion uint8
	MinorVersion uint8
	PacketVersion uint8
	PacketId uint8
	SessionUid uint64
	SessionTime float32
	FrameIdentifier uint32
	PlayerCarIndex uint8
	SecondaryPlayerCarIndex uint8

}
type RaceInfo struct {
	TrackName string
	Class string
}
func DecodeRaceInfo(data [1500]uint8) *RaceInfo {
	var i = &RaceInfo{}
	i.TrackName = strconv.Itoa(int(int8(data[31])))
	var formula = data[32]
	if formula == 0 {
		i.Class = "Modern F1"
	}
	if formula == 1 {
		i.Class = "Classic F1"
	}
	if formula == 2 {
		i.Class = "F2"
	}
	return i
}
func DecodeTelemetry(entry *models.CarTelemetry, data [1500]uint8)  {
	entry.Speed = float32(binary.LittleEndian.Uint16(data[24:26]))
	entry.Throttle = math.Float32frombits(binary.LittleEndian.Uint32(data[26:30]))
	entry.Steering = math.Float32frombits(binary.LittleEndian.Uint32(data[30:34]))
	entry.Brake = math.Float32frombits(binary.LittleEndian.Uint32(data[34:38]))
	entry.Clutch = data[38]
	entry.Gear = int8(data[39])
	entry.EngineRPM = binary.LittleEndian.Uint16(data[40:42])
	entry.EngineTorque = 0

	//brake temp
	entry.BrakeTemperature = []uint16{binary.LittleEndian.Uint16(data[43:45]),binary.LittleEndian.Uint16(data[45:47]),binary.LittleEndian.Uint16(data[47:49]),binary.LittleEndian.Uint16(data[49:51])}
	entry.TyreTemperature = []uint16{uint16(data[52]), uint16( data[53]),uint16( data[54]), uint16(data[55])}

}
func DecodeTiming(entry *models.CarTelemetry, data [1500]uint8)  {
	entry.CurrentTime = math.Float32frombits(binary.LittleEndian.Uint32(data[28:32]))
	entry.CurrentLap = data[24 + 45]
	entry.Sector = data[24 + 47]
}
func DecodeStatus(entry *models.CarTelemetry, data [1500]uint8)  {
	entry.TyreWear = []uint16{uint16(data[49]), uint16(data[50]), uint16( data[51]), uint16(data[52])}
	entry.FuelLevel =  math.Float32frombits(binary.LittleEndian.Uint32(data[29:33]))
}
func DecodeMotion(entry *models.CarTelemetry, data [1500]uint8)  {
	var pos = models.Position{
		X: math.Float32frombits(binary.LittleEndian.Uint32(data[24:28])),
		Y:  math.Float32frombits(binary.LittleEndian.Uint32(data[28:32])),
		Z:  math.Float32frombits(binary.LittleEndian.Uint32(data[32:36])),
	}
	entry.Pos = pos
}

func DecodeHeader(data [1500]uint8) PacketHeader {
	var header = PacketHeader{}
	header.PacketFormat = binary.LittleEndian.Uint16(data[0:2])
	header.MajorVersion =data[2]
	header.MinorVersion =data[3]
	header.PacketVersion =data[4]
	header.PacketId = data[5]
	header.SessionUid = binary.LittleEndian.Uint64(data[6:14])
	sessionTimeBytes := binary.LittleEndian.Uint32(data[14:18])
	header.SessionTime = math.Float32frombits(sessionTimeBytes)
	header.FrameIdentifier = binary.LittleEndian.Uint32(data[18:22])

	return header
}