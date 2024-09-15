package main

import (
    "io"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "os"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/upload", UploadHandler).Methods("POST")
    r.HandleFunc("/files/{filename:[a-zA-Z0-9_\\.]+}", DownloadHandler).Methods("GET")
    r.HandleFunc("/register", RegisterHandler).Methods("POST")

    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    file, _, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Unable to get file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    out, err := os.Create("./uploads/" + r.FormValue("filename"))
    if err != nil {
        http.Error(w, "Unable to save file", http.StatusInternalServerError)
        return
    }
    defer out.Close()

    _, err = io.Copy(out, file)
    if err != nil {
        http.Error(w, "Unable to save file", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("File uploaded successfully"))
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    filename := vars["filename"]

    http.ServeFile(w, r, "./uploads/"+filename)
}

// RegisterHandler handles user registration requests
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    // Simple response for demonstration purposes
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Register endpoint hit"))
}
