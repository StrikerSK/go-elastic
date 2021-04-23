package src

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"go-elastic/src/response"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	todoIndexUrl    = HOST_URL + "/" + TODOS_INDEX
	todoDocumentUrl = todoIndexUrl + "/_doc"
	todoSearchUrl   = todoIndexUrl + "/_search"
)

func createData(object CustomInterface) response.RequestResponse {
	marshalledObject, _ := object.MarshalItem()
	payload := strings.NewReader(string(marshalledObject))

	client := &http.Client{}
	req, err := http.NewRequest("POST", todoDocumentUrl, payload)

	if err != nil {
		fmt.Println(err)
		return response.RequestResponse{
			Data:   err,
			Status: "Error",
			Code:   400,
		}
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return response.RequestResponse{
			Data:   err,
			Status: "Error",
			Code:   http.StatusInternalServerError,
		}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return response.RequestResponse{
			Data:   err,
			Status: "Error",
			Code:   http.StatusInternalServerError,
		}
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

	client := &http.Client{}
	req, err := http.NewRequest("GET", todoDocumentUrl+"/"+todoID, nil)

	if err != nil {
		fmt.Println(err)
		return response.RequestResponse{
			Data:   err,
			Status: "Error",
			Code:   http.StatusInternalServerError,
		}
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return response.RequestResponse{
			Data:   err,
			Status: "Request error",
			Code:   http.StatusInternalServerError,
		}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return response.RequestResponse{
			Data:   err,
			Status: "Response body error",
			Code:   http.StatusInternalServerError,
		}
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

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", todoDocumentUrl+"/"+todoID, nil)

	if err != nil {
		fmt.Println(err)
		return response.RequestResponse{
			Data:   err,
			Status: "Error",
			Code:   http.StatusInternalServerError,
		}
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return response.RequestResponse{
			Data:   err,
			Status: "Request error",
			Code:   http.StatusInternalServerError,
		}
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return response.RequestResponse{
			Data:   err,
			Status: "Response body error",
			Code:   http.StatusInternalServerError,
		}
	}

	if res.StatusCode == http.StatusNotFound {
		log.Printf("Todo not found")
	}

	return response.RequestResponse{
		Data:   nil,
		Status: "Todo deleted",
		Code:   http.StatusOK,
	}
}
