package util

import (
	"regexp"
	"strconv"
)

// TargetMode represents a kind of Target
type TargetMode int

// Enum of TargetMode
const (
	TargetModeUnknown TargetMode = iota
	TargetModeUserID
	TargetModeDeviceToken
	TargetModeEndpointArn
)

// Target represents a target to receive push notification
type Target struct {
	Mode            TargetMode
	String          string
	ApplicationName string
	UserID          int64
	DeviceToken     string
	EndpointArn     string
}

var (
	targetUserIDPattern      = regexp.MustCompile(`^([\w-]+)/(\d+)$`)
	targetDeviceTokenPattern = regexp.MustCompile(`^([\w-]+)/([\w-]+)$`)
	targetEndpointPattern    = regexp.MustCompile(`^arn:aws:sns:[\w-]+:\d+:endpoint/.+$`)
)

// ParseTarget parses string-represented identifier into Target
func ParseTarget(id string) (t Target) {
	t.String = id

	if targetEndpointPattern.MatchString(id) {
		t.Mode = TargetModeEndpointArn
		t.EndpointArn = id
		return
	}

	if m := targetUserIDPattern.FindStringSubmatch(id); len(m) == 3 {
		t.Mode = TargetModeUserID
		t.ApplicationName = m[1]
		t.UserID, _ = strconv.ParseInt(m[2], 10, 64)
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
