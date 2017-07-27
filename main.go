package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/creasty/aws-sns-push/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "aws-sns-push"
	app.Usage = "Send SNS push notifications painlessly"
	app.UsageText = "aws-sns-push [options]"
	app.Version = fmt.Sprintf("%s (%s)", version.Version, version.Revision)

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "app, a",
			Usage: "Name of a platform application",
		},
		cli.IntFlag{
			Name:  "user, u",
			Usage: "Target user ID",
		},
		cli.StringFlag{
			Name:  "token, t",
			Usage: "Filter by a device token prefix",
		},
		cli.BoolFlag{
			Name:  "yes, y",
			Usage: "Send without confirmation",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println(c.GlobalString("app"))
		fmt.Println(c.GlobalString("token"))
		fmt.Println(c.GlobalInt("user"))
		return nil
	}

	app.Run(os.Args)
}
