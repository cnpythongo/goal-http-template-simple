package main

import (
	"goal-app/cmd/api/cmd"
	_ "goal-app/docs/api"
)

// @title 前台接口文档
// @version 1.0
//
// @description http状态码是200，code是0时表示正常返回；code不是200时表示有业务错误。返回的JSON数据由下面的结构包裹
// @description {
// @description     "code": 0, // 0是成功，其他是失败
// @description     "data": {object},  // 接口返回的成功数据
// @description     "msg": "ok"    // ok 或其他失败信息
// @description }
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	cmd.Execute()
}
