package storage

import (
    "context"
    "fmt"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
    "github.com/noom1009/appcore/config"
)

func CreateBucket() {
    cfg := config.AppConfig.S3
    bucketName := cfg.Bucket

    client, err := minio.New("localhost:9000", &minio.Options{
        Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
        Secure: false,
    })
    if err != nil {
        fmt.Printf("❌ Failed to create MinIO client: %v", err)
    }

    exists, err := client.BucketExists(context.Background(), bucketName)
    if err != nil {
        fmt.Printf("❌ Failed to check if bucket exists: %v", err)
    }

    if !exists {
        err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
        if err != nil {
            fmt.Printf("❌ Failed to create bucket: %v", err)
        }
        fmt.Printf("✅ Bucket %s created successfully.", bucketName)
    } else {
        fmt.Printf("ℹ️ Bucket %s already exists.", bucketName)
    }
}

