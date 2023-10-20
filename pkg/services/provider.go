package services

import (
	"fmt"

	shared "github.com/regulatory-transparency-monitor/commons/models"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/api"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/httpwrapper"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/interfaces"
	m "github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/models"
)

// OpenStackPlugin is a struct holding Keystone and Nova service
type OpenStackPlugin struct {
	Keystone interfaces.KeystoneAPI
	Nova     interfaces.NovaAPI
	Cinder   interfaces.CinderAPI
	Config   map[string]interface{}
}
type APIContext struct {
	ProjectID string
}

func (provider *OpenStackPlugin) Initialize(config map[string]interface{}) error {
	provider.Config = config
	apiAccess, ok := config["api_access"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("openStack api_access configuration is missing or invalid")
	}
	baseURL, ok := apiAccess["base_url"].(string)
	if !ok || baseURL == "" {
		return fmt.Errorf("openStack base_url configuration is missing or invalid")
	}
	identityAPI, ok := apiAccess["identity_api"].(string)
	if !ok || identityAPI == "" {
		return fmt.Errorf("OpenStack identity_api configuration is missing or invalid")
	}
	compute_api, ok := apiAccess["compute_api"].(string)
	if !ok || identityAPI == "" {
		return fmt.Errorf("OpenStack compute_api configuration is missing or invalid")
	}
	storage_api, ok := apiAccess["storage_api"].(string)
	if !ok || identityAPI == "" {
		return fmt.Errorf("OpenStack storage_api configuration is missing or invalid")
	}
	httpClient := httpwrapper.NewClient(baseURL)

	provider.Keystone = &api.KeystoneService{
		BaseURL: identityAPI,
		Client:  httpClient,
	}
	provider.Nova = &api.NovaService{
		BaseURL: compute_api,
		Client:  httpClient,
	}
	provider.Cinder = &api.CinderService{
		BaseURL: storage_api,
		Client:  httpClient,
	}

	return nil

}

func (provider *OpenStackPlugin) FetchData() (shared.RawData, error) {
	data := make(shared.RawData)
	// Initialize the context for shared paramters
	ctx, err := provider.initContext()
	if err != nil {
		return data, err
	}

	keystoneData, err := provider.FetchKeystoneData(ctx)
	if err != nil {
		return data, err
	}

	novaData, err := provider.FetchNovaData(ctx)
	if err != nil {
		return data, err
	}

	cinderData, err := provider.FetchCinderData(ctx)
	if err != nil {
		return data, err
	}

	cinderSnapshots, err := provider.FetchCinderSnapshots(ctx)
	if err != nil {
		return data, err
	}

	//PrintAsJSON(cinderSnapshots)
	data["os_project"] = []interface{}{keystoneData}
	data["os_instance"] = novaData
	data["os_volume"] = cinderData
	data["os_snapshot"] = cinderSnapshots

	return data, nil
}
func (provider *OpenStackPlugin) FetchKeystoneData(ctx *APIContext) (*m.ProjectDetails, error) {
	credentials := provider.Config["credentials"].(map[string]interface{})

	projectID, err := provider.Keystone.Authenticate(credentials)
	if err != nil {
		fmt.Println(err)
	}

	projectModel, err := provider.Keystone.GetProjectDetailsByID(projectID)
	if err != nil {
		fmt.Println(err)
	}

	ctx.ProjectID = projectID

	return &projectModel, nil
}

func (provider *OpenStackPlugin) FetchNovaData(ctx *APIContext) ([]interface{}, error) {
	var serverDetailsList []interface{}

	projectID := ctx.ProjectID

	serverList, err := provider.Nova.GetServerListByProjectID(projectID)
	if err != nil {
		return serverDetailsList, err
	}

	for _, server := range serverList.Servers {
		serverDetails, err := provider.Nova.GetServerByID(server.ID)
		if err != nil {
			return serverDetailsList, err
		}
		serverDetailsList = append(serverDetailsList, serverDetails)
	}

	return serverDetailsList, nil
}

func (provider *OpenStackPlugin) FetchCinderData(ctx *APIContext) ([]interface{}, error) {
	var volumeDetailsList []interface{}
	// Use the project ID from the context
	projectID := ctx.ProjectID
	// List of hard-coded volume IDs
	volumeIDs := []string{
		"b5d527e9-0441-4db8-b2f1-b0de5f0ca4ea",
		"8227897a-61d2-408b-a22d-c98705fcb39b",
	}
	for _, volumeID := range volumeIDs {
		volumeDetails, err := provider.Cinder.GetVolumeByID(volumeID, projectID)

		if err != nil {
			fmt.Println(err)
		}
		volumeDetailsList = append(volumeDetailsList, volumeDetails)
	}

	return volumeDetailsList, nil
}
func (provider *OpenStackPlugin) FetchCinderSnapshots(ctx *APIContext) ([]interface{}, error) {
	var snapshotDetailsList []interface{}
	// Use the project ID from the context
	projectID := ctx.ProjectID

	snapshotList, err := provider.Cinder.GetSnapshots(projectID)
	if err != nil {
		return snapshotDetailsList, err
	}

	for _, snapshot := range snapshotList {
		snapshotDetailsList = append(snapshotDetailsList, snapshot)

	}

	return snapshotDetailsList, nil
}

func (provider *OpenStackPlugin) initContext() (*APIContext, error) {
	ctx := &APIContext{}
	return ctx, nil
}
