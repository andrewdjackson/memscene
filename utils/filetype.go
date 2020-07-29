package utils

import (
	"bufio"
	"os"
	"strings"
)

const (
	ReadMemsFile  = "readmems"
	MemsRoscoFile = "memsrosco"
	MemsFCRFile   = "memsfcr"
	Unknown       = "unknown"
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
			return MemsFCRFile
		}
	}

	return Unknown
}
