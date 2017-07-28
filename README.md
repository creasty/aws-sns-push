aws-sns-push
============

[![Build Status](https://travis-ci.org/creasty/aws-sns-push.svg?branch=master)](https://travis-ci.org/creasty/aws-sns-push)
[![GoDoc](https://godoc.org/github.com/creasty/aws-sns-push?status.svg)](https://godoc.org/github.com/creasty/aws-sns-push)
[![GitHub release](https://img.shields.io/github/release/creasty/aws-sns-push.svg)](https://github.com/creasty/aws-sns-push/releases)
[![License](https://img.shields.io/github/license/creasty/aws-sns-push.svg)](./LICENSE)

Send SNS push notifications painlessly.


```sh-session
$ aws-sns-push -h
Send SNS push notifications painlessly.

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
```
