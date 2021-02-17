package scenarios

import (
	"os"

	"github.com/andrewdjackson/memscene/utils"
	"github.com/gocarina/gocsv"
)

// Scenario represents the scenario data
type Scenario struct {
	// Memsdata log
	Memsdata []*MemsFCRData
	// Position in the log
	Position int
	// Count of items in the log
	Count int
}

// NewScenario creates a new scenario
func NewScenario() *Scenario {
	scenario := &Scenario{}
	// initialise the log
	scenario.Memsdata = []*MemsFCRData{}
	// start at the beginning
	scenario.Position = 0
	// no items in the log
	scenario.Count = 0

	return scenario
}

// SaveCSVFile saves the Memdata to a CSV file
func (scenario *Scenario) SaveCSVFile(filepath string) {
	file, _ := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)

	err := gocsv.MarshalFile(&scenario.Memsdata, file)
	if err != nil {
		utils.LogI.Printf("error saving csv file %s", err)
	}
}
