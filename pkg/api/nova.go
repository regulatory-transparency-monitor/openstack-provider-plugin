package api

import (
	"collector-service/pkg/client"
	"collector-service/pkg/models"
	"fmt"
)

type NovaService struct {
	authToken string
	projectID string
}

// b.3 Nova GET server by projectID, returns 200
func (n *NovaService) GetServers(tokenInstance *models.Token) (models.NovaResponse, error) {
	n.authToken = tokenInstance.HeaderToken
	fmt.Printf("TOKEN RETRIEVED: %+v \n ", n.authToken)
	n.projectID = tokenInstance.ProjectID


	fmt.Sprintf("Get server details for projectID: [%v]", n.projectID)
	// Construct URL with the project ID ad path parameter
	fullUrl := fmt.Sprintf("https://api.pub1.infomaniak.cloud/compute/v2.1/servers?project_id=%s", n.projectID) 

	// Set header
	headers := map[string]string{
		"X-AUTH-TOKEN": n.authToken,
	}
	var response models.NovaResponse

	response, _, err := client.MakeHTTPRequest(fullUrl, "GET", headers, nil, nil, response)
	// Handle the response
	fmt.Printf("Server Details: %+v", response)
	if err != nil {
		fmt.Errorf("Error while converting: %s", err)
	}

	// getServerByID(response, appLogger, headers)
	// getServerIP(response, appLogger, headers)
	// getMetadata(response, appLogger, headers )
	return response, err

}

/*
func (k NovaService) GetProjects() ([]Project, error) {
    // ... your implementation ...
}

func (k NovaService) GetUsers() ([]User, error) {
    // ... your implementation ...
}

// Do the same for the KeystoneService

// b.3 Nova GET server by projectID, returns 200
func sendAuthRequest(appLogger *logger.APIlogger, xAuthToken, projectID string) {
	appLogger.Infof("Get server details for projectID: [%v]", projectID)
	// Construct URL with the project ID ad path parameter
	fullUrl := fmt.Sprintf("https://pub1.infomaniak.cloud/compute/v2.1/servers?project_id=%s", projectID)
	// Set header
	headers := map[string]string{
		"X-AUTH-TOKEN": xAuthToken,
	}
	var response models.NovaResponse

	response, _, err := client.MakeHTTPRequest( fullUrl, "GET", headers, nil, nil, response)
	// Handle the response
	appLogger.Infof("Server Details: %+v", response)
	if err != nil {
		appLogger.Fatalf("Error while coverting: %s", err)

	}

	getServerByID(response, appLogger, headers)
	getServerIP(response, appLogger, headers)
	//getMetadata(response, appLogger, headers )

}

func getServerByID(response models.NovaResponse, appLogger *logger.APIlogger, headers map[string]string) {
	// Iterate over each server and make the API call
	for _, server := range response.Servers {

		apiEndpoint := fmt.Sprintf("https://pub1.infomaniak.cloud/compute/v2.1/servers/%s", server.ID)

		var serverResp models.NovaServerPayload
		// Make the API call to the endpoint using the server ID

		serverResp, _, err := client.MakeHTTPRequest( apiEndpoint, "GET", headers, nil, nil, serverResp)

		// Handle the response
		appLogger.Infof("Response: %+s", &serverResp)
		if err != nil {
			appLogger.Fatalf("Error while coverting: %s", err)

		}

	}
}

// Doenst return anthing so far
func getServerIP(response models.NovaResponse, appLogger *logger.APIlogger, headers map[string]string) {
	// Iterate over each server and make the API call
	for _, server := range response.Servers {

		apiEndpoint := fmt.Sprintf("https://pub1.infomaniak.cloud/compute/v2.1/servers/%s/ips", server.ID)

		var response models.Addresses
		// Make the API call to the endpoint using the server ID

		response, _, err := client.MakeHTTPRequest( apiEndpoint, "GET", headers, nil, nil, response)

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

		response, _, err := client.MakeHTTPRequest( apiEndpoint, "GET", headers, nil, nil, response)

		// Handle the response
		appLogger.Infof("Response: %+s", response)
		if err != nil {
			appLogger.Fatalf("Error while coverting: %s", err)

		}

	}
}*/
