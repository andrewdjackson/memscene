package utils

import (
	"encoding/csv"
	"io"
	"math"

	"github.com/gocarina/gocsv"
)

// ConvertBooltoInt converts boolean to integer
func ConvertBooltoInt(b bool) uint8 {
	if b {
		return 1
	}

	return 0
}

// RoundTo2DecimalPoints rounds to 2 dp
func RoundTo2DecimalPoints(x float32) float32 {
	return float32(math.Round(float64(x)*100) / 100)
}

// NewLineSkipDecoder skips lines in the file
func NewLineSkipDecoder(r io.Reader, LinesToSkip int) (gocsv.SimpleDecoder, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = -1
	for i := 0; i < LinesToSkip; i++ {
		if _, err := reader.Read(); err != nil {
			return nil, err
		}
	}
	reader.FieldsPerRecord = 0
	return gocsv.NewSimpleDecoderFromCSVReader(reader), nil
}
