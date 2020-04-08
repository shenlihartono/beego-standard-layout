// Package service contains the response type returned by service.
package response

type Response map[string]interface{}

func Success(body interface{}) Response {
	return Response{"result": body}
}

func Error(msg string) Response {
	return Response{"error": msg}
}
