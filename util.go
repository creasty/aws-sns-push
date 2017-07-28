package main

import (
	"os"
	"regexp"
	"strconv"
)

// TargetMode represents a kind of Target
type TargetMode int

// Enum of TargetMode
const (
	TargetModeUnknown TargetMode = iota
	TargetModeUser
	TargetModeDeviceToken
	TargetModeEndpointArn
)

// Target represents a target to receive push notification
type Target struct {
	Mode            TargetMode
	ApplicationName string
	UserID          int64
	DeviceToken     string
	EndpointArn     string
}

var (
	targetUserPattern        = regexp.MustCompile(`^([\w-]+)/(\d+)(?:/([\w-]+))?$`)
	targetDeviceTokenPattern = regexp.MustCompile(`^([\w-]+)/([\w-]+)$`)
	targetEndpointPattern    = regexp.MustCompile(`^arn:aws:sns:[\w-]+:\d+:endpoint/.+$`)
)

// ParseTarget parses string-represented identifier into Target
func ParseTarget(id string) (t Target) {
	if targetEndpointPattern.MatchString(id) {
		t.Mode = TargetModeEndpointArn
		t.EndpointArn = id
		return
	}

	if m := targetUserPattern.FindStringSubmatch(id); len(m) == 4 {
		t.Mode = TargetModeUser
		t.ApplicationName = m[1]
		t.UserID, _ = strconv.ParseInt(m[2], 10, 64)
		t.DeviceToken = m[3]
		return
	}

	if m := targetDeviceTokenPattern.FindStringSubmatch(id); len(m) == 3 {
		t.Mode = TargetModeDeviceToken
		t.ApplicationName = m[1]
		t.DeviceToken = m[2]
		return
	}

	return
}

// IsPiped returns true when data is piped into Stdin
func IsPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
