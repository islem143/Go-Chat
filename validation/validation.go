package validation

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (g *GlobalErrorHandlerResp) Error() string {
	return g.Message
}
