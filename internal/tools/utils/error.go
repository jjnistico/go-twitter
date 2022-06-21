package utils

import "gotwitter/internal/types"

func OneOffErrorResponse(message string, title string) types.ApiResponse {
	errors := []types.Error{{Title: title, Message: message, Detail: "", Type: ""}}
	return types.ApiResponse{Errors: errors, Data: nil}
}

func ErrorResponse(messages []string, title string) types.ApiResponse {
	errors := []types.Error{}
	for _, message := range messages {
		errors = append(errors, types.Error{Title: title, Message: message, Detail: "", Type: "missing query parameter"})
	}
	return types.ApiResponse{Errors: errors, Data: nil}
}

func ApiResponseFromData(data interface{}) types.ApiResponse {
	return types.ApiResponse{Errors: nil, Data: data}
}
