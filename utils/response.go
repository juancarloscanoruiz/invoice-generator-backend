package utils

type APIResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  *string     `json:"error"`
}
