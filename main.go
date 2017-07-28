package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/creasty/aws-sns-push/aws"
	"github.com/creasty/aws-sns-push/version"
)

var flagYes bool
var flagVersion bool

func init() {
	flag.BoolVar(&flagYes, "y", false, "Send without confirmation")
	flag.BoolVar(&flagVersion, "v", false, "Show version")
	flag.Usage = showHelp
	flag.Parse()
}

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	if flagVersion {
		showVersion()
		return
	}

	target := ParseTarget(os.Args[1])
	if target.Mode == TargetModeUnknown {
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
       e.g., sample-production/12345

    2. {application-name}/{device-token}
       e.g., sample-production/ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff

    3. {endpoint-arn}
       e.g., arn:aws:sns:ap-northeast-1:000000000000:endpoint/sample-production/ffffffff-ffff-ffff-ffff-ffffffffffff

OPTIONS:
    -y    Send without confirmation
    -h    Show help
`)
	os.Exit(1)
}

func doSend(target Target) {
	s := aws.NewSNS()

	endpoints := make([]string, 0)

	switch target.Mode {
	case TargetModeEndpointArn:
		endpoints = append(endpoints, target.EndpointArn)
	case TargetModeUserID, TargetModeDeviceToken:
		eps, err := s.FindEndpointsFor(target.ApplicationName, target.UserID, target.DeviceToken)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		endpoints = append(endpoints, eps...)
	}

	if !IsPiped() {
		fmt.Println("Enter JSON message (Ctrl-D):")
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	payload := string(bytes)
	if payload == "" {
		fmt.Fprintf(os.Stderr, "No message given\n")
		os.Exit(1)
	}

	if err := s.Send(endpoints, payload); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
