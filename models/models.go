package models

//Response for request response
type Response struct {
	Status int         `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
