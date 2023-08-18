package main

import (
	"fmt"
	"log"
	"os"

	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/config"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/api"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/logger"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/services"
)

func main() {

	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)
	dirPath := "../../pkg/logger/"
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, os.ModePerm) // Creates the directory along with any necessary parents
	}
	// create a new config
	cfg, err := Init()

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	// create logger
	appLogger := logger.NewAPIlogger(cfg)
	appLogger.InitLogger()

	// Create a new KeystoneService
	keystoneService := &api.KeystoneService{}
	novaService := &api.NovaService{}

	// Create a new Provider
	t := &services.OpenStackPlugin{Keystone: keystoneService, Nova: novaService}
	t.Initialize()
	_, err = t.Scan()
	if err != nil {
		appLogger.Errorf("Scan Error: %s", err)
	}

}

func Init() (*config.Config, error) {
	// Get config path
	cfgPath := getConfigPath(os.Getenv("../config"))

	log.Printf("cfgPath: %s", cfgPath)
	// Load config file
	cfgFile, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	// Parse config file
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	return cfg, err
}

/*
// a.1 Keystone retrieve unscoped Token from X-Subject Header,returns status code 201
func getKeystoneUnscopedToken(appLogger *logger.APIlogger) {
	// the url to call
	fullUrl := "https://pub1.infomaniak.cloud/identity/v3/auth/tokens"

	// the headers to pass - none required for this test
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	// the query parameters to pass - none required but example commented out here
	//queryParameters.Has("false")
	queryParameters := url.Values{}
	// Initialize auth payload function.

	authReq := models.AuthPayload()
	// the type to unmarshal the response into
	// Convert the request payload to JSON (like json_encode)
	reqBody, err := json.Marshal(authReq)
	if err != nil {
		appLogger.Fatalf("Error while coverting: %s", err)
	} else {
		appLogger.Infof("Payload converted, JSON: %s", reqBody)
	}

	body := bytes.NewReader(reqBody)

	var rsp map[string]interface{}
	// call the function
	response, responseHeaders, err := client.MakeHTTPRequest(fullUrl, "POST", headers, queryParameters, body, rsp)

	if err != nil {
		panic(err)
	}
	// do something with the response
	appLogger.Infof("response: %+v", response)
	appLogger.Infof("Header response: %+v", responseHeaders)
} */

/* // b.1 Keystone retrieve application credentials, returns status code 201
func applicationCredentialToken(appLogger *logger.APIlogger) {
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
	response, responseHeaders, err := client.MakeHTTPRequest(fullUrl, "POST", headers, queryParameters, body, rsp)
	if err != nil {
		panic(err)
	}
	// Retrieve project ID from the response map
	projectID := response["token"].(map[string]interface{})["project"].(map[string]interface{})["id"].(string)
	appLogger.Infof("Project ID: %s", projectID)
	xSubjectToken := responseHeaders.Get("X-Subject-Token")
	appLogger.Infof("X-Subject-Token header: %s", xSubjectToken)
	getProjectDetailsByID(appLogger, xSubjectToken, projectID)
	//appLogger.Infof("Response headers: %s", responseHeaders)
	//headersValue := responseHeaders.Get("Content-Type")
	//appLogger.Infof("Content-Type header: %s", headersValue)
	//appLogger.Infof("Application Credential response: %+v", response)

}

// b.2 get project detials by project Id, returns status code 200
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

// Get config path for local or docker
func getConfigPath(configPath string) string {
	if configPath == "docker" {
		return "../config/config-docker"
	}
	return "../config/config-local"
}