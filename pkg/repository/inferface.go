package repository

import (
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/models"
)

type KeystoneRepository interface {
	Authenticate() (string, error)
	GetProjectDetailsByID(projectID string) (*models.ProjectDetails, error)
}

type NovaRepository interface {
	GetServerListByProjectID(projectID string) (*models.NovaResponse, error)
	// GetServerByID(token string, serverID string) (*models.Server, error)
	//
}
