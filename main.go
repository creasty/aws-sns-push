package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/creasty/aws-sns-push/version"
	"github.com/k0kubun/pp"
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
	pp.Println(target)
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

    2. {application-name}/{user-id}/{device-token}
       e.g., sample-production/12345/ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff

    3. {application-name}/{device-token}
       e.g., sample-production/ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff

    4. {endpoint-arn}
       e.g., arn:aws:sns:ap-northeast-1:000000000000:endpoint/sample-production/ffffffff-ffff-ffff-ffff-ffffffffffff

OPTIONS:
    -y    Send without confirmation
    -h    Show help
`)
	os.Exit(1)
}
