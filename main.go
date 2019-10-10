package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tamanobi/go-tutorial/calc"
	"github.com/tamanobi/go-tutorial/checkdigit"
)

// CheckDigitResponse はレスポンス
type CheckDigitResponse struct {
	Error bool   `json:"error"`
	Input string `json:"input"`
	Valid bool   `json:"valid"`
}

type NameResponse struct {
	Name string `json:"name"`
}

func nameHandler(w http.ResponseWriter, req *http.Request) {

}

func main() {
	rootHandler := func(w http.ResponseWriter, req *http.Request) {
		q := req.URL.Query()
		numbers, ok := q["numbers"]
		if q == nil || !ok || len(numbers) == 0 {
			http.Error(w, "numbers required", http.StatusBadRequest)
			return
		}

		input := numbers[0]
		m, err := json.Marshal(CheckDigitResponse{
			true,
			input,
			checkdigit.Luhn(input),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(m)
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/simple", func(w http.ResponseWriter, req *http.Request) {
		// レスポンスがJSONであると、クライアントへ教える
		w.Header().Set("Content-Type", "application/json")
		// 空の構造体(struct{}{})をJSONに変換する
		m, err := json.Marshal(struct{}{})
		// 変換処理のエラーチェック
		if err != nil {
			// エラーだったら、InternalServerErrorを返す
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(m) // JSONをサーバーから返す
	})

	// サーバーを実行する
	log.Fatal(http.ListenAndServe(":8080", nil))
}
