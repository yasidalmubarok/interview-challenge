package helper

type Meta struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

type Data struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func APIResponse(status, message string, code int, data interface{}) Data {
	return Data{
		Meta: Meta{
			Status:  status,
			Message: message,
			Code:    code,
		},
		Data: data,
	}
}