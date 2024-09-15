# 21BCE9151_Backend

This is a backend project for managing user authentication, file uploads, and file management. It is implemented in Go and uses PostgreSQL for database management.

## Project Structure

- `main.go`: Main entry point of the application.
- `auth.go`: Handles user authentication.
- `db.go`: Database connection and initialization.
- `background.go`: Background job processing.
- `file_retrieval.go`: File retrieval and management logic.
- `upload.go`: Handles file uploads and integration with S3.
- `search.go`: Search functionality for files.
- `share.go`: File sharing logic.
- `main_test.go`: Unit tests for various handlers.

## Prerequisites

- Go 1.18 or later
- PostgreSQL
- (Optional) AWS S3 for file storage

## Setup

### 1. Clone the Repository


```git clone https://github.com/ShashankByalla/21BCE9151_Backend.git```
```cd 21BCE9151_Backend ```

2. Install Dependencies

Make sure you have Go installed. Then run:
go mod tidy

3. Configure PostgreSQL

Ensure PostgreSQL is installed and running. Create a database and a user for the project:
CREATE DATABASE file_management_db;
CREATE USER file_user WITH ENCRYPTED PASSWORD 'file_password';
GRANT ALL PRIVILEGES ON DATABASE file_management_db TO file_user;

Update the db.go file with your database credentials if they differ from the defaults.

4. Update Configuration

Edit upload.go to set the correct AWS S3 bucket name and region, if you are using S3:

// Update with your S3 bucket name and region
Bucket: aws.String("your-bucket-name"),
Region: aws.String("us-west-2"),

5. Run the Application

Start the application using:
go run main.go

The server will start on port 8080 by default. You can change this in the main.go file if needed.

Endpoints

User Authentication

	•	POST /register: Register a new user.
	•	POST /login: Login a user.

File Management

	•	POST /upload: Upload a file. Requires a file in the form-data under the key file.
	•	GET /files: Retrieve a list of files.
	•	GET /files/search: Search for files by name.
	•	POST /files/share: Share a file with another user.

Testing

To run unit tests:
go test ./...

Troubleshooting

If you encounter issues:

	•	Ensure PostgreSQL is running and accessible.
	•	Check that AWS credentials and S3 bucket details are correctly configured.
	•	Verify that all required environment variables are set.

Contributing

Feel free to fork the repository, create a pull request, and contribute to improving the project.

License

This project is licensed under the MIT License - see the LICENSE file for details.

Contact

For any questions, please contact Shashank Byalla.
