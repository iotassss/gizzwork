package main

import (
	"encoding/json"
	"net/http"
)

// HelloWorld はレスポンスに使用するメッセージを保持します
type HelloWorld struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloWorld{Message: "Hello World from app2"}

	// Content-Typeをapplication/jsonに設定
	w.Header().Set("Content-Type", "application/json")

	// HelloWorld構造体をJSONにエンコードしてレスポンスに書き込む
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// / パスへのリクエストをhelloHandler関数にルーティング
	http.HandleFunc("/", helloHandler)

	// サーバーを80ポートで起動
	// 注意: 80ポートを使用するには管理者権限が必要です
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
