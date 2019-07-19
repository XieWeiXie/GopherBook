package make_request

import (
	"encoding/json"
	"net/http"
)

func BindJson(request *http.Request, param interface{}) error {
	err := json.NewDecoder(request.Body).Decode(param)
	if err != nil {
		return err
	}
	return nil
}

func Query(request *http.Request, key string) string {
	return request.URL.Query().Get(key)
}

func QueryAndDefault(request *http.Request, key string, defaultValue string) string {
	value := Query(request, key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Vars(request *http.Request) map[string]string {
	all := request.URL.Query()
	var results = make(map[string]string)
	for k, i := range all {
		results[k] = i[0]
	}
	return results
}

func Param(request *http.Request, key string) (interface{}, error) {
	if err := request.ParseForm(); err != nil {
		return nil, err
	}
	return request.PostFormValue(key), nil
}

func Params(request *http.Request) (map[string][]string, error) {
	if err := request.ParseForm(); err != nil {
		return nil, err
	}
	return request.PostForm, nil
}
