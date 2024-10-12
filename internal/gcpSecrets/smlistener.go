package smlistener

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

// SecretManagerService defines an interface for interacting with Google Cloud Secret Manager
type SecretManagerService interface {
	GetSecret(ctx context.Context, projectID, secretID string) (string, error)
}

// GCPSecretManagerService is an implementation of SecretManagerService
type GCPSecretManagerService struct {
	client *secretmanager.Client
}

// NewGCPSecretManagerService initializes a new GCPSecretManagerService
func NewGCPSecretManagerService(client *secretmanager.Client) *GCPSecretManagerService {
	return &GCPSecretManagerService{client: client}
}

// GetSecret retrieves the secret value from Google Cloud Secret Manager
func (s *GCPSecretManagerService) GetSecret(ctx context.Context, projectID, secretID string) (string, error) {
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretID),
	}

	result, err := s.client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", err
	}

	// Return the secret payload
	return string(result.Payload.Data), nil
}
