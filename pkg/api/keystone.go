package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/httpwrapper"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/models"
)

type KeystoneService struct {
	BaseURL string
	Client  *httpwrapper.HTTPClient
}

func (k *KeystoneService) Authenticate(credential map[string]interface{}) (string, error) {
	endpoint := fmt.Sprintf("%sauth/tokens", k.BaseURL)
	// Prepare the payload
	credentials := CredentialPayload(credential)
	requestBody, err := json.Marshal(credentials)
	if err != nil {
		return "", fmt.Errorf("error marshaling credentials: %w", err)
	}

	// Create the HTTP request
	request, err := k.Client.NewRequest("POST", endpoint, nil, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error marshaling credentials: %w", err)
	}

	// Send the request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("error marshaling credentials: %w", err)
	}

	defer response.Body.Close()

	// Decode the response body
	var decodedResponse map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&decodedResponse); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	// Extract the project ID from the decoded response
	projectID, valid := decodedResponse["token"].(map[string]interface{})["project"].(map[string]interface{})["id"].(string)
	if !valid {
		return "", fmt.Errorf("could not extract project ID from response: %w", err)
	}

	k.Client.SetToken(response.Header.Get("X-Subject-Token"))

	return projectID, err
}

// GetProject by ID, returns project Details
func (k *KeystoneService) GetProjectDetailsByID(projectID string) (models.ProjectDetails, error) {
	endpoint := fmt.Sprintf("%sprojects/%v", k.BaseURL, projectID)

	// Construct the GET request.
	req, err := k.Client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return models.ProjectDetails{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Execute the request.
	res, err := k.Client.Do(req)
	if err != nil {
		return models.ProjectDetails{}, fmt.Errorf("execute request failed: %w", err)
	}
	defer res.Body.Close()

	// Read and decode the response.
	var projectDetails models.ProjectDetails
	if err := json.NewDecoder(res.Body).Decode(&projectDetails); err != nil {
		return models.ProjectDetails{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return projectDetails, nil
}

func CredentialPayload(credentials map[string]interface{}) models.AuthRequest {
    appCredentialsID, _ := credentials["app_credentials_id"].(string)
    appCredentialsSecret, _ := credentials["app_credentials_secret"].(string)

    authReq := models.AuthRequest{
        Auth: models.AuthIdentity{
            Identity: models.IdentityData{
                Methods: []string{"application_credential"},
                ApplicationCredential: models.ApplicationCredential{
                    ID:     appCredentialsID,
                    Secret: appCredentialsSecret,
                },
            },
        },
    }
    return authReq
}