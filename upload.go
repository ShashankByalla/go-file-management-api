package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "net/http"
    "os"
    "time"
)

func uploadFileToS3(filePath string, fileName string) (string, error) {
    sess, _ := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2")},
    )
    svc := s3.New(sess)

    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    uploadInput := &s3.PutObjectInput{
        Bucket: aws.String("your-bucket-name"),
        Key:    aws.String(fileName),
        Body:   file,
    }

    _, err = svc.PutObject(uploadInput)
    if err != nil {
        return "", err
    }

    s3URL := "https://your-bucket-name.s3.amazonaws.com/" + fileName
    return s3URL, nil
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    file, header, err := r.FormFile("file")
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    defer file.Close()

    fileName := header.Filename
    filePath := "/tmp/" + fileName

    out, err := os.Create(filePath)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer out.Close()

    _, err = out.ReadFrom(file)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    s3URL, err := uploadFileToS3(filePath, fileName)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    _, err = db.Exec("INSERT INTO files (user_id, file_name, upload_date, file_size, s3_url) VALUES ($1, $2, $3, $4, $5)",
        1, fileName, time.Now(), header.Size, s3URL)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Write([]byte("File uploaded successfully: " + s3URL))
}
