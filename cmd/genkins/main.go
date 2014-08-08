package main

import (
	"fmt"
	"os"

	"github.com/johnweldon/genkins"

	"github.com/codegangsta/cli"
)

func main() {

	commonFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "url, u",
			EnvVar: "GENKINS_URL",
			Value:  "https://ci.jenkins-ci.org/api/json",
			Usage:  "url for the jenkins server json api",
		},
		cli.BoolFlag{
			Name:   "full, f",
			EnvVar: "GENKINS_FULL",
			Usage:  "Show full build info",
		},
	}

	app := cli.NewApp()
	app.Name = "genkins"
	app.Usage = "get jenkins updates"
	app.Action = doIt
	app.Commands = []cli.Command{
		{
			Name:      "all",
			ShortName: "a",
			Usage:     "show status of all jobs",
			Action:    showAll,
			Flags:     commonFlags,
		},
		{
			Name:      "bad",
			ShortName: "b",
			Usage:     "show bad jobs",
			Action:    showBad,
			Flags:     commonFlags,
		},
	}
	app.Run(os.Args)
}

func doIt(c *cli.Context) { fmt.Fprintf(os.Stdout, "%s\n", c.Command.Usage) }

func showAll(c *cli.Context) {
	node, err := genkins.GetInfo(c.String("url"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	for _, job := range node.AllJobs() {
		if c.Bool("full") {
			fmt.Fprintf(os.Stdout, "%-45s : %s\n", job, job.URL)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", job)
		}
	}
}

func showBad(c *cli.Context) {
	node, err := genkins.GetInfo(c.String("url"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	for _, job := range node.BadJobs() {
		if c.Bool("full") {
			fmt.Fprintf(os.Stdout, "%-45s : %s\n", job, job.URL)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", job)
		}
	}
}
