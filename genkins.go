package genkins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

// Color represents the current build status
type Color string

// Known colors
const (
	Red         Color = "red"
	Blue        Color = "blue"
	RedRunning  Color = "red_anime"
	BlueRunning Color = "blue_anime"
	Disabled    Color = "disabled"
)

// Job represents basic job information
type Job struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Color Color  `json:"color"`
}

// ByColor is used to sort []Job by color
type ByColor []Job

func (b ByColor) Len() int      { return len(b) }
func (b ByColor) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByColor) Less(i, j int) bool {
	if b[i].colorPrecedence() == b[j].colorPrecedence() {
		return b[i].Name < b[j].Name
	}
	return b[i].colorPrecedence() < b[j].colorPrecedence()
}

// String implements Stringer
func (j Job) String() string {
	var state string
	switch j.Color {
	case Blue, BlueRunning:
		state = "good"
	case Red, RedRunning:
		state = "bad"
	default:
		state = "unknown"
	}

	return fmt.Sprintf("%8s : %s", state, j.Name)
}

var colorPrecedence map[Color]int = map[Color]int{
	Red: 0, RedRunning: 1,
	Blue: 10, BlueRunning: 11,
}

func (j Job) colorPrecedence() int {
	p, ok := colorPrecedence[j.Color]
	if ok {
		return p
	}
	return 100
}

// BaseNode shows root level information about this Jenkins server
type BaseNode struct {
	Mode            string `json:"mode"`
	NodeDescription string `json:"nodeDescription"`
	NodeName        string `json:"nodeName"`
	Executors       int    `json:"numExecutors"`
	Description     string `json:"description"`
	Jobs            []Job  `json:"jobs"`
}

// BadJobs returns the list of jobs filtered by Color and only returning
// Jobs that are Red or RedRunning
func (n *BaseNode) BadJobs() []Job {
	retval := []Job{}
	for _, job := range n.Jobs {
		switch job.Color {
		case Red, RedRunning:
			retval = append(retval, job)
		}
	}
	sort.Sort(ByColor(retval))
	return retval
}

// AllJobs returns all the jobs sorted by color, then name
func (n *BaseNode) AllJobs() []Job {
	retval := make([]Job, len(n.Jobs))
	for i, job := range n.Jobs {
		retval[i] = job
	}
	sort.Sort(ByColor(retval))
	return retval
}

// GetInfo reads the json api endpoint of a Jenkins server and returns
// BaseNode information
func GetInfo(url string) (BaseNode, error) {
	r, err := http.Get(url)
	if err != nil {
		return BaseNode{}, err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return BaseNode{}, err
	}

	node := BaseNode{}
	err = json.Unmarshal(body, &node)
	if err != nil {
		return BaseNode{}, err
	}

	return node, nil
}
