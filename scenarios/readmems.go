package scenarios

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"math"
	"os"
	"strings"
	"time"

	"github.com/andrewdjackson/memscene/utils"
)

// ReadMems logs are piped output from the readmems console and
// in the format:
// ECUID:
// 7D: 00 00 ...
// 80: 00 00 ...
//

// ReadMems structure
type ReadMems struct {
	scenario *Scenario
	memsdata *MemsFCRData
}

// NewReadMems create a new ReadMems instance
func NewReadMems() *ReadMems {
	readmems := &ReadMems{}
	readmems.scenario = NewScenario()

	return readmems
}

// Convert takes Readmems Log file and converts into MemsFCR format
func (readmems *ReadMems) Convert(filepath string) *Scenario {
	lines := readmems.readResponseFile(filepath)
	startTime := time.Now()

	// convert to a compress byte string line by line
	// from 80: 00 00...
	// to 800000...
	// skip header lines
	for i := 3; i < len(lines); i++ {
		line := lines[i]

		if readmems.isACommandResponse(line) {
			line = readmems.cleanCommandResponse(line)

			if strings.HasPrefix(line, "80") {
				readmems.memsdata = &MemsFCRData{}
				readmems.memsdata.Dataframe80 = line
			}
			if strings.HasPrefix(line, "7D") {
				readmems.memsdata.Dataframe7d = line
			}

			// convert to a struct, response order is 80, 7d so
			// make sure we don't increment until we have both dataframes
			if readmems.memsdata.Dataframe80 != "" && readmems.memsdata.Dataframe7d != "" {
				startTime = startTime.Add(1 * time.Second)
				readmems.memsdata.Time = startTime.Format("15:04:05")
				readmems.calculateMemsData(readmems.memsdata)
				readmems.scenario.Memsdata = append(readmems.scenario.Memsdata, readmems.memsdata)
				readmems.scenario.Count++
			}
		}
	}

	return readmems.scenario
}

// returns true if the line starts with an 80 or 75 which
// indicate command responses
func (readmems *ReadMems) isACommandResponse(line string) bool {
	if strings.HasPrefix(line, "80") {
		return true
	}

	if strings.HasPrefix(line, "7D") {
		return true
	}

	return false
}

// returns compacted and cleaned response
func (readmems *ReadMems) cleanCommandResponse(line string) string {
	// remove the : character
	line = strings.ReplaceAll(line, ":", "")
	// remove all the spaces
	line = strings.ReplaceAll(line, " ", "")
	// remove all the LF
	line = strings.ReplaceAll(line, "\n", "")
	line = strings.ReplaceAll(line, "\r", "")

	return line
}

// ReadMems logs are piped output from the readmems console and
// in the format:
// 7D: 00 00 ...
// 80: 00 00 ...
// this function reads the response file into an array of strings for
// processing
func (readmems *ReadMems) readResponseFile(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		utils.LogE.Printf("unable to open %s", err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// calculateMemsData reads the raw dataframes and returns structured data
func (readmems *ReadMems) calculateMemsData(memsdata *MemsFCRData) {
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

	t := time.Now()

	// calculate IAC postion, 0 closed - 180 fully open
	// convert to %
	//iac := math.Round(float64(df80.IacPosition) / 1.8)
	//if iac > 100 {
	// if value is > 100% then cap it
	//	iac = 100
	//}

	// build the Mems Data frame using the raw data and applying the relevant
	// adjustments and calculations
	readmems.memsdata = &MemsFCRData{
		Time:                     t.Format("15:04:05.000"),
		EngineRPM:                int(df80.EngineRpm),
		CoolantTemp:              int(df80.CoolantTemp) - 55,
		AmbientTemp:              int(df80.AmbientTemp) - 55,
		IntakeAirTemp:            int(df80.IntakeAirTemp) - 55,
		FuelTemp:                 int(df80.FuelTemp) - 55,
		ManifoldAbsolutePressure: float32(df80.ManifoldAbsolutePressure),
		BatteryVoltage:           float32(df80.BatteryVoltage) / 10,
		ThrottlePotSensor:        roundTo2DecimalPoints(float32(df80.ThrottlePotSensor) * 0.02),
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
		CoilTime:                 roundTo2DecimalPoints(float32(df80.CoilTime) * 0.002),
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
		IdleSpeedOffset:          int(df7d.IdleSpeedOffset), // - 128) * 25,
		DTC5:                     df7d.Dtc5,
		JackCount:                int(df7d.JackCount),
		Dataframe80:              hex.EncodeToString(d80),
		Dataframe7d:              hex.EncodeToString(d7d),
	}

	utils.LogI.Printf("%s built mems dataframe %v", utils.ECUCommandTrace, readmems.memsdata)
}

func roundTo2DecimalPoints(x float32) float32 {
	return float32(math.Round(float64(x)*100) / 100)
}
