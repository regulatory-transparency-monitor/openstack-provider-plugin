package api

import (
	"encoding/json"
	"fmt"

	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/httpwrapper"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/models"
)

type CinderService struct {
	BaseURL string
	Client  *httpwrapper.HTTPClient
}

// Shows details for a volume.
func (c *CinderService) GetVolumeByID(volumeID string, projectID string) (models.Volume, error) {
	endpoint := fmt.Sprintf("%s%s/volumes/%s", c.BaseURL, projectID, volumeID)

	// Construct the GET request.
	req, err := c.Client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return models.Volume{}, fmt.Errorf("failed to construct GetVolumeByID request: %w", err)
	}

	// Execute the request.
	res, err := c.Client.Do(req)
	if err != nil {
		return models.Volume{}, fmt.Errorf("failed to execute GetVolumeByID request: %w", err)
	}
	defer res.Body.Close()

	/* bodyBytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Raw HTTP Response:", string(bodyBytes))
	res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) */

	var responseWrapper models.VolumeResponse
	if err := json.NewDecoder(res.Body).Decode(&responseWrapper); err != nil {
		return models.Volume{}, fmt.Errorf("failed to decode GetVolumeByID response: %w", err)
	}

	response := responseWrapper.Volume
	return response, err
}

func (c *CinderService) GetSnapshots(projectID string) ([]models.Snapshot, error) {
    endpoint := fmt.Sprintf("%s%s/snapshots/detail", c.BaseURL, projectID)

    // Construct the GET request.
    req, err := c.Client.NewRequest("GET", endpoint, nil, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to construct GetSnapshots request: %w", err)
    }

    // Execute the request.
    res, err := c.Client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to execute GetSnapshots request: %w", err)
    }
    defer res.Body.Close()

    var responseWrapper models.SnapshotsResponse
    if err := json.NewDecoder(res.Body).Decode(&responseWrapper); err != nil {
        return nil, fmt.Errorf("failed to decode GetSnapshots response: %w", err)
    }

    return responseWrapper.Snapshots, nil
}

func (c *CinderService) GetVolumesByProjectID(projectID string) ([]interface{}, error) {
	endpoint := "compute/v2.1/servers/detail"

	// Construct the GET request.
	req, err := c.Client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return []interface{}{}, fmt.Errorf("failed to construct GetVolumesByProjectID request: %w", err)
	}

	// Execute the request.
	res, err := c.Client.Do(req)
	if err != nil {
		return []interface{}{}, fmt.Errorf("failed to execute GetVolumesByProjectID request: %w", err)
	}
	defer res.Body.Close()

	/* body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read the GetAttachedVolumes response body: %v", err)
	} */

	//fmt.Println("Response:", string(body))
	// Decode the response.
	//var response models.ServerDetails

	/* if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return []interface{}{},, fmt.Errorf("failed to decode GetServerByID response: %w", err)
	}
	fmt.Println("Response:", response) */

	return []interface{}{}, err

}
