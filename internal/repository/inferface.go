package repository

import (
	"collector-service/pkg/models"
	"context"
)

type KeystoneRepository interface {
	Authenticate(ctx context.Context) (*models.Token, error)
	GetProjects() (models.ProjectDetails, error)
	// GetKeystoneServiceInstance() KeystoneRepository
	//IdentityToken() error
}

type NovaRepository interface {
	GetServers(*models.Token) (models.NovaResponse, error)
	// ... other methods as necessary ...
}

// .. Rest of the apis plus other interfaces
