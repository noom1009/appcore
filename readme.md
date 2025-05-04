### 🔧 โครงสร้างโฟลเดอร์ appcore
```
appcore/
│
├── config/               # โหลดและจัดการ config เช่น .env, yaml
│   └── config.go
│
├── database/             # เชื่อมต่อกับฐานข้อมูล เช่น PostgreSQL, MongoDB
│   ├── postgres.go
│   └── mongo.go
│
├── redis/                # Client สำหรับ Redis
│   └── redis.go
│
├── logger/               # ระบบ log เช่น zap หรือ zerolog
│   └── logger.go
│
├── middleware/           # Middleware สำหรับ HTTP เช่น JWT, CORS
│   ├── auth.go
│   └── cors.go
│
├── utils/                # ฟังก์ชันที่ใช้ทั่วไป เช่น validate, random, time
│   ├── validator.go
│   └── helper.go
│
├── httpclient/           # Wrapper สำหรับ HTTP call ภายนอก
│   └── client.go
│
├── grpcclient/           # Wrapper สำหรับ gRPC client
│   └── client.go
│
├── messagequeue/         # Kafka, RabbitMQ, NATS
│   └── kafka.go
│
├── storage/              # จัดการ S3, MinIO ฯลฯ
│   └── s3.go
│
├── observability/        # SigNoz, OpenTelemetry
│   └── tracing.go
│
├── security/           # JWT, Password Hashing, Auth Helpers
│   ├── jwt.go
│   └── hash.go
│
├── appcore.go            # ฟังก์ชันหลักสำหรับ init ทั้งหมด
└── go.mod
```

✅ PostgreSQL

✅ Redis

✅ RabbitMQ

✅ Kafka (พร้อม UI เช่น Kafka UI หรือ Kafdrop)

✅ MinIO (แทน S3 สำหรับ local dev ใช้ร่วมกับ SDK AWS ได้)

✅ Network และ Environment ที่ match กับ .env.yaml

```
docker-compose up -d
```

URLs และ Credentials ที่คุณใช้ได้ใน Local
Service	URL/Port	Username/Password
PostgreSQL	localhost:5432	postgres / password
Redis	localhost:6379	–
RabbitMQ	localhost:5672	guest / guest
RabbitMQ UI	http://localhost:15672	guest / guest
Kafka Broker	localhost:9092	–
Kafka UI	http://localhost:8081	–
MinIO API	http://localhost:9000	your-access-key / your-secret-key
MinIO UI	http://localhost:9001	your-access-key / your-secret-key


วิธีที่ 1: สร้าง Bucket ผ่าน MinIO Console UI
เมื่อคุณรัน MinIO แล้วสามารถเข้าถึง Console UI ผ่าน http://localhost:9001.

เข้าระบบโดยใช้ MINIO_ROOT_USER และ MINIO_ROOT_PASSWORD ที่คุณตั้งค่าในไฟล์ docker-compose.yml.

เมื่อเข้าสู่ระบบแล้ว คุณจะเห็นหน้าหลักของ MinIO Console.

คลิกที่ปุ่ม + Create Bucket.

กรอกชื่อ bucket ที่คุณต้องการสร้าง เช่น mybucket แล้วกด Create.

หลังจากนั้น bucket ใหม่ก็จะถูกสร้างขึ้นใน MinIO และคุณสามารถเริ่มใช้งานได้ทันที

วิธีที่ 2: สร้าง Bucket ผ่าน MinIO Client (mc)
ติดตั้ง MinIO Client (mc) ถ้ายังไม่ได้ติดตั้ง:

คุณสามารถดาวน์โหลดได้จากที่ MinIO Client (mc)

หรือถ้าใช้ Docker ก็สามารถใช้คำสั่ง mc ผ่าน Docker ได้

กำหนด alias สำหรับ MinIO:

bash
Copy
Edit
mc alias set myminio http://localhost:9000 your-access-key your-secret-key
myminio เป็นชื่อ alias ที่ตั้งขึ้น

http://localhost:9000 คือ URL ของ MinIO

your-access-key และ your-secret-key คือค่าที่คุณกำหนดในไฟล์ docker-compose.yml

สร้าง bucket:

bash
Copy
Edit
mc mb myminio/mybucket
mybucket คือชื่อของ bucket ที่คุณต้องการสร้าง

ตรวจสอบว่า bucket ถูกสร้างขึ้นหรือไม่:

bash
Copy
Edit
mc ls myminio

