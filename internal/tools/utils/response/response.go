package response

import "gotwitter/internal/types"

func OneOffErrorResponse(message string, title string) types.ApiResponse {
	errors := []types.Error{{Title: title, Message: message, Detail: "", Type: "internal exception"}}
	return types.ApiResponse{Errors: errors, Data: nil}
}

func ErrorResponse(message string, title string, detail string, err_type string) types.ApiResponse {
	errors := []types.Error{{Message: message, Title: title, Detail: detail, Type: err_type}}
	return types.ApiResponse{Errors: errors, Data: nil}
}

func ApiResponseFromData(data interface{}) types.ApiResponse {
	return types.ApiResponse{Errors: nil, Data: data}
}
