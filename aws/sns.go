package aws

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// SNS is a client of AWS SNS
type SNS struct {
	client *sns.SNS
}

// NewSNS initializes aws session and creates SNS
func NewSNS() *SNS {
	sess := session.Must(session.NewSession())

	return &SNS{
		client: sns.New(sess),
	}
}

func (s *SNS) getApplication(name string) (applicationArn string, err error) {
	params := &sns.ListPlatformApplicationsInput{}

	err = s.client.ListPlatformApplicationsPages(params, func(page *sns.ListPlatformApplicationsOutput, lastPage bool) bool {
		for _, app := range page.PlatformApplications {
			arn := aws.StringValue(app.PlatformApplicationArn)

			if strings.HasSuffix(arn, "/"+name) {
				applicationArn = arn
				return false
			}
		}
		return true
	})
	if err != nil {
		return
	}
	if applicationArn == "" {
		err = fmt.Errorf("No application found for %q", name)
		return
	}
	return
}

// FindEndpointsFor returns a array of endpoint arn matching with appName+userID or with specific deviceToken
func (s *SNS) FindEndpointsFor(appName string, userID int64, deviceToken string) (endpoints []string, err error) {
	applicationArn, err := s.getApplication(appName)
	if err != nil {
		return
	}

	params := &sns.ListEndpointsByPlatformApplicationInput{
		PlatformApplicationArn: aws.String(applicationArn),
	}

	err = s.client.ListEndpointsByPlatformApplicationPages(params, func(page *sns.ListEndpointsByPlatformApplicationOutput, lastPage bool) bool {
		for _, ep := range page.Endpoints {
			if aws.StringValue(ep.Attributes["Enabled"]) != "true" {
				continue
			}

			arn := aws.StringValue(ep.EndpointArn)
			token := aws.StringValue(ep.Attributes["Token"])

			if token == deviceToken {
				endpoints = append(endpoints, arn)
				return false
			}

			rawData := aws.StringValue(ep.Attributes["CustomUserData"])
			data := struct {
				UserID int64 `json:"user_id"`
			}{}
			json.Unmarshal([]byte(rawData), &data)

			if data.UserID == userID {
				endpoints = append(endpoints, arn)
				// return false
			}
		}
		return true
	})

	if err != nil {
		return
	}
	if len(endpoints) == 0 {
		if deviceToken != "" {
			err = fmt.Errorf("No endpoint found for deviceToken=%q", deviceToken)
		} else {
			err = fmt.Errorf("No endpoint found for user_id=%d", userID)
		}
	}
	return
}

// Send publishes push notifications
func (s *SNS) Send(endpoints []string, message string) error {
	for _, ep := range endpoints {
		params := &sns.PublishInput{
			Message:          aws.String(message),
			TargetArn:        aws.String(ep),
			MessageStructure: aws.String("json"),
		}

		if _, err := s.client.Publish(params); err != nil {
			return err
		}
	}

	return nil
}
