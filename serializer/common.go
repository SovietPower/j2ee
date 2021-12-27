package serializer

// Response 基本回应结构
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataWithTotal 带有总数的Data结构
type DataWithTotal struct {
	Data  interface{} `json:"data"`
	Total uint        `json:"total"`
}

// TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// BuildDataWithTotal 带有总数的列表构建器
func BuildDataWithTotal(data interface{}, total uint) DataWithTotal {
	return DataWithTotal{
		Data:  data,
		Total: total,
	}
}
