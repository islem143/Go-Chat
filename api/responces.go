package api

type ApiResponseList struct {
	List any `json:"list"`
}

type ApiResponse struct {
	Message string `json:"message"`
}
