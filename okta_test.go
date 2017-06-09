package okta

import (
	"testing"

	"github.com/apex/log"
)

func TestAuthenticate(t *testing.T) {
	var err error
	var awsConfig *AwsConfig

	log.SetLevel(log.DebugLevel)

	username := "sloth@segment.com"
	password := "\\tszkGpO)h7%qWZVI20?bRg}`\"SLOr"
	profile := "stage"

	awsConfig, err = NewAwsConfig()
	awsConfig.Load()

	roleArn := awsConfig.GetProfileArn(profile)
	if roleArn == "" {
		t.Errorf("%s profile not found in AWS configuration", profile)
	}

	client := NewOktaClient(
		OktaOrganization,
		username,
		password,
	)

	err = client.Authenticate(roleArn, profile)
	if err != nil {
		t.Error(err)
	}
}
