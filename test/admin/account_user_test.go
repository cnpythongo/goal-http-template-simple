package admin

import (
	"encoding/json"
	"fmt"
	"github.com/cnpythongo/goal/admin/types"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/test/utils"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	payload := map[string]interface{}{
		"phone":    "13800138000",
		"password": "123456",
	}
	r := GetRouter()
	w := utils.DoRequest(r, "POST", "/api/v1/account/login", payload)
	assert.Equal(t, http.StatusOK, w.Code)

	response := utils.ParseResponseToJSON(w)
	result, ok := response["code"]
	// fmt.Printf("%v\n", response)
	result = result.(float64)
	assert.Equal(t, ok, true)
	assert.Equal(t, result, float64(0))

	data, _ := json.Marshal(response["data"])
	var res types.RespAdminAuth
	_ = json.Unmarshal(data, &res)
	fmt.Printf("%v\n", res.Token)
}

func TestCreateUser(t *testing.T) {
	payload := model.User{
		Phone:    "13800138007",
		Password: "123123",
		Email:    "aaabbbddd@qq.com",
		Avatar:   "http://www.qq.com/aaa.jpg",
	}
	r := GetRouter()
	w := utils.DoRequest(r, "POST", "/api/v1/account/user/create", payload)
	assert.Equal(t, http.StatusOK, w.Code)

	response := utils.ParseResponseToJSON(w)
	result, ok := response["code"]
	fmt.Printf("%v\n", response)
	result = result.(float64)
	assert.Equal(t, ok, true)
	assert.Equal(t, result, float64(0))
}

func TestGetUserByUuid(t *testing.T) {
	r := GetRouter()
	uid := "3610b2e5ab0a43c6b909eece0cb1c167"
	w := utils.DoRequest(r, "GET", fmt.Sprintf("/api/v1/user/detail?uuid=%s", uid), nil)
	assert.Equal(t, http.StatusOK, w.Code)
	response := utils.ParseResponseToJSON(w)
	fmt.Printf("%v\n", response)
	result, ok := response["code"]
	result = result.(float64)
	assert.Equal(t, ok, true)
	assert.Equal(t, result, float64(0))
}
