package scenarios

import (
	"os"

	"github.com/andrewdjackson/memscene/utils"
	"github.com/gocarina/gocsv"
)

// MemsData is the mems information computed from dataframes 0x80 and 0x7d
type MemsData struct {
	Time                     string  `csv:"#time"`
	EngineRPM                uint16  `csv:"80x01-02_engine-rpm"`
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
	CrankshaftPositionSensor bool    `csv:"80x19_crankshaft_position_sensor"`
	Uk801a                   int     `csv:"80x1A_uk4"`
	Uk801b                   int     `csv:"80x1B_uk5"`
	IgnitionSwitch           bool    `csv:"7dx01_ignition_switch"`
	ThrottleAngle            int     `csv:"7dx02_throttle_angle"`
	Uk7d03                   int     `csv:"7dx03_uk6"`
	AirFuelRatio             float32 `csv:"7dx04_air_fuel_ratio"`
	DTC2                     int     `csv:"7dx05_dtc2"`
	LambdaVoltage            int     `csv:"7dx06_lambda_voltage"`
	LambdaFrequency          int     `csv:"7dx07_lambda_sensor_frequency"`
	LambdaDutycycle          int     `csv:"7dx08_lambda_sensor_dutycycle"`
	LambdaStatus             int     `csv:"7dx09_lambda_sensor_status"`
	ClosedLoop               bool    `csv:"7dx0A_closed_loop"`
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
	Dataframe7d              string  `csv:"0x7d_raw"`
	Dataframe80              string  `csv:"0x80_raw"`
}

// RawData represents the raw data from the log file
type RawData struct {
	Dataframe7d string `csv:"0x7d_raw"`
	Dataframe80 string `csv:"0x80_raw"`
}

// Scenario represents the scenario data
type Scenario struct {
	file *os.File
	// Memsdata log
	Memsdata []*MemsData
	// Rawdata from log
	Rawdata []*RawData
	// Position in the log
	Position int
	// Count of items in the log
	Count int
}

// NewScenario creates a new scenario
func NewScenario() *Scenario {
	scenario := &Scenario{}
	// initialise the log
	scenario.Memsdata = []*MemsData{}
	scenario.Rawdata = []*RawData{}
	// start at the beginning
	scenario.Position = 0
	// no items in the log
	scenario.Count = 0

	return scenario
}

// Open the CSV scenario file
func (scenario *Scenario) openFile(filepath string) {
	var err error

	scenario.file, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		utils.LogE.Printf("unable to open %s", err)
	}
}

// Load the scenario
func (scenario *Scenario) Load(filepath string) {
	scenario.openFile(filepath)

	if err := gocsv.Unmarshal(scenario.file, &scenario.Rawdata); err != nil {
		utils.LogE.Printf("unable to parse file %s", err)
	} else {
		scenario.Count = len(scenario.Rawdata)
		utils.LogI.Printf("loaded scenario %s (%d dataframes)", filepath, scenario.Count)
	}
}

// Next provides the next item in the log
func (scenario *Scenario) Next() *RawData {
	item := scenario.Rawdata[scenario.Position]
	scenario.Position = scenario.Position + 1

	// if we pass the end, loop back to the start
	if scenario.Position >= scenario.Count {
		utils.LogW.Printf("reached end of scenario, restarting from beginning")
		scenario.Position = 0
	}

	return item
}

// SaveCSVFile saves the Memdata to a CSV file
func (scenario *Scenario) SaveCSVFile(filepath string) {
	file, _ := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)

	err := gocsv.MarshalFile(&scenario.Memsdata, file)
	if err != nil {
		utils.LogI.Printf("error saving csv file %s", err)
	}
}
