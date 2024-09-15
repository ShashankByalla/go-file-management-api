package main

import (
    "time"
    "log"
)

func deleteExpiredFiles() {
    for {
        rows, err := db.Query("SELECT id, s3_url FROM files WHERE is_shared=true AND upload_date < NOW() - INTERVAL '30 days'")
        if err != nil {
            log.Println("Error fetching expired files:", err)
            time.Sleep(1 * time.Hour)
            continue
        }
        defer rows.Close()

        var fileID int
        var s3URL string
        for rows.Next() {
            err = rows.Scan(&fileID, &s3URL)
            if err != nil {
                log.Println("Error scanning file:", err)
                continue
            }

            // Delete file from S3
            err = deleteFileFromS3(s3URL)
            if err != nil {
                log.Println("Error deleting file from S3:", err)
                continue
            }

            // Remove metadata from the database
            _, err = db.Exec("DELETE FROM files WHERE id=$1", fileID)
            if err != nil {
                log.Println("Error deleting file from DB:", err)
            }
        }

        time.Sleep(24 * time.Hour) // Run the job daily
    }
}

func deleteFileFromS3(s3URL string) error {
    // Implement S3 file deletion logic here
    return nil
}

func initBackgroundJobs() {
    go deleteExpiredFiles() // Run file deletion in background
}
