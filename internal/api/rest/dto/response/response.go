package response

type SuccessResponse[T any] struct {
	Data T `json:"data"`
}

type ErrorResponse struct {
	Error error `json:"error"`
}
