package models

type Frame struct {
	ProjectID     string `json:"projectID"`
	ProjectName   string `json:"projectName"`
	PrincipalID   string `json:"principalID"`
	PrincipalName string `json:"principalName"`
}
