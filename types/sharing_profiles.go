package types

// GuacSharingProfile base guacamole sharingProfile info
type GuacSharingProfile struct {
	Name                        string                       `json:"name"`
	Identifier                  string                       `json:"identifier,omitempty"`
	PrimaryConnectionIdentifier string                       `json:"primaryConnectionIdentifier"`
	Attributes                  GuacSharingProfileAttributes `json:"attributes"`
	Parameters                  GuacSharingProfileParameters `json:"parameters"`
}

// GuacSharingProfileAttributes sharingProfile attributes
type GuacSharingProfileAttributes struct {
}

// GuacSharingProfileParameters sharingProfile parameters
type GuacSharingProfileParameters struct {
	ReadOnly string `json:"read-only"`
}
