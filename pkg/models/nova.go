package models



type NovaResponse struct {
    Servers []Server `json:"servers"`
}
type NovaServerPayload struct {
	Server Server `json:"server"`
}
 type Server struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Status            string            `json:"status"`
	TenantID          string            `json:"tenant_id"`
	UserID            string            `json:"user_id"`
	Metadata          ServerMetadata    `json:"metadata"`
	HostID            string            `json:"hostId"`
	Image             ServerImage       `json:"image"`
	Flavor            ServerFlavor      `json:"flavor"`
	Created           string            `json:"created"`
	Updated           string            `json:"updated"`
	Addresses         ServerAddresses   `json:"addresses"`
	AccessIPv4        string            `json:"accessIPv4"`
	AccessIPv6        string            `json:"accessIPv6"`
	Links             []ServerLink      `json:"links"`
	DiskConfig        string            `json:"OS-DCF:diskConfig"`
	Progress          int               `json:"progress"`
	AvailabilityZone  string            `json:"OS-EXT-AZ:availability_zone"`
	ConfigDrive       string            `json:"config_drive"`
	KeyName           string            `json:"key_name"`
	LaunchedAt        string            `json:"OS-SRV-USG:launched_at"`
	TerminatedAt      *string           `json:"OS-SRV-USG:terminated_at"`
	SecurityGroups    []ServerSecurityGroup `json:"security_groups"`
	TaskState         *string           `json:"OS-EXT-STS:task_state"`
	VMState           string            `json:"OS-EXT-STS:vm_state"`
	PowerState        int               `json:"OS-EXT-STS:power_state"`
	VolumesAttached   []string          `json:"os-extended-volumes:volumes_attached"`
}

type ServerMetadata struct {
	DependsOn       string `json:"depends_on"`
	KubesprayGroups string `json:"kubespray_groups"`
	SSHUser         string `json:"ssh_user"`
	UseAccessIP     string `json:"use_access_ip"`
}

type ServerImage struct {
	ID    string         `json:"id"`
	Links []ServerLink   `json:"links"`
}

type ServerFlavor struct {
	ID    string         `json:"id"`
	Links []ServerLink   `json:"links"`
}

type ServerAddresses struct {
	DemoCluster []ServerAddress `json:"demo-cluster"`
}

type ServerAddress struct {
	Version         int    `json:"version"`
	Address         string `json:"addr"`
	Type            string `json:"OS-EXT-IPS:type"`
	MACAddress      string `json:"OS-EXT-IPS-MAC:mac_addr"`
}

type ServerLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type ServerSecurityGroup struct {
	Name string `json:"name"`
}
 
type ServerIpAdress struct {
	Addresses Addresses `json:"addresses"`
}

type Addresses struct {
	Addresses map[string]interface{} `json:"addresses"`
}



type Metadata struct {
	Metadata map[string]interface{} `json:"metadata"`
}