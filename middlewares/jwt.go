// middlewares/jwt.go
package middlewares

import (
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") // สามารถเปลี่ยนเป็นการดึงจาก config ได้

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
            return
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrInvalidKeyType
            }
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}

// package middlewares

// import (
//     "net/http"
//     "strings"

//     "github.com/golang-jwt/jwt/v5"
// )

// var jwtSecret = []byte("your-secret-key") // ใช้ config ได้เช่นกัน

// func JWTMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         authHeader := r.Header.Get("Authorization")
//         if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
//             http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
//             return
//         }

//         tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

//         token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
//             if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//                 return nil, jwt.ErrInvalidKeyType
//             }
//             return jwtSecret, nil
//         })

//         if err != nil || !token.Valid {
//             http.Error(w, "Invalid token", http.StatusUnauthorized)
//             return
//         }

//         next.ServeHTTP(w, r)
//     })
// }
