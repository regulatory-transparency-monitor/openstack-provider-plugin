package interfaces

import (
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/models"
)

type KeystoneAPI interface {
	Authenticate() (string, error)
	GetProjectDetailsByID(projectID string) (models.ProjectDetails, error)
}

type NovaAPI interface {
	GetServerListByProjectID(projectID string) (models.ServerList, error)
	GetServerByID(serverID string) (models.ServerDetails, error)
	//
}
