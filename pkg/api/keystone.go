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

func (k *KeystoneService) Authenticate() (string, error) {
	endpoint := fmt.Sprintf("%sauth/tokens", k.BaseURL)
	// Prepare the payload
	credentials := models.CredentialPayload()
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

/* // Implements the KeystoneRepository interface
func (k *KeystoneService) Authenticate() error {
	// set url
	endpoint := "https://api.pub1.infomaniak.cloud/identity/v3/auth/tokens"
	// TODO pass as paramter to function
	payload := models.CredentialPayload()

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	queryParameters := url.Values{}
	requestBody, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	body := bytes.NewReader(requestBody)
	var rsp map[string]interface{}

	// http request
	response, responseHeaders, err := client.MakeHTTPRequest(endpoint, "POST", headers, queryParameters, body, rsp)
	//fmt.Printf("Payload converted, JSON: %s\n", rsp)
	if err != nil {
		panic(err)
	}

	token := responseHeaders.Get("X-Subject-Token")
	//k.Client.SetToken(token)

	return err
} */

/* func (k *KeystoneService) GetToken() (string, error) {
	// Check if the token is expired. If it is, refresh it.
	if k.isTokenExpired() {
		fmt.Printf("Token Expired, refreshing token...%s\n", k.expirationTime)
		if err := k.refreshToken(); err != nil {
			return "", err
		}
	}
	return k.Token, nil
}
n time.Now
func (k *KeystoneService) isTokenExpired() bool {
	t, err := time.Parse(time.RFC3339, k.expirationTime)

	if err != nil {
		fmt.Println("Error parsing expiry time:", err)
		return false // or return true to signal that the token is expired if there's an error
	}
	// Check is token expired, true when t <= current time
	retur().After(t)
}

func (k *KeystoneService) refreshToken() error {
	// Refresh the token
	//ctx := k.ctx
	return nil
}
*/

/* // b.2 get project detials by project Id, returns status code 200
func getProjectDetailsByID(appLogger *logger.APIlogger, xAuthToken, projectID string) {
	appLogger.Infof("Get project details for projectID: [%v]", projectID)
	// Construct URL with the project ID ad path parameter
	url := fmt.Sprintf("https://pub1.infomaniak.cloud/identity/v3/projects/%s", projectID)
	// Set header
	headers := map[string]string{
		"X-AUTH-TOKEN": xAuthToken,
	}

	var response models.ProjectDetails
	// Make the GET request
	response, _, err := client.MakeHTTPRequest( url, "GET", headers, nil, nil, response)
	// Handle the response
	appLogger.Infof("Project Details: %+v", response)
	//appLogger.Infof("Response headers: %s", responseHeaders)
	if err != nil {
		appLogger.Fatalf("Error while coverting: %s", err)

	}
	sendAuthRequest(appLogger, xAuthToken, projectID)
} */

/*
 // b.1 Keystone retrieve application credentials, returns status code 201
func (k KeystoneService) AuthenticateToken() (Token, error) {
	endpoint := "https://pub1.infomaniak.cloud/identity/v3/auth/tokens"
	// the headers to pass - none required for this test
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	queryParameters := url.Values{}
	authReq := models.CredentialPayload()
	// Convert the request payload to JSON (like json_encode)
	requestBody, err := json.Marshal(authReq)
	if err != nil {
		appLogger.Fatalf("Error while coverting: %s", err)
	} else {
		appLogger.Infof("Payload converted, JSON: %s", requestBody)
	}
	body := bytes.NewReader(requestBody)

	var rsp map[string]interface{}
	// http request
	response, responseHeaders, err := MakeHTTPRequest(endpoint, "POST", headers, queryParameters, body, rsp)
	if err != nil {
		panic(err)
	}
	// Retrieve project ID from the response map
	projectID := response["token"].(map[string]interface{})["project"].(map[string]interface{})["id"].(string)
	appLogger.Infof("Project ID: %s", projectID)
	xSubjectToken := responseHeaders.Get("X-Subject-Token")
	appLogger.Infof("X-Subject-Token header: %s", xSubjectToken)
	//getProjectDetailsByID(appLogger, xSubjectToken, projectID)
	//appLogger.Infof("Response headers: %s", responseHeaders)
	//headersValue := responseHeaders.Get("Content-Type")
	//appLogger.Infof("Content-Type header: %s", headersValue)
	//appLogger.Infof("Application Credential response: %+v", response)
	return Token{}, nil
}  */

/* func GetKeystoneServiceInstance() *KeystoneService {
	once.Do(func() {
		instance = &KeystoneService{}
		// Add initialization logic if necessary
	})
	return instance
}  */
