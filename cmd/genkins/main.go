package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/johnweldon/genkins"
)

func main() {

	commonFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "url, u",
			EnvVar: "GENKINS_URL",
			Value:  "https://ci.jenkins.io/api/json",
			Usage:  "url for the jenkins server json api",
		},
		cli.BoolFlag{
			Name:   "full, f",
			EnvVar: "GENKINS_FULL",
			Usage:  "Show full build info",
		},
		cli.StringFlag{
			Name:   "id, i",
			EnvVar: "GENKINS_ID",
			Value:  "",
			Usage:  "user id for the jenkins server json api",
		},
		cli.StringFlag{
			Name:   "token, t",
			EnvVar: "GENKINS_TOKEN",
			Value:  "",
			Usage:  "user token for the jenkins server json api",
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
	node, err := genkins.GetInfo(c.String("url"), c.String("id"), c.String("token"))
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
	node, err := genkins.GetInfo(c.String("url"), c.String("id"), c.String("token"))
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
