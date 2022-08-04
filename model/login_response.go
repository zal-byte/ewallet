package model

type LoginResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
	Token  interface{} `json:"token"`
}
