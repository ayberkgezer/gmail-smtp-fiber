package base

type BaseResponse struct {
	Code    int    `json:"statusCode"`
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
}

type BaseResponseWithData struct {
	BaseResponse
	Data any `json:"data"`
}

type ErrorResponse struct {
	BaseResponse
	ErrorMessage string `json:"errorMessage"`
}

func NewBaseResponse(code int, message string) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: message,
		Success: true,
	}
}

func NewBaseResponseWithData(code int, message string, success bool, data any) *BaseResponseWithData {
	return &BaseResponseWithData{
		BaseResponse: BaseResponse{
			Code:    code,
			Message: message,
			Success: true,
		},
		Data: data,
	}
}

func NewErrorResponse(code int, errorMessage string) *ErrorResponse {
	return &ErrorResponse{
		BaseResponse: BaseResponse{
			Code:    code,
			Success: false,
		},
		ErrorMessage: errorMessage,
	}
}
