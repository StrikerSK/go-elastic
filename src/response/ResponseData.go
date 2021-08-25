package response

type RequestResponse struct {
	Data       interface{}
	StatusCode int
	Headers    map[string]string
}

func NewRequestResponse(statusCode int, inputData interface{}) *RequestResponse {
	return &RequestResponse{
		Data:       inputData,
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func (rr RequestResponse) GetData() interface{} {
	return map[string]interface{}{
		"data": rr.Data,
	}
}

func (rr RequestResponse) GetStatusCode() int {
	return rr.StatusCode
}

func (rr RequestResponse) GetHeaders() map[string]string {
	return rr.Headers
}
