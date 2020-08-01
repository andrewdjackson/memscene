package scenarios

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"math"
	"os"

	"github.com/andrewdjackson/memscene/utils"
	"github.com/gocarina/gocsv"
)

// MemsFCR structure
type MemsFCR struct {
	scenario *Scenario
	file     *os.File
	data     []*MemsFCRRawData
}

// NewMemsFCR create a new MemsRosco instance
func NewMemsFCR() *MemsFCR {
	memsfcr := &MemsFCR{}
	memsfcr.scenario = NewScenario()

	return memsfcr
}

// Convert takes Readmems Log files and converts them into MemsFCR format
func (memsfcr *MemsFCR) Convert(filepath string) *Scenario {
	// open the file
	memsfcr.file = utils.OpenFile(filepath)

	if err := gocsv.Unmarshal(memsfcr.file, &memsfcr.data); err != nil {
		utils.LogE.Printf("unable to parse file %s", err)
	} else {
		memsfcr.scenario.Count = len(memsfcr.data)
		utils.LogI.Printf("loaded scenario %s (%d dataframes)", filepath, memsfcr.scenario.Count)

		// recreate the Dataframes from the CSV values
		for _, m := range memsfcr.data {
			memsfcr.calculateMemsData(m)
		}
	}

	i, _ := json.Marshal(memsfcr.data)
	json.Unmarshal(i, &memsfcr.scenario.Memsdata)

	return memsfcr.scenario
}

func (memsfcr *MemsFCR) calculateMemsData(memsdata *MemsFCRRawData) {
	d80, _ := hex.DecodeString(memsdata.Dataframe80)
	d7d, _ := hex.DecodeString(memsdata.Dataframe7d)

	utils.LogI.Printf("%s getting x7d and x80 dataframes", utils.ECUCommandTrace)

	// populate the DataFrame structure for command 0x80
	r := bytes.NewReader(d80)
	var df80 DataFrame80

	if err := binary.Read(r, binary.BigEndian, &df80); err != nil {
		utils.LogE.Printf("%s dataframe x80 binary.Read failed: %v", utils.ECUCommandTrace, err)
	}

	// populate the DataFrame structure for command 0x7d
	r = bytes.NewReader(d7d)
	var df7d DataFrame7d

	if err := binary.Read(r, binary.BigEndian, &df7d); err != nil {
		utils.LogE.Printf("%s dataframe x7d binary.Read failed: %v", utils.ECUCommandTrace, err)
	}

	// calculate IAC postion, 0 closed - 180 fully open
	// convert to %
	iac := math.Round(float64(df80.IacPosition) / 1.8)
	if iac > 100 {
		// if value is > 100% then cap it
		iac = 100
	}

	// build the Mems Data frame using the raw data and applying the relevant
	// adjustments and calculations
	m := &MemsFCRRawData{
		Time:                     memsdata.Time,
		EngineRPM:                int(df80.EngineRpm),
		CoolantTemp:              int(df80.CoolantTemp) - 55,
		AmbientTemp:              int(df80.AmbientTemp) - 55,
		IntakeAirTemp:            int(df80.IntakeAirTemp) - 55,
		FuelTemp:                 int(df80.FuelTemp) - 55,
		ManifoldAbsolutePressure: float32(df80.ManifoldAbsolutePressure),
		BatteryVoltage:           float32(df80.BatteryVoltage) / 10,
		ThrottlePotSensor:        utils.RoundTo2DecimalPoints(float32(df80.ThrottlePotSensor) * 0.02),
		IdleSwitch:               bool(df80.IdleSwitch&IdleSwitchActive != 0),
		AirconSwitch:             bool(df80.AirconSwitch != 0),
		ParkNeutralSwitch:        bool(df80.ParkNeutralSwitch != 0),
		DTC0:                     df80.Dtc0,
		DTC1:                     df80.Dtc1,
		IdleSetPoint:             int(df80.IdleSetPoint),
		IdleHot:                  int(df80.IdleHot) - 35,
		IACPosition:              int(df80.IacPosition),
		IdleSpeedDeviation:       int(df80.IdleSpeedDeviation),
		IgnitionAdvanceOffset80:  int(df80.IgnitionAdvanceOffset80),
		IgnitionAdvance:          (float32(df80.IgnitionAdvance) / 2) - 24,
		CoilTime:                 utils.RoundTo2DecimalPoints(float32(df80.CoilTime) * 0.002),
		CrankshaftPositionSensor: bool(df80.CrankshaftPositionSensor != 0),
		IgnitionSwitch:           bool(df7d.IgnitionSwitch != 0),
		ThrottleAngle:            int(math.Round(float64(df7d.ThrottleAngle) * 6 / 10)),
		AirFuelRatio:             float32(df7d.AirFuelRatio) / 10,
		DTC2:                     df7d.Dtc2,
		LambdaVoltage:            int(df7d.LambdaVoltage) * 5,
		LambdaFrequency:          int(df7d.LambdaFrequency),
		LambdaDutycycle:          int(df7d.LambdaDutyCycle),
		LambdaStatus:             int(df7d.LambdaStatus),
		ClosedLoop:               bool(df7d.LoopIndicator != 0),
		LongTermFuelTrim:         int(df7d.LongTermFuelTrim) - 128,
		ShortTermFuelTrim:        int(df7d.ShortTermFuelTrim),
		CarbonCanisterPurgeValve: int(df7d.CarbonCanisterPurgeValve),
		DTC3:                     df7d.Dtc3,
		IdleBasePosition:         int(df7d.IdleBasePos),
		DTC4:                     df7d.Dtc4,
		IgnitionAdvanceOffset7d:  int(df7d.IgnitionAdvanceOffset7d) - 48,
		IdleSpeedOffset:          (int(df7d.IdleSpeedOffset) - 128) * 25,
		DTC5:                     df7d.Dtc5,
		JackCount:                int(df7d.JackCount),
		Dataframe80:              hex.EncodeToString(d80),
		Dataframe7d:              hex.EncodeToString(d7d),
	}

	*memsdata = *m
	utils.LogI.Printf("%s built mems dataframe %v", utils.ECUCommandTrace, memsdata)
}
