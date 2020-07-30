package utils

import (
	"bufio"
	"os"
	"strings"
)

const (
	// ReadMemsFile a data file from readmems
	ReadMemsFile = "readmems"
	// MemsRoscoFile a data file from mems-rosco (LeopoldG readmems)
	MemsRoscoFile = "memsrosco"
	// MemsRoscoFilev2 a data file from mems-rosco latest version (LeopoldG readmems)
	MemsRoscoFilev2 = "memsroscov2"
	// MemsDiagFile a data file from mems-diag (Haro?)
	MemsDiagFile = "memsdiag"
	// MemsFCRFile that's mine
	MemsFCRFile = "memsfcr"
	// Unknown eh?
	Unknown = "unknown"
)

// GetFileType determines the file type
func GetFileType(path string) string {
	file, err := os.Open(path)

	if err != nil {
		LogE.Printf("unable to open %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "80: 1C") {
			return ReadMemsFile
		}

		if strings.HasPrefix(line, "#time,engine-rpm,coolant_temp,ambient_temp,") {
			return MemsRoscoFile
		}

		if strings.HasPrefix(line, "#time,80x01-02_engine-rpm,80x03_coolant_temp,") {
			if strings.HasSuffix(line, "0x7d_raw,0x80_raw") {
				// memsfcr data file
				return MemsFCRFile
			}

			// mems-rosco v2 is similar to memsfcr but without the raw data
			return MemsRoscoFilev2
		}

		if strings.HasPrefix(line, "Time,RPM,IdleError,IdlePos(Steps),") {
			// memsdiag data file
			return MemsDiagFile
		}
	}

	return Unknown
}

// OpenFile opens the  file
func OpenFile(filepath string) *os.File {
	file, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)

	if err != nil {
		LogE.Printf("unable to open %s", err)
	}

	return file
}
