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
func NewGCPSecretManagerService(ctx context.Context) (*GCPSecretManagerService, error) {
	secretManagerClient, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create Secret Manager client: %w", err)
	}
	return &GCPSecretManagerService{client: secretManagerClient}, nil
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
