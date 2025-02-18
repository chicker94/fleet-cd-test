package main

import (
	"fmt"
	"net/http"
)

// 환경 설정
const ENV = "dev" // "prod"로 변경 가능

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Rancher CI/CD Test (%s)</h1>", ENV)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

