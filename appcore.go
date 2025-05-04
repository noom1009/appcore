package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "github.com/golang-jwt/jwt/v5" 
    "github.com/noom1009/appcore/config"
    "github.com/noom1009/appcore/database"
    "github.com/noom1009/appcore/logger"
    "github.com/noom1009/appcore/redis"
    "github.com/noom1009/appcore/storage"
    "github.com/noom1009/appcore/messagequeue"
    "github.com/noom1009/appcore/observability"
    "github.com/noom1009/appcore/healthcheck"
    "github.com/noom1009/appcore/middlewares"
)

var jwtSecret []byte

// ฟังก์ชันสร้าง JWT Token สำหรับทดสอบ
func GenerateJWT() (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": 123,
        "role":    "admin",
        "exp":     time.Now().Add(time.Hour * 8).Unix(), 
    })

    return token.SignedString(jwtSecret)
}

// Handler สำหรับการทดสอบ JWT
func protectedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "🎉 Access granted to protected route!")
}

func main() {
    // โหลด config
    config.LoadConfig(".env.yaml")

    // โหลดค่า JWT_SECRET จาก .env.yaml
    jwtSecret = []byte(config.AppConfig.JWTSecret)

    // เริ่มต้น logger
    logger.InitLogger()
    fmt.Println("Initializing appcore...")

    // เชื่อมต่อ Redis
    redis.InitRedis(config.AppConfig.Redis.Addr)

    // สร้าง DSN แล้วเชื่อมต่อ Postgres
    dsn := "postgres://" +
        config.AppConfig.DB.User + ":" +
        config.AppConfig.DB.Pass + "@" +
        config.AppConfig.DB.Host + ":" +
        fmt.Sprint(config.AppConfig.DB.Port) + "/" +
        config.AppConfig.DB.Name +
        "?sslmode=disable"

    database.InitPostgres(dsn)

    messagequeue.InitKafka("localhost:9092", "app-topic")
    observability.InitTracer("AppCore")
    healthcheck.StartHealthCheck("8081")
    // เรียกใช้ฟังก์ชันสร้าง Bucket
    storage.CreateBucket()

    log.Println("✅ MinIO bucket creation process completed.")

    // สร้าง JWT Token (ตัวอย่างการทดสอบ)
    token, err := GenerateJWT()
    if err != nil {
        log.Fatalf("❌ Failed to generate JWT: %v", err)
    }
    log.Printf("✅ Generated JWT: %s", token)

    // ตั้งค่า handler สำหรับทดสอบ JWT middleware
    http.Handle("/protected", middlewares.JWTMiddleware(http.HandlerFunc(protectedHandler)))

    fmt.Println("✅ AppCore initialized successfully")
    log.Println("✅ Listening on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}