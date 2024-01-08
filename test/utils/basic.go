package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func DoRequest(r http.Handler, method, path string, data interface{}) *httptest.ResponseRecorder {
	var reqBody *bytes.Buffer
	if data != nil {
		body, _ := json.Marshal(data)
		reqBody = bytes.NewBuffer(body)
	} else {
		body, _ := json.Marshal(make(map[string]string))
		reqBody = bytes.NewBuffer(body)
	}
	req, _ := http.NewRequest(method, path, reqBody)
	w := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AuthToken))
	r.ServeHTTP(w, req)
	return w
}

func ParseResponseToJSON(w *httptest.ResponseRecorder) map[string]interface{} {
	var response map[string]interface{}
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	return response
}
