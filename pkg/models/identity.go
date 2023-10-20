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

type Token struct {
	HeaderToken string
	ProjectID   string
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
