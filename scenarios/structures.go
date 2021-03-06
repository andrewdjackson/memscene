package scenarios

// MemsFCRData is the mems information computed from dataframes 0x80 and 0x7d
type MemsFCRData struct {
	Time                     string  `csv:"#time"`
	EngineRPM                int     `csv:"80x01-02_engine-rpm"`
	CoolantTemp              int     `csv:"80x03_coolant_temp"`
	AmbientTemp              int     `csv:"80x04_ambient_temp"`
	IntakeAirTemp            int     `csv:"80x05_intake_air_temp"`
	FuelTemp                 int     `csv:"80x06_fuel_temp"`
	ManifoldAbsolutePressure float32 `csv:"80x07_map_kpa"`
	BatteryVoltage           float32 `csv:"80x08_battery_voltage"`
	ThrottlePotSensor        float32 `csv:"80x09_throttle_pot"`
	IdleSwitch               bool    `csv:"80x0A_idle_switch"`
	AirconSwitch             bool    `csv:"80x0B_uk1"`
	ParkNeutralSwitch        bool    `csv:"80x0C_park_neutral_switch"`
	DTC0                     uint8   `csv:"80x0D-0E_fault_codes"`
	DTC1                     uint8   `csv:"-"`
	IdleSetPoint             int     `csv:"80x0F_idle_set_point"`
	IdleHot                  int     `csv:"80x10_idle_hot"`
	Uk8011                   int     `csv:"80x11_uk2"`
	IACPosition              int     `csv:"80x12_iac_position"`
	IdleSpeedDeviation       int     `csv:"80x13-14_idle_error"`
	IgnitionAdvanceOffset80  int     `csv:"80x15_ignition_advance_offset"`
	IgnitionAdvance          float32 `csv:"80x16_ignition_advance"`
	CoilTime                 float32 `csv:"80x17-18_coil_time"`
	CrankshaftPositionSensor bool    `csv:"80x19_crankshaft_position_sensor"`
	Uk801a                   int     `csv:"80x1A_uk4"`
	Uk801b                   int     `csv:"80x1B_uk5"`
	IgnitionSwitch           bool    `csv:"7dx01_ignition_switch"`
	ThrottleAngle            int     `csv:"7dx02_throttle_angle"`
	Uk7d03                   int     `csv:"7dx03_uk6"`
	AirFuelRatio             float32 `csv:"7dx04_air_fuel_ratio"`
	DTC2                     uint8   `csv:"7dx05_dtc2"`
	LambdaVoltage            int     `csv:"7dx06_lambda_voltage"`
	LambdaFrequency          int     `csv:"7dx07_lambda_sensor_frequency"`
	LambdaDutycycle          int     `csv:"7dx08_lambda_sensor_dutycycle"`
	LambdaStatus             int     `csv:"7dx09_lambda_sensor_status"`
	ClosedLoop               bool    `csv:"7dx0A_closed_loop"`
	LongTermFuelTrim         int     `csv:"7dx0B_long_term_fuel_trim"`
	ShortTermFuelTrim        int     `csv:"7dx0C_short_term_fuel_trim"`
	CarbonCanisterPurgeValve int     `csv:"7dx0D_carbon_canister_dutycycle"`
	DTC3                     uint8   `csv:"7dx0E_dtc3"`
	IdleBasePosition         int     `csv:"7dx0F_idle_base_pos"`
	Uk7d10                   int     `csv:"7dx10_uk7"`
	DTC4                     uint8   `csv:"7dx11_dtc4"`
	IgnitionAdvanceOffset7d  int     `csv:"7dx12_ignition_advance2"`
	IdleSpeedOffset          int     `csv:"7dx13_idle_speed_offset"`
	Uk7d14                   int     `csv:"7dx14_idle_error2"`
	Uk7d15                   int     `csv:"7dx14-15_uk10"`
	DTC5                     uint8   `csv:"7dx16_dtc5"`
	Uk7d17                   int     `csv:"7dx17_uk11"`
	Uk7d18                   int     `csv:"7dx18_uk12"`
	Uk7d19                   int     `csv:"7dx19_uk13"`
	Uk7d1a                   int     `csv:"7dx1A_uk14"`
	Uk7d1b                   int     `csv:"7dx1B_uk15"`
	Uk7d1c                   int     `csv:"7dx1C_uk16"`
	Uk7d1d                   int     `csv:"7dx1D_uk17"`
	Uk7d1e                   int     `csv:"7dx1E_uk18"`
	JackCount                int     `csv:"7dx1F_uk19"`
	Dataframe7d              string  `csv:"0x7d_raw"`
	Dataframe80              string  `csv:"0x80_raw"`
}

// MemsFCRRawData structure used for reprocessing raw data
type MemsFCRRawData struct {
	Time                     string  `csv:"#time"`
	EngineRPM                int     `csv:"-"`
	CoolantTemp              int     `csv:"-"`
	AmbientTemp              int     `csv:"-"`
	IntakeAirTemp            int     `csv:"-"`
	FuelTemp                 int     `csv:"-"`
	ManifoldAbsolutePressure float32 `csv:"-"`
	BatteryVoltage           float32 `csv:"-"`
	ThrottlePotSensor        float32 `csv:"-"`
	IdleSwitch               bool    `csv:"-"`
	AirconSwitch             bool    `csv:"-"`
	ParkNeutralSwitch        bool    `csv:"-"`
	DTC0                     uint8   `csv:"-"`
	DTC1                     uint8   `csv:"-"`
	IdleSetPoint             int     `csv:"-"`
	IdleHot                  int     `csv:"-"`
	Uk8011                   int     `csv:"-"`
	IACPosition              int     `csv:"-"`
	IdleSpeedDeviation       int     `csv:"-"`
	IgnitionAdvanceOffset80  int     `csv:"-"`
	IgnitionAdvance          float32 `csv:"-"`
	CoilTime                 float32 `csv:"-"`
	CrankshaftPositionSensor bool    `csv:"-"`
	Uk801a                   int     `csv:"-"`
	Uk801b                   int     `csv:"-"`
	IgnitionSwitch           bool    `csv:"-"`
	ThrottleAngle            int     `csv:"-"`
	Uk7d03                   int     `csv:"-"`
	AirFuelRatio             float32 `csv:"-"`
	DTC2                     uint8   `csv:"-"`
	LambdaVoltage            int     `csv:"-"`
	LambdaFrequency          int     `csv:"-"`
	LambdaDutycycle          int     `csv:"-"`
	LambdaStatus             int     `csv:"-"`
	ClosedLoop               bool    `csv:"-"`
	LongTermFuelTrim         int     `csv:"-"`
	ShortTermFuelTrim        int     `csv:"-"`
	CarbonCanisterPurgeValve int     `csv:"-"`
	DTC3                     uint8   `csv:"-"`
	IdleBasePosition         int     `csv:"-"`
	Uk7d10                   int     `csv:"-"`
	DTC4                     uint8   `csv:"-"`
	IgnitionAdvanceOffset7d  int     `csv:"-"`
	IdleSpeedOffset          int     `csv:"-"`
	Uk7d14                   int     `csv:"-"`
	Uk7d15                   int     `csv:"-"`
	DTC5                     uint8   `csv:"-"`
	Uk7d17                   int     `csv:"-"`
	Uk7d18                   int     `csv:"-"`
	Uk7d19                   int     `csv:"-"`
	Uk7d1a                   int     `csv:"-"`
	Uk7d1b                   int     `csv:"-"`
	Uk7d1c                   int     `csv:"-"`
	Uk7d1d                   int     `csv:"-"`
	Uk7d1e                   int     `csv:"-"`
	JackCount                int     `csv:"-"`
	Dataframe7d              string  `csv:"0x7d_raw"`
	Dataframe80              string  `csv:"0x80_raw"`
}

// MemsRoscoData mems-rosco logs are in CSV format with a .TXT extension in the format:
//
// Ecu Id:
// #time,engine-rpm,coolant_temp,ambient_temp,intake_air_temp,fuel_temp,map_kpa,battery_voltage,throttle_pot_voltage,idle_switch,uk1,park_neutral_switch,fault_codes
//    idle_set_point,idle_hot,uk2,iac_position,idle_error,ignition_advance_offset,ignition_advance,coil_time,crancs,uk4,uk5,ignition_switch,
//    throttle_angle,uk6,air_fuel_ratio,fault_code0,lambda_voltage_mv,lambda_sensor_frequency,lambda_sensor_dutycycle,lambda_sensor_status,closed_loop,
//    long_term_fuel_trim,short_term_fuel_trim,carbon_canister_dutycycle,fault_code1,idle_base_pos,uk7,uk8,ignition_advance2,uk9,idle_error2,uk10,
//    fault_code4,uk11,uk12,uk13,uk14,uk15,uk16,uk17,uk18,uk19
//
type MemsRoscoData struct {
	Time                     string  `csv:"#time"`
	EngineRPM                uint16  `csv:"engine-rpm"`
	CoolantTemp              int     `csv:"coolant_temp"`
	AmbientTemp              int     `csv:"ambient_temp"`
	IntakeAirTemp            int     `csv:"intake_air_temp"`
	FuelTemp                 int     `csv:"fuel_temp"`
	ManifoldAbsolutePressure float32 `csv:"map_kpa"`
	BatteryVoltage           float32 `csv:"battery_voltage"`
	ThrottlePotSensor        float32 `csv:"throttle_pot_voltage"`
	IdleSwitch               int     `csv:"idle_switch"`
	AirconSwitch             int     `csv:"uk1"`
	ParkNeutralSwitch        int     `csv:"park_neutral_switch"`
	DTC0                     int     `csv:"fault_codes"`
	DTC1                     int     `csv:"-"`
	IdleSetPoint             int     `csv:"idle_set_point"`
	IdleHot                  int     `csv:"idle_hot"`
	Uk8011                   int     `csv:"uk2"`
	IACPosition              int     `csv:"iac_position"`
	IdleSpeedDeviation       uint16  `csv:"idle_error"`
	IgnitionAdvanceOffset80  int     `csv:"ignition_advance_offset"`
	IgnitionAdvance          float32 `csv:"ignition_advance"`
	CoilTime                 float32 `csv:"coil_time"`
	CrankshaftPositionSensor int     `csv:"crancs"`
	Uk801a                   int     `csv:"uk4"`
	Uk801b                   int     `csv:"uk5"`
	IgnitionSwitch           int     `csv:"ignition_switch"`
	ThrottleAngle            int     `csv:"throttle_angle"`
	Uk7d03                   int     `csv:"uk6"`
	AirFuelRatio             float32 `csv:"air_fuel_ratio"`
	DTC2                     int     `csv:"fault_code0"`
	LambdaVoltage            int     `csv:"lambda_voltage_mv"`
	LambdaFrequency          int     `csv:"lambda_sensor_frequency"`
	LambdaDutycycle          int     `csv:"lambda_sensor_dutycycle"`
	LambdaStatus             int     `csv:"lambda_sensor_status"`
	ClosedLoop               int     `csv:"closed_loop"`
	LongTermFuelTrim         int     `csv:"long_term_fuel_trim"`
	ShortTermFuelTrim        int     `csv:"short_term_fuel_trim"`
	CarbonCanisterPurgeValve int     `csv:"carbon_canister_dutycycle"`
	DTC3                     int     `csv:"fault_code1"`
	IdleBasePosition         int     `csv:"idle_base_pos"`
	Uk7d10                   int     `csv:"uk7"`
	DTC4                     int     `csv:"uk8"`
	IgnitionAdvanceOffset7d  int     `csv:"ignition_advance2"`
	IdleSpeedOffset          int     `csv:"uk9"`
	Uk7d14                   int     `csv:"idle_error2"`
	Uk7d15                   int     `csv:"uk10"`
	DTC5                     int     `csv:"fault_code4"`
	Uk7d17                   int     `csv:"uk11"`
	Uk7d18                   int     `csv:"uk12"`
	Uk7d19                   int     `csv:"uk13"`
	Uk7d1a                   int     `csv:"uk14"`
	Uk7d1b                   int     `csv:"uk15"`
	Uk7d1c                   int     `csv:"uk16"`
	Uk7d1d                   int     `csv:"uk17"`
	Uk7d1e                   int     `csv:"uk18"`
	JackCount                int     `csv:"uk19"`
	Dataframe7d              string  `csv:"-"`
	Dataframe80              string  `csv:"-"`
}

// MemsRoscoV2Data mems-rosco version 2 are in CSV format with a .TXT extension in the format:
//
// #time,80x01-02_engine-rpm,80x03_coolant_temp,80x04_ambient_temp,80x05_intake_air_temp,80x06_fuel_temp,80x07_map_kpa,80x08_battery_voltage,
// 80x09_throttle_pot,80x0A_idle_switch,80x0B_uk1,80x0C_park_neutral_switch,80x0D-0E_fault_codes,80x0F_idle_set_point,80x10_idle_hot,
// 80x11_uk2,80x12_iac_position,80x13-14_idle_error,80x15_ignition_advance_offset,80x16_ignition_advance,80x17-18_coil_time,
// 80x19_crankshaft_position_sensor,80x1A_uk4,80x1B_uk5,7dx01_ignition_switch,7dx02_throttle_angle,7dx03_uk6,7dx04_air_fuel_ratio,7dx05_dtc2,
// 7dx06_lambda_voltage,7dx07_lambda_sensor_frequency,7dx08_lambda_sensor_dutycycle,7dx09_lambda_sensor_status,7dx0A_closed_loop,
// 7dx0B_long_term_fuel_trim,7dx0C_short_term_fuel_trim,7dx0D_carbon_canister_dutycycle,7dx0E_dtc3,7dx0F_idle_base_pos,7dx10_uk7,
// 7dx11_dtc4,7dx12_ignition_advance2,7dx13_idle_speed_offset,7dx14_idle_error2,7dx14-15_uk10,7dx16_dtc5,7dx17_uk11,7dx18_uk12,7dx19_uk13,
// 7dx1A_uk14,7dx1B_uk15,7dx1C_uk16,7dx1D_uk17,7dx1E_uk18,7dx1F_uk19
//
type MemsRoscoV2Data struct {
	Time                     string  `csv:"#time"`
	EngineRPM                uint16  `csv:"80x01-02_engine-rpm"`
	CoolantTemp              int     `csv:"80x03_coolant_temp"`
	AmbientTemp              int     `csv:"80x04_ambient_temp"`
	IntakeAirTemp            int     `csv:"80x05_intake_air_temp"`
	FuelTemp                 int     `csv:"80x06_fuel_temp"`
	ManifoldAbsolutePressure float32 `csv:"80x07_map_kpa"`
	BatteryVoltage           float32 `csv:"80x08_battery_voltage"`
	ThrottlePotSensor        float32 `csv:"80x09_throttle_pot"`
	IdleSwitch               int     `csv:"80x0A_idle_switch"`
	AirconSwitch             int     `csv:"80x0B_uk1"`
	ParkNeutralSwitch        int     `csv:"80x0C_park_neutral_switch"`
	DTC0                     int     `csv:"80x0D-0E_fault_codes"`
	DTC1                     int     `csv:"-"`
	IdleSetPoint             int     `csv:"80x0F_idle_set_point"`
	IdleHot                  int     `csv:"80x10_idle_hot"`
	Uk8011                   int     `csv:"80x11_uk2"`
	IACPosition              int     `csv:"80x12_iac_position"`
	IdleSpeedDeviation       uint16  `csv:"80x13-14_idle_error"`
	IgnitionAdvanceOffset80  int     `csv:"80x15_ignition_advance_offset"`
	IgnitionAdvance          float32 `csv:"80x16_ignition_advance"`
	CoilTime                 float32 `csv:"80x17-18_coil_time"`
	CrankshaftPositionSensor int     `csv:"80x19_crankshaft_position_sensor"`
	Uk801a                   int     `csv:"80x1A_uk4"`
	Uk801b                   int     `csv:"80x1B_uk5"`
	IgnitionSwitch           int     `csv:"7dx01_ignition_switch"`
	ThrottleAngle            int     `csv:"7dx02_throttle_angle"`
	Uk7d03                   int     `csv:"7dx03_uk6"`
	AirFuelRatio             float32 `csv:"7dx04_air_fuel_ratio"`
	DTC2                     int     `csv:"7dx05_dtc2"`
	LambdaVoltage            int     `csv:"7dx06_lambda_voltage"`
	LambdaFrequency          int     `csv:"7dx07_lambda_sensor_frequency"`
	LambdaDutycycle          int     `csv:"7dx08_lambda_sensor_dutycycle"`
	LambdaStatus             int     `csv:"7dx09_lambda_sensor_status"`
	ClosedLoop               int     `csv:"7dx0A_closed_loop"`
	LongTermFuelTrim         int     `csv:"7dx0B_long_term_fuel_trim"`
	ShortTermFuelTrim        int     `csv:"7dx0C_short_term_fuel_trim"`
	CarbonCanisterPurgeValve int     `csv:"7dx0D_carbon_canister_dutycycle"`
	DTC3                     int     `csv:"7dx0E_dtc3"`
	IdleBasePosition         int     `csv:"7dx0F_idle_base_pos"`
	Uk7d10                   int     `csv:"7dx10_uk7"`
	DTC4                     int     `csv:"7dx11_dtc4"`
	IgnitionAdvanceOffset7d  int     `csv:"7dx12_ignition_advance2"`
	IdleSpeedOffset          int     `csv:"7dx13_idle_speed_offset"`
	Uk7d14                   int     `csv:"7dx14_idle_error2"`
	Uk7d15                   int     `csv:"7dx14-15_uk10"`
	DTC5                     int     `csv:"7dx16_dtc5"`
	Uk7d17                   int     `csv:"7dx17_uk11"`
	Uk7d18                   int     `csv:"7dx18_uk12"`
	Uk7d19                   int     `csv:"7dx19_uk13"`
	Uk7d1a                   int     `csv:"7dx1A_uk14"`
	Uk7d1b                   int     `csv:"7dx1B_uk15"`
	Uk7d1c                   int     `csv:"7dx1C_uk16"`
	Uk7d1d                   int     `csv:"7dx1D_uk17"`
	Uk7d1e                   int     `csv:"7dx1E_uk18"`
	JackCount                int     `csv:"7dx1F_uk19"`
	Dataframe7d              string  `csv:"-"`
	Dataframe80              string  `csv:"-"`
}

type (
	// DataFrame7d data sequence returned by the ECU in reply to the command 0x7D.
	// This structure represents the raw data from the ECU
	//
	DataFrame7d struct {
		Command                  uint8
		BytesinFrame             uint8 // 7dx00
		IgnitionSwitch           uint8 // 7dx01
		ThrottleAngle            uint8 // 7dx03
		Uk7d03                   uint8 // 7dx03
		AirFuelRatio             uint8 // 7dx04
		Dtc2                     uint8 // 7dx05
		LambdaVoltage            uint8 // 7dx06
		LambdaFrequency          uint8 // 7dx07
		LambdaDutyCycle          uint8 // 7dx08
		LambdaStatus             uint8 // 7dx09
		LoopIndicator            uint8 // 7dx0A
		LongTermFuelTrim         uint8 // 7dx0B
		ShortTermFuelTrim        uint8 // 7dx0C
		CarbonCanisterPurgeValve uint8 // 7dx0D
		Dtc3                     uint8 // 7dx0E
		IdleBasePos              uint8 // 7dx0F
		Uk7d10                   uint8 // 7dx10
		Dtc4                     uint8 // 7dx11
		IgnitionAdvanceOffset7d  uint8 // 7dx12
		IdleSpeedOffset          uint8 // 7dx13
		Uk7d14                   uint8 // 7dx14
		Uk7d15                   uint8 // 7dx15
		Dtc5                     uint8 // 7dx16
		Uk7d17                   uint8 // 7dx17
		Uk7d18                   uint8 // 7dx18
		Uk7d19                   uint8 // 7dx19
		Uk7d1a                   uint8 // 7dx1a
		Uk7d1b                   uint8 // 7dx1b
		Uk7d1c                   uint8 // 7dx1c
		Uk7d1d                   uint8 // 7dx1d
		Uk7d1e                   uint8 // 7dx1e
		JackCount                uint8 // 7dx1f
	}
)

type (
	// DataFrame80 data sequence returned by the ECU in reply to the command 0x80.
	// This structure represents the raw data from the ECU
	//
	DataFrame80 struct {
		Command                  uint8
		BytesinFrame             uint8  // 80x00
		EngineRpm                uint16 // 80x01 - 80x02
		CoolantTemp              uint8  // 80x03
		AmbientTemp              uint8  // 80x04
		IntakeAirTemp            uint8  // 80x05
		FuelTemp                 uint8  // 80x06
		ManifoldAbsolutePressure uint8  // 80x07
		BatteryVoltage           uint8  // 80x08
		ThrottlePotSensor        uint8  // 80x09
		IdleSwitch               uint8  // 80x0A
		AirconSwitch             uint8  // 80x0B
		ParkNeutralSwitch        uint8  // 80x0C
		Dtc0                     uint8  // 80x0D  Bit 0: Coolant temp sensor fault (Code 1) Bit 1: Inlet air temp sensor fault (Code 2)
		Dtc1                     uint8  // 80x0E  Bit 1: Fuel pump circuit fault (Code 10)  Bit 7: Throttle pot circuit fault (Code 16)
		IdleSetPoint             uint8  // 80x0F
		IdleHot                  uint8  // 80x10
		Uk8011                   uint8  // 80x11
		IacPosition              uint8  // 80x12
		IdleSpeedDeviation       uint16 // 80x13 - 80x14
		IgnitionAdvanceOffset80  uint8  // 80x15
		IgnitionAdvance          uint8  // 80x16
		CoilTime                 uint16 // 80x17 - 80x18
		CrankshaftPositionSensor uint8  // 80x19
		Uk801a                   uint8  // 80x1A
		Uk801b                   uint8  // 80x1B
	}
)

const (
	// AirSensorFaultCode 0x80 DTC0 Fault
	AirSensorFaultCode = byte(0b00000001)
	// CoolantSensorFaultCode 0x80 DTC0 Fault
	CoolantSensorFaultCode = byte(0b00000010)
	// FuelPumpFaultCode 0x80 DTC1 Fault
	FuelPumpFaultCode = byte(0b00000001)
	// ThrottlePotFaultCode 0x80 DTC1 Fault
	ThrottlePotFaultCode = byte(0b01000000)
	// IdleSwitchActive flag
	IdleSwitchActive = byte(0b00001000)
)
