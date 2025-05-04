package healthcheck

import (
    "fmt"
    "net/http"
)

func StartHealthCheck(port string) {
    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "ok")
    })
    go http.ListenAndServe(":"+port, nil)
    fmt.Println("✅ Healthcheck running on /healthz at port", port)
}
