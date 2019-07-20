package error_votes

import (
	"fmt"
	"net/http"
)

type ErrorForVotes struct {
	Code      int    `json:"code"`
	Detail    string `json:"detail"`
	Message   string `json:"message"`
	MessageZh string `json:"message_zh"`
}

func (e ErrorForVotes) Error() string {
	return fmt.Sprintf("Code: %d, Detail: %s, Message: %s, MessageZh: %s",
		e.Code, e.Detail, e.Message, e.MessageZh)
}

var (
	ErrorParam = ErrorForVotes{
		Code:      http.StatusBadRequest,
		Detail:    "param fail",
		Message:   "param invalid",
		MessageZh: "参数校验失效"}
	ErrorInsert = ErrorForVotes{
		Code:      http.StatusBadRequest,
		Detail:    "insert data fail",
		Message:   "insert data fail",
		MessageZh: "记录插入失败"}
)
