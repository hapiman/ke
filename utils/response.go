package utils

type RespError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}
type RespSucc struct {
	Error   interface{} `json:"error"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type RespFail struct {
	Error   RespError   `json:"error"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
