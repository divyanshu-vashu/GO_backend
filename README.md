
# Go Backend Project

## Overview

This project is a backend service built with Go, designed to handle file uploads, management, and metadata management. It provides APIs for uploading, retrieving, and sharing files, along with search and caching capabilities. The service is connected to a PostgreSQL database and can utilize S3 for file storage.
<img width="1280" alt="Screenshot 2024-09-15 at 11 12 47" src="https://github.com/user-attachments/assets/44e3e3ff-3a05-42c0-8c64-af668068353d">


## Features

1. **File Upload & Management**
   - **Task:** Allows users to upload files (e.g., documents, images) to S3 or local storage and manage file metadata.
   - **Endpoints:**
     - `POST /upload`: Uploads files and saves metadata to PostgreSQL. Returns a public URL to access the file.
   - **Concurrency:** Uses goroutines to handle large uploads efficiently.

2. **File Retrieval & Sharing**
   - **Task:** Enables users to retrieve metadata for uploaded files and share files via public links.
   - **Endpoints:**
     - `GET /files`: Retrieves metadata for uploaded files.
     - `GET /share/:file_id`: Provides a public link to share the file.

3. **File Search**
   - **Task:** Allows users to search for files based on metadata.
   - **Requirements:** Users can search by name, upload date, or file type with optimized performance for large datasets.

4. **Caching Layer for File Metadata**
   - **Task:** Implements a caching mechanism using Redis to reduce database load.
   - **Requirements:**
     - Cache file metadata on retrieval.
     - Invalidate cache when metadata is updated.
     - Ensure cache refreshes automatically after expiry (e.g., 5 minutes).

5. **Database Interaction**
   - **Task:** Manages file metadata, S3 locations, and user data in PostgreSQL.
   - **Requirements:**
     - Create tables for users and files.
     - Ensure efficient queries for retrieving user-specific files.
     - Handle database transactions for critical operations (e.g., file upload).
<img width="1280" alt="Screenshot 2024-09-15 at 11 13 03" src="https://github.com/user-attachments/assets/3e1552ef-46bc-4f47-b77c-810b1a693971">
<img width="1280" alt="Screenshot 2024-09-15 at 11 12 40" src="https://github.com/user-attachments/assets/4d57359e-3da6-4ee4-ae2d-b799686f417b">
## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/divyanshu-vashu/go-backend.git
   ```

2. **Navigate to the Project Directory:**

   ```bash
   cd go-backend
   ```

3. **Install Dependencies:**

   ```bash
   go mod tidy
   ```

4. **Run the Application:**

   ```bash
   go run main.go
   ```

5. **Setup the Database:**

   - Ensure you have PostgreSQL installed and running.
   - Configure the database connection details in your environment variables or configuration file.
   - Create the necessary tables and schemas in PostgreSQL based on the provided schema files or migration scripts.

## Configuration

The application requires certain environment variables to be set. Create a `.env` file in the root of the project and configure it as follows:

```env
DATABASE_URL=postgres://user:password@localhost:5432/dbname
S3_BUCKET_NAME=mybucket
REDIS_URL=redis://localhost:6379
# Add other environment variables as needed
```

Replace placeholders with your actual configuration values.

## Running Tests

To run tests for the project, use:

```bash
go test ./...
```

## Building for Production

To build the project for production, compile the Go code with:

```bash
go build -o myapp
```

## Contributing

If you wish to contribute to this project, please fork the repository and submit a pull request with your changes. Ensure to follow the code style and include tests where applicable.

## License

This project is licensed under the [MIT License](LICENSE).

---

Feel free to adjust or expand on any sections based on your specific needs!
