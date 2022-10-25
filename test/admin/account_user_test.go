package admin

import (
	"fmt"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/test/utils"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {
	payload := model.User{
		Username: "lyh333555",
		Password: "123123",
		Email:    "aaabbbddd@qq.com",
		Avatar:   "http://www.qq.com/aaa.jpg",
	}
	r := GetRouter()
	w := utils.DoRequest(r, "POST", "/api/account/users", payload)
	assert.Equal(t, http.StatusOK, w.Code)

	response := utils.ParseResponseToJSON(w)
	result, ok := response["code"]
	fmt.Printf("%v\n", response)
	result = result.(float64)
	assert.Equal(t, ok, true)
	assert.Equal(t, result, float64(1000))
}

func TestGetUserByUuid(t *testing.T) {
	r := GetRouter()
	uid := "3610b2e5ab0a43c6b909eece0cb1c167"
	w := utils.DoRequest(r, "GET", fmt.Sprintf("/api/users/%s", uid), nil)
	assert.Equal(t, http.StatusOK, w.Code)
	response := utils.ParseResponseToJSON(w)
	fmt.Printf("%v\n", response)
	result, ok := response["code"]
	result = result.(float64)
	assert.Equal(t, ok, true)
	assert.Equal(t, result, float64(1000))
}
