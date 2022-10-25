package admin

import (
	"github.com/cnpythongo/goal/test/utils"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	r := GetRouter()
	w := utils.DoRequest(r, "GET", "/api/ping", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	response := utils.ParseResponseToJSON(w)
	result, ok := response["code"]
	result = result.(float64)
	assert.Equal(t, ok, true)
	assert.Equal(t, result, float64(0))
}
