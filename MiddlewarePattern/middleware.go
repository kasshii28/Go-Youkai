package middlewarepattern

import (
    "log"
    "net/http"
    "time"
)

// MyMiddleWare はHTTPリクエストの処理時間を計測するミドルウェア
// 各リクエストの開始時刻と処理にかかった時間をログに出力する
func MyMiddleWare(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // リクエスト開始時刻を記録
        start := time.Now()
        // 後続のハンドラを実行
        h.ServeHTTP(w, r)
        // 処理時間を計算してログ出力
        d := time.Now().Sub(start).Milliseconds()
        log.Printf("end %s(%d ms)\n", start.Format(time.RFC3339), d)
    })
}