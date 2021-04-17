package src

type RequestResponse struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}
