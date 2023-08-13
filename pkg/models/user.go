package models

/*
type Domain struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Domain            Domain `json:"domain"`
	ID                string `json:"id"`
	Name              string `json:"name"`
	PasswordExpiresAt string `json:"password_expires_at"`
}

type AuthRequest struct {
	Auth struct {
		Identity struct {
			Methods               []string              `json:"methods"`
			Password              PasswordDetails       `json:"password"`
			ApplicationCredential ApplicationCredential `json:"application_credential"`
		} `json:"identity"`
	} `json:"auth"`
}

type PasswordDetails struct {
	User UserDetails `json:"user"`
}

type UserDetails struct {
	Name     string `json:"name"`
	Domain   Domain `json:"domain"`
	Password string `json:"password"`
}

type ApplicationCredential struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
}

type Response struct {
	Token   Token    `json:"token"`
	Servers []Server `json:"servers"`
}

type Token struct {
	Methods   []string `json:"methods"`
	User      User     `json:"user"`
	AuditIDs  []string `json:"audit_ids"`
	ExpiresAt string   `json:"expires_at"`
	IssuedAt  string   `json:"issued_at"`
}

type TokenResponse struct {
	Token struct {
		Project struct {
			ID string `json:"id"`
		} `json:"project"`
	} `json:"token"`
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

// Prepare auth payload for authentication.
func AuthPayload() AuthRequest {
	return AuthRequest{
		Auth: struct {
			Identity: struct {
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
}

// Prepare auth payload using application credentials.
func CredentialPayload() AuthRequest {
	return AuthRequest{
		Auth: struct {
			Identity: struct {
				Methods: []string{"application_credential"},
				ApplicationCredential: ApplicationCredential{
					ID:     "197c82bd809e413a86467caa678352a4",
					Secret: "xT1W-_pTjrLHbaCdstpL3dgXvgxmLWGggSXOu7x-elN8kmqRpam-agbW1KIwz3rNoa7zkUP101QCEE5t-Pq4Fw",
				},
			},
		},
	}
}
*/
