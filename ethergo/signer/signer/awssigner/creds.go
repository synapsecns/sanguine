package awssigner

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
)

// credentialProvider provides credentials to the kms signer.
type credentialProvider struct {
	// awsAccessKey is the aws access key used for kms
	awsAccessKey string
	// awsSecretAccessKey is the secret key used for kms
	awsSecretAccessKey string
}

// newCredentialProvider creates a new credential provider for aws.
func newCredentialProvider(awsAccessKey, awsSecretAccessKey string) *credentialProvider {
	return &credentialProvider{
		awsAccessKey:       awsAccessKey,
		awsSecretAccessKey: awsSecretAccessKey,
	}
}

// Retrieve retreives credentials from the local provider.
func (c credentialProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     c.awsAccessKey,
		SecretAccessKey: c.awsSecretAccessKey,
		Source:          "config",
	}, nil
}

var _ aws.CredentialsProvider = &credentialProvider{}
