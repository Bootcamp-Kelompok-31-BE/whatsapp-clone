package models

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type BasePaginationDataResponse struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Size  int `json:"size"`
}

type BasePaginationResponse struct {
	BaseResponse

	Pagination BasePaginationDataResponse `json:"pagination"`
}

func NewOkResponse(data any) BaseResponse {
	return BaseResponse{
		Success: true,
		Message: "success",
		Data:    data,
	}
}

func NewErrResponse(message string) BaseResponse {
	return BaseResponse{
		Success: false,
		Message: message,
		Data:    nil,
	}
}

func NewOkPaginationResponse(data any, total int, page int, size int) BasePaginationResponse {
	return BasePaginationResponse{
		BaseResponse: BaseResponse{
			Success: true,
			Message: "success",
			Data:    data,
		},
		Pagination: BasePaginationDataResponse{
			Total: total,
			Page:  page,
			Size:  size,
		},
	}
}

func NewErrPaginationResponse(message string) BasePaginationResponse {
	return BasePaginationResponse{
		BaseResponse: BaseResponse{
			Success: false,
			Message: message,
			Data:    nil,
		},
		Pagination: BasePaginationDataResponse{
			Total: 0,
			Page:  0,
			Size:  0,
		},
	}
}
