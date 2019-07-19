package make_response

import (
	"encoding/json"
	"log"
	"net/http"
)

func Response(w http.ResponseWriter, code int, data interface{}) {
	var results = make(map[string]interface{})
	results["code"] = code
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	if code == http.StatusOK {
		results["data"] = data
	} else {
		results["error"] = data
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "")
	err := enc.Encode(results)
	if err != nil {
		log.Println("err : ", err.Error())
		return
	}

}
