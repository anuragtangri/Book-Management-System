package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	err := json.Unmarshal(body, &x)
	fmt.Println(x)
	if err == nil {
		return
	}
}
