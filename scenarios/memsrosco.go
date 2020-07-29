package scenarios

import (
	"encoding/csv"
	"fmt"
	"github.com/andrewdjackson/memscene/utils"
	"github.com/gocarina/gocsv"
	"io"
	"math"
	"os"
	"strings"
)

// Mems-Rosco logs are in CSV format with a .TXT extension
// in the format:
//
// Ecu Id:
// #time,engine-rpm,coolant_temp,ambient_temp,intake_air_temp,fuel_temp,map_kpa,battery_voltage,throttle_pot_voltage,idle_switch,uk1,park_neutral_switch,fault_codes
//    idle_set_point,idle_hot,uk2,iac_position,idle_error,ignition_advance_offset,ignition_advance,coil_time,crancs,uk4,uk5,ignition_switch,
//    throttle_angle,uk6,air_fuel_ratio,fault_code0,lambda_voltage_mv,lambda_sensor_frequency,lambda_sensor_dutycycle,lambda_sensor_status,closed_loop,
//    long_term_fuel_trim,short_term_fuel_trim,carbon_canister_dutycycle,fault_code1,idle_base_pos,uk7,uk8,ignition_advance2,uk9,idle_error2,uk10,
//    fault_code4,uk11,uk12,uk13,uk14,uk15,uk16,uk17,uk18,uk19
//

// MemsRoscoData is the mems information computed from dataframes 0x80 and 0x7d
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
	IdleSwitch               bool    `csv:"idle_switch"`
	AirconSwitch             bool    `csv:"uk1"`
	ParkNeutralSwitch        bool    `csv:"park_neutral_switch"`
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
	CrankshaftPositionSensor bool    `csv:"crancs"`
	Uk801a                   int     `csv:"uk4"`
	Uk801b                   int     `csv:"uk5"`
	IgnitionSwitch           bool    `csv:"ignition_switch"`
	ThrottleAngle            int     `csv:"throttle_angle"`
	Uk7d03                   int     `csv:"uk6"`
	AirFuelRatio             float32 `csv:"air_fuel_ratio"`
	DTC2                     int     `csv:"fault_code0"`
	LambdaVoltage            int     `csv:"lambda_voltage_mv"`
	LambdaFrequency          int     `csv:"lambda_sensor_frequency"`
	LambdaDutycycle          int     `csv:"lambda_sensor_dutycycle"`
	LambdaStatus             int     `csv:"lambda_sensor_status"`
	ClosedLoop               bool    `csv:"closed_loop"`
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

// MemsRosco structure
type MemsRosco struct {
	scenario  *Scenario
	file      *os.File
	memsData  []*MemsData
	roscoData []*MemsRoscoDataStructure
}

// NewMemsRosco create a new MemsRosco instance
func NewMemsRosco() *MemsRosco {
	memsrosco := &MemsRosco{}
	memsrosco.scenario = NewScenario()

	return memsrosco
}

// Convert takes Readmems Log files and converts them into MemsFCR format
func (memsrosco *MemsRosco) Convert(filepath string) *Scenario {
	scenario := NewScenario()

	// open the file
	memsrosco.openFile(filepath)

	// marshall into the correct format
	roscofile, _ := newLineSkipDecoder(memsrosco.file, 3)

	if err := gocsv.UnmarshalDecoder(roscofile, &memsrosco.roscoData); err != nil {
		utils.LogE.Printf("unable to parse file %s", err)
	} else {
		scenario.Count = len(memsrosco.roscoData)
		utils.LogI.Printf("loaded scenario %s (%d dataframes)", filepath, scenario.Count)
	}

	return memsrosco.scenario
}

// Open the CSV  file
func (memsrosco *MemsRosco) openFile(filepath string) {
	var err error

	memsrosco.file, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		utils.LogE.Printf("unable to open %s", err)
	}
}

// ConvertCSVToMemsFCR takes Readmems CSV files and converts them into MemsFCR format
/*
func (scenario *Scenario) ConvertCSVToMemsFCR(filepath string) {
	// load the Readmems CSV
	scenario.Load(filepath)

	// recreate the Dataframes from the CSV values
	for _, m := range scenario.Memsdata {
		scenario.recreateDataframes(m)
	}
}
*/

// Recreate the Dataframe HEX data from the parameters
// The CSV data fields are calculated from the raw data, we need to undo
// those computations
func (memsrosco *MemsRosco) recreateDataframes(data *roscoData) {
	// undo all the computations and put all data back into integer/hex format
	df80 := fmt.Sprintf("801C"+
		"%04x%02x%02x%02x%02x%02x%02x%02x%02x%02x"+
		"%02x%02x%02x%02x%02x%02x%02x%04x%02x%02x%04x"+
		"%02x%02x%02x",
		uint16(data.EngineRPM),
		uint8(data.CoolantTemp+55),
		uint8(data.AmbientTemp+55),
		uint8(data.IntakeAirTemp+55),
		uint8(data.FuelTemp+55),
		uint8(data.ManifoldAbsolutePressure),
		uint8(data.BatteryVoltage*10),
		uint8(data.ThrottlePotSensor/0.02),
		utils.ConvertBooltoInt(data.IdleSwitch),
		utils.ConvertBooltoInt(data.AirconSwitch),
		utils.ConvertBooltoInt(data.ParkNeutralSwitch),
		uint8(data.DTC0),
		uint8(data.DTC1),
		uint8(data.IdleSetPoint),
		uint8(data.IdleHot),
		uint8(data.Uk8011),
		uint8(math.Round(float64(data.IACPosition)*1.8)),
		uint16(data.IdleSpeedDeviation),
		uint8(data.IgnitionAdvanceOffset80),
		uint8((data.IgnitionAdvance*2)+24),
		uint16(data.CoilTime/0.002),
		utils.ConvertBooltoInt(data.CrankshaftPositionSensor),
		uint8(data.Uk801a),
		uint8(data.Uk801b),
	)

	df7d := fmt.Sprintf("7D20"+
		"%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x"+
		"%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x"+
		"%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x",
		utils.ConvertBooltoInt(data.IgnitionSwitch),
		uint8(data.ThrottleAngle/6*10),
		uint8(data.Uk7d03),
		uint8(data.AirFuelRatio*10),
		uint8(data.DTC2),
		uint8(data.LambdaVoltage/5),
		uint8(data.LambdaFrequency),
		uint8(data.LambdaDutycycle),
		uint8(data.LambdaStatus),
		utils.ConvertBooltoInt(data.ClosedLoop),
		uint8(data.LongTermFuelTrim),
		uint8(data.ShortTermFuelTrim),
		uint8(data.CarbonCanisterPurgeValve),
		uint8(data.DTC3),
		uint8(data.IdleBasePosition),
		uint8(data.Uk7d10),
		uint8(data.DTC4),
		uint8(data.IgnitionAdvanceOffset7d+48),
		uint8((data.IdleSpeedOffset+128)/25),
		uint8(data.Uk7d14),
		uint8(data.Uk7d15),
		uint8(data.DTC5),
		uint8(data.Uk7d17),
		uint8(data.Uk7d18),
		uint8(data.Uk7d19),
		uint8(data.Uk7d1a),
		uint8(data.Uk7d1b),
		uint8(data.Uk7d1c),
		uint8(data.Uk7d1d),
		uint8(data.Uk7d1e),
		uint8(data.JackCount),
	)

	data.Dataframe7d = strings.ToUpper(df7d)
	data.Dataframe80 = strings.ToUpper(df80)

	fmt.Printf("0x80: %s\n", data.Dataframe80)
	fmt.Printf("0x7d: %s\n", data.Dataframe7d)
}

func newLineSkipDecoder(r io.Reader, LinesToSkip int) (gocsv.SimpleDecoder, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = -1
	for i := 0; i < LinesToSkip; i++ {
		if _, err := reader.Read(); err != nil {
			return nil, err
		}
	}
	reader.FieldsPerRecord = 0
	return gocsv.NewSimpleDecoderFromCSVReader(reader), nil
}
