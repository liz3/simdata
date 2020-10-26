package pc2

import (
	"encoding/binary"
	"math"
	"simdata/models"
)

type Header struct {
	PacketNumber         uint32
	CategoryPacketNumber uint32
	PartialPacketIndex   uint8
	PartialPacketNumber  uint8
	PacketType           uint8
	PacketVersion        uint8
}

type VehicleInfo struct {
	Index uint16
	Class uint32
	Name string
}
type RaceInfo struct {
	TrackName string

}

func DecodeVehicleNameInfo(data [1500]uint8) VehicleInfo {
	var info = VehicleInfo{}
	info.Index = binary.LittleEndian.Uint16(data[14:16])
	info.Class = binary.LittleEndian.Uint32(data[16:20])
	var name = ""
	for i := 0; i < 64; i++ {
		var curr = 20 + i
		if data[curr] == 0 {
			break
		}
		name += string(data[curr])
	}
	info.Name = name
	return info
}

func DecodeHeader(data [1500]uint8) Header {
	var header = Header{}
	header.PacketNumber = binary.LittleEndian.Uint32(data[0:4])
	header.CategoryPacketNumber = binary.LittleEndian.Uint32(data[4:8])
	header.PartialPacketIndex = data[8]
	header.PartialPacketNumber = data[9]
	header.PacketType = data[10]
	header.PacketVersion = data[11]

	return header
}
func DecodeTiming(entry *models.CarTelemetry, data [1500]uint8) {
//	var amount = data[12]
	entry.CurrentLap = data[33 + 21]
	entry.CurrentTime = math.Float32frombits(binary.LittleEndian.Uint32(data[33 + 22:33 + 22 + 4]))
	entry.CurrentSectorTime = math.Float32frombits(binary.LittleEndian.Uint32(data[33 + 26:33 + 26 + 4]))
	entry.Sector = data[33 + 15]
}
func DecodeRaceInfo(data [1500]uint8) *RaceInfo {
	var i = &RaceInfo{}
	var n = ""
	for i := 0; i < 64; i++ {
		var curr = 176 + i
		if data[curr] == 0 {
			break
		}
		n += string(data[curr])
	}
	n += " "
	for i := 0; i < 64; i++ {
		var curr = 240 + i
		if data[curr] == 0 {
			break
		}
		n += string(data[curr])
	}
	i.TrackName = n
	return i
}
func DecodeTelemetry(entry *models.CarTelemetry, data [1500]byte) {
	entry.EngineRPM = binary.LittleEndian.Uint16(data[40:42])
	entry.Speed = math.Float32frombits(binary.LittleEndian.Uint32(data[36:40]))
	entry.Steering = float32(int8(data[44]))
	entry.Clutch = data[31]
	entry.Brake = float32(data[29])
	entry.Throttle = float32(data[30])
	entry.Gear = int8(data[45])
	entry.Pos = models.Position{
		X: math.Float32frombits(binary.LittleEndian.Uint32(data[542:546])),
		Y: math.Float32frombits(binary.LittleEndian.Uint32(data[546:550])),
		Z: math.Float32frombits(binary.LittleEndian.Uint32(data[550:554])),
	}
	entry.FuelLevel = math.Float32frombits(binary.LittleEndian.Uint32(data[32:36]))
	entry.EngineTorque = uint16(math.Float32frombits(binary.LittleEndian.Uint32(data[364:368])))
	entry.TyreTemperature = []uint16{uint16(data[176]), uint16( data[177]), uint16(data[178]), uint16( data[179])}
	entry.TyreWear = []uint16{uint16(data[196]), uint16(data[197]), uint16( data[198]),uint16( data[199])}
	entry.TyreRPS = []float32{math.Float32frombits(binary.LittleEndian.Uint32(data[160:164])),
		math.Float32frombits(binary.LittleEndian.Uint32(data[164:168])),
		math.Float32frombits(binary.LittleEndian.Uint32(data[168:172])),
		math.Float32frombits(binary.LittleEndian.Uint32(data[172:176]))}
	entry.BrakeTemperature = []uint16{binary.LittleEndian.Uint16(data[208:210]), binary.LittleEndian.Uint16(data[210:212]), binary.LittleEndian.Uint16(data[212:214]), binary.LittleEndian.Uint16(data[214:216])}

}
