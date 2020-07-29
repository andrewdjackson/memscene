package main

import (
	"flag"
	"fmt"

	"github.com/andrewdjackson/memscene/scenarios"
	"github.com/andrewdjackson/memscene/utils"
)

func main() {
	file := flag.String("file", "", "file to convert")
	convert := flag.String("type", "none", "use 'readmems' to convert from readmems log,\n'memsrosco' to convert from mems-rosco csv")
	output := flag.String("output", "", "destination file")
	flag.Parse()

	utils.LogI.Printf("using command line, file: %s, type: %s", *file, *convert)

	scenario := scenarios.NewScenario()

	// convert the file and exit
	switch *convert {
	case "readmems":
		utils.LogI.Printf("converting from readmems file to MemsFCR")
		r := scenarios.NewReadMems()
		scenario = r.Convert(*file)
	case "memsrosco":
		utils.LogI.Printf("converting from memsrosco file to MemsFCR")
		r := scenarios.NewMemsRosco()
		scenario = r.Convert(*file)
	}

	save := fmt.Sprintf("%s", *output)
	scenario.SaveCSVFile(save)
}
