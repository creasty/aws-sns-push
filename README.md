aws-sns-push
============

[![Build Status](https://travis-ci.org/creasty/aws-sns-push.svg?branch=master)](https://travis-ci.org/creasty/aws-sns-push)
[![GoDoc](https://godoc.org/github.com/creasty/aws-sns-push?status.svg)](https://godoc.org/github.com/creasty/aws-sns-push)
[![GitHub release](https://img.shields.io/github/release/creasty/aws-sns-push.svg)](https://github.com/creasty/aws-sns-push/releases)
[![License](https://img.shields.io/github/license/creasty/aws-sns-push.svg)](./LICENSE)

Send SNS push notifications painlessly.


Installation
------------

```sh
$ brew install creasty/tools/aws-sns-push
```


Usage
-----

```sh-session
$ aws-sns-push
Send SNS push notifications painlessly.

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
```

### Send

```sh-session
$ aws-sns-push sample-prod/12345
==> Endpoints
- arn:aws:sns:ap-northeast-1:000000000000:endpoint/APNS/sample-prod/ffffffff-ffff-ffff-ffff-ffffffffffff
==> Enter JSON message (Ctrl-D)
{
  "APNS": "{\"aps\":{\"alert\":{\"title\":\"Hello from aws-sns-push\",\"body\":\"This is awesome\"}}}"
}
> Proceed to send [y] y
```

### Advance usage with `jq`

```sh-session
$ cat | jq '{ APNS: (. | tojson) }' | aws-sns-push -y sample-prod/12345
{
  "aps": {
    "alert": {
      "title": "Hello from aws-sns-push",
      "body": "This is awesome"
    }
  }
}
^D
```
