package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/creasty/aws-sns-push/aws"
	"github.com/creasty/aws-sns-push/util"
	"github.com/creasty/aws-sns-push/version"
)

var (
	flagYes     bool
	flagVersion bool
)

func init() {
	flag.BoolVar(&flagYes, "y", false, "Send without confirmation")
	flag.BoolVar(&flagVersion, "v", false, "Show version")
	flag.Usage = showHelp
	flag.Parse()
}

func main() {
	args := flag.Args()

	if len(args) < 1 {
		showHelp()
		return
	}

	if flagVersion {
		showVersion()
		return
	}

	target := util.ParseTarget(args[0])
	if target.Mode == util.TargetModeUnknown {
		fmt.Fprintf(os.Stderr, "Invalid target: %q\n", target.String)
		showHelp()
	}

	doSend(target)
}

func showVersion() {
	fmt.Printf("%s (%s)\n", version.Version, version.Revision)
}

func showHelp() {
	fmt.Printf(`Send SNS push notifications painlessly.

USAGE:
    aws-sns-push [OPTIONS] TARGET

TARGET:
    1. {application-name}/{user-id}
       e.g., sample-prod/12345

    2. {application-name}/{device-token}
       e.g., sample-prod/ffffff

    3. {endpoint-arn}
       e.g., arn:aws:sns:ap-northeast-1:0000:endpoint/APNS/sample-prod/ffffff

OPTIONS:
    -y    Send without confirmation
    -h    Show help
`)
	os.Exit(1)
}

func doSend(target util.Target) {
	s := aws.NewSNS()

	endpoints := make([]string, 0)
	isPiped := util.IsPiped()

	switch target.Mode {
	case util.TargetModeEndpointArn:
		endpoints = append(endpoints, target.EndpointArn)
	case util.TargetModeUserID, util.TargetModeDeviceToken:
		eps, err := s.FindEndpointsFor(target.ApplicationName, target.UserID, target.DeviceToken)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		endpoints = append(endpoints, eps...)
	}

	if !flagYes {
		fmt.Println("==> Endpoints")
		for _, ep := range endpoints {
			fmt.Printf("- %s\n", ep)
		}
	}

	if !isPiped {
		fmt.Println("==> Enter JSON message (Ctrl-D)")
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	message := string(bytes)
	if message == "" {
		fmt.Fprintf(os.Stderr, "No message given\n")
		os.Exit(1)
	}

	if !flagYes {
		if isPiped {
			fmt.Printf("==> Message\n%s\n", message)
		}

		if f, err := util.AskBool("Proceed to send"); !f {
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			} else {
				fmt.Fprintf(os.Stderr, "canceled\n")
			}
			os.Exit(1)
		}
	}

	if err := s.Send(endpoints, message); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
