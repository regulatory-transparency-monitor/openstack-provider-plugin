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
	GetAttachedVolumes(serverID string) (models.ServerVolumeAttachments, error)
}

type CinderAPI interface {
	GetVolumesByProjectID(projectID string) ([]interface{}, error) // not working provider issue
	GetVolumeByID(volumeID string, projectID string) (models.Volume, error)
	GetSnapshots(projectID string) ([]models.Snapshot, error)
}

type APItest interface {
	Test() ([]interface{}, error)
}
