package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var limit int

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "limit",
			Value: 10,
			Usage: "number of results to return",
			Destination: &limit,
		},
	}

	app.Name = "gm"
	app.Usage = "deal with information related to George Mason High School"
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Giving %d results\n", limit)
		return nil
	}

	app.Run(os.Args)
}
