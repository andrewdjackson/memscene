package scenarios

import (
	"encoding/json"
	"os"

	"github.com/andrewdjackson/memscene/utils"
	"github.com/gocarina/gocsv"
)

// MemsRoscoV2 structure
type MemsRoscoV2 struct {
	scenario  *Scenario
	file      *os.File
	roscoData []*MemsRoscoV2Data
}

// NewMemsRoscoV2 create a new MemsRosco instance
func NewMemsRoscoV2() *MemsRoscoV2 {
	memsrosco := &MemsRoscoV2{}
	memsrosco.scenario = NewScenario()

	return memsrosco
}

// Convert takes Readmems Log files and converts them into MemsFCR format
func (memsrosco *MemsRoscoV2) Convert(filepath string) *Scenario {
	// open the file
	memsrosco.file = utils.OpenFile(filepath)

	// marshall into the correct format
	roscofile, _ := utils.NewLineSkipDecoder(memsrosco.file, 1)

	if err := gocsv.UnmarshalDecoder(roscofile, &memsrosco.roscoData); err != nil {
		utils.LogE.Printf("unable to parse file %s", err)
	} else {
		memsrosco.scenario.Count = len(memsrosco.roscoData)
		utils.LogI.Printf("loaded scenario %s (%d dataframes)", filepath, memsrosco.scenario.Count)

		// recreate the Dataframes from the CSV values
		for _, m := range memsrosco.roscoData {
			memsrosco.recreateDataframes(m)
		}
	}

	i, _ := json.Marshal(memsrosco.roscoData)
	json.Unmarshal(i, &memsrosco.scenario.Memsdata)

	return memsrosco.scenario
}
