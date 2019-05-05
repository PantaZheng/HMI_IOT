package models

type Frame struct{
	ProjectID string `json:"project_id"`
	ProjectName string `json:"project_name"`
	PrincipalID	string	   	`json:"principal_id"`
	PrincipalName string	`json:"principal_name"`
}

