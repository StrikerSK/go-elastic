package response

type RequestResponse struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
	Code   int         `json:"code"`
}

func NewRequestResponse(status string, statusCode int, inputData interface{}) *RequestResponse {
	return &RequestResponse{
		Data:   inputData,
		Status: status,
		Code:   statusCode,
	}
}

func (receiver RequestResponse) GetData() interface{} {
	return receiver.Data
}
