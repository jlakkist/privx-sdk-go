//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package authorizer

// Params query params definition
type Params struct {
	ResponseType  string `json:"response_type,omitempty"`
	ClientID      string `json:"client_id,omitempty"`
	State         string `json:"state,omitempty"`
	RedirectURI   string `json:"redirect_uri,omitempty"`
	UserAgent     string `json:"user_agent,omitempty"`
	OidcID        string `json:"oidc_id,omitempty"`
	AccessGroupID string `json:"access_group_id,omitempty"`
	KeyID         string `json:"key_id,omitempty"`
	Filter        string `json:"filter,omitempty"`
	Service       string `json:"service,omitempty"`
	Sortkey       string `json:"sortkey,omitempty"`
	Sortdir       string `json:"sortdir,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty"`
}

// ClientTime client time definition
type ClientTime struct {
	ClientTime string `json:"client_time,omitempty"`
}

// ConfigurationError configuration error definition
type ConfigurationError struct {
	Errors []string `json:"errors,omitempty"`
}

// Session session definition
type Session struct {
	SessionID string `json:"session_id"`
}

// Signature signature  definition
type Signature struct {
	Signature string `json:"signature"`
}

// Credential end user authentication credentials definition
type Credential struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// PrincipalKeyImportRequest principal key import definition
type PrincipalKeyImportRequest struct {
	Algorithm string `json:"algorithm"`
	Data      string `json:"data"`
}

// AuthorizationRequest end user authorization request definition
type AuthorizationRequest struct {
	PublicKey string `json:"public_key,omitempty"`
	HostID    string `json:"host_id,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	Username  string `json:"username,omitempty"`
	Service   string `json:"service,omitempty"`
	RoleID    string `json:"role_id,omitempty"`
}

// Principal principal definition
type Principal struct {
	ID              string `json:"id"`
	GroupID         string `json:"group_id,omitempty"`
	Type            string `json:"type,omitempty"`
	Comment         string `json:"comment,omitempty"`
	PublicKey       string `json:"public_key,omitempty"`
	PublicKeyString string `json:"public_key_string,omitempty"`
	Size            int    `json:"size,omitempty"`
}

// CA is root certificate representation
type CA struct {
	ID        string `json:"id"`
	GroupID   string `json:"group_id"`
	Type      string `json:"type"`
	Size      int    `json:"size"`
	PublicKey string `json:"public_key"`
	X509      string `json:"x509_certificate"`
}

// AccessGroup access group definition
type AccessGroup struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Comment   string `json:"comment,omitempty"`
	CAID      string `json:"ca_id,omitempty"`
	Author    string `json:"author,omitempty"`
	Created   string `json:"created,omitempty"`
	Updated   string `json:"updated,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	Default   bool   `json:"default,omitempty"`
}
