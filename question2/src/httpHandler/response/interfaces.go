package response

type httpResponse struct {
	Message string      `json:"message,omitempty"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}
