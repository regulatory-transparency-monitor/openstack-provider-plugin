package models

type VolumeResponse struct {
	Volume Volume `json:"volume"`
}
type Volume struct {
	Attachments        []Attachment           `json:"attachments"`
	AvailabilityZone   string                 `json:"availability_zone"`
	Bootable           string                 `json:"bootable"`
	ConsistencyGroupID *string                `json:"consistencygroup_id"`
	CreatedAt          string                 `json:"created_at"`
	Description        string                 `json:"description"`
	Encrypted          bool                   `json:"encrypted"`
	ID                 string                 `json:"id"`
	Links              []Link                 `json:"links"`
	Metadata           map[string]interface{} `json:"metadata"`
	Multiattach        bool                   `json:"multiattach"`
	Name               string                 `json:"name"`
	TenantID           string                 `json:"os-vol-tenant-attr:tenant_id"`
	ReplicationStatus  *string                `json:"replication_status"`
	Size               int                    `json:"size"`
	SnapshotID         string                 `json:"snapshot_id"`
	SourceVolid        *string                `json:"source_volid"`
	Status             string                 `json:"status"`
	UpdatedAt          string                 `json:"updated_at"`
	UserID             string                 `json:"user_id"`
	VolumeType         string                 `json:"volume_type"`
}

type Attachment struct {
	AttachedAt   string  `json:"attached_at"`
	AttachmentID string  `json:"attachment_id"`
	Device       string  `json:"device"`
	HostName     *string `json:"host_name"`
	ID           string  `json:"id"`
	ServerID     string  `json:"server_id"`
	VolumeID     string  `json:"volume_id"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type ServerVolumeAttachments struct {
	VolumeAttachments []VolumeAttachment `json:"volumeAttachments"`
}

type VolumeAttachment struct {
	ID       string `json:"id"`
	VolumeID string `json:"volumeId"`
	ServerID string `json:"serverId"`
	Device   string `json:"device"`
}

type Snapshot struct {
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Metadata    struct {
		Key string `json:"key"`
	} `json:"metadata"`
	Name                                  string  `json:"name"`
	OSExtendedSnapshotAttributesProgress  string  `json:"os-extended-snapshot-attributes:progress"`
	OSExtendedSnapshotAttributesProjectID string  `json:"os-extended-snapshot-attributes:project_id"`
	Size                                  int     `json:"size"`
	Status                                string  `json:"status"`
	UpdatedAt                             string  `json:"updated_at"` // Using pointer for nullable time
	VolumeID                              string  `json:"volume_id"`
	GroupSnapshotID                       *string `json:"group_snapshot_id"` // Using pointer for nullable string
	UserID                                string  `json:"user_id"`
	ConsumesQuota                         bool    `json:"consumes_quota"`
}

type SnapshotsResponse struct {
	Snapshots []Snapshot `json:"snapshots"`
}
