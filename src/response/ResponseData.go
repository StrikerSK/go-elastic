package response

type ResponseInterface interface {
	GetData() interface{}
}

type RequestResponse struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
	Code   int         `json:"code"`
}

func (receiver RequestResponse) GetData() interface{} {
	return receiver.Data
}

type ErrorResponse struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

func (receiver ErrorResponse) GetData() interface{} {
	return receiver.Error
}
