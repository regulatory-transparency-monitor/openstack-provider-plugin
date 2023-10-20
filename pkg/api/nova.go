package api

import (
	"encoding/json"
	"fmt"

	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/httpwrapper"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/models"
)

type NovaService struct {
	BaseURL string
	Client  *httpwrapper.HTTPClient
}

// b.3 Nova GET server list by projectID, returns 200
func (n *NovaService) GetServerListByProjectID(projectID string) (models.ServerList, error) {
	endpoint := fmt.Sprintf("%sservers?project_id=%s", n.BaseURL, projectID)

	// Construct the GET request.
	req, err := n.Client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return models.ServerList{}, fmt.Errorf("failed to create GetServerListByProjectID request: %w", err)
	}

	// Execute the request.
	res, err := n.Client.Do(req)
	if err != nil {
		return models.ServerList{}, fmt.Errorf("request GetServerListByProjectID failed: %w", err)
	}
	defer res.Body.Close()

	// Decode the response.
	var response models.ServerList
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return models.ServerList{}, fmt.Errorf("failed to decode GetServerListByProjectID response: %w", err)
	}

	return response, err

}

// Get server details by ID
func (n *NovaService) GetServerByID(serverID string) (models.ServerDetails, error) {
	// Iterate over each server and make the API call
	endpoint := fmt.Sprintf("%sservers/%s", n.BaseURL, serverID)

	// Construct the GET request.
	req, err := n.Client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return models.ServerDetails{}, fmt.Errorf("failed to construct GetServerByID request: %w", err)
	}

	// Execute the request.
	res, err := n.Client.Do(req)
	if err != nil {
		return models.ServerDetails{}, fmt.Errorf("failed to execute GetServerByID request: %w", err)
	}
	defer res.Body.Close()

	// Decode the response.
	var response models.ServerDetails

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return models.ServerDetails{}, fmt.Errorf("failed to decode GetServerByID response: %w", err)
	}

	return response, err

}

// Get volumes attached to server
func (n *NovaService) GetAttachedVolumes(serverID string) (models.ServerVolumeAttachments, error) {

	// Iterate over each server and make the API call
	endpoint := fmt.Sprintf("%sservers/%s/os-volume_attachments", n.BaseURL, serverID)

	// Construct the GET request.
	req, err := n.Client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return models.ServerVolumeAttachments{}, fmt.Errorf("failed to construct GetAttachedVolumes request: %w", err)
	}

	// Execute the request.
	res, err := n.Client.Do(req)
	if err != nil {
		return models.ServerVolumeAttachments{}, fmt.Errorf("failed to execute GetAttachedVolumes request: %w", err)
	}
	defer res.Body.Close()

	// Decode the response.
	var response models.ServerVolumeAttachments

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return response, fmt.Errorf("failed to decode GetAttachedVolumes response: %w", err)
	}

	return response, err
}
