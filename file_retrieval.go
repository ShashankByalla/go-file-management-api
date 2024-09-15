package main

import (
    "net/http"
    "log"
    "encoding/json"
    "github.com/go-redis/redis/v8"
    "context"
    "time"
)

var redisClient *redis.Client
var ctx = context.Background()

func initRedis() {
    redisClient = redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis address
    })
}

func getFilesHandler(w http.ResponseWriter, r *http.Request) {
    userID := 1 // replace with actual user ID from JWT claims

    cacheKey := "files_user_" + string(userID)
    cachedFiles, err := redisClient.Get(ctx, cacheKey).Result()
    if err == redis.Nil {
        rows, err := db.Query("SELECT id, file_name, upload_date, file_size, s3_url FROM files WHERE user_id=$1", userID)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var files []map[string]interface{}
        for rows.Next() {
            var file map[string]interface{}
            var fileName, s3URL string
            var uploadDate time.Time
            var fileSize int64
            var id int
            err = rows.Scan(&id, &fileName, &uploadDate, &fileSize, &s3URL)
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            file = map[string]interface{}{
                "id":         id,
                "file_name":  fileName,
                "upload_date": uploadDate,
                "file_size":  fileSize,
                "s3_url":     s3URL,
            }
            files = append(files, file)
        }

        // Cache the result in Redis for 5 minutes
        cachedFilesBytes, _ := json.Marshal(files)
        redisClient.Set(ctx, cacheKey, string(cachedFilesBytes), 5*time.Minute)

        json.NewEncoder(w).Encode(files)
    } else if err == nil {
        w.Write([]byte(cachedFiles)) // Serve cached data
    } else {
        w.WriteHeader(http.StatusInternalServerError)
    }
}

func shareFileHandler(w http.ResponseWriter, r *http.Request) {
    fileID := r.URL.Query().Get("file_id")
    userID := 1 // replace with actual user ID from JWT claims

    var fileURL string
    err := db.QueryRow("SELECT s3_url FROM files WHERE id=$1 AND user_id=$2", fileID, userID).Scan(&fileURL)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    sharedURL := "https://your-app.com/share/" + fileID
    _, err = db.Exec("UPDATE files SET shared_url=$1, is_shared=true WHERE id=$2", sharedURL, fileID)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Write([]byte("File shared successfully: " + sharedURL))
}
