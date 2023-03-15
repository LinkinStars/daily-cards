package handler

type RespBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrRespBody(msg string) *RespBody {
	return &RespBody{
		Code:    -1,
		Message: msg,
	}
}

func SucceedRespBody(msg string) *RespBody {
	return &RespBody{
		Code:    200,
		Message: msg,
	}
}

func SucceedRespBodyAndData(msg string, data interface{}) *RespBody {
	return &RespBody{
		Code:    200,
		Message: msg,
		Data:    data,
	}
}

func NewRespBody(code int, msg string, data interface{}) *RespBody {
	return &RespBody{
		Code:    200,
		Message: msg,
		Data:    data,
	}
}
