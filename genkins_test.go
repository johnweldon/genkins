package genkins_test

import (
	"testing"

	"github.com/johnweldon/genkins"
)

func TestGetInfo(t *testing.T) {
	node, err := genkins.GetInfo("https://ci.jenkins.io/api/json", "", "")
	if err != nil {
		t.Error(err)
	}

	for _, j := range node.AllJobs() {
		t.Logf("%s\n", j)
	}

	for _, j := range node.BadJobs() {
		t.Logf("%s\n", j)
	}
}
