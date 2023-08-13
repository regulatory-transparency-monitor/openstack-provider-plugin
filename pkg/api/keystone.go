// keystone.go (Keystone Provider Plugin)
package api

import (
	"bytes"
	"collector-service/pkg/client"
	"collector-service/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type KeystoneService struct {
	Token          string
	projectID      string
	expirationTime string
	ctx            context.Context
}

/* func GetKeystoneServiceInstance() *KeystoneService {
	once.Do(func() {
		instance = &KeystoneService{}
		// Add initialization logic if necessary
	})
	return instance
} */

func (k *KeystoneService) GetToken(ctx context.Context) (string, error) {
	// Check if the token is expired. If it is, refresh it.
	if k.isTokenExpired() {
		fmt.Printf("Token Expired, refreshing token...%s\n", k.expirationTime)
		if err := k.refreshToken(); err != nil {
			return "", err
		}
	}
	return k.Token, nil
}

func (k *KeystoneService) isTokenExpired() bool {
	t, err := time.Parse(time.RFC3339, k.expirationTime)

	if err != nil {
		fmt.Println("Error parsing expiry time:", err)
		return false // or return true to signal that the token is expired if there's an error
	}
	// Check is token expired, true when t <= current time
	return time.Now().After(t)
}

func (k *KeystoneService) refreshToken() error {
	// Refresh the token

	//ctx := k.ctx
	return nil
}

// Get AuthToken using applciation credentials
func (k *KeystoneService) Authenticate(ctx context.Context) (*models.Token, error) {
	// set url
	fullUrl := "https://api.pub1.infomaniak.cloud/identity/v3/auth/tokens"
	// request payload containing application credentials
	// TODO pass from a configuration file, transparency UI instead of hardcoding
	payload := models.CredentialPayload()

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	queryParameters := url.Values{}
	reqBody, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	body := bytes.NewReader(reqBody)
	var rsp map[string]interface{}

	// http request
	response, responseHeaders, err := client.MakeHTTPRequest(fullUrl, "POST", headers, queryParameters, body, rsp)
	//fmt.Printf("Payload converted, JSON: %s\n", rsp)
	if err != nil {
		panic(err)
	}

	// Retrieve project ID from the response map
	k.projectID = response["token"].(map[string]interface{})["project"].(map[string]interface{})["id"].(string)
	// fmt.Printf("Project ID: %s\n", k.projectID)

	// Get expiration time from the response map
	k.expirationTime = response["token"].(map[string]interface{})["expires_at"].(string)
	// fmt.Printf("Expiration time: %s \n", k.expirationTime)

	// Create an instance of the Token struct
	token := &models.Token{
		
	}

	// Retrieve project ID from the response map
	token.ProjectID = response["token"].(map[string]interface{})["project"].(map[string]interface{})["id"].(string)

	// Get the X-Subject-Token from the response headers
	token.HeaderToken = responseHeaders.Get("X-Subject-Token")

	fmt.Printf("X-Subject-Token header: %s\n", token.HeaderToken)

	return token, nil

}

// GetProject by ID, returns project Details
func (k *KeystoneService) GetProjects() (models.ProjectDetails, error) {
	url := fmt.Sprintf("https://api.pub1.infomaniak.cloud/identity/v3/projects/%s", k.projectID)
	// Set header
	headers := map[string]string{
		"X-AUTH-TOKEN": k.Token,
	}

	var response models.ProjectDetails
	// Make the GET request
	response, _, err := client.MakeHTTPRequest(url, "GET", headers, nil, nil, response)
	if err != nil {
		panic(err)
	}
	return response, err
}

// Authenticate using application credentials
func (k *KeystoneService) IdentityToken() error {
	// General authentication when the
	fullUrl := "https://api.pub1.infomaniak.cloud/identity/v3/auth/tokens"

	payload := models.CredentialPayload()

	// Marshal payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	projectID := result["token"].(map[string]interface{})["project"].(map[string]interface{})["id"].(string)
	fmt.Printf("Project ID: %s", projectID)

	fmt.Printf("Body: %s\n", body)
	// Now let's get the X-Subject-Token from the response headers
	token := resp.Header.Get("X-Subject-Token")

	fmt.Printf("X-Subject-Token header: %s\n", token)

	k.Token = token
	return err
}

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
	fullUrl := "https://pub1.infomaniak.cloud/identity/v3/auth/tokens"
	// the headers to pass - none required for this test
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	queryParameters := url.Values{}
	authReq := models.CredentialPayload()
	// Convert the request payload to JSON (like json_encode)
	reqBody, err := json.Marshal(authReq)
	if err != nil {
		appLogger.Fatalf("Error while coverting: %s", err)
	} else {
		appLogger.Infof("Payload converted, JSON: %s", reqBody)
	}
	body := bytes.NewReader(reqBody)

	var rsp map[string]interface{}
	// http request
	response, responseHeaders, err := MakeHTTPRequest(fullUrl, "POST", headers, queryParameters, body, rsp)
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
