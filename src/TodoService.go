package src

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"go-elastic/src/response"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	todoIndexUrl    = HostUrl + "/" + TodosIndex
	todoDocumentUrl = todoIndexUrl + "/_doc"
	todoSearchUrl   = todoIndexUrl + "/_search"
)

func createData(object CustomInterface) response.RequestResponse {
	marshalledObject, _ := object.MarshalItem()
	payload := strings.NewReader(string(marshalledObject))

	res, body, err := sendRequest(todoDocumentUrl, "POST", payload)
	if err != nil {
		log.Printf("createData() read failed: %s\n", err)
	}

	if res.StatusCode == http.StatusCreated {
		log.Print("Custom object created successfully!")
		value := gjson.Get(string(body), "_id")
		return response.RequestResponse{
			Data:   value.String(),
			Status: "Data created",
			Code:   http.StatusCreated,
		}
	} else {
		value := gjson.Get(string(body), "error.reason")
		fmt.Printf("Retrieved error message: %s", value.String())
		return response.RequestResponse{
			Data:   body,
			Status: res.Status,
			Code:   res.StatusCode,
		}
	}
}

func getTodo(todoID string) response.RequestResponse {

	res, body, err := sendRequest(todoDocumentUrl+"/"+todoID, "GET", nil)
	if err != nil {
		log.Printf("getTodo() [%s] read failed: %s\n", todoID, err)
	}

	switch res.StatusCode {
	case http.StatusOK:
		m := make(map[string]interface{})
		_ = json.Unmarshal(body, &m)
		return response.RequestResponse{
			Data:   m["_source"].(map[string]interface{}),
			Status: "Todo retrieved",
			Code:   http.StatusOK,
		}
	case http.StatusNotFound:
		log.Print("Todo not found")
		return response.RequestResponse{
			Data:   "Cannot find requested todo " + todoID,
			Status: "Todo not found",
			Code:   http.StatusNotFound,
		}
	//case http.StatusBadRequest:
	//	value := gjson.Get(string(body), "error.reason")
	//	log.Println("Something went wrong at getTodo()")
	//	fmt.Printf("Retrieved id: %s\n", value.String())
	//	return response.RequestResponse{
	//		Data:   "Something went wrong",
	//		Status: "Error",
	//		Code:   http.StatusBadRequest,
	//	}
	default:
		log.Printf("Unexpected error occurred \n")
		return response.RequestResponse{
			Data:   "Unexpected error occurred",
			Status: "Error",
			Code:   http.StatusBadRequest,
		}
	}
}

func deleteTodo(todoID string) response.RequestResponse {

	res, _, err := sendRequest(todoDocumentUrl+"/"+todoID, "DELETE", nil)
	if err != nil {
		log.Printf("deleteTodo() [%s] failed: %s\n", todoID, err)
	}

	if res.StatusCode == http.StatusNotFound {
		log.Printf("Todo [%s] not found\n", todoID)
	}

	return response.RequestResponse{
		Data:   nil,
		Status: "Todo deleted",
		Code:   http.StatusOK,
	}
}

func sendRequest(requestUrl, requestMethod string, requestBody io.Reader) (*http.Response, []byte, error) {
	client := &http.Client{}
	newRequest, err := http.NewRequest(requestMethod, requestUrl, requestBody)

	if err != nil {
		return nil, nil, err
	}

	newRequest.Header.Add("Content-Type", "application/json")

	serverResponse, err := client.Do(newRequest)
	if err != nil {
		return nil, nil, err
	}

	responseBody, err := ioutil.ReadAll(serverResponse.Body)
	if err != nil {
		return nil, nil, err
	}

	return serverResponse, responseBody, err
}
