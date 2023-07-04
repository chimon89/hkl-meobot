package errors

type ErrorResponse struct {
	Ok           bool   `json:"ok"`
	ErrorMessage string `json:"error_message"`
}

type ErrorResponseDev struct {
	Ok           bool        `json:"ok"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}
