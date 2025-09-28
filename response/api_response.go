package response

// APIResponse общий формат ответа
type APIResponse struct {
	Status  bool        `json:"status" example:"true"`
	Message string      `json:"message" example:"ok"`
	Result  interface{} `json:"result,omitempty"`
}

// ErrorResponse общий формат ошибки
type ErrorResponse struct {
	Status  bool   `json:"status" example:"false"`
	Message string `json:"message" example:"something went wrong"`
}
