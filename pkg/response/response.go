package response

// BaseResponse is the standard response format for all API endpoints
type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Object  interface{} `json:"object,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

// PaginatedResponse is used for endpoints that return paginated data
type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Object     interface{} `json:"object,omitempty"`
	PageNumber int         `json:"pageNumber"`
	PageSize   int         `json:"pageSize"`
	TotalSize  int         `json:"totalSize"`
	Errors     []string    `json:"errors,omitempty"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(message string, data interface{}) BaseResponse {
	return BaseResponse{
		Success: true,
		Message: message,
		Object:  data,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(message string, errors []string) BaseResponse {
	return BaseResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	}
}

// NewPaginatedResponse creates a new paginated response
func NewPaginatedResponse(message string, data interface{}, page, pageSize, totalSize int) PaginatedResponse {
	return PaginatedResponse{
		Success:    true,
		Message:    message,
		Object:     data,
		PageNumber: page,
		PageSize:   pageSize,
		TotalSize:  totalSize,
	}
}
