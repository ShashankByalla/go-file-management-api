# 🗂️ Go File Management API

![Go Build & Test](https://github.com/ShashankByalla/go-file-management-api/actions/workflows/go-test.yml/badge.svg?branch=main)
![Language](https://img.shields.io/badge/Language-Go-00ADD8)
![Database](https://img.shields.io/badge/Database-PostgreSQL-336791)
![Storage](https://img.shields.io/badge/Storage-AWS%20S3-FF9900)

A high-performance RESTful backend API built in Go for secure 
file management — featuring user authentication, AWS S3 file 
storage, full-text search, file sharing, and background processing.

## ✨ Features
- 🔐 JWT-based user authentication (register/login)
- 📁 File upload & retrieval with AWS S3
- 🔍 File search by name
- 🤝 File sharing between users
- ⚙️ Background job processing
- 🧪 Unit tested with Go test suite
- 🗄️ PostgreSQL database

## 🛠️ Tech Stack
| Layer | Technology |
|-------|-----------|
| Language | Go 1.18+ |
| Database | PostgreSQL |
| File Storage | AWS S3 |
| Auth | JWT |
| Testing | Go test |

## 📡 API Endpoints

### Authentication
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/register` | POST | Register new user |
| `/login` | POST | Login & get JWT token |

### File Management
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/upload` | POST | Upload file to S3 |
| `/files` | GET | List all files |
| `/files/search` | GET | Search files by name |
| `/files/share` | POST | Share file with user |

## 🚀 Getting Started

### Prerequisites
- Go 1.18+
- PostgreSQL
- AWS account (for S3)

### Setup
```bash
git clone https://github.com/ShashankByalla/go-file-management-api.git
cd go-file-management-api
go mod tidy
go run main.go
```

### Run Tests
```bash
go test ./...
```
