package models

type Message struct{
	Code int `json:"code"`
	Message Message
}