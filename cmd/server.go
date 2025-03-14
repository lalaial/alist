package main

import (
	"log"
	"net/http"
	"github.com/Xhofe/alist/server"
)

func main() {
	// 初始化 Alist
	if err := server.Init(); err != nil {
		log.Fatalf("Failed to initialize Alist: %v", err)
	}

	// 启动 HTTP 服务器
	http.HandleFunc("/", server.HandleRequest)
	log.Println("Alist is running on Vercel")
	http.ListenAndServe(":3000", nil)
}
