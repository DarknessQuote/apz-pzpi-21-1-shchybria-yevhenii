package network

type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}