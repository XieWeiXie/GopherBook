package error_v1

var (

	// database
	ErrorDatabase       = ErrorV1{Code: 400, Detail: "数据库错误", Message: "database error"}
	ErrorRecordNotFound = ErrorV1{Code: 400, Detail: "记录不存在", Message: "record not found"}

	// body
	ErrorBodyJson   = ErrorV1{Code: 400, Detail: "请求消息体失败", Message: "read json body fail"}
	ErrorBodyIsNull = ErrorV1{Code: 400, Detail: "参数为空", Message: "body is null"}
)
