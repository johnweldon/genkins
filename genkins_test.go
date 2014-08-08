package genkins_test

import (
	"github.com/johnweldon/genkins"
	"testing"
)

func TestGetInfo(t *testing.T) {
	node, err := genkins.GetInfo("https://ci.jenkins-ci.org/api/json")
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
