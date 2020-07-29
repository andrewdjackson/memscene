package scenarios

import (
	"bufio"
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
	memsdata *MemsData
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
				readmems.memsdata = &MemsData{}
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
				readmems.scenario.Memsdata = append(readmems.scenario.Memsdata, readmems.memsdata)
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
