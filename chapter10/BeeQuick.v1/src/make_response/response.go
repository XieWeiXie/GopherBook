package make_response

func MakeResponse(code int, value interface{}, isError bool) map[string]interface{} {
	result := make(map[string]interface{})
	result["code"] = code
	if isError {
		result["error"] = value
	} else {
		result["data"] = value
	}
	return result
}
