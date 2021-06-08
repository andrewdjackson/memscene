package tests

import (
	"github.com/andrewdjackson/memscene/scenarios"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"path/filepath"
	"testing"
)

func getFilePath(filename string) string {
	return filepath.FromSlash(filename)
}

func TestReadMemsData(t *testing.T) {
	file := getFilePath("../data/readmems.data")
	r := scenarios.NewReadMems()
	scenario := r.Convert(file)
	then.AssertThat(t, scenario.Count, is.GreaterThan(0))
}
