package models

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type CarTelemetry struct {
	PacketNumber uint32
	Speed float32 `json:"speed"`
	Gear int8 `json:"gear"`
	EngineRPM uint16 `json:"engine_rpm"`
	EngineTorque uint16 `json:"engine_torque"`
	
	Clutch uint8 `json:"clutch"`
	Brake float32 `json:"brake"`
	Throttle float32 `json:"throttle"`
	
	Steering float32 `json:"steering"`

	BrakeTemperature []uint16 `json:"brake_temperature"`
	TyreTemperature []uint16 `json:"tyre_temperature"`
	TyreWear []uint16 `json:"tyre_wear"`
	TyreRPS []float32 `json:"tyre_rps"`

	FuelLevel float32 `json:"fuel_level"`
	
	Pos Position `json:"pos"`
	
	CurrentLap uint8 `json:"current_lap"`
	CurrentTime float32 `json:"current_time"`
	CurrentSectorTime float32 `json:"current_sector_time"`
	Sector uint8 `json:"sector"`
}