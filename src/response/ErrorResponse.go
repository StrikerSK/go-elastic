package response

type ResponseInterface interface {
	GetData() interface{}
}

type ErrorResponse struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

func NewErrorResponse(statusCode, errorInput string) *ErrorResponse {
	return &ErrorResponse{
		Error:  errorInput,
		Status: statusCode,
	}
}

func (receiver ErrorResponse) GetData() interface{} {
	return receiver.Error
}
