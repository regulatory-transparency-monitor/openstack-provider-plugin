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
		return models.ServerList{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Execute the request.
	res, err := n.Client.Do(req)
	if err != nil {
		return models.ServerList{}, fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	// Decode the response.
	var response models.ServerList
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return models.ServerList{}, fmt.Errorf("failed to decode response: %w", err)
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
		return models.ServerDetails{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Execute the request.
	res, err := n.Client.Do(req)
	if err != nil {
		return models.ServerDetails{}, fmt.Errorf("failed to create request: %w", err)
	}
	defer res.Body.Close()

	// Decode the response.
	var response models.ServerDetails
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return models.ServerDetails{}, fmt.Errorf("failed to create request: %w", err)
	}

	return response, err

}

/*
// Doenst return anthing so far
func getServerIP(response models.ServerList, appLogger *logger.APIlogger, headers map[string]string) {
	// Iterate over each server and make the API call
	for _, server := range response.Servers {

		apiEndpoint := fmt.Sprintf("https://pub1.infomaniak.cloud/compute/v2.1/servers/%s/ips", server.ID)

		var response models.Addresses
		// Make the API call to the endpoint using the server ID

		response, _, err := client.MakeHTTPRequest(apiEndpoint, "GET", headers, nil, nil, response)

		// Handle the response
		appLogger.Infof("Response: %+s", response)
		if err != nil {
			appLogger.Fatalf("Error while coverting: %s", err)

		}

	}
}

func getMetadata(response models.ServerList, appLogger *logger.APIlogger, headers map[string]string) {
	// Iterate over each server and make the API call
	for _, server := range response.Servers {

		apiEndpoint := fmt.Sprintf("https://pub1.infomaniak.cloud/compute/v2.1/servers/%s/metadata", server.ID)

		var response models.Metadata
		// Make the API call to the endpoint using the server ID

		response, _, err := client.MakeHTTPRequest(apiEndpoint, "GET", headers, nil, nil, response)

		// Handle the response
		appLogger.Infof("Response: %+s", response)
		if err != nil {
			appLogger.Fatalf("Error while coverting: %s", err)

		}

	}
}
*/
