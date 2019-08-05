package error_for_project

import "errors"

var (
	NotFound   = errors.New("field not found")
	ParamField = errors.New("forget post params ")
)
