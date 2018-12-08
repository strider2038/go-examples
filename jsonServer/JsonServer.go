package jsonServer

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"log"
	"net/http"
)

type RequestBody struct {
	Protocol string                 `json:"protocol"`
	Id       int                    `json:"id"`
	Message  map[string]interface{} `json:"message"`
}

type EnvelopeHeaders struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

type EnvelopeMessage struct {
	Headers EnvelopeHeaders `json:"headers"`
	Body    map[string]interface{}
}

type LogMessage struct {
	Body string `json:"body"`
}

type ResponseBody struct {
	Log []string `json:"log"`
}

func EndpointHandler(writer http.ResponseWriter, request *http.Request) {
	var requests []RequestBody
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&requests)

	if err != nil {
		panic(err)
	}

	response := processRequests(requests)

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(response)
	if err != nil {
		panic(err)
	}
}

func processRequests(requests []RequestBody) ResponseBody {
	response := ResponseBody{}

	for _, requestBody := range requests {
		logMessage := fmt.Sprintf("Message %d with protocol \"%s\" processed.", requestBody.Id, requestBody.Protocol)
		response.Log = append(response.Log, logMessage)

		if requestBody.Protocol == "envelope" {
			processEnvelopeMessage(requestBody, &response)
		} else if requestBody.Protocol == "log" {
			processLogMessage(requestBody, &response)
		}
	}

	return response
}

func processEnvelopeMessage(requestBody RequestBody, response *ResponseBody) {
	var message EnvelopeMessage

	err := mapstructure.Decode(requestBody.Message, &message)
	if err != nil {
		panic(err)
	}

	logMessage := fmt.Sprintf(
		"Routing envelope message from sender \"%s\" to receiver \"%s\".",
		message.Headers.Sender,
		message.Headers.Receiver,
	)

	response.Log = append(response.Log, logMessage)
}

func processLogMessage(requestBody RequestBody, response *ResponseBody) {
	var message LogMessage

	err := mapstructure.Decode(requestBody.Message, &message)
	if err != nil {
		panic(err)
	}

	logMessage := fmt.Sprintf("Logging message \"%s\".", message.Body)
	response.Log = append(response.Log, logMessage)
}

func RunJsonServer() {
	router := mux.NewRouter()
	// Routes consist of a path and a handler function.
	router.HandleFunc("/endpoint", EndpointHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":1234", router))
}
