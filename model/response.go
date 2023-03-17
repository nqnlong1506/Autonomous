package model

type Response struct {
	Status   string      `json:"status,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Messsage string      `json:"message,omitempty"`
}
