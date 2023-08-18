package api

import (
	"encoding/json"
	"fmt"

	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/client"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/models"
)

type NovaService struct {
	BaseURL string
	Client  *client.HTTPClient
}

// b.3 Nova GET server list by projectID, returns 200
func (n *NovaService) GetServerListByProjectID(projectID string) (*models.NovaResponse, error) {
	endpoint := fmt.Sprintf("%sservers?project_id=%s", n.BaseURL, projectID)

	// Construct the GET request.
	req, err := n.Client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Execute the request.
	res, err := n.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	// Decode the response.
	var response models.NovaResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, err

}

/*
// b.3 Nova GET server by projectID, returns 200
func sendAuthRequest(xAuthToken, projectID string) {
	// Construct URL with the project ID ad path parameter
	fullUrl := fmt.Sprintf("https://pub1.infomaniak.cloud/compute/v2.1/servers?project_id=%s", projectID)
	// Set header
	header := n.Client.GetHeader()
	var response models.NovaResponse

	response, _, err := client.MakeHTTPRequest(fullUrl, "GET", headers, nil, nil, response)
	// Handle the response
	if err != nil {
		fmt.Errorf("Error while converting: %s", err)
	}

} */

/* func (n *NovaService) GetServerByID(response models.NovaResponse) {
	// Iterate over each server and make the API call
	for _, server := range response.Servers {
		apiEndpoint := fmt.Sprintf("https://pub1.infomaniak.cloud/compute/v2.1/servers/%s", server.ID)

		var serverResp models.NovaServerPayload
		header := n.Client.GetHeader()
		// Make the API call to the endpoint using the server ID
		serverResp, _, err := client.MakeHTTPRequest(apiEndpoint, "GET", header, nil, nil, serverResp)
		if err != nil {
			fmt.Errorf("Error while coverting: %s", err)

		}
	}

} */

/*
// Doenst return anthing so far
func getServerIP(response models.NovaResponse, appLogger *logger.APIlogger, headers map[string]string) {
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

func getMetadata(response models.NovaResponse, appLogger *logger.APIlogger, headers map[string]string) {
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
