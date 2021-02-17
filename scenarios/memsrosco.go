package scenarios

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/andrewdjackson/memscene/utils"
	"github.com/gocarina/gocsv"
)

// MemsRosco structure
type MemsRosco struct {
	scenario *Scenario
	file     *os.File
	data     []*MemsRoscoData
}

// NewMemsRosco create a new MemsRosco instance
func NewMemsRosco() *MemsRosco {
	memsrosco := &MemsRosco{}
	memsrosco.scenario = NewScenario()

	return memsrosco
}

// Convert takes Readmems Log files and converts them into MemsFCR format
func (memsrosco *MemsRosco) Convert(filepath string) *Scenario {
	// open the file
	memsrosco.file = utils.OpenFile(filepath)

	// marshall into the correct format
	roscofile, _ := utils.NewLineSkipDecoder(memsrosco.file, 1)

	if err := gocsv.UnmarshalDecoder(roscofile, &memsrosco.data); err != nil {
		utils.LogE.Printf("unable to parse file %s", err)
	} else {
		memsrosco.scenario.Count = len(memsrosco.data)
		utils.LogI.Printf("loaded scenario %s (%d dataframes)", filepath, memsrosco.scenario.Count)

		// recreate the Dataframes from the CSV values
		for _, m := range memsrosco.data {
			memsrosco.recreateDataframes(m)
		}
	}

	i, _ := json.Marshal(memsrosco.data)
	_ = json.Unmarshal(i, &memsrosco.scenario.Memsdata)

	return memsrosco.scenario
}

// Recreate the Dataframe HEX data from the parameters
// The CSV data fields are calculated from the raw data, we need to undo
// those computations
func (memsrosco *MemsRosco) recreateDataframes(data *MemsRoscoData) {
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
		uint8(data.IdleSwitch),
		uint8(data.AirconSwitch),
		uint8(data.ParkNeutralSwitch),
		uint8(data.DTC0),
		uint8(data.DTC1),
		uint8(data.IdleSetPoint),
		uint8(data.IdleHot+35),
		uint8(data.Uk8011),
		uint8(math.Round(float64(data.IACPosition)*1.8)),
		uint16(data.IdleSpeedDeviation),
		uint8(data.IgnitionAdvanceOffset80),
		uint8((data.IgnitionAdvance*2)+24),
		uint16(data.CoilTime/0.002),
		uint8(data.CrankshaftPositionSensor),
		uint8(data.Uk801a),
		uint8(data.Uk801b),
	)

	df7d := fmt.Sprintf("7D20"+
		"%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x"+
		"%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x"+
		"%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x",
		uint8(data.IgnitionSwitch),
		uint8(data.ThrottleAngle/60),
		uint8(data.Uk7d03),
		uint8(data.AirFuelRatio*10),
		uint8(data.DTC2),
		uint8(data.LambdaVoltage/5),
		uint8(data.LambdaFrequency),
		uint8(data.LambdaDutycycle),
		uint8(data.LambdaStatus),
		uint8(data.ClosedLoop),
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

	//fmt.Printf("0x80: %s\n", data.Dataframe80)
	//fmt.Printf("0x7d: %s\n", data.Dataframe7d)
}
