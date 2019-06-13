package error_v1

var (
	// 200

	// 300

	// 400

	// 500

	// database
	ErrorDatabase       = ErrorV1{Code: 400, Detail: "数据库错误", Message: "database error"}
	ErrorRecordNotFound = ErrorV1{Code: 400, Detail: "记录不存在", Message: "record not found"}
)
