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
	if rr.Data == nil {
		return nil
	} else {
		dataMap := make(map[string]interface{}, 1)

		switch rr.Data.(type) {
		case error:
			dataMap["error"] = rr.Data.(error)
		default:
			dataMap["data"] = rr.Data
		}

		return dataMap
	}
}

func (rr RequestResponse) GetStatusCode() int {
	return rr.StatusCode
}

func (rr RequestResponse) GetHeaders() map[string]string {
	return rr.Headers
}
