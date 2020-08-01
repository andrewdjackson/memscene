package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/andrewdjackson/memscene/scenarios"
	"github.com/andrewdjackson/memscene/utils"
)

func main() {
	var file string
	var output string

	flag.StringVar(&file, "file", "", "file to convert")
	flag.StringVar(&output, "output", "", "destination file")
	flag.Parse()

	if output == "" {
		_, filename := filepath.Split(file)
		output = fmt.Sprintf("%s.output.csv", filename)
	}

	filetype := utils.GetFileType(file)
	utils.LogI.Printf("file identified as '%s' type", filetype)

	if filetype == utils.Unknown {
		utils.LogE.Fatalf("Unknown file type")
	}

	scenario := scenarios.NewScenario()

	// convert the file and exit
	switch filetype {
	case utils.ReadMemsFile:
		utils.LogI.Printf("converting from readmems file to MemsFCR")
		r := scenarios.NewReadMems()
		scenario = r.Convert(file)
	case utils.MemsRoscoFile:
		utils.LogI.Printf("converting from memsrosco file to MemsFCR")
		r := scenarios.NewMemsRosco()
		scenario = r.Convert(file)
	case utils.MemsRoscoFilev2:
		utils.LogI.Printf("converting from memsrosco version 2 file to MemsFCR")
		r := scenarios.NewMemsRoscoV2()
		scenario = r.Convert(file)
	case utils.MemsDiagFile:
		utils.LogW.Printf("unable to process memsdiag files, not enough data")
	case utils.MemsFCRFile:
		utils.LogI.Printf("reprocessing MemsFCR file")
		r := scenarios.NewMemsFCR()
		scenario = r.Convert(file)
	}

	if scenario.Count > 0 {
		save := fmt.Sprintf("%s", output)
		scenario.SaveCSVFile(save)
	}
}
