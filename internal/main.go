package main

import "OJ/internal/router"

func main() {
	r := router.Router()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
