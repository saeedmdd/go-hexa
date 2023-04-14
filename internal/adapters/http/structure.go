package http

const apiVersion = "1.0.0"

type Response struct {
	Data       interface{} `json:"data"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	ApiVersion string      `json:"api_version"`
}

func newResponse(data interface{}, status string, message string) Response {
	return Response{
		Data:       data,
		Status:     status,
		Message:    message,
		ApiVersion: apiVersion,
	}
}
