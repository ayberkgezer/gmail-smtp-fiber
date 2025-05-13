package base

type Response struct {
	Code    int    `json:"statusCode"`
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
}

type ResponseWithData struct {
	Response
	Data any `json:"data"`
}

type ErrorResponse struct {
	Response
	ErrorMessage string `json:"errorMessage"`
}

func NewBaseResponse(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Success: true,
	}
}

func NewBaseResponseWithData(code int, message string, success bool, data any) *ResponseWithData {
	return &ResponseWithData{
		Response: Response{
			Code:    code,
			Message: message,
			Success: true,
		},
		Data: data,
	}
}

func NewErrorResponse(code int, errorMessage string) *ErrorResponse {
	return &ErrorResponse{
		Response: Response{
			Code:    code,
			Success: false,
		},
		ErrorMessage: errorMessage,
	}
}
