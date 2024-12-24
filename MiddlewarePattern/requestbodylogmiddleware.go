package middlewarepattern

import (
    "bytes"
    "io"
    "log"
    "net/http"

    "go.uber.org/zap"
)

// RequestBodyLogMiddleware はHTTPリクエストのボディをログに記録するミドルウェア
// リクエストボディを読み取った後、再度読み取り可能な状態に復元する
func RequestBodyLogMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // リクエストボディを読み取り
        body, err := io.ReadAll(r.Body)
        if err != nil {
            log.Printf("Failed to log request body: %v", zap.Error(err))
            http.Error(w, "Failed to log request body", http.StatusBadRequest)
            return 
        }
        defer r.Body.Close()
        
        // ボディを再度読み取り可能な状態に復元
        // io.NopCloserを使用して新しいReadCloserを作成
        r.Body = io.NopCloser(bytes.NewBuffer(body))
        
        // 後続のハンドラにリクエストを渡す
        next.ServeHTTP(w, r)
    })
}