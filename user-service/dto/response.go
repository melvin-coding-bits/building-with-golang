//Package dto will has the dto model defnition and its utlities.
//Clients use dto to interact with the service. Also the service sents out
//data in dto model. This abstracts the underlying data model from the client.
package dto

import "net/http"

//Response is the web request response struct to be sent to the client
type Response struct {
	//Data to be sent as the response
	Data interface{} `json:"data,omitempty"`
	//Ok will be true if the response denotes a success request
	Ok bool `json:"ok"`
	//Message for the client
	Message string `json:"message,omitempty"`
	//Code is the http status code
	Code int `json:"code"`
}

//Success will return a successfull response
func Success(data interface{}) Response {
	return Response{
		Data:    data,
		Ok:      true,
		Message: "",
		Code:    http.StatusOK,
	}
}

//Error will return a failed response
func Error(err error, code int) Response {
	return Response{
		Data:    nil,
		Ok:      false,
		Message: err.Error(),
		Code:    code,
	}
}
