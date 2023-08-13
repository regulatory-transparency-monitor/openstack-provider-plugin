package models

type AuthRequest struct {
	Auth AuthIdentity `json:"auth"`
}

type AuthIdentity struct {
	Identity IdentityData `json:"identity"`
}

type IdentityData struct {
	Methods               []string              `json:"methods"`
	Password              PasswordDetails       `json:"password"`
	ApplicationCredential ApplicationCredential `json:"application_credential"`
}

type ApplicationCredential struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
}

type PasswordDetails struct {
	User UserDetails `json:"user"`
}

type UserDetails struct {
	Name     string `json:"name"`
	Domain   Domain `json:"domain"`
	Password string `json:"password"`
}

// Token struct for the authentication request
type Response struct {
	Token   Token    `json:"token"`
	Servers []Server `json:"servers"`
}
// Token struct for the authentication request
type AuthToken struct {
	Token string
}


/* type Token struct {
	Methods   []string `json:"methods"`
	User      User     `json:"user"`
	AuditIDs  []string `json:"audit_ids"`
	ExpiresAt string   `json:"expires_at"`
	IssuedAt  string   `json:"issued_at"`
} */

type Token struct {
	HeaderToken string 
	ProjectID string
}
type User struct {
	Domain            Domain `json:"domain"`
	ID                string `json:"id"`
	Name              string `json:"name"`
	PasswordExpiresAt string `json:"password_expires_at"`
}

type Domain struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TokenResponse struct {
	Token struct {
		Project struct {
			ID string `json:"id"`
		} `json:"project"`
	} `json:"token"`
}

// DoAuth performs authentication
func AuthPayload() AuthRequest {
	authReq := AuthRequest{
		Auth: AuthIdentity{
			Identity: IdentityData{
				Methods: []string{"password"},
				Password: PasswordDetails{
					User: UserDetails{
						Name:     "PCU-JMY29J4",
						Domain:   Domain{Name: "Default"},
						Password: "rIhhym-xescud-2kyxwe",
					},
				},
			},
		},
	}
	return authReq
}

// performs authentication
func CredentialPayload() AuthRequest {
	authReq := AuthRequest{
		Auth: AuthIdentity{
			Identity: IdentityData{
				Methods: []string{"application_credential"},
				ApplicationCredential: ApplicationCredential{
					ID:     "d4ac29bc2a0449419754693baa1c9e3f",
					Secret: "J9zApG1iGomwo2463ImstALe3QA8A-Rwl_NrpmA1LsD-ua9Wgm95hOKQ_pp7EfJ3dYDFMQ-LlC5l4hIAMtEkLg",
				},
			},
		},
	}
	return authReq
}

type ProjectDetails struct {
	Project struct {
		ID          string                 `json:"id"`
		Name        string                 `json:"name"`
		DomainID    string                 `json:"domain_id"`
		Description string                 `json:"description"`
		Enabled     bool                   `json:"enabled"`
		ParentID    string                 `json:"parent_id"`
		IsDomain    bool                   `json:"is_domain"`
		Tags        []string               `json:"tags"`
		Options     map[string]interface{} `json:"options"`
		Links       struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"project"`
}
