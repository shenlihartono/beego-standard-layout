// Package service contains the response type returned by service.
package response

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func New(status int, body interface{}) Response {
	return Response{Status: status, Result: body}
}
