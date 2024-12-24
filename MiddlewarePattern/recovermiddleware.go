package middlewarepattern

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// RecoveryMiddleware はパニックを捕捉して適切なエラーレスポンスを返すミドルウェア
// パニックが発生した場合、500 Internal Server Errorとともにエラー内容をJSONで返す
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            // パニックの捕捉
            if r := recover(); r != nil {
                // エラーメッセージをJSON形式に変換
                jsonBody, _ := json.Marshal(map[string]string{
                    "error": fmt.Sprintf("%v", r),
                })

                // JSONレスポンスを設定して返す
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                w.Write(jsonBody)
            }
        }()
        // 後続のハンドラを実行
        next.ServeHTTP(w, r)
    })
}