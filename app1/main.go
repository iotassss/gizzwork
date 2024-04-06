package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// HelloWorld はレスポンスに使用するメッセージを保持します
type HelloWorld struct {
	Message string `json:"message"`
}

func fetchMessageFromApp2() (string, error) {
	// app2からのレスポンスを受け取る
	resp, err := http.Get("http://app2:80")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// レスポンスボディを読み取る
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// ここでbodyを必要に応じて処理する
	// 今回はそのまま文字列として返します
	return string(body), nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// app2からメッセージを受け取る
	message, err := fetchMessageFromApp2()
	if err != nil {
		http.Error(w, "Failed to fetch message from app2", http.StatusInternalServerError)
		return
	}

	// 受け取ったメッセージをHelloWorld構造体にセット
	response := HelloWorld{Message: message}

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
