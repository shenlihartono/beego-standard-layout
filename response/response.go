// Package service contains the response type returned by service.
package response

type Response map[string]interface{}

func New(status int, body interface{}) Response {
	return Response{"status": status, "result": body}
}
