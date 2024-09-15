package main

import (
    "net/http"
    "encoding/json"
    "time"
)

func searchFilesHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("query")
    userID := 1 // replace with actual user ID from JWT claims

    rows, err := db.Query("SELECT id, file_name, upload_date, file_size, s3_url FROM files WHERE user_id=$1 AND file_name ILIKE $2", userID, "%"+query+"%")
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

    json.NewEncoder(w).Encode(files)
}
